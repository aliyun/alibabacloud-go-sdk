// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *StartRecordContentRequest
	GetClientId() *string
	SetClientOS(v string) *StartRecordContentRequest
	GetClientOS() *string
	SetClientVersion(v string) *StartRecordContentRequest
	GetClientVersion() *string
	SetDesktopId(v string) *StartRecordContentRequest
	GetDesktopId() *string
	SetFilePath(v string) *StartRecordContentRequest
	GetFilePath() *string
	SetLoginToken(v string) *StartRecordContentRequest
	GetLoginToken() *string
	SetRegionId(v string) *StartRecordContentRequest
	GetRegionId() *string
	SetSessionId(v string) *StartRecordContentRequest
	GetSessionId() *string
}

type StartRecordContentRequest struct {
	// This parameter is required.
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StartRecordContentRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StartRecordContentRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StartRecordContentRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StartRecordContentRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StartRecordContentRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *StartRecordContentRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *StartRecordContentRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StartRecordContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartRecordContentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartRecordContentRequest) SetClientId(v string) *StartRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StartRecordContentRequest) SetClientOS(v string) *StartRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StartRecordContentRequest) SetClientVersion(v string) *StartRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartRecordContentRequest) SetDesktopId(v string) *StartRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StartRecordContentRequest) SetFilePath(v string) *StartRecordContentRequest {
	s.FilePath = &v
	return s
}

func (s *StartRecordContentRequest) SetLoginToken(v string) *StartRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StartRecordContentRequest) SetRegionId(v string) *StartRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StartRecordContentRequest) SetSessionId(v string) *StartRecordContentRequest {
	s.SessionId = &v
	return s
}

func (s *StartRecordContentRequest) Validate() error {
	return dara.Validate(s)
}
