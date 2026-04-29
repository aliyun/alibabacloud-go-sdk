// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawCronJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawCronJobsShrinkRequest
	GetApplicationId() *string
	SetIncludeDisabled(v bool) *DescribePolarClawCronJobsShrinkRequest
	GetIncludeDisabled() *bool
	SetIncludeRuns(v bool) *DescribePolarClawCronJobsShrinkRequest
	GetIncludeRuns() *bool
	SetJobIdListShrink(v string) *DescribePolarClawCronJobsShrinkRequest
	GetJobIdListShrink() *string
	SetRunLimit(v int32) *DescribePolarClawCronJobsShrinkRequest
	GetRunLimit() *int32
}

type DescribePolarClawCronJobsShrinkRequest struct {
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
	JobIdListShrink *string `json:"JobIdList,omitempty" xml:"JobIdList,omitempty"`
	// example:
	//
	// 10
	RunLimit *int32 `json:"RunLimit,omitempty" xml:"RunLimit,omitempty"`
}

func (s DescribePolarClawCronJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawCronJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawCronJobsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawCronJobsShrinkRequest) GetIncludeDisabled() *bool {
	return s.IncludeDisabled
}

func (s *DescribePolarClawCronJobsShrinkRequest) GetIncludeRuns() *bool {
	return s.IncludeRuns
}

func (s *DescribePolarClawCronJobsShrinkRequest) GetJobIdListShrink() *string {
	return s.JobIdListShrink
}

func (s *DescribePolarClawCronJobsShrinkRequest) GetRunLimit() *int32 {
	return s.RunLimit
}

func (s *DescribePolarClawCronJobsShrinkRequest) SetApplicationId(v string) *DescribePolarClawCronJobsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawCronJobsShrinkRequest) SetIncludeDisabled(v bool) *DescribePolarClawCronJobsShrinkRequest {
	s.IncludeDisabled = &v
	return s
}

func (s *DescribePolarClawCronJobsShrinkRequest) SetIncludeRuns(v bool) *DescribePolarClawCronJobsShrinkRequest {
	s.IncludeRuns = &v
	return s
}

func (s *DescribePolarClawCronJobsShrinkRequest) SetJobIdListShrink(v string) *DescribePolarClawCronJobsShrinkRequest {
	s.JobIdListShrink = &v
	return s
}

func (s *DescribePolarClawCronJobsShrinkRequest) SetRunLimit(v int32) *DescribePolarClawCronJobsShrinkRequest {
	s.RunLimit = &v
	return s
}

func (s *DescribePolarClawCronJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
