// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAvatarsResponseBodyData) *ListAvatarsResponseBody
	GetData() *ListAvatarsResponseBodyData
	SetRequestId(v string) *ListAvatarsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvatarsResponseBody
	GetSuccess() *bool
}

type ListAvatarsResponseBody struct {
	// The data returned.
	Data *ListAvatarsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAvatarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvatarsResponseBody) GetData() *ListAvatarsResponseBodyData {
	return s.Data
}

func (s *ListAvatarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvatarsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvatarsResponseBody) SetData(v *ListAvatarsResponseBodyData) *ListAvatarsResponseBody {
	s.Data = v
	return s
}

func (s *ListAvatarsResponseBody) SetRequestId(v string) *ListAvatarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvatarsResponseBody) SetSuccess(v bool) *ListAvatarsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvatarsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAvatarsResponseBodyData struct {
	// The queried digital humans.
	AvatarList []*ListAvatarsResponseBodyDataAvatarList `json:"AvatarList,omitempty" xml:"AvatarList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvatarsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvatarsResponseBodyData) GetAvatarList() []*ListAvatarsResponseBodyDataAvatarList {
	return s.AvatarList
}

func (s *ListAvatarsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAvatarsResponseBodyData) SetAvatarList(v []*ListAvatarsResponseBodyDataAvatarList) *ListAvatarsResponseBodyData {
	s.AvatarList = v
	return s
}

func (s *ListAvatarsResponseBodyData) SetTotalCount(v int64) *ListAvatarsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAvatarsResponseBodyData) Validate() error {
	if s.AvatarList != nil {
		for _, item := range s.AvatarList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvatarsResponseBodyDataAvatarList struct {
	// The description of the digital human.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// The ID of the digital human.
	//
	// example:
	//
	// Avatar-XXX
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The name of the digital human.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The type of the digital human.
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
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
	// Indicates whether the digital human image supports the alpha channels.
	//
	// example:
	//
	// true
	Transparent *bool `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
}

func (s ListAvatarsResponseBodyDataAvatarList) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarsResponseBodyDataAvatarList) GoString() string {
	return s.String()
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetAvatarId() *string {
	return s.AvatarId
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetAvatarName() *string {
	return s.AvatarName
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetAvatarType() *string {
	return s.AvatarType
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetPortrait() *string {
	return s.Portrait
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *ListAvatarsResponseBodyDataAvatarList) GetTransparent() *bool {
	return s.Transparent
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetAvatarDescription(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.AvatarDescription = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetAvatarId(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.AvatarId = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetAvatarName(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.AvatarName = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetAvatarType(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.AvatarType = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetPortrait(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.Portrait = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetThumbnail(v string) *ListAvatarsResponseBodyDataAvatarList {
	s.Thumbnail = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) SetTransparent(v bool) *ListAvatarsResponseBodyDataAvatarList {
	s.Transparent = &v
	return s
}

func (s *ListAvatarsResponseBodyDataAvatarList) Validate() error {
	return dara.Validate(s)
}
