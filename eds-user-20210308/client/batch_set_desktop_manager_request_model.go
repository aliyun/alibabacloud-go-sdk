// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDesktopManagerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDesktopManager(v string) *BatchSetDesktopManagerRequest
	GetIsDesktopManager() *string
	SetUsers(v []*string) *BatchSetDesktopManagerRequest
	GetUsers() []*string
}

type BatchSetDesktopManagerRequest struct {
	// Whether the convenience account has the local administrator permissions on cloud computers.
	//
	// Valid values:
	//
	// 	- 0: no
	//
	// 	- 1 (default): yes
	//
	// example:
	//
	// 1
	IsDesktopManager *string `json:"IsDesktopManager,omitempty" xml:"IsDesktopManager,omitempty"`
	// The convenience accounts.
	//
	// This parameter is required.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s BatchSetDesktopManagerRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDesktopManagerRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerRequest) GetIsDesktopManager() *string {
	return s.IsDesktopManager
}

func (s *BatchSetDesktopManagerRequest) GetUsers() []*string {
	return s.Users
}

func (s *BatchSetDesktopManagerRequest) SetIsDesktopManager(v string) *BatchSetDesktopManagerRequest {
	s.IsDesktopManager = &v
	return s
}

func (s *BatchSetDesktopManagerRequest) SetUsers(v []*string) *BatchSetDesktopManagerRequest {
	s.Users = v
	return s
}

func (s *BatchSetDesktopManagerRequest) Validate() error {
	return dara.Validate(s)
}
