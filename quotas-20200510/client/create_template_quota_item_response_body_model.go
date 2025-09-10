// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateQuotaItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateTemplateQuotaItemResponseBody
	GetId() *string
	SetRequestId(v string) *CreateTemplateQuotaItemResponseBody
	GetRequestId() *string
}

type CreateTemplateQuotaItemResponseBody struct {
	// The ID of the quota template.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D47B3A10-CDAC-5412-B2EE-EC9A3DBE9053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateQuotaItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateTemplateQuotaItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateQuotaItemResponseBody) SetId(v string) *CreateTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTemplateQuotaItemResponseBody) SetRequestId(v string) *CreateTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateQuotaItemResponseBody) Validate() error {
	return dara.Validate(s)
}
