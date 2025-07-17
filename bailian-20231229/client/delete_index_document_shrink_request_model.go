// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIdsShrink(v string) *DeleteIndexDocumentShrinkRequest
	GetDocumentIdsShrink() *string
	SetIndexId(v string) *DeleteIndexDocumentShrinkRequest
	GetIndexId() *string
}

type DeleteIndexDocumentShrinkRequest struct {
	// The list of the primary key IDs of the documents.
	//
	// This parameter is required.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentShrinkRequest) GetDocumentIdsShrink() *string {
	return s.DocumentIdsShrink
}

func (s *DeleteIndexDocumentShrinkRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *DeleteIndexDocumentShrinkRequest) SetDocumentIdsShrink(v string) *DeleteIndexDocumentShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *DeleteIndexDocumentShrinkRequest) SetIndexId(v string) *DeleteIndexDocumentShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *DeleteIndexDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
