// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInstructionUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v string) *PbcInstructionUpdateCmd
	GetDocument() *string
	SetPbcVersionId(v int64) *PbcInstructionUpdateCmd
	GetPbcVersionId() *int64
}

type PbcInstructionUpdateCmd struct {
	Document     *string `json:"document,omitempty" xml:"document,omitempty"`
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
}

func (s PbcInstructionUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcInstructionUpdateCmd) GoString() string {
	return s.String()
}

func (s *PbcInstructionUpdateCmd) GetDocument() *string {
	return s.Document
}

func (s *PbcInstructionUpdateCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcInstructionUpdateCmd) SetDocument(v string) *PbcInstructionUpdateCmd {
	s.Document = &v
	return s
}

func (s *PbcInstructionUpdateCmd) SetPbcVersionId(v int64) *PbcInstructionUpdateCmd {
	s.PbcVersionId = &v
	return s
}

func (s *PbcInstructionUpdateCmd) Validate() error {
	return dara.Validate(s)
}
