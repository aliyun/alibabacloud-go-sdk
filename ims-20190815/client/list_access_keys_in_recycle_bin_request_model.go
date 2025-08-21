// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListAccessKeysInRecycleBinRequest
	GetUserId() *string
}

type ListAccessKeysInRecycleBinRequest struct {
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAccessKeysInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListAccessKeysInRecycleBinRequest) SetUserId(v string) *ListAccessKeysInRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *ListAccessKeysInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
