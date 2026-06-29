// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkNodeWorkforceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RemoveWorkNodeWorkforceResponseBody
	GetCode() *int32
	SetDetails(v string) *RemoveWorkNodeWorkforceResponseBody
	GetDetails() *string
	SetErrorCode(v string) *RemoveWorkNodeWorkforceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RemoveWorkNodeWorkforceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveWorkNodeWorkforceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveWorkNodeWorkforceResponseBody
	GetSuccess() *bool
}

type RemoveWorkNodeWorkforceResponseBody struct {
	// Return code, default is 0 indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// Returned a business error code when Success is false,
	//
	// Returned as empty when Success is true.
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
	// Whether it was successful, possible values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveWorkNodeWorkforceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkNodeWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetDetails() *string {
	return s.Details
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveWorkNodeWorkforceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetCode(v int32) *RemoveWorkNodeWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetDetails(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetErrorCode(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetMessage(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetRequestId(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetSuccess(v bool) *RemoveWorkNodeWorkforceResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) Validate() error {
	return dara.Validate(s)
}
