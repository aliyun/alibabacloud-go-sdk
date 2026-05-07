// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersion(v int32) *UpdateJobTemplateResponseBody
	GetDefaultVersion() *int32
	SetGmtModifyTime(v string) *UpdateJobTemplateResponseBody
	GetGmtModifyTime() *string
	SetRequestId(v string) *UpdateJobTemplateResponseBody
	GetRequestId() *string
	SetVersion(v int32) *UpdateJobTemplateResponseBody
	GetVersion() *int32
	SetVersionCreated(v bool) *UpdateJobTemplateResponseBody
	GetVersionCreated() *bool
}

type UpdateJobTemplateResponseBody struct {
	// example:
	//
	// 2
	DefaultVersion *int32 `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 本次请求的 ID，用于诊断和答疑。
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// true
	VersionCreated *bool `json:"VersionCreated,omitempty" xml:"VersionCreated,omitempty"`
}

func (s UpdateJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateResponseBody) GetDefaultVersion() *int32 {
	return s.DefaultVersion
}

func (s *UpdateJobTemplateResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *UpdateJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobTemplateResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateJobTemplateResponseBody) GetVersionCreated() *bool {
	return s.VersionCreated
}

func (s *UpdateJobTemplateResponseBody) SetDefaultVersion(v int32) *UpdateJobTemplateResponseBody {
	s.DefaultVersion = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) SetGmtModifyTime(v string) *UpdateJobTemplateResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) SetRequestId(v string) *UpdateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) SetVersion(v int32) *UpdateJobTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) SetVersionCreated(v bool) *UpdateJobTemplateResponseBody {
	s.VersionCreated = &v
	return s
}

func (s *UpdateJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
