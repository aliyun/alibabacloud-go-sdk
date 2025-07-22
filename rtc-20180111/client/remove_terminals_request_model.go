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
	SetOwnerId(v int64) *RemoveTerminalsRequest
	GetOwnerId() *int64
	SetTerminalIds(v []*string) *RemoveTerminalsRequest
	GetTerminalIds() []*string
}

type RemoveTerminalsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1811xxxx
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

func (s *RemoveTerminalsRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *RemoveTerminalsRequest) SetOwnerId(v int64) *RemoveTerminalsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetTerminalIds(v []*string) *RemoveTerminalsRequest {
	s.TerminalIds = v
	return s
}

func (s *RemoveTerminalsRequest) Validate() error {
	return dara.Validate(s)
}
