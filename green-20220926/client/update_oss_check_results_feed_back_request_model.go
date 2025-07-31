// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFeedBackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeedback(v string) *UpdateOssCheckResultsFeedBackRequest
	GetFeedback() *string
	SetQueryRequestId(v string) *UpdateOssCheckResultsFeedBackRequest
	GetQueryRequestId() *string
	SetRegionId(v string) *UpdateOssCheckResultsFeedBackRequest
	GetRegionId() *string
	SetServiceCode(v string) *UpdateOssCheckResultsFeedBackRequest
	GetServiceCode() *string
	SetTaskId(v string) *UpdateOssCheckResultsFeedBackRequest
	GetTaskId() *string
}

type UpdateOssCheckResultsFeedBackRequest struct {
	// example:
	//
	// misreport
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	QueryRequestId *string `json:"QueryRequestId,omitempty" xml:"QueryRequestId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// P_7SCUK8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateOssCheckResultsFeedBackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFeedBackRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFeedBackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *UpdateOssCheckResultsFeedBackRequest) GetQueryRequestId() *string {
	return s.QueryRequestId
}

func (s *UpdateOssCheckResultsFeedBackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateOssCheckResultsFeedBackRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *UpdateOssCheckResultsFeedBackRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateOssCheckResultsFeedBackRequest) SetFeedback(v string) *UpdateOssCheckResultsFeedBackRequest {
	s.Feedback = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackRequest) SetQueryRequestId(v string) *UpdateOssCheckResultsFeedBackRequest {
	s.QueryRequestId = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackRequest) SetRegionId(v string) *UpdateOssCheckResultsFeedBackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackRequest) SetServiceCode(v string) *UpdateOssCheckResultsFeedBackRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackRequest) SetTaskId(v string) *UpdateOssCheckResultsFeedBackRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateOssCheckResultsFeedBackRequest) Validate() error {
	return dara.Validate(s)
}
