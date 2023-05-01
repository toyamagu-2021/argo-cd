package e2e

import (
	"testing"

	"github.com/argoproj/argo-cd/v2/common"
	. "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	. "github.com/argoproj/argo-cd/v2/test/e2e/fixture/app"
	"github.com/argoproj/gitops-engine/pkg/health"
)

const (
	failureManifestPath = "failure-during-sync"
)

func TestDontSkipHealthCheck(t *testing.T) {
	Given(t).
		Path(failureManifestPath).
		When().
		// app should have no status
		CreateFromFile(func(app *Application) {
			app.Annotations = map[string]string{common.AnnotationKeySkipHealthCheck: "not-skip"}
		}).
		Then().
		Expect(HealthIs(health.HealthStatusDegraded))
}
