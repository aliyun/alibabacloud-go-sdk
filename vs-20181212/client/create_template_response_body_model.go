// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
}

type CreateTemplateResponseBody struct {
	// example:
	//
	// 399*****735-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) SetId(v string) *CreateTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
