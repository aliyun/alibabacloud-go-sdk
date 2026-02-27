// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LoginInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *LoginInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *LoginInstanceResponseBody
	GetRequestId() *string
	SetRoot(v *LoginInstanceResponseBodyRoot) *LoginInstanceResponseBody
	GetRoot() *LoginInstanceResponseBodyRoot
	SetSuccess(v string) *LoginInstanceResponseBody
	GetSuccess() *string
}

type LoginInstanceResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// abc-123
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *LoginInstanceResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true/false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *LoginInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LoginInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginInstanceResponseBody) GetRoot() *LoginInstanceResponseBodyRoot {
	return s.Root
}

func (s *LoginInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *LoginInstanceResponseBody) SetCode(v string) *LoginInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *LoginInstanceResponseBody) SetMessage(v string) *LoginInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRequestId(v string) *LoginInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRoot(v *LoginInstanceResponseBodyRoot) *LoginInstanceResponseBody {
	s.Root = v
	return s
}

func (s *LoginInstanceResponseBody) SetSuccess(v string) *LoginInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *LoginInstanceResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceResponseBodyRoot struct {
	DisposableAccount     *LoginInstanceResponseBodyRootDisposableAccount       `json:"DisposableAccount,omitempty" xml:"DisposableAccount,omitempty" type:"Struct"`
	InstanceLoginInfoList []*LoginInstanceResponseBodyRootInstanceLoginInfoList `json:"InstanceLoginInfoList,omitempty" xml:"InstanceLoginInfoList,omitempty" type:"Repeated"`
	SessionControl        *LoginInstanceResponseBodyRootSessionControl          `json:"SessionControl,omitempty" xml:"SessionControl,omitempty" type:"Struct"`
}

func (s LoginInstanceResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRoot) GetDisposableAccount() *LoginInstanceResponseBodyRootDisposableAccount {
	return s.DisposableAccount
}

func (s *LoginInstanceResponseBodyRoot) GetInstanceLoginInfoList() []*LoginInstanceResponseBodyRootInstanceLoginInfoList {
	return s.InstanceLoginInfoList
}

func (s *LoginInstanceResponseBodyRoot) GetSessionControl() *LoginInstanceResponseBodyRootSessionControl {
	return s.SessionControl
}

func (s *LoginInstanceResponseBodyRoot) SetDisposableAccount(v *LoginInstanceResponseBodyRootDisposableAccount) *LoginInstanceResponseBodyRoot {
	s.DisposableAccount = v
	return s
}

func (s *LoginInstanceResponseBodyRoot) SetInstanceLoginInfoList(v []*LoginInstanceResponseBodyRootInstanceLoginInfoList) *LoginInstanceResponseBodyRoot {
	s.InstanceLoginInfoList = v
	return s
}

func (s *LoginInstanceResponseBodyRoot) SetSessionControl(v *LoginInstanceResponseBodyRootSessionControl) *LoginInstanceResponseBodyRoot {
	s.SessionControl = v
	return s
}

func (s *LoginInstanceResponseBodyRoot) Validate() error {
	if s.DisposableAccount != nil {
		if err := s.DisposableAccount.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceLoginInfoList != nil {
		for _, item := range s.InstanceLoginInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SessionControl != nil {
		if err := s.SessionControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceResponseBodyRootDisposableAccount struct {
	// example:
	//
	// abc
	LoginFormActionUrl *string `json:"LoginFormActionUrl,omitempty" xml:"LoginFormActionUrl,omitempty"`
	// example:
	//
	// abc
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootDisposableAccount) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBodyRootDisposableAccount) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) GetLoginFormActionUrl() *string {
	return s.LoginFormActionUrl
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) SetLoginFormActionUrl(v string) *LoginInstanceResponseBodyRootDisposableAccount {
	s.LoginFormActionUrl = &v
	return s
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) SetLoginUrl(v string) *LoginInstanceResponseBodyRootDisposableAccount {
	s.LoginUrl = &v
	return s
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) Validate() error {
	return dara.Validate(s)
}

type LoginInstanceResponseBodyRootInstanceLoginInfoList struct {
	// example:
	//
	// i-abc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 134
	InstanceLoginToken *string                                                              `json:"InstanceLoginToken,omitempty" xml:"InstanceLoginToken,omitempty"`
	InstanceLoginView  *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView `json:"InstanceLoginView,omitempty" xml:"InstanceLoginView,omitempty" type:"Struct"`
	// example:
	//
	// true
	LoginSuccess *bool `json:"LoginSuccess,omitempty" xml:"LoginSuccess,omitempty"`
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoList) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoList) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) GetInstanceLoginToken() *string {
	return s.InstanceLoginToken
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) GetInstanceLoginView() *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView {
	return s.InstanceLoginView
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) GetLoginSuccess() *bool {
	return s.LoginSuccess
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceId(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceLoginToken(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceLoginToken = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceLoginView(v *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceLoginView = v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetLoginSuccess(v bool) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.LoginSuccess = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) Validate() error {
	if s.InstanceLoginView != nil {
		if err := s.InstanceLoginView.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView struct {
	// example:
	//
	// abc
	DefaultViewUrl *string `json:"DefaultViewUrl,omitempty" xml:"DefaultViewUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) GetDefaultViewUrl() *string {
	return s.DefaultViewUrl
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) SetDefaultViewUrl(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView {
	s.DefaultViewUrl = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) Validate() error {
	return dara.Validate(s)
}

type LoginInstanceResponseBodyRootSessionControl struct {
	// example:
	//
	// abc
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootSessionControl) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBodyRootSessionControl) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootSessionControl) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *LoginInstanceResponseBodyRootSessionControl) SetBaseUrl(v string) *LoginInstanceResponseBodyRootSessionControl {
	s.BaseUrl = &v
	return s
}

func (s *LoginInstanceResponseBodyRootSessionControl) Validate() error {
	return dara.Validate(s)
}
