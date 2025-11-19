// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DeleteDesktopImageRequest
	GetApiKey() *string
	SetImageId(v string) *DeleteDesktopImageRequest
	GetImageId() *string
	SetIsCleanImageCode(v bool) *DeleteDesktopImageRequest
	GetIsCleanImageCode() *bool
	SetSessionId(v string) *DeleteDesktopImageRequest
	GetSessionId() *string
}

type DeleteDesktopImageRequest struct {
	// This parameter is required.
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	IsCleanImageCode *bool   `json:"IsCleanImageCode,omitempty" xml:"IsCleanImageCode,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteDesktopImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopImageRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeleteDesktopImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteDesktopImageRequest) GetIsCleanImageCode() *bool {
	return s.IsCleanImageCode
}

func (s *DeleteDesktopImageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteDesktopImageRequest) SetApiKey(v string) *DeleteDesktopImageRequest {
	s.ApiKey = &v
	return s
}

func (s *DeleteDesktopImageRequest) SetImageId(v string) *DeleteDesktopImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteDesktopImageRequest) SetIsCleanImageCode(v bool) *DeleteDesktopImageRequest {
	s.IsCleanImageCode = &v
	return s
}

func (s *DeleteDesktopImageRequest) SetSessionId(v string) *DeleteDesktopImageRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteDesktopImageRequest) Validate() error {
	return dara.Validate(s)
}
