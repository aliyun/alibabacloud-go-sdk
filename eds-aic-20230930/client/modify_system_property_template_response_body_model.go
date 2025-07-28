// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemPropertyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySystemPropertyTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *ModifySystemPropertyTemplateResponseBody
	GetTemplateId() *string
}

type ModifySystemPropertyTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ap-angyvganxlf****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifySystemPropertyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySystemPropertyTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifySystemPropertyTemplateResponseBody) SetRequestId(v string) *ModifySystemPropertyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySystemPropertyTemplateResponseBody) SetTemplateId(v string) *ModifySystemPropertyTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ModifySystemPropertyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
