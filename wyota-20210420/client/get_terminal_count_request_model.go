// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerminalCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v int32) *GetTerminalCountRequest
	GetClientType() *int32
}

type GetTerminalCountRequest struct {
	// The terminal type. Valid values:
	//
	// - 1: hardware terminal.
	//
	// - 2: software terminal.
	//
	// - 3: secure browser plug-in.
	//
	// - 4: GuestOS application.
	//
	// - 5: DingTalk Wuying plug-in.
	//
	// - 6: cloud application component.
	//
	// - 7: Cloud Hub.
	//
	// - 8: H5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
}

func (s GetTerminalCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTerminalCountRequest) GoString() string {
	return s.String()
}

func (s *GetTerminalCountRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetTerminalCountRequest) SetClientType(v int32) *GetTerminalCountRequest {
	s.ClientType = &v
	return s
}

func (s *GetTerminalCountRequest) Validate() error {
	return dara.Validate(s)
}
