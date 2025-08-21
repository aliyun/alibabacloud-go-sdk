// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginInfo(v *ManageLoginResponseBodyLoginInfo) *ManageLoginResponseBody
	GetLoginInfo() *ManageLoginResponseBodyLoginInfo
	SetRequestId(v string) *ManageLoginResponseBody
	GetRequestId() *string
}

type ManageLoginResponseBody struct {
	LoginInfo *ManageLoginResponseBodyLoginInfo `json:"LoginInfo,omitempty" xml:"LoginInfo,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageLoginResponseBody) GoString() string {
	return s.String()
}

func (s *ManageLoginResponseBody) GetLoginInfo() *ManageLoginResponseBodyLoginInfo {
	return s.LoginInfo
}

func (s *ManageLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageLoginResponseBody) SetLoginInfo(v *ManageLoginResponseBodyLoginInfo) *ManageLoginResponseBody {
	s.LoginInfo = v
	return s
}

func (s *ManageLoginResponseBody) SetRequestId(v string) *ManageLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageLoginResponseBody) Validate() error {
	return dara.Validate(s)
}

type ManageLoginResponseBodyLoginInfo struct {
	AdbLoginPort *int32 `json:"AdbLoginPort,omitempty" xml:"AdbLoginPort,omitempty"`
	// example:
	//
	// 12.10.4.10
	LoginHostname *string `json:"LoginHostname,omitempty" xml:"LoginHostname,omitempty"`
	// example:
	//
	// 10004
	LoginPort *int32 `json:"LoginPort,omitempty" xml:"LoginPort,omitempty"`
}

func (s ManageLoginResponseBodyLoginInfo) String() string {
	return dara.Prettify(s)
}

func (s ManageLoginResponseBodyLoginInfo) GoString() string {
	return s.String()
}

func (s *ManageLoginResponseBodyLoginInfo) GetAdbLoginPort() *int32 {
	return s.AdbLoginPort
}

func (s *ManageLoginResponseBodyLoginInfo) GetLoginHostname() *string {
	return s.LoginHostname
}

func (s *ManageLoginResponseBodyLoginInfo) GetLoginPort() *int32 {
	return s.LoginPort
}

func (s *ManageLoginResponseBodyLoginInfo) SetAdbLoginPort(v int32) *ManageLoginResponseBodyLoginInfo {
	s.AdbLoginPort = &v
	return s
}

func (s *ManageLoginResponseBodyLoginInfo) SetLoginHostname(v string) *ManageLoginResponseBodyLoginInfo {
	s.LoginHostname = &v
	return s
}

func (s *ManageLoginResponseBodyLoginInfo) SetLoginPort(v int32) *ManageLoginResponseBodyLoginInfo {
	s.LoginPort = &v
	return s
}

func (s *ManageLoginResponseBodyLoginInfo) Validate() error {
	return dara.Validate(s)
}
