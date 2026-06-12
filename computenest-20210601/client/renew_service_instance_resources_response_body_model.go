// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailureDetails(v []*RenewServiceInstanceResourcesResponseBodyFailureDetails) *RenewServiceInstanceResourcesResponseBody
	GetFailureDetails() []*RenewServiceInstanceResourcesResponseBodyFailureDetails
	SetRenewalResult(v *RenewServiceInstanceResourcesResponseBodyRenewalResult) *RenewServiceInstanceResourcesResponseBody
	GetRenewalResult() *RenewServiceInstanceResourcesResponseBodyRenewalResult
	SetRequestId(v string) *RenewServiceInstanceResourcesResponseBody
	GetRequestId() *string
}

type RenewServiceInstanceResourcesResponseBody struct {
	// The details of renewal failures.
	FailureDetails []*RenewServiceInstanceResourcesResponseBodyFailureDetails `json:"FailureDetails,omitempty" xml:"FailureDetails,omitempty" type:"Repeated"`
	// The renewal result.
	RenewalResult *RenewServiceInstanceResourcesResponseBodyRenewalResult `json:"RenewalResult,omitempty" xml:"RenewalResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBody) GetFailureDetails() []*RenewServiceInstanceResourcesResponseBodyFailureDetails {
	return s.FailureDetails
}

func (s *RenewServiceInstanceResourcesResponseBody) GetRenewalResult() *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	return s.RenewalResult
}

func (s *RenewServiceInstanceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewServiceInstanceResourcesResponseBody) SetFailureDetails(v []*RenewServiceInstanceResourcesResponseBodyFailureDetails) *RenewServiceInstanceResourcesResponseBody {
	s.FailureDetails = v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBody) SetRenewalResult(v *RenewServiceInstanceResourcesResponseBodyRenewalResult) *RenewServiceInstanceResourcesResponseBody {
	s.RenewalResult = v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBody) SetRequestId(v string) *RenewServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBody) Validate() error {
	if s.FailureDetails != nil {
		for _, item := range s.FailureDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RenewalResult != nil {
		if err := s.RenewalResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewServiceInstanceResourcesResponseBodyFailureDetails struct {
	// The error code.
	//
	// example:
	//
	// InvalidPeriod
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Renewal failure reason.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBodyFailureDetails) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBodyFailureDetails) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetErrorCode(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ErrorCode = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetErrorMessage(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ErrorMessage = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetResourceArn(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ResourceArn = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) Validate() error {
	return dara.Validate(s)
}

type RenewServiceInstanceResourcesResponseBodyRenewalResult struct {
	// The number of resources that failed to be renewed.
	//
	// example:
	//
	// 1
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of successfully renewed resources.
	//
	// example:
	//
	// 9
	Succeeded *int32 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	// The number of resources to be renewed.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBodyRenewalResult) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBodyRenewalResult) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) GetFailed() *int32 {
	return s.Failed
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) GetSucceeded() *int32 {
	return s.Succeeded
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetFailed(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.Failed = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetSucceeded(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.Succeeded = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetTotalCount(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.TotalCount = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) Validate() error {
	return dara.Validate(s)
}
