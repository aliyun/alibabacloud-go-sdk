// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteChannelRequest
	GetAppId() *string
	SetChannelId(v string) *DeleteChannelRequest
	GetChannelId() *string
	SetOwnerId(v int64) *DeleteChannelRequest
	GetOwnerId() *int64
}

type DeleteChannelRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the joined channel. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DeleteChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChannelRequest) SetAppId(v string) *DeleteChannelRequest {
	s.AppId = &v
	return s
}

func (s *DeleteChannelRequest) SetChannelId(v string) *DeleteChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DeleteChannelRequest) SetOwnerId(v int64) *DeleteChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChannelRequest) Validate() error {
	return dara.Validate(s)
}
