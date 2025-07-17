// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIds(v []*string) *DeleteIndexDocumentRequest
	GetDocumentIds() []*string
	SetIndexId(v string) *DeleteIndexDocumentRequest
	GetIndexId() *string
}

type DeleteIndexDocumentRequest struct {
	// The list of the primary key IDs of the documents.
	//
	// This parameter is required.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *DeleteIndexDocumentRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *DeleteIndexDocumentRequest) SetDocumentIds(v []*string) *DeleteIndexDocumentRequest {
	s.DocumentIds = v
	return s
}

func (s *DeleteIndexDocumentRequest) SetIndexId(v string) *DeleteIndexDocumentRequest {
	s.IndexId = &v
	return s
}

func (s *DeleteIndexDocumentRequest) Validate() error {
	return dara.Validate(s)
}
