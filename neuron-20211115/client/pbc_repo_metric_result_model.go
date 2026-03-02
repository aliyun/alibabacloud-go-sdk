// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcRepoMetricResult interface {
	dara.Model
	String() string
	GoString() string
	SetRepoMetrics(v []*RepoMetric) *PbcRepoMetricResult
	GetRepoMetrics() []*RepoMetric
	SetRequestId(v string) *PbcRepoMetricResult
	GetRequestId() *string
}

type PbcRepoMetricResult struct {
	RepoMetrics []*RepoMetric `json:"repoMetrics,omitempty" xml:"repoMetrics,omitempty" type:"Repeated"`
	RequestId   *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcRepoMetricResult) String() string {
	return dara.Prettify(s)
}

func (s PbcRepoMetricResult) GoString() string {
	return s.String()
}

func (s *PbcRepoMetricResult) GetRepoMetrics() []*RepoMetric {
	return s.RepoMetrics
}

func (s *PbcRepoMetricResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcRepoMetricResult) SetRepoMetrics(v []*RepoMetric) *PbcRepoMetricResult {
	s.RepoMetrics = v
	return s
}

func (s *PbcRepoMetricResult) SetRequestId(v string) *PbcRepoMetricResult {
	s.RequestId = &v
	return s
}

func (s *PbcRepoMetricResult) Validate() error {
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
