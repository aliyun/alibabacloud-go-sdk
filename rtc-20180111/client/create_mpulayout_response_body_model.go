// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMPULayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v int64) *CreateMPULayoutResponseBody
	GetLayoutId() *int64
	SetRequestId(v string) *CreateMPULayoutResponseBody
	GetRequestId() *string
}

type CreateMPULayoutResponseBody struct {
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMPULayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutResponseBody) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *CreateMPULayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMPULayoutResponseBody) SetLayoutId(v int64) *CreateMPULayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *CreateMPULayoutResponseBody) SetRequestId(v string) *CreateMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMPULayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
