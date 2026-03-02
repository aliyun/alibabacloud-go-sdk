// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcDeveloperRepoMetricResult interface {
	dara.Model
	String() string
	GoString() string
	SetDeveloperRepoMetrics(v []*ReposDeveloperGroupMetric) *PbcDeveloperRepoMetricResult
	GetDeveloperRepoMetrics() []*ReposDeveloperGroupMetric
	SetRequestId(v string) *PbcDeveloperRepoMetricResult
	GetRequestId() *string
}

type PbcDeveloperRepoMetricResult struct {
	DeveloperRepoMetrics []*ReposDeveloperGroupMetric `json:"developerRepoMetrics,omitempty" xml:"developerRepoMetrics,omitempty" type:"Repeated"`
	RequestId            *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcDeveloperRepoMetricResult) String() string {
	return dara.Prettify(s)
}

func (s PbcDeveloperRepoMetricResult) GoString() string {
	return s.String()
}

func (s *PbcDeveloperRepoMetricResult) GetDeveloperRepoMetrics() []*ReposDeveloperGroupMetric {
	return s.DeveloperRepoMetrics
}

func (s *PbcDeveloperRepoMetricResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcDeveloperRepoMetricResult) SetDeveloperRepoMetrics(v []*ReposDeveloperGroupMetric) *PbcDeveloperRepoMetricResult {
	s.DeveloperRepoMetrics = v
	return s
}

func (s *PbcDeveloperRepoMetricResult) SetRequestId(v string) *PbcDeveloperRepoMetricResult {
	s.RequestId = &v
	return s
}

func (s *PbcDeveloperRepoMetricResult) Validate() error {
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
