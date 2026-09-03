// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttributePassingSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetAttributePassingSettingRequest
	GetDirectoryId() *string
}

type GetAttributePassingSettingRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetAttributePassingSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttributePassingSettingRequest) GoString() string {
	return s.String()
}

func (s *GetAttributePassingSettingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetAttributePassingSettingRequest) SetDirectoryId(v string) *GetAttributePassingSettingRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetAttributePassingSettingRequest) Validate() error {
	return dara.Validate(s)
}
