// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyTemplateResponseBody
	GetRequestId() *string
}

type ModifyTemplateResponseBody struct {
	// example:
	//
	// 323*****998-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateResponseBody) SetId(v string) *ModifyTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyTemplateResponseBody) SetRequestId(v string) *ModifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
