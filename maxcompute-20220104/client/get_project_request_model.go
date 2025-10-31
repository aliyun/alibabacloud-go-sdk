// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetProjectRequest
	GetVerbose() *bool
	SetWithQuotaProductType(v bool) *GetProjectRequest
	GetWithQuotaProductType() *bool
	SetWithStorageTierInfo(v bool) *GetProjectRequest
	GetWithStorageTierInfo() *bool
}

type GetProjectRequest struct {
	// Specifies whether to use additional information.
	//
	// example:
	//
	// true
	Verbose *bool `json:"verbose,omitempty" xml:"verbose,omitempty"`
	// example:
	//
	// true
	WithQuotaProductType *bool `json:"withQuotaProductType,omitempty" xml:"withQuotaProductType,omitempty"`
	// example:
	//
	// true
	WithStorageTierInfo *bool `json:"withStorageTierInfo,omitempty" xml:"withStorageTierInfo,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetProjectRequest) GetWithQuotaProductType() *bool {
	return s.WithQuotaProductType
}

func (s *GetProjectRequest) GetWithStorageTierInfo() *bool {
	return s.WithStorageTierInfo
}

func (s *GetProjectRequest) SetVerbose(v bool) *GetProjectRequest {
	s.Verbose = &v
	return s
}

func (s *GetProjectRequest) SetWithQuotaProductType(v bool) *GetProjectRequest {
	s.WithQuotaProductType = &v
	return s
}

func (s *GetProjectRequest) SetWithStorageTierInfo(v bool) *GetProjectRequest {
	s.WithStorageTierInfo = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
