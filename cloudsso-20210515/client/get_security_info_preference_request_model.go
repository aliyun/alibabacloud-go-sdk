// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetSecurityInfoPreferenceRequest
	GetDirectoryId() *string
}

type GetSecurityInfoPreferenceRequest struct {
	// The ID of the directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetSecurityInfoPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoPreferenceRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoPreferenceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetSecurityInfoPreferenceRequest) SetDirectoryId(v string) *GetSecurityInfoPreferenceRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSecurityInfoPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
