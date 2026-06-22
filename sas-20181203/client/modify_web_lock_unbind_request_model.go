// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUnbindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *ModifyWebLockUnbindRequest
	GetUuid() *string
}

type ModifyWebLockUnbindRequest struct {
	// The UUID of the server from which you want to remove the protection directories.
	//
	// > Call [DescribeWebLockBindList](~~DescribeWebLockBindList~~) to obtain the UUID of the server.
	//
	// example:
	//
	// 8d217d3f-6999-43a6-a435-c7a6854180e9
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockUnbindRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUnbindRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockUnbindRequest) SetUuid(v string) *ModifyWebLockUnbindRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockUnbindRequest) Validate() error {
	return dara.Validate(s)
}
