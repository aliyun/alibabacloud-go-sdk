// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLdpsColumnarIndexStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckLdpsColumnarIndexStatusResponseBody
	GetAccessDeniedDetail() *string
	SetMode(v string) *CheckLdpsColumnarIndexStatusResponseBody
	GetMode() *string
	SetOpened(v bool) *CheckLdpsColumnarIndexStatusResponseBody
	GetOpened() *bool
	SetRequestId(v string) *CheckLdpsColumnarIndexStatusResponseBody
	GetRequestId() *string
}

type CheckLdpsColumnarIndexStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Columnar index mode. Valid values: LAKE and OLAP.
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Opened    *bool   `json:"Opened,omitempty" xml:"Opened,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) GetMode() *string {
	return s.Mode
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) GetOpened() *bool {
	return s.Opened
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetAccessDeniedDetail(v string) *CheckLdpsColumnarIndexStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetMode(v string) *CheckLdpsColumnarIndexStatusResponseBody {
	s.Mode = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetOpened(v bool) *CheckLdpsColumnarIndexStatusResponseBody {
	s.Opened = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetRequestId(v string) *CheckLdpsColumnarIndexStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
