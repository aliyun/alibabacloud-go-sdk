// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAuditRequestResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAuditRequestResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UpdateAuditRequestResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UpdateAuditRequestResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAuditRequestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAuditRequestResponseBody
	GetSuccess() *bool
}

type UpdateAuditRequestResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 14111561****
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAuditRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAuditRequestResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAuditRequestResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpdateAuditRequestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAuditRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuditRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAuditRequestResponseBody) SetAccessDeniedDetail(v string) *UpdateAuditRequestResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAuditRequestResponseBody) SetCode(v string) *UpdateAuditRequestResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAuditRequestResponseBody) SetData(v map[string]interface{}) *UpdateAuditRequestResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAuditRequestResponseBody) SetMessage(v string) *UpdateAuditRequestResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAuditRequestResponseBody) SetRequestId(v string) *UpdateAuditRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuditRequestResponseBody) SetSuccess(v bool) *UpdateAuditRequestResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAuditRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
