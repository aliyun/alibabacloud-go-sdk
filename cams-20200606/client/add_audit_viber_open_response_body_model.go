// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditViberOpenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddAuditViberOpenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddAuditViberOpenResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *AddAuditViberOpenResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *AddAuditViberOpenResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAuditViberOpenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuditViberOpenResponseBody
	GetSuccess() *bool
}

type AddAuditViberOpenResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 122
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2121-112
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAuditViberOpenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddAuditViberOpenResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAuditViberOpenResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddAuditViberOpenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAuditViberOpenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuditViberOpenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuditViberOpenResponseBody) SetAccessDeniedDetail(v string) *AddAuditViberOpenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddAuditViberOpenResponseBody) SetCode(v string) *AddAuditViberOpenResponseBody {
	s.Code = &v
	return s
}

func (s *AddAuditViberOpenResponseBody) SetData(v map[string]interface{}) *AddAuditViberOpenResponseBody {
	s.Data = v
	return s
}

func (s *AddAuditViberOpenResponseBody) SetMessage(v string) *AddAuditViberOpenResponseBody {
	s.Message = &v
	return s
}

func (s *AddAuditViberOpenResponseBody) SetRequestId(v string) *AddAuditViberOpenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuditViberOpenResponseBody) SetSuccess(v bool) *AddAuditViberOpenResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuditViberOpenResponseBody) Validate() error {
	return dara.Validate(s)
}
