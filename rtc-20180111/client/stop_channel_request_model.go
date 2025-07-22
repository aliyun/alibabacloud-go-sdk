// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopChannelRequest
	GetAppId() *string
	SetChannelId(v string) *StopChannelRequest
	GetChannelId() *string
}

type StopChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s StopChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s StopChannelRequest) GoString() string {
	return s.String()
}

func (s *StopChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopChannelRequest) SetAppId(v string) *StopChannelRequest {
	s.AppId = &v
	return s
}

func (s *StopChannelRequest) SetChannelId(v string) *StopChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *StopChannelRequest) Validate() error {
	return dara.Validate(s)
}
