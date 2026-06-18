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
	// The list of file IDs.
	//
	// This parameter is required.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
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
