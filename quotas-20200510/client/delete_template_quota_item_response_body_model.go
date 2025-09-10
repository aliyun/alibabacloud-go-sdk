// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateQuotaItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteTemplateQuotaItemResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteTemplateQuotaItemResponseBody
	GetRequestId() *string
}

type DeleteTemplateQuotaItemResponseBody struct {
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

func (s DeleteTemplateQuotaItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateQuotaItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteTemplateQuotaItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateQuotaItemResponseBody) SetId(v string) *DeleteTemplateQuotaItemResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteTemplateQuotaItemResponseBody) SetRequestId(v string) *DeleteTemplateQuotaItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateQuotaItemResponseBody) Validate() error {
	return dara.Validate(s)
}
