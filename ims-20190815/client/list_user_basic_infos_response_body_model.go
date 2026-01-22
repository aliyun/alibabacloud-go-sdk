// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBasicInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUserBasicInfosResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListUserBasicInfosResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListUserBasicInfosResponseBody
	GetRequestId() *string
	SetUserBasicInfos(v *ListUserBasicInfosResponseBodyUserBasicInfos) *ListUserBasicInfosResponseBody
	GetUserBasicInfos() *ListUserBasicInfosResponseBodyUserBasicInfos
}

type ListUserBasicInfosResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF2B25FD-CADE-445B-BE4D-E082E0FF1A0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The basic information about the RAM users.
	UserBasicInfos *ListUserBasicInfosResponseBodyUserBasicInfos `json:"UserBasicInfos,omitempty" xml:"UserBasicInfos,omitempty" type:"Struct"`
}

func (s ListUserBasicInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUserBasicInfosResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListUserBasicInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserBasicInfosResponseBody) GetUserBasicInfos() *ListUserBasicInfosResponseBodyUserBasicInfos {
	return s.UserBasicInfos
}

func (s *ListUserBasicInfosResponseBody) SetIsTruncated(v bool) *ListUserBasicInfosResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetMarker(v string) *ListUserBasicInfosResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetRequestId(v string) *ListUserBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetUserBasicInfos(v *ListUserBasicInfosResponseBodyUserBasicInfos) *ListUserBasicInfosResponseBody {
	s.UserBasicInfos = v
	return s
}

func (s *ListUserBasicInfosResponseBody) Validate() error {
	if s.UserBasicInfos != nil {
		if err := s.UserBasicInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserBasicInfosResponseBodyUserBasicInfos struct {
	UserBasicInfo []*ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo `json:"UserBasicInfo,omitempty" xml:"UserBasicInfo,omitempty" type:"Repeated"`
}

func (s ListUserBasicInfosResponseBodyUserBasicInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosResponseBodyUserBasicInfos) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfos) GetUserBasicInfo() []*ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	return s.UserBasicInfo
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfos) SetUserBasicInfo(v []*ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) *ListUserBasicInfosResponseBodyUserBasicInfos {
	s.UserBasicInfo = v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfos) Validate() error {
	if s.UserBasicInfo != nil {
		for _, item := range s.UserBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the RAM user.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GetUserId() *string {
	return s.UserId
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetDisplayName(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.DisplayName = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetStatus(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.Status = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetUserId(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.UserId = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetUserPrincipalName(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) Validate() error {
	return dara.Validate(s)
}
