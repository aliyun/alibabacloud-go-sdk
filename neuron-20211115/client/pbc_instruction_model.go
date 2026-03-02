// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInstruction interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v string) *PbcInstruction
	GetDocument() *string
	SetId(v int64) *PbcInstruction
	GetId() *int64
	SetPbcVersionId(v int64) *PbcInstruction
	GetPbcVersionId() *int64
	SetRequestId(v string) *PbcInstruction
	GetRequestId() *string
}

type PbcInstruction struct {
	// This parameter is required.
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcInstruction) String() string {
	return dara.Prettify(s)
}

func (s PbcInstruction) GoString() string {
	return s.String()
}

func (s *PbcInstruction) GetDocument() *string {
	return s.Document
}

func (s *PbcInstruction) GetId() *int64 {
	return s.Id
}

func (s *PbcInstruction) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcInstruction) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcInstruction) SetDocument(v string) *PbcInstruction {
	s.Document = &v
	return s
}

func (s *PbcInstruction) SetId(v int64) *PbcInstruction {
	s.Id = &v
	return s
}

func (s *PbcInstruction) SetPbcVersionId(v int64) *PbcInstruction {
	s.PbcVersionId = &v
	return s
}

func (s *PbcInstruction) SetRequestId(v string) *PbcInstruction {
	s.RequestId = &v
	return s
}

func (s *PbcInstruction) Validate() error {
	return dara.Validate(s)
}
