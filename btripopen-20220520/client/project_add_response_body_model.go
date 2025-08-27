// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProjectAddResponseBody
	GetCode() *string
	SetMessage(v string) *ProjectAddResponseBody
	GetMessage() *string
	SetModule(v int64) *ProjectAddResponseBody
	GetModule() *int64
	SetMorePage(v bool) *ProjectAddResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *ProjectAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ProjectAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ProjectAddResponseBody
	GetTraceId() *string
}

type ProjectAddResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 530404
	Module *int64 `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ProjectAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProjectAddResponseBody) GoString() string {
	return s.String()
}

func (s *ProjectAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProjectAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProjectAddResponseBody) GetModule() *int64 {
	return s.Module
}

func (s *ProjectAddResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *ProjectAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProjectAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ProjectAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ProjectAddResponseBody) SetCode(v string) *ProjectAddResponseBody {
	s.Code = &v
	return s
}

func (s *ProjectAddResponseBody) SetMessage(v string) *ProjectAddResponseBody {
	s.Message = &v
	return s
}

func (s *ProjectAddResponseBody) SetModule(v int64) *ProjectAddResponseBody {
	s.Module = &v
	return s
}

func (s *ProjectAddResponseBody) SetMorePage(v bool) *ProjectAddResponseBody {
	s.MorePage = &v
	return s
}

func (s *ProjectAddResponseBody) SetRequestId(v string) *ProjectAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProjectAddResponseBody) SetSuccess(v bool) *ProjectAddResponseBody {
	s.Success = &v
	return s
}

func (s *ProjectAddResponseBody) SetTraceId(v string) *ProjectAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *ProjectAddResponseBody) Validate() error {
	return dara.Validate(s)
}
