// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v int64) *ReadMessageRequest
	GetMessageId() *int64
	SetUserInfo(v *ReadMessageRequestUserInfo) *ReadMessageRequest
	GetUserInfo() *ReadMessageRequestUserInfo
}

type ReadMessageRequest struct {
	// example:
	//
	// 12345
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// This parameter is required.
	UserInfo *ReadMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ReadMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) GetMessageId() *int64 {
	return s.MessageId
}

func (s *ReadMessageRequest) GetUserInfo() *ReadMessageRequestUserInfo {
	return s.UserInfo
}

func (s *ReadMessageRequest) SetMessageId(v int64) *ReadMessageRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageRequest) SetUserInfo(v *ReadMessageRequestUserInfo) *ReadMessageRequest {
	s.UserInfo = v
	return s
}

func (s *ReadMessageRequest) Validate() error {
	return dara.Validate(s)
}

type ReadMessageRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123L
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ReadMessageRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ReadMessageRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ReadMessageRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ReadMessageRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ReadMessageRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ReadMessageRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ReadMessageRequestUserInfo) SetEncodeKey(v string) *ReadMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetEncodeType(v string) *ReadMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetId(v string) *ReadMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetIdType(v string) *ReadMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ReadMessageRequestUserInfo) SetOrganizationId(v string) *ReadMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ReadMessageRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
