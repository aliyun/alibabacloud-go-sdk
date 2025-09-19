// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *ModifyWebLockRefreshRequest
	GetUuid() *string
}

type ModifyWebLockRefreshRequest struct {
	// The UUID of the server for which you want to refresh the status of the web tamper proofing feature.
	//
	// >  You can call the [DescribeWebLockBindList](~~DescribeWebLockBindList~~) operation to query the servers for which the web tamper proofing feature is enabled.
	//
	// example:
	//
	// 55c0f41b-3093-47a7-8eae-02d3a584****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockRefreshRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockRefreshRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockRefreshRequest) SetUuid(v string) *ModifyWebLockRefreshRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockRefreshRequest) Validate() error {
	return dara.Validate(s)
}
