// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupIdByGroupNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *QueryGroupIdByGroupNameRequest
	GetGroupName() *string
	SetSourceIp(v string) *QueryGroupIdByGroupNameRequest
	GetSourceIp() *string
}

type QueryGroupIdByGroupNameRequest struct {
	// The name of the asset group.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestGroupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 10.12.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s QueryGroupIdByGroupNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupIdByGroupNameRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupIdByGroupNameRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryGroupIdByGroupNameRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *QueryGroupIdByGroupNameRequest) SetGroupName(v string) *QueryGroupIdByGroupNameRequest {
	s.GroupName = &v
	return s
}

func (s *QueryGroupIdByGroupNameRequest) SetSourceIp(v string) *QueryGroupIdByGroupNameRequest {
	s.SourceIp = &v
	return s
}

func (s *QueryGroupIdByGroupNameRequest) Validate() error {
	return dara.Validate(s)
}
