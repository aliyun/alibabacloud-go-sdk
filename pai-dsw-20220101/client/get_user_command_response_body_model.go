// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserCommandResponseBody
	GetCode() *string
	SetMessage(v string) *GetUserCommandResponseBody
	GetMessage() *string
	SetOnStart(v *GetUserCommandResponseBodyOnStart) *GetUserCommandResponseBody
	GetOnStart() *GetUserCommandResponseBodyOnStart
	SetRequestId(v string) *GetUserCommandResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserCommandResponseBody
	GetSuccess() *bool
	SetUserCommandId(v string) *GetUserCommandResponseBody
	GetUserCommandId() *string
	SetAccessDeniedDetail(v map[string]interface{}) *GetUserCommandResponseBody
	GetAccessDeniedDetail() map[string]interface{}
}

type GetUserCommandResponseBody struct {
	// example:
	//
	// ValidationError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OnStart *GetUserCommandResponseBodyOnStart `json:"OnStart,omitempty" xml:"OnStart,omitempty" type:"Struct"`
	// example:
	//
	// BEBDF2EE-642E-5992-8907-D2011A7ACEFE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1900
	UserCommandId *string `json:"UserCommandId,omitempty" xml:"UserCommandId,omitempty"`
	// example:
	//
	// "AccessDeniedDetail": {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "xxx",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAA....bwhg==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "xxx",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "ram:GetUserCommand"
	//
	//   }
	AccessDeniedDetail map[string]interface{} `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
}

func (s GetUserCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserCommandResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCommandResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserCommandResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserCommandResponseBody) GetOnStart() *GetUserCommandResponseBodyOnStart {
	return s.OnStart
}

func (s *GetUserCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserCommandResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserCommandResponseBody) GetUserCommandId() *string {
	return s.UserCommandId
}

func (s *GetUserCommandResponseBody) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

func (s *GetUserCommandResponseBody) SetCode(v string) *GetUserCommandResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserCommandResponseBody) SetMessage(v string) *GetUserCommandResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserCommandResponseBody) SetOnStart(v *GetUserCommandResponseBodyOnStart) *GetUserCommandResponseBody {
	s.OnStart = v
	return s
}

func (s *GetUserCommandResponseBody) SetRequestId(v string) *GetUserCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserCommandResponseBody) SetSuccess(v bool) *GetUserCommandResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserCommandResponseBody) SetUserCommandId(v string) *GetUserCommandResponseBody {
	s.UserCommandId = &v
	return s
}

func (s *GetUserCommandResponseBody) SetAccessDeniedDetail(v map[string]interface{}) *GetUserCommandResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetUserCommandResponseBody) Validate() error {
	if s.OnStart != nil {
		if err := s.OnStart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserCommandResponseBodyOnStart struct {
	// example:
	//
	// apt update
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetUserCommandResponseBodyOnStart) String() string {
	return dara.Prettify(s)
}

func (s GetUserCommandResponseBodyOnStart) GoString() string {
	return s.String()
}

func (s *GetUserCommandResponseBodyOnStart) GetContent() *string {
	return s.Content
}

func (s *GetUserCommandResponseBodyOnStart) SetContent(v string) *GetUserCommandResponseBodyOnStart {
	s.Content = &v
	return s
}

func (s *GetUserCommandResponseBodyOnStart) Validate() error {
	return dara.Validate(s)
}
