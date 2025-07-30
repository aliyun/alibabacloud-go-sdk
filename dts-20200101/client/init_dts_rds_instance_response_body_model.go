// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitDtsRdsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminAccount(v string) *InitDtsRdsInstanceResponseBody
	GetAdminAccount() *string
	SetAdminPassword(v string) *InitDtsRdsInstanceResponseBody
	GetAdminPassword() *string
	SetErrCode(v string) *InitDtsRdsInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *InitDtsRdsInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *InitDtsRdsInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *InitDtsRdsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *InitDtsRdsInstanceResponseBody
	GetSuccess() *string
}

type InitDtsRdsInstanceResponseBody struct {
	// The built-in account that is used by DTS to connect to the node.
	//
	// example:
	//
	// rdsdt_dtsacct
	AdminAccount *string `json:"AdminAccount,omitempty" xml:"AdminAccount,omitempty"`
	// The password of the built-in account.
	//
	// example:
	//
	// 1jecpqrtc****
	AdminPassword *string `json:"AdminPassword,omitempty" xml:"AdminPassword,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7796ECE0-9C17-5E4D-9CE6-B7EC825A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitDtsRdsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitDtsRdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceResponseBody) GetAdminAccount() *string {
	return s.AdminAccount
}

func (s *InitDtsRdsInstanceResponseBody) GetAdminPassword() *string {
	return s.AdminPassword
}

func (s *InitDtsRdsInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *InitDtsRdsInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *InitDtsRdsInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *InitDtsRdsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitDtsRdsInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *InitDtsRdsInstanceResponseBody) SetAdminAccount(v string) *InitDtsRdsInstanceResponseBody {
	s.AdminAccount = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetAdminPassword(v string) *InitDtsRdsInstanceResponseBody {
	s.AdminPassword = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetErrCode(v string) *InitDtsRdsInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetErrMessage(v string) *InitDtsRdsInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetHttpStatusCode(v string) *InitDtsRdsInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetRequestId(v string) *InitDtsRdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetSuccess(v string) *InitDtsRdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
