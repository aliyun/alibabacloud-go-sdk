// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenMetaKnowledgeAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *GenMetaKnowledgeAssetRequest
	GetDbId() *int32
}

type GenMetaKnowledgeAssetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1860****
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
}

func (s GenMetaKnowledgeAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s GenMetaKnowledgeAssetRequest) GoString() string {
	return s.String()
}

func (s *GenMetaKnowledgeAssetRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GenMetaKnowledgeAssetRequest) SetDbId(v int32) *GenMetaKnowledgeAssetRequest {
	s.DbId = &v
	return s
}

func (s *GenMetaKnowledgeAssetRequest) Validate() error {
	return dara.Validate(s)
}
