// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *GetAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *GetAccessConfigurationRequest
	GetDirectoryId() *string
}

type GetAccessConfigurationRequest struct {
	// The access configuration ID.
	//
	// example:
	//
	// ac-00ccule7tadaijxc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *GetAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetAccessConfigurationRequest) SetAccessConfigurationId(v string) *GetAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetAccessConfigurationRequest) SetDirectoryId(v string) *GetAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
