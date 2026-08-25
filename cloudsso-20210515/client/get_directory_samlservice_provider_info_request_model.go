// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectorySAMLServiceProviderInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetDirectorySAMLServiceProviderInfoRequest
	GetDirectoryId() *string
}

type GetDirectorySAMLServiceProviderInfoRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectorySAMLServiceProviderInfoRequest) SetDirectoryId(v string) *GetDirectorySAMLServiceProviderInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoRequest) Validate() error {
	return dara.Validate(s)
}
