// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefenseTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v int64) *CreateDefenseTemplateResponseBody
	GetTemplateId() *int64
}

type CreateDefenseTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F48ABDF7-D777-5F26-892A-57349765D7A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the protection rule template.
	//
	// example:
	//
	// 2212
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefenseTemplateResponseBody) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateDefenseTemplateResponseBody) SetRequestId(v string) *CreateDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseTemplateResponseBody) SetTemplateId(v int64) *CreateDefenseTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateDefenseTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
