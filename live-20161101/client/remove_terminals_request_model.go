// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTerminalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RemoveTerminalsRequest
	GetAppId() *string
	SetChannelId(v string) *RemoveTerminalsRequest
	GetChannelId() *string
	SetTerminalIds(v []*string) *RemoveTerminalsRequest
	GetTerminalIds() []*string
}

type RemoveTerminalsRequest struct {
	// The ID of the application. You can specify only one application ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel. You can specify only one channel ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The IDs of the users that you want to remove.
	//
	// This parameter is required.
	TerminalIds []*string `json:"TerminalIds,omitempty" xml:"TerminalIds,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsRequest) GetAppId() *string {
	return s.AppId
}

func (s *RemoveTerminalsRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RemoveTerminalsRequest) GetTerminalIds() []*string {
	return s.TerminalIds
}

func (s *RemoveTerminalsRequest) SetAppId(v string) *RemoveTerminalsRequest {
	s.AppId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetChannelId(v string) *RemoveTerminalsRequest {
	s.ChannelId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetTerminalIds(v []*string) *RemoveTerminalsRequest {
	s.TerminalIds = v
	return s
}

func (s *RemoveTerminalsRequest) Validate() error {
	return dara.Validate(s)
}
