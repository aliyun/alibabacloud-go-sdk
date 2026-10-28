// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDesktopManagerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *BatchSetDesktopManagerRequest
	GetBusinessChannel() *string
	SetIsDesktopManager(v string) *BatchSetDesktopManagerRequest
	GetIsDesktopManager() *string
	SetUsers(v []*string) *BatchSetDesktopManagerRequest
	GetUsers() []*string
}

type BatchSetDesktopManagerRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// Specifies whether the convenience account has local administrator permissions on the cloud desktop.
	//
	// example:
	//
	// 1
	IsDesktopManager *string `json:"IsDesktopManager,omitempty" xml:"IsDesktopManager,omitempty"`
	// The list of convenience accounts.
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

func (s *BatchSetDesktopManagerRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *BatchSetDesktopManagerRequest) GetIsDesktopManager() *string {
	return s.IsDesktopManager
}

func (s *BatchSetDesktopManagerRequest) GetUsers() []*string {
	return s.Users
}

func (s *BatchSetDesktopManagerRequest) SetBusinessChannel(v string) *BatchSetDesktopManagerRequest {
	s.BusinessChannel = &v
	return s
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
