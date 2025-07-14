// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceCodeRepoBranch interface {
	dara.Model
	String() string
	GoString() string
	SetCommitId(v string) *SourceCodeRepoBranch
	GetCommitId() *string
	SetName(v string) *SourceCodeRepoBranch
	GetName() *string
}

type SourceCodeRepoBranch struct {
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SourceCodeRepoBranch) String() string {
	return dara.Prettify(s)
}

func (s SourceCodeRepoBranch) GoString() string {
	return s.String()
}

func (s *SourceCodeRepoBranch) GetCommitId() *string {
	return s.CommitId
}

func (s *SourceCodeRepoBranch) GetName() *string {
	return s.Name
}

func (s *SourceCodeRepoBranch) SetCommitId(v string) *SourceCodeRepoBranch {
	s.CommitId = &v
	return s
}

func (s *SourceCodeRepoBranch) SetName(v string) *SourceCodeRepoBranch {
	s.Name = &v
	return s
}

func (s *SourceCodeRepoBranch) Validate() error {
	return dara.Validate(s)
}
