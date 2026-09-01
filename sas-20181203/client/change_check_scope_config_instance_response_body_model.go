// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckScopeConfigInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeCheckScopeConfigInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *ChangeCheckScopeConfigInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeCheckScopeConfigInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeCheckScopeConfigInstanceResponseBody
	GetSuccess() *bool
}

type ChangeCheckScopeConfigInstanceResponseBody struct {
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// There was an error with your request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeCheckScopeConfigInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckScopeConfigInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) SetCode(v string) *ChangeCheckScopeConfigInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) SetMessage(v string) *ChangeCheckScopeConfigInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) SetRequestId(v string) *ChangeCheckScopeConfigInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) SetSuccess(v bool) *ChangeCheckScopeConfigInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
