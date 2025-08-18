// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDebugAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetDebugAppInstanceResponseBody
	GetAppId() *string
	SetAppInstanceGroupId(v string) *GetDebugAppInstanceResponseBody
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *GetDebugAppInstanceResponseBody
	GetAppInstanceId() *string
	SetAppVersion(v string) *GetDebugAppInstanceResponseBody
	GetAppVersion() *string
	SetAuthCode(v string) *GetDebugAppInstanceResponseBody
	GetAuthCode() *string
	SetRequestId(v string) *GetDebugAppInstanceResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetDebugAppInstanceResponseBody
	GetUserId() *string
}

type GetDebugAppInstanceResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// __DEBUG_APP
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-7ybdeiyoeh5e****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the application version.
	//
	// example:
	//
	// 1.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The authorization code. This authorization code is valid for 3 minutes and can be used only once, regardless of whether the authentication succeeds. If multiple authentication codes are generated for a user, only the latest authentication code takes effect.
	//
	// example:
	//
	// e4e169bea1cc48e8afac53**********
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// __debug__
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDebugAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDebugAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetDebugAppInstanceResponseBody) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetDebugAppInstanceResponseBody) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetDebugAppInstanceResponseBody) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetDebugAppInstanceResponseBody) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetDebugAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDebugAppInstanceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetDebugAppInstanceResponseBody) SetAppId(v string) *GetDebugAppInstanceResponseBody {
	s.AppId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceGroupId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppVersion(v string) *GetDebugAppInstanceResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAuthCode(v string) *GetDebugAppInstanceResponseBody {
	s.AuthCode = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetRequestId(v string) *GetDebugAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetUserId(v string) *GetDebugAppInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
