// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmarttagTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSmarttagTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *AddSmarttagTemplateResponseBody
	GetTemplateId() *string
}

type AddSmarttagTemplateResponseBody struct {
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175EDAS3Q
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 05de22f255284c7a8d2aab535dde****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddSmarttagTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSmarttagTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmarttagTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSmarttagTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddSmarttagTemplateResponseBody) SetRequestId(v string) *AddSmarttagTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmarttagTemplateResponseBody) SetTemplateId(v string) *AddSmarttagTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddSmarttagTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
