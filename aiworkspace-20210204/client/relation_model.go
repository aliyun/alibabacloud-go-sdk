// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelation interface {
	dara.Model
	String() string
	GoString() string
	SetErrMsg(v string) *Relation
	GetErrMsg() *string
	SetLineageRelation(v *LineageRelation) *Relation
	GetLineageRelation() *LineageRelation
	SetResult(v bool) *Relation
	GetResult() *bool
}

type Relation struct {
	ErrMsg          *string          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	LineageRelation *LineageRelation `json:"LineageRelation,omitempty" xml:"LineageRelation,omitempty"`
	Result          *bool            `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s Relation) String() string {
	return dara.Prettify(s)
}

func (s Relation) GoString() string {
	return s.String()
}

func (s *Relation) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *Relation) GetLineageRelation() *LineageRelation {
	return s.LineageRelation
}

func (s *Relation) GetResult() *bool {
	return s.Result
}

func (s *Relation) SetErrMsg(v string) *Relation {
	s.ErrMsg = &v
	return s
}

func (s *Relation) SetLineageRelation(v *LineageRelation) *Relation {
	s.LineageRelation = v
	return s
}

func (s *Relation) SetResult(v bool) *Relation {
	s.Result = &v
	return s
}

func (s *Relation) Validate() error {
	return dara.Validate(s)
}
