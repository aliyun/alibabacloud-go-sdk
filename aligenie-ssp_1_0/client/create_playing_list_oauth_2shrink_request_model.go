// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListOAuth2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *CreatePlayingListOAuth2ShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenCreatePlayingListRequestShrink(v string) *CreatePlayingListOAuth2ShrinkRequest
	GetOpenCreatePlayingListRequestShrink() *string
}

type CreatePlayingListOAuth2ShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenCreatePlayingListRequestShrink *string `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty"`
}

func (s CreatePlayingListOAuth2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2ShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CreatePlayingListOAuth2ShrinkRequest) GetOpenCreatePlayingListRequestShrink() *string {
	return s.OpenCreatePlayingListRequestShrink
}

func (s *CreatePlayingListOAuth2ShrinkRequest) SetDeviceInfoShrink(v string) *CreatePlayingListOAuth2ShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreatePlayingListOAuth2ShrinkRequest) SetOpenCreatePlayingListRequestShrink(v string) *CreatePlayingListOAuth2ShrinkRequest {
	s.OpenCreatePlayingListRequestShrink = &v
	return s
}

func (s *CreatePlayingListOAuth2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
