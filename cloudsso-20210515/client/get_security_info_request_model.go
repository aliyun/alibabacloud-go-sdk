// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetSecurityInfoRequest
	GetDirectoryId() *string
	SetUserId(v string) *GetSecurityInfoRequest
	GetUserId() *string
}

type GetSecurityInfoRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSecurityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetSecurityInfoRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetSecurityInfoRequest) SetDirectoryId(v string) *GetSecurityInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSecurityInfoRequest) SetUserId(v string) *GetSecurityInfoRequest {
	s.UserId = &v
	return s
}

func (s *GetSecurityInfoRequest) Validate() error {
	return dara.Validate(s)
}
