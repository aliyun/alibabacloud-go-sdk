// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStagingRoutineCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDescription(v string) *UploadStagingRoutineCodeRequest
	GetCodeDescription() *string
	SetName(v string) *UploadStagingRoutineCodeRequest
	GetName() *string
}

type UploadStagingRoutineCodeRequest struct {
	// The description of the version.
	//
	// example:
	//
	// desc
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The name of the routine. The name needs to be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UploadStagingRoutineCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadStagingRoutineCodeRequest) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *UploadStagingRoutineCodeRequest) GetName() *string {
	return s.Name
}

func (s *UploadStagingRoutineCodeRequest) SetCodeDescription(v string) *UploadStagingRoutineCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *UploadStagingRoutineCodeRequest) SetName(v string) *UploadStagingRoutineCodeRequest {
	s.Name = &v
	return s
}

func (s *UploadStagingRoutineCodeRequest) Validate() error {
	return dara.Validate(s)
}
