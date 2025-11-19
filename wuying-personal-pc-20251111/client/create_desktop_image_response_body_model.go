// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDesktopImageResponseBody
	GetCode() *string
	SetData(v string) *CreateDesktopImageResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateDesktopImageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDesktopImageResponseBody
	GetRequestId() *string
	SetTraceId(v string) *CreateDesktopImageResponseBody
	GetTraceId() *string
}

type CreateDesktopImageResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId        *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateDesktopImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDesktopImageResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateDesktopImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDesktopImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDesktopImageResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateDesktopImageResponseBody) SetCode(v string) *CreateDesktopImageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDesktopImageResponseBody) SetData(v string) *CreateDesktopImageResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDesktopImageResponseBody) SetHttpStatusCode(v int32) *CreateDesktopImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDesktopImageResponseBody) SetRequestId(v string) *CreateDesktopImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopImageResponseBody) SetTraceId(v string) *CreateDesktopImageResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateDesktopImageResponseBody) Validate() error {
	return dara.Validate(s)
}
