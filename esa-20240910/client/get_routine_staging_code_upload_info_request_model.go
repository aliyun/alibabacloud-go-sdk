// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineStagingCodeUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDescription(v string) *GetRoutineStagingCodeUploadInfoRequest
	GetCodeDescription() *string
	SetName(v string) *GetRoutineStagingCodeUploadInfoRequest
	GetName() *string
}

type GetRoutineStagingCodeUploadInfoRequest struct {
	// The code description.
	//
	// example:
	//
	// the description of code
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The routine name.
	//
	// This parameter is required.
	//
	// example:
	//
	// GetRoutineStagingCodeUploadInfo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRoutineStagingCodeUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineStagingCodeUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineStagingCodeUploadInfoRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *GetRoutineStagingCodeUploadInfoRequest) GetName() *string {
	return s.Name
}

func (s *GetRoutineStagingCodeUploadInfoRequest) SetCodeDescription(v string) *GetRoutineStagingCodeUploadInfoRequest {
	s.CodeDescription = &v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoRequest) SetName(v string) *GetRoutineStagingCodeUploadInfoRequest {
	s.Name = &v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}
