// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDocumentChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunks(v string) *UpsertDocumentChunksRequest
	GetChunks() *string
	SetDocumentName(v string) *UpsertDocumentChunksRequest
	GetDocumentName() *string
	SetKbUuid(v string) *UpsertDocumentChunksRequest
	GetKbUuid() *string
}

type UpsertDocumentChunksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "Id": "2e12aeb5-52cd-4834-bcd8-****",
	//
	//     "Content": "test1"
	//
	//   },
	//
	//   {
	//
	//     "Id": "2fdnwefi5-dsad-4t35-bcd8-****",
	//
	//     "Content": "test2"
	//
	//   }
	//
	// ]
	Chunks *string `json:"Chunks,omitempty" xml:"Chunks,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s UpsertDocumentChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertDocumentChunksRequest) GoString() string {
	return s.String()
}

func (s *UpsertDocumentChunksRequest) GetChunks() *string {
	return s.Chunks
}

func (s *UpsertDocumentChunksRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *UpsertDocumentChunksRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *UpsertDocumentChunksRequest) SetChunks(v string) *UpsertDocumentChunksRequest {
	s.Chunks = &v
	return s
}

func (s *UpsertDocumentChunksRequest) SetDocumentName(v string) *UpsertDocumentChunksRequest {
	s.DocumentName = &v
	return s
}

func (s *UpsertDocumentChunksRequest) SetKbUuid(v string) *UpsertDocumentChunksRequest {
	s.KbUuid = &v
	return s
}

func (s *UpsertDocumentChunksRequest) Validate() error {
	return dara.Validate(s)
}
