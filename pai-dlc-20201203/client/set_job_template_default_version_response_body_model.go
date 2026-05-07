// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobTemplateDefaultVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersion(v int32) *SetJobTemplateDefaultVersionResponseBody
	GetDefaultVersion() *int32
	SetGmtModifyTime(v string) *SetJobTemplateDefaultVersionResponseBody
	GetGmtModifyTime() *string
	SetRequestId(v string) *SetJobTemplateDefaultVersionResponseBody
	GetRequestId() *string
}

type SetJobTemplateDefaultVersionResponseBody struct {
	// 设置后的默认版本号
	//
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
}

func (s SetJobTemplateDefaultVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetJobTemplateDefaultVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetJobTemplateDefaultVersionResponseBody) GetDefaultVersion() *int32 {
	return s.DefaultVersion
}

func (s *SetJobTemplateDefaultVersionResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *SetJobTemplateDefaultVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetJobTemplateDefaultVersionResponseBody) SetDefaultVersion(v int32) *SetJobTemplateDefaultVersionResponseBody {
	s.DefaultVersion = &v
	return s
}

func (s *SetJobTemplateDefaultVersionResponseBody) SetGmtModifyTime(v string) *SetJobTemplateDefaultVersionResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *SetJobTemplateDefaultVersionResponseBody) SetRequestId(v string) *SetJobTemplateDefaultVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetJobTemplateDefaultVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
