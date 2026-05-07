// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJobTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DeleteJobTemplateResponseBody
	GetTemplateId() *string
}

type DeleteJobTemplateResponseBody struct {
	// 本次请求的 ID，用于诊断和答疑。
	//
	// example:
	//
	// 78F6FCE2-278F-4C4A-A6B7-DD8ECEA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// tplwk80096dw****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteJobTemplateResponseBody) SetRequestId(v string) *DeleteJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobTemplateResponseBody) SetTemplateId(v string) *DeleteJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
