// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetLoginPreferenceRequest
	GetDirectoryId() *string
}

type GetLoginPreferenceRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetLoginPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceRequest) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetLoginPreferenceRequest) SetDirectoryId(v string) *GetLoginPreferenceRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetLoginPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
