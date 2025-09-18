// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarUrl(v string) *GetUserBasicInfoResponseBody
	GetAvatarUrl() *string
	SetNickname(v string) *GetUserBasicInfoResponseBody
	GetNickname() *string
	SetOpenId(v string) *GetUserBasicInfoResponseBody
	GetOpenId() *string
	SetRequestId(v string) *GetUserBasicInfoResponseBody
	GetRequestId() *string
	SetUnionIds(v []*GetUserBasicInfoResponseBodyUnionIds) *GetUserBasicInfoResponseBody
	GetUnionIds() []*GetUserBasicInfoResponseBodyUnionIds
}

type GetUserBasicInfoResponseBody struct {
	// example:
	//
	// https://xxxxxx
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// xxxxxx
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// 3hPPBd9YuhfJQCzZ/07AAWdoO3K8zCb/KAqW96zPHXPiFkzjB/JfcWuuFHQQDaGZ4wVbNMV6wYuj075p/rhVLg==
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnionIds  []*GetUserBasicInfoResponseBodyUnionIds `json:"UnionIds,omitempty" xml:"UnionIds,omitempty" type:"Repeated"`
}

func (s GetUserBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponseBody) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetUserBasicInfoResponseBody) GetNickname() *string {
	return s.Nickname
}

func (s *GetUserBasicInfoResponseBody) GetOpenId() *string {
	return s.OpenId
}

func (s *GetUserBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserBasicInfoResponseBody) GetUnionIds() []*GetUserBasicInfoResponseBodyUnionIds {
	return s.UnionIds
}

func (s *GetUserBasicInfoResponseBody) SetAvatarUrl(v string) *GetUserBasicInfoResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetNickname(v string) *GetUserBasicInfoResponseBody {
	s.Nickname = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetOpenId(v string) *GetUserBasicInfoResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetRequestId(v string) *GetUserBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBasicInfoResponseBody) SetUnionIds(v []*GetUserBasicInfoResponseBodyUnionIds) *GetUserBasicInfoResponseBody {
	s.UnionIds = v
	return s
}

func (s *GetUserBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserBasicInfoResponseBodyUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUserBasicInfoResponseBodyUnionIds) String() string {
	return dara.Prettify(s)
}

func (s GetUserBasicInfoResponseBodyUnionIds) GoString() string {
	return s.String()
}

func (s *GetUserBasicInfoResponseBodyUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUserBasicInfoResponseBodyUnionIds) GetUnionId() *string {
	return s.UnionId
}

func (s *GetUserBasicInfoResponseBodyUnionIds) SetOrganizationId(v string) *GetUserBasicInfoResponseBodyUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *GetUserBasicInfoResponseBodyUnionIds) SetUnionId(v string) *GetUserBasicInfoResponseBodyUnionIds {
	s.UnionId = &v
	return s
}

func (s *GetUserBasicInfoResponseBodyUnionIds) Validate() error {
	return dara.Validate(s)
}
