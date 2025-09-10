// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateQuotaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteTemplateQuotaItemRequest
	GetId() *string
}

type DeleteTemplateQuotaItemRequest struct {
	// The ID of the quota template.
	//
	// You can call the [ListQuotaApplicationTemplates](https://help.aliyun.com/document_detail/450403.html) operation to obtain the ID of a quota template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTemplateQuotaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateQuotaItemRequest) GetId() *string {
	return s.Id
}

func (s *DeleteTemplateQuotaItemRequest) SetId(v string) *DeleteTemplateQuotaItemRequest {
	s.Id = &v
	return s
}

func (s *DeleteTemplateQuotaItemRequest) Validate() error {
	return dara.Validate(s)
}
