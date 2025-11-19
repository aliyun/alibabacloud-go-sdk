// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDesktopResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteDesktopResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DeleteDesktopResponseBody
	GetTraceId() *string
}

type DeleteDesktopResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteDesktopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDesktopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDesktopResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteDesktopResponseBody) SetCode(v string) *DeleteDesktopResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDesktopResponseBody) SetRequestId(v string) *DeleteDesktopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDesktopResponseBody) SetTraceId(v string) *DeleteDesktopResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteDesktopResponseBody) Validate() error {
	return dara.Validate(s)
}
