// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *ListAccessKeysRequest
	GetUserName() *string
}

type ListAccessKeysRequest struct {
	// The name of the RAM user. If a RAM user calls this operation and does not specify this parameter, the AccessKey pairs of the RAM user are returned.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAccessKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListAccessKeysRequest) SetUserName(v string) *ListAccessKeysRequest {
	s.UserName = &v
	return s
}

func (s *ListAccessKeysRequest) Validate() error {
	return dara.Validate(s)
}
