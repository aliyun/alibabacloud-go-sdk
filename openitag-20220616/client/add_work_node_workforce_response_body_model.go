// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkNodeWorkforceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddWorkNodeWorkforceResponseBody
	GetCode() *int32
	SetDetails(v string) *AddWorkNodeWorkforceResponseBody
	GetDetails() *string
	SetErrorCode(v string) *AddWorkNodeWorkforceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AddWorkNodeWorkforceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddWorkNodeWorkforceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddWorkNodeWorkforceResponseBody
	GetSuccess() *bool
}

type AddWorkNodeWorkforceResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// - When Success is false, a business error code is returned.
	//
	// - When Success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorkNodeWorkforceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorkNodeWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddWorkNodeWorkforceResponseBody) GetDetails() *string {
	return s.Details
}

func (s *AddWorkNodeWorkforceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddWorkNodeWorkforceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddWorkNodeWorkforceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorkNodeWorkforceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWorkNodeWorkforceResponseBody) SetCode(v int32) *AddWorkNodeWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetDetails(v string) *AddWorkNodeWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetErrorCode(v string) *AddWorkNodeWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetMessage(v string) *AddWorkNodeWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetRequestId(v string) *AddWorkNodeWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetSuccess(v bool) *AddWorkNodeWorkforceResponseBody {
	s.Success = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) Validate() error {
	return dara.Validate(s)
}
