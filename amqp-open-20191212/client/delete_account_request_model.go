// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestamp(v int64) *DeleteAccountRequest
	GetCreateTimestamp() *int64
	SetUserName(v string) *DeleteAccountRequest
	GetUserName() *string
}

type DeleteAccountRequest struct {
	// The timestamp that indicates when the pair of static username and password that you want to delete was created. Unit: milliseconds.
	//
	// You can call the [ListAccounts](https://help.aliyun.com/document_detail/472730.html) operation to view the timestamp.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The pair of username and password that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ5cjc3djAwMzpMVEFJNXQ4YmVNbVZNMWVSWnRFSjZ2Zm1=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DeleteAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteAccountRequest) SetCreateTimestamp(v int64) *DeleteAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *DeleteAccountRequest) SetUserName(v string) *DeleteAccountRequest {
	s.UserName = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
