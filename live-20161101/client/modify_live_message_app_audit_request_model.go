// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppAuditRequest
	GetAppId() *string
	SetAuditType(v int32) *ModifyLiveMessageAppAuditRequest
	GetAuditType() *int32
	SetAuditUrl(v string) *ModifyLiveMessageAppAuditRequest
	GetAuditUrl() *string
	SetDataCenter(v string) *ModifyLiveMessageAppAuditRequest
	GetDataCenter() *string
}

type ModifyLiveMessageAppAuditRequest struct {
	// The ID of the interactive messaging application to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The security audit method. Valid values:
	//
	// - 0: No security audit.
	//
	// - 1: Built-in security audit.
	//
	// - 2: Custom security audit.
	//
	// example:
	//
	// 2
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The security audit URL required when custom security audit is used (AuditType=2). The URL must start with http:// or https://, must not contain private IP addresses, and must not include port numbers.
	//
	// example:
	//
	// http://example.aliyundoc.com/exampleaudit
	AuditUrl *string `json:"AuditUrl,omitempty" xml:"AuditUrl,omitempty"`
	// The data center. This value must be the same as the data center specified in [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html). Valid values: cn-shanghai (Shanghai) and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
}

func (s ModifyLiveMessageAppAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppAuditRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppAuditRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppAuditRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *ModifyLiveMessageAppAuditRequest) GetAuditUrl() *string {
	return s.AuditUrl
}

func (s *ModifyLiveMessageAppAuditRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageAppAuditRequest) SetAppId(v string) *ModifyLiveMessageAppAuditRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppAuditRequest) SetAuditType(v int32) *ModifyLiveMessageAppAuditRequest {
	s.AuditType = &v
	return s
}

func (s *ModifyLiveMessageAppAuditRequest) SetAuditUrl(v string) *ModifyLiveMessageAppAuditRequest {
	s.AuditUrl = &v
	return s
}

func (s *ModifyLiveMessageAppAuditRequest) SetDataCenter(v string) *ModifyLiveMessageAppAuditRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageAppAuditRequest) Validate() error {
	return dara.Validate(s)
}
