// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitStagingRoutineCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDescription(v string) *CommitStagingRoutineCodeRequest
	GetCodeDescription() *string
	SetName(v string) *CommitStagingRoutineCodeRequest
	GetName() *string
}

type CommitStagingRoutineCodeRequest struct {
	// The description of the code version.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello World
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CommitStagingRoutineCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitStagingRoutineCodeRequest) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *CommitStagingRoutineCodeRequest) GetName() *string {
	return s.Name
}

func (s *CommitStagingRoutineCodeRequest) SetCodeDescription(v string) *CommitStagingRoutineCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *CommitStagingRoutineCodeRequest) SetName(v string) *CommitStagingRoutineCodeRequest {
	s.Name = &v
	return s
}

func (s *CommitStagingRoutineCodeRequest) Validate() error {
	return dara.Validate(s)
}
