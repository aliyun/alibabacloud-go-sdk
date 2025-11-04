// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAvatarResponseBodyData) *GetAvatarResponseBody
	GetData() *GetAvatarResponseBodyData
	SetRequestId(v string) *GetAvatarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAvatarResponseBody
	GetSuccess() *bool
}

type GetAvatarResponseBody struct {
	// The data returned.
	Data *GetAvatarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvatarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvatarResponseBody) GetData() *GetAvatarResponseBodyData {
	return s.Data
}

func (s *GetAvatarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvatarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAvatarResponseBody) SetData(v *GetAvatarResponseBodyData) *GetAvatarResponseBody {
	s.Data = v
	return s
}

func (s *GetAvatarResponseBody) SetRequestId(v string) *GetAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvatarResponseBody) SetSuccess(v bool) *GetAvatarResponseBody {
	s.Success = &v
	return s
}

func (s *GetAvatarResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAvatarResponseBodyData struct {
	// The information about the digital human.
	Avatar *GetAvatarResponseBodyDataAvatar `json:"Avatar,omitempty" xml:"Avatar,omitempty" type:"Struct"`
}

func (s GetAvatarResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvatarResponseBodyData) GetAvatar() *GetAvatarResponseBodyDataAvatar {
	return s.Avatar
}

func (s *GetAvatarResponseBodyData) SetAvatar(v *GetAvatarResponseBodyDataAvatar) *GetAvatarResponseBodyData {
	s.Avatar = v
	return s
}

func (s *GetAvatarResponseBodyData) Validate() error {
	if s.Avatar != nil {
		if err := s.Avatar.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAvatarResponseBodyDataAvatar struct {
	// The description of the digital human.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// The ID of the digital human.
	//
	// example:
	//
	// Avatar-XXXX
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The name of the digital human.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The type of the digital human.
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	// The height of the digital human image in pixels.
	//
	// example:
	//
	// 1920
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The media asset ID of the portrait image.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Portrait *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/thumbnail.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// Indicates whether the digital human supports alpha channels.
	//
	// example:
	//
	// true
	Transparent *bool `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
	// The width of the digital human image in pixels.
	//
	// example:
	//
	// 1080
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetAvatarResponseBodyDataAvatar) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarResponseBodyDataAvatar) GoString() string {
	return s.String()
}

func (s *GetAvatarResponseBodyDataAvatar) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *GetAvatarResponseBodyDataAvatar) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetAvatarResponseBodyDataAvatar) GetAvatarName() *string {
	return s.AvatarName
}

func (s *GetAvatarResponseBodyDataAvatar) GetAvatarType() *string {
	return s.AvatarType
}

func (s *GetAvatarResponseBodyDataAvatar) GetHeight() *int32 {
	return s.Height
}

func (s *GetAvatarResponseBodyDataAvatar) GetPortrait() *string {
	return s.Portrait
}

func (s *GetAvatarResponseBodyDataAvatar) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetAvatarResponseBodyDataAvatar) GetTransparent() *bool {
	return s.Transparent
}

func (s *GetAvatarResponseBodyDataAvatar) GetWidth() *int32 {
	return s.Width
}

func (s *GetAvatarResponseBodyDataAvatar) SetAvatarDescription(v string) *GetAvatarResponseBodyDataAvatar {
	s.AvatarDescription = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetAvatarId(v string) *GetAvatarResponseBodyDataAvatar {
	s.AvatarId = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetAvatarName(v string) *GetAvatarResponseBodyDataAvatar {
	s.AvatarName = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetAvatarType(v string) *GetAvatarResponseBodyDataAvatar {
	s.AvatarType = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetHeight(v int32) *GetAvatarResponseBodyDataAvatar {
	s.Height = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetPortrait(v string) *GetAvatarResponseBodyDataAvatar {
	s.Portrait = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetThumbnail(v string) *GetAvatarResponseBodyDataAvatar {
	s.Thumbnail = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetTransparent(v bool) *GetAvatarResponseBodyDataAvatar {
	s.Transparent = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) SetWidth(v int32) *GetAvatarResponseBodyDataAvatar {
	s.Width = &v
	return s
}

func (s *GetAvatarResponseBodyDataAvatar) Validate() error {
	return dara.Validate(s)
}
