// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisplayKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *UpdateDisplayKnowledgeRequest
	GetRagId() *string
	SetRecords(v []*UpdateDisplayKnowledgeRequestRecords) *UpdateDisplayKnowledgeRequest
	GetRecords() []*UpdateDisplayKnowledgeRequestRecords
}

type UpdateDisplayKnowledgeRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The records to update. A maximum of 500 records can be specified.
	//
	// This parameter is required.
	Records []*UpdateDisplayKnowledgeRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s UpdateDisplayKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeRequest) GetRagId() *string {
	return s.RagId
}

func (s *UpdateDisplayKnowledgeRequest) GetRecords() []*UpdateDisplayKnowledgeRequestRecords {
	return s.Records
}

func (s *UpdateDisplayKnowledgeRequest) SetRagId(v string) *UpdateDisplayKnowledgeRequest {
	s.RagId = &v
	return s
}

func (s *UpdateDisplayKnowledgeRequest) SetRecords(v []*UpdateDisplayKnowledgeRequestRecords) *UpdateDisplayKnowledgeRequest {
	s.Records = v
	return s
}

func (s *UpdateDisplayKnowledgeRequest) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDisplayKnowledgeRequestRecords struct {
	// The HTTP or HTTPS URL of the product image.
	//
	// example:
	//
	// https://xxxx/96f21ebdaeeba27a.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 1489638581
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The SKU ID. If this parameter is not empty, the original value is overwritten.
	//
	// example:
	//
	// sku-001
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The SKU name.
	//
	// example:
	//
	// Product front image
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
}

func (s UpdateDisplayKnowledgeRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeRequestRecords) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeRequestRecords) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateDisplayKnowledgeRequestRecords) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateDisplayKnowledgeRequestRecords) GetSkuId() *string {
	return s.SkuId
}

func (s *UpdateDisplayKnowledgeRequestRecords) GetSkuName() *string {
	return s.SkuName
}

func (s *UpdateDisplayKnowledgeRequestRecords) SetImageUrl(v string) *UpdateDisplayKnowledgeRequestRecords {
	s.ImageUrl = &v
	return s
}

func (s *UpdateDisplayKnowledgeRequestRecords) SetRecordId(v int64) *UpdateDisplayKnowledgeRequestRecords {
	s.RecordId = &v
	return s
}

func (s *UpdateDisplayKnowledgeRequestRecords) SetSkuId(v string) *UpdateDisplayKnowledgeRequestRecords {
	s.SkuId = &v
	return s
}

func (s *UpdateDisplayKnowledgeRequestRecords) SetSkuName(v string) *UpdateDisplayKnowledgeRequestRecords {
	s.SkuName = &v
	return s
}

func (s *UpdateDisplayKnowledgeRequestRecords) Validate() error {
	return dara.Validate(s)
}
