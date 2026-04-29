// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawCronJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawCronJobsRequest
	GetApplicationId() *string
	SetIncludeDisabled(v bool) *DescribePolarClawCronJobsRequest
	GetIncludeDisabled() *bool
	SetIncludeRuns(v bool) *DescribePolarClawCronJobsRequest
	GetIncludeRuns() *bool
	SetJobIdList(v []*string) *DescribePolarClawCronJobsRequest
	GetJobIdList() []*string
	SetRunLimit(v int32) *DescribePolarClawCronJobsRequest
	GetRunLimit() *int32
}

type DescribePolarClawCronJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// true
	IncludeDisabled *bool `json:"IncludeDisabled,omitempty" xml:"IncludeDisabled,omitempty"`
	// example:
	//
	// true
	IncludeRuns *bool `json:"IncludeRuns,omitempty" xml:"IncludeRuns,omitempty"`
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e,1ee00f56-f467-4d41-858c-ca4ede2c770f
	JobIdList []*string `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	RunLimit *int32 `json:"RunLimit,omitempty" xml:"RunLimit,omitempty"`
}

func (s DescribePolarClawCronJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawCronJobsRequest) GetIncludeDisabled() *bool {
	return s.IncludeDisabled
}

func (s *DescribePolarClawCronJobsRequest) GetIncludeRuns() *bool {
	return s.IncludeRuns
}

func (s *DescribePolarClawCronJobsRequest) GetJobIdList() []*string {
	return s.JobIdList
}

func (s *DescribePolarClawCronJobsRequest) GetRunLimit() *int32 {
	return s.RunLimit
}

func (s *DescribePolarClawCronJobsRequest) SetApplicationId(v string) *DescribePolarClawCronJobsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawCronJobsRequest) SetIncludeDisabled(v bool) *DescribePolarClawCronJobsRequest {
	s.IncludeDisabled = &v
	return s
}

func (s *DescribePolarClawCronJobsRequest) SetIncludeRuns(v bool) *DescribePolarClawCronJobsRequest {
	s.IncludeRuns = &v
	return s
}

func (s *DescribePolarClawCronJobsRequest) SetJobIdList(v []*string) *DescribePolarClawCronJobsRequest {
	s.JobIdList = v
	return s
}

func (s *DescribePolarClawCronJobsRequest) SetRunLimit(v int32) *DescribePolarClawCronJobsRequest {
	s.RunLimit = &v
	return s
}

func (s *DescribePolarClawCronJobsRequest) Validate() error {
	return dara.Validate(s)
}
