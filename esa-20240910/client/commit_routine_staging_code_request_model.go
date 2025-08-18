// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitRoutineStagingCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDescription(v string) *CommitRoutineStagingCodeRequest
	GetCodeDescription() *string
	SetName(v string) *CommitRoutineStagingCodeRequest
	GetName() *string
}

type CommitRoutineStagingCodeRequest struct {
	// The description of the code version.
	//
	// example:
	//
	// description of this code ver
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The routine name.
	//
	// This parameter is required.
	//
	// example:
	//
	// CommitRoutineStagingCode
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CommitRoutineStagingCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitRoutineStagingCodeRequest) GoString() string {
	return s.String()
}

func (s *CommitRoutineStagingCodeRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *CommitRoutineStagingCodeRequest) GetName() *string {
	return s.Name
}

func (s *CommitRoutineStagingCodeRequest) SetCodeDescription(v string) *CommitRoutineStagingCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *CommitRoutineStagingCodeRequest) SetName(v string) *CommitRoutineStagingCodeRequest {
	s.Name = &v
	return s
}

func (s *CommitRoutineStagingCodeRequest) Validate() error {
	return dara.Validate(s)
}
