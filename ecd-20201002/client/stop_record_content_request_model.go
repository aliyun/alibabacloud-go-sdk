// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *StopRecordContentRequest
	GetClientId() *string
	SetClientOS(v string) *StopRecordContentRequest
	GetClientOS() *string
	SetClientVersion(v string) *StopRecordContentRequest
	GetClientVersion() *string
	SetDesktopId(v string) *StopRecordContentRequest
	GetDesktopId() *string
	SetLoginToken(v string) *StopRecordContentRequest
	GetLoginToken() *string
	SetRegionId(v string) *StopRecordContentRequest
	GetRegionId() *string
	SetSessionId(v string) *StopRecordContentRequest
	GetSessionId() *string
}

type StopRecordContentRequest struct {
	// This parameter is required.
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopRecordContentRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StopRecordContentRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StopRecordContentRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StopRecordContentRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StopRecordContentRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *StopRecordContentRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StopRecordContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopRecordContentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopRecordContentRequest) SetClientId(v string) *StopRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StopRecordContentRequest) SetClientOS(v string) *StopRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StopRecordContentRequest) SetClientVersion(v string) *StopRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopRecordContentRequest) SetDesktopId(v string) *StopRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StopRecordContentRequest) SetLoginToken(v string) *StopRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StopRecordContentRequest) SetRegionId(v string) *StopRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StopRecordContentRequest) SetSessionId(v string) *StopRecordContentRequest {
	s.SessionId = &v
	return s
}

func (s *StopRecordContentRequest) Validate() error {
	return dara.Validate(s)
}
