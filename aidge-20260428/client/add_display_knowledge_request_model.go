// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisplayKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *AddDisplayKnowledgeRequest
	GetRagId() *string
	SetRecords(v []*AddDisplayKnowledgeRequestRecords) *AddDisplayKnowledgeRequest
	GetRecords() []*AddDisplayKnowledgeRequestRecords
}

type AddDisplayKnowledgeRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The records to add. A maximum of 500 records are supported.
	//
	// This parameter is required.
	Records []*AddDisplayKnowledgeRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s AddDisplayKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeRequest) GetRagId() *string {
	return s.RagId
}

func (s *AddDisplayKnowledgeRequest) GetRecords() []*AddDisplayKnowledgeRequestRecords {
	return s.Records
}

func (s *AddDisplayKnowledgeRequest) SetRagId(v string) *AddDisplayKnowledgeRequest {
	s.RagId = &v
	return s
}

func (s *AddDisplayKnowledgeRequest) SetRecords(v []*AddDisplayKnowledgeRequestRecords) *AddDisplayKnowledgeRequest {
	s.Records = v
	return s
}

func (s *AddDisplayKnowledgeRequest) Validate() error {
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

type AddDisplayKnowledgeRequestRecords struct {
	// The HTTP(S) URL of the product image.
	//
	// example:
	//
	// https://xxxx/img_1784599244373_9c2309a0.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The SKU ID. If this parameter is left empty, the platform generates the ID.
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

func (s AddDisplayKnowledgeRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeRequestRecords) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeRequestRecords) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *AddDisplayKnowledgeRequestRecords) GetSkuId() *string {
	return s.SkuId
}

func (s *AddDisplayKnowledgeRequestRecords) GetSkuName() *string {
	return s.SkuName
}

func (s *AddDisplayKnowledgeRequestRecords) SetImageUrl(v string) *AddDisplayKnowledgeRequestRecords {
	s.ImageUrl = &v
	return s
}

func (s *AddDisplayKnowledgeRequestRecords) SetSkuId(v string) *AddDisplayKnowledgeRequestRecords {
	s.SkuId = &v
	return s
}

func (s *AddDisplayKnowledgeRequestRecords) SetSkuName(v string) *AddDisplayKnowledgeRequestRecords {
	s.SkuName = &v
	return s
}

func (s *AddDisplayKnowledgeRequestRecords) Validate() error {
	return dara.Validate(s)
}
