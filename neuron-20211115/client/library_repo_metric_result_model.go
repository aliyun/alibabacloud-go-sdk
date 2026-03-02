// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryRepoMetricResult interface {
	dara.Model
	String() string
	GoString() string
	SetRepoMetrics(v []*RepoMetric) *LibraryRepoMetricResult
	GetRepoMetrics() []*RepoMetric
	SetRequestId(v string) *LibraryRepoMetricResult
	GetRequestId() *string
}

type LibraryRepoMetricResult struct {
	RepoMetrics []*RepoMetric `json:"repoMetrics,omitempty" xml:"repoMetrics,omitempty" type:"Repeated"`
	RequestId   *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LibraryRepoMetricResult) String() string {
	return dara.Prettify(s)
}

func (s LibraryRepoMetricResult) GoString() string {
	return s.String()
}

func (s *LibraryRepoMetricResult) GetRepoMetrics() []*RepoMetric {
	return s.RepoMetrics
}

func (s *LibraryRepoMetricResult) GetRequestId() *string {
	return s.RequestId
}

func (s *LibraryRepoMetricResult) SetRepoMetrics(v []*RepoMetric) *LibraryRepoMetricResult {
	s.RepoMetrics = v
	return s
}

func (s *LibraryRepoMetricResult) SetRequestId(v string) *LibraryRepoMetricResult {
	s.RequestId = &v
	return s
}

func (s *LibraryRepoMetricResult) Validate() error {
	if s.RepoMetrics != nil {
		for _, item := range s.RepoMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
