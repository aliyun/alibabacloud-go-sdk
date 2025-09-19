// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRefreshProcessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *ModifyRefreshProcessInfoRequest
	GetUuid() *string
}

type ModifyRefreshProcessInfoRequest struct {
	// The UUID of the server.
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 0f3b8f76-90e5-4455-a5aa-23ce30b5****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyRefreshProcessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRefreshProcessInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyRefreshProcessInfoRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyRefreshProcessInfoRequest) SetUuid(v string) *ModifyRefreshProcessInfoRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyRefreshProcessInfoRequest) Validate() error {
	return dara.Validate(s)
}
