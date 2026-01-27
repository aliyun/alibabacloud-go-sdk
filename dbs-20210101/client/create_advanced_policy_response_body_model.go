// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAdvancedPolicyResponseBody
	GetCode() *string
	SetData(v bool) *CreateAdvancedPolicyResponseBody
	GetData() *bool
	SetErrCode(v string) *CreateAdvancedPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateAdvancedPolicyResponseBody
	GetErrMessage() *string
	SetMessage(v string) *CreateAdvancedPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAdvancedPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateAdvancedPolicyResponseBody
	GetSuccess() *string
}

type CreateAdvancedPolicyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the advanced backup policy takes effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// instanceName can not be empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1EFBAC73-4A72-5AD0-BE27-932491FCB848
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAdvancedPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdvancedPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAdvancedPolicyResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateAdvancedPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAdvancedPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateAdvancedPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAdvancedPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdvancedPolicyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateAdvancedPolicyResponseBody) SetCode(v string) *CreateAdvancedPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetData(v bool) *CreateAdvancedPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetErrCode(v string) *CreateAdvancedPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetErrMessage(v string) *CreateAdvancedPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetMessage(v string) *CreateAdvancedPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetRequestId(v string) *CreateAdvancedPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) SetSuccess(v string) *CreateAdvancedPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAdvancedPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
