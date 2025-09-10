// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateQuotaItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyTemplateQuotaItemResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyTemplateQuotaItemResponseBody
	GetRequestId() *string
}

type ModifyTemplateQuotaItemResponseBody struct {
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateQuotaItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyTemplateQuotaItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTemplateQuotaItemResponseBody) SetId(v string) *ModifyTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyTemplateQuotaItemResponseBody) SetRequestId(v string) *ModifyTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTemplateQuotaItemResponseBody) Validate() error {
	return dara.Validate(s)
}
