// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCacheAnalysisJobRequest
	GetInstanceId() *string
	SetJobId(v string) *DescribeCacheAnalysisJobRequest
	GetJobId() *string
}

type DescribeCacheAnalysisJobRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the cache analysis task. You can obtain the task ID from the response parameters of the [CreateCacheAnalysisJob](https://help.aliyun.com/document_detail/180982.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// sf79-sd99-sa37-****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeCacheAnalysisJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCacheAnalysisJobRequest) SetInstanceId(v string) *DescribeCacheAnalysisJobRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobRequest) SetJobId(v string) *DescribeCacheAnalysisJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeCacheAnalysisJobRequest) Validate() error {
	return dara.Validate(s)
}
