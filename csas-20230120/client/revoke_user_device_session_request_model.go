// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserDeviceSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTag(v string) *RevokeUserDeviceSessionRequest
	GetDevTag() *string
	SetSaseUserId(v string) *RevokeUserDeviceSessionRequest
	GetSaseUserId() *string
}

type RevokeUserDeviceSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2987b3e0-8108-2f99-4d18-3b4f1c1c36d7
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// su_xxxxxxxx
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s RevokeUserDeviceSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserDeviceSessionRequest) GoString() string {
	return s.String()
}

func (s *RevokeUserDeviceSessionRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *RevokeUserDeviceSessionRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *RevokeUserDeviceSessionRequest) SetDevTag(v string) *RevokeUserDeviceSessionRequest {
	s.DevTag = &v
	return s
}

func (s *RevokeUserDeviceSessionRequest) SetSaseUserId(v string) *RevokeUserDeviceSessionRequest {
	s.SaseUserId = &v
	return s
}

func (s *RevokeUserDeviceSessionRequest) Validate() error {
	return dara.Validate(s)
}
