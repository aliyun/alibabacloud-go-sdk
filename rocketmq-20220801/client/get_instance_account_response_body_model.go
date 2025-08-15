// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceAccountResponseBody
	GetCode() *string
	SetData(v *GetInstanceAccountResponseBodyData) *GetInstanceAccountResponseBody
	GetData() *GetInstanceAccountResponseBodyData
	SetDynamicCode(v string) *GetInstanceAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetInstanceAccountResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceAccountResponseBody
	GetSuccess() *bool
}

type GetInstanceAccountResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetInstanceAccountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request\\"s ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// B5C59E80-FCFC-5796-ABE4-D39EAAE578E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceAccountResponseBody) GetData() *GetInstanceAccountResponseBodyData {
	return s.Data
}

func (s *GetInstanceAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetInstanceAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetInstanceAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceAccountResponseBody) SetCode(v string) *GetInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetData(v *GetInstanceAccountResponseBodyData) *GetInstanceAccountResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceAccountResponseBody) SetDynamicCode(v string) *GetInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetDynamicMessage(v string) *GetInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetHttpStatusCode(v int32) *GetInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetMessage(v string) *GetInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetRequestId(v string) *GetInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetSuccess(v bool) *GetInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceAccountResponseBodyData struct {
	// The status of the account.
	//
	// Valid values:
	//
	// 	- DISABLE
	//
	// 	- ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The password of the account.
	//
	// example:
	//
	// *************
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponseBodyData) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *GetInstanceAccountResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetInstanceAccountResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *GetInstanceAccountResponseBodyData) SetAccountStatus(v string) *GetInstanceAccountResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *GetInstanceAccountResponseBodyData) SetPassword(v string) *GetInstanceAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetInstanceAccountResponseBodyData) SetUsername(v string) *GetInstanceAccountResponseBodyData {
	s.Username = &v
	return s
}

func (s *GetInstanceAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
