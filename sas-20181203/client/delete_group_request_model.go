// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteGroupRequest
	GetGroupId() *int64
	SetSourceIp(v string) *DeleteGroupRequest
	GetSourceIp() *string
}

type DeleteGroupRequest struct {
	// The ID of the server group that you want to delete.
	//
	// >  To delete a server group, you must provide the ID of the server group. You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9454789
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.172.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteGroupRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteGroupRequest) SetGroupId(v int64) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupRequest) SetSourceIp(v string) *DeleteGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
