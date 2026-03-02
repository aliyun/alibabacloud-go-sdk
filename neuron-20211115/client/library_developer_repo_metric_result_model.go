// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryDeveloperRepoMetricResult interface {
	dara.Model
	String() string
	GoString() string
	SetDeveloperRepoMetrics(v []*ReposDeveloperGroupMetric) *LibraryDeveloperRepoMetricResult
	GetDeveloperRepoMetrics() []*ReposDeveloperGroupMetric
}

type LibraryDeveloperRepoMetricResult struct {
	DeveloperRepoMetrics []*ReposDeveloperGroupMetric `json:"developerRepoMetrics,omitempty" xml:"developerRepoMetrics,omitempty" type:"Repeated"`
}

func (s LibraryDeveloperRepoMetricResult) String() string {
	return dara.Prettify(s)
}

func (s LibraryDeveloperRepoMetricResult) GoString() string {
	return s.String()
}

func (s *LibraryDeveloperRepoMetricResult) GetDeveloperRepoMetrics() []*ReposDeveloperGroupMetric {
	return s.DeveloperRepoMetrics
}

func (s *LibraryDeveloperRepoMetricResult) SetDeveloperRepoMetrics(v []*ReposDeveloperGroupMetric) *LibraryDeveloperRepoMetricResult {
	s.DeveloperRepoMetrics = v
	return s
}

func (s *LibraryDeveloperRepoMetricResult) Validate() error {
	if s.DeveloperRepoMetrics != nil {
		for _, item := range s.DeveloperRepoMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
