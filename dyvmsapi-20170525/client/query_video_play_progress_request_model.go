// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoPlayProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *QueryVideoPlayProgressRequest
	GetCallId() *string
	SetCalledNumber(v string) *QueryVideoPlayProgressRequest
	GetCalledNumber() *string
	SetOwnerId(v int64) *QueryVideoPlayProgressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryVideoPlayProgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVideoPlayProgressRequest
	GetResourceOwnerId() *int64
}

type QueryVideoPlayProgressRequest struct {
	// example:
	//
	// 116004767703^102806****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 159****0000
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryVideoPlayProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoPlayProgressRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressRequest) GetCallId() *string {
	return s.CallId
}

func (s *QueryVideoPlayProgressRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryVideoPlayProgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVideoPlayProgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVideoPlayProgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVideoPlayProgressRequest) SetCallId(v string) *QueryVideoPlayProgressRequest {
	s.CallId = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetCalledNumber(v string) *QueryVideoPlayProgressRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetOwnerId(v int64) *QueryVideoPlayProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetResourceOwnerAccount(v string) *QueryVideoPlayProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetResourceOwnerId(v int64) *QueryVideoPlayProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) Validate() error {
	return dara.Validate(s)
}
