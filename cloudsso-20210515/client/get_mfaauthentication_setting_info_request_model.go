// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetMFAAuthenticationSettingInfoRequest
	GetDirectoryId() *string
}

type GetMFAAuthenticationSettingInfoRequest struct {
	// The directory ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetMFAAuthenticationSettingInfoRequest) SetDirectoryId(v string) *GetMFAAuthenticationSettingInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoRequest) Validate() error {
	return dara.Validate(s)
}
