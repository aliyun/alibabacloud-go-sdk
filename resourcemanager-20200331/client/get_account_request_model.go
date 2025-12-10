// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetAccountRequest
	GetAccountId() *string
	SetIncludeTags(v bool) *GetAccountRequest
	GetIncludeTags() *bool
}

type GetAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 181761095690****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
}

func (s GetAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *GetAccountRequest) SetAccountId(v string) *GetAccountRequest {
	s.AccountId = &v
	return s
}

func (s *GetAccountRequest) SetIncludeTags(v bool) *GetAccountRequest {
	s.IncludeTags = &v
	return s
}

func (s *GetAccountRequest) Validate() error {
	return dara.Validate(s)
}
