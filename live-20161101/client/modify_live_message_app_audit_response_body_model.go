// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppAuditResponseBody
	GetAppId() *string
	SetAppSign(v string) *ModifyLiveMessageAppAuditResponseBody
	GetAppSign() *string
	SetAuditNeedAuthentication(v bool) *ModifyLiveMessageAppAuditResponseBody
	GetAuditNeedAuthentication() *bool
	SetAuditType(v int32) *ModifyLiveMessageAppAuditResponseBody
	GetAuditType() *int32
	SetAuditUrl(v string) *ModifyLiveMessageAppAuditResponseBody
	GetAuditUrl() *string
	SetRequestId(v string) *ModifyLiveMessageAppAuditResponseBody
	GetRequestId() *string
}

type ModifyLiveMessageAppAuditResponseBody struct {
	// The ID of the interactive messaging application to modify.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The signature of the interactive messaging application. This information is required by the interactive messaging service SDK.
	//
	// example:
	//
	// **************************************************************************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// Indicates whether call authentication is enabled. If custom security audit is used, this parameter is set to true by default to enable call authentication.
	//
	// example:
	//
	// true
	AuditNeedAuthentication *bool `json:"AuditNeedAuthentication,omitempty" xml:"AuditNeedAuthentication,omitempty"`
	// The security audit method.
	//
	// example:
	//
	// 2
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The security audit URL provided when AuditType is set to 2.
	//
	// example:
	//
	// http: //example.aliyundoc.com/exampleaudit
	AuditUrl *string `json:"AuditUrl,omitempty" xml:"AuditUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C4E8440-3838-1831-9BDE-AFC15803****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveMessageAppAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppAuditResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetAuditNeedAuthentication() *bool {
	return s.AuditNeedAuthentication
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetAuditType() *int32 {
	return s.AuditType
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetAuditUrl() *string {
	return s.AuditUrl
}

func (s *ModifyLiveMessageAppAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetAppId(v string) *ModifyLiveMessageAppAuditResponseBody {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetAppSign(v string) *ModifyLiveMessageAppAuditResponseBody {
	s.AppSign = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetAuditNeedAuthentication(v bool) *ModifyLiveMessageAppAuditResponseBody {
	s.AuditNeedAuthentication = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetAuditType(v int32) *ModifyLiveMessageAppAuditResponseBody {
	s.AuditType = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetAuditUrl(v string) *ModifyLiveMessageAppAuditResponseBody {
	s.AuditUrl = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) SetRequestId(v string) *ModifyLiveMessageAppAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
