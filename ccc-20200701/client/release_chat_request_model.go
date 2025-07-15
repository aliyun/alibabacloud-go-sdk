// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseChatRequest
	GetInstanceId() *string
	SetJobId(v string) *ReleaseChatRequest
	GetJobId() *string
	SetToken(v string) *ReleaseChatRequest
	GetToken() *string
	SetUserId(v string) *ReleaseChatRequest
	GetUserId() *string
	SetUserType(v string) *ReleaseChatRequest
	GetUserType() *string
}

type ReleaseChatRequest struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chat-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 9XYGTGWtq2wFi_Bpg7aUnIoYi_vG_rO3bjEn0YtsxbHRHrYHlz1LDBLJAyZcLxieRQR4h_6AnWvTjJeNU5jg************Hwej7WgWrmA=
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// edbcf95a-ef9f-4296-a0a6-985ac9e36db3
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ReleaseChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseChatRequest) GoString() string {
	return s.String()
}

func (s *ReleaseChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseChatRequest) GetJobId() *string {
	return s.JobId
}

func (s *ReleaseChatRequest) GetToken() *string {
	return s.Token
}

func (s *ReleaseChatRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReleaseChatRequest) GetUserType() *string {
	return s.UserType
}

func (s *ReleaseChatRequest) SetInstanceId(v string) *ReleaseChatRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseChatRequest) SetJobId(v string) *ReleaseChatRequest {
	s.JobId = &v
	return s
}

func (s *ReleaseChatRequest) SetToken(v string) *ReleaseChatRequest {
	s.Token = &v
	return s
}

func (s *ReleaseChatRequest) SetUserId(v string) *ReleaseChatRequest {
	s.UserId = &v
	return s
}

func (s *ReleaseChatRequest) SetUserType(v string) *ReleaseChatRequest {
	s.UserType = &v
	return s
}

func (s *ReleaseChatRequest) Validate() error {
	return dara.Validate(s)
}
