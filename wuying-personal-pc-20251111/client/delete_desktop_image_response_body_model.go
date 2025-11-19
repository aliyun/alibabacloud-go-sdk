// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDesktopImageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDesktopImageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteDesktopImageResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DeleteDesktopImageResponseBody
	GetTraceId() *string
}

type DeleteDesktopImageResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId        *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteDesktopImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDesktopImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDesktopImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDesktopImageResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteDesktopImageResponseBody) SetCode(v string) *DeleteDesktopImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDesktopImageResponseBody) SetHttpStatusCode(v int32) *DeleteDesktopImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDesktopImageResponseBody) SetRequestId(v string) *DeleteDesktopImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDesktopImageResponseBody) SetTraceId(v string) *DeleteDesktopImageResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteDesktopImageResponseBody) Validate() error {
	return dara.Validate(s)
}
