// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskWorkforceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskWorkforceResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskWorkforceResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskWorkforceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskWorkforceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskWorkforceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskWorkforceResponseBody
	GetSuccess() *bool
	SetWorkforce(v []*Workforce) *GetTaskWorkforceResponseBody
	GetWorkforce() []*Workforce
}

type GetTaskWorkforceResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
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
	// - When Success is false, returns a business error code.
	//
	// - When Success is true, returns an empty value.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message of the request.
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
	// List of workforce.
	Workforce []*Workforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s GetTaskWorkforceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskWorkforceResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskWorkforceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskWorkforceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskWorkforceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskWorkforceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskWorkforceResponseBody) GetWorkforce() []*Workforce {
	return s.Workforce
}

func (s *GetTaskWorkforceResponseBody) SetCode(v int32) *GetTaskWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetDetails(v string) *GetTaskWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetErrorCode(v string) *GetTaskWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetMessage(v string) *GetTaskWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetRequestId(v string) *GetTaskWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetSuccess(v bool) *GetTaskWorkforceResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetWorkforce(v []*Workforce) *GetTaskWorkforceResponseBody {
	s.Workforce = v
	return s
}

func (s *GetTaskWorkforceResponseBody) Validate() error {
	if s.Workforce != nil {
		for _, item := range s.Workforce {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
