// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunId(v string) *GetUserInfoResponseBody
	GetAliyunId() *string
	SetNickName(v string) *GetUserInfoResponseBody
	GetNickName() *string
	SetPhone(v string) *GetUserInfoResponseBody
	GetPhone() *string
	SetRequestId(v string) *GetUserInfoResponseBody
	GetRequestId() *string
	SetUnionId(v string) *GetUserInfoResponseBody
	GetUnionId() *string
}

type GetUserInfoResponseBody struct {
	AliyunId  *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	NickName  *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnionId   *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) GetAliyunId() *string {
	return s.AliyunId
}

func (s *GetUserInfoResponseBody) GetNickName() *string {
	return s.NickName
}

func (s *GetUserInfoResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserInfoResponseBody) GetUnionId() *string {
	return s.UnionId
}

func (s *GetUserInfoResponseBody) SetAliyunId(v string) *GetUserInfoResponseBody {
	s.AliyunId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetNickName(v string) *GetUserInfoResponseBody {
	s.NickName = &v
	return s
}

func (s *GetUserInfoResponseBody) SetPhone(v string) *GetUserInfoResponseBody {
	s.Phone = &v
	return s
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetUnionId(v string) *GetUserInfoResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
