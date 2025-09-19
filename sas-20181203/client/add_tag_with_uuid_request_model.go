// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagWithUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagName(v string) *AddTagWithUuidRequest
	GetTagName() *string
	SetUuidList(v string) *AddTagWithUuidRequest
	GetUuidList() *string
}

type AddTagWithUuidRequest struct {
	// The name of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// InternetIp
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The UUIDs of the servers. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 71f5313e-4355-4c59-86d1-557dda7b****,71f5313e-4355-4c59-86d1-557dda7b****
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s AddTagWithUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagWithUuidRequest) GoString() string {
	return s.String()
}

func (s *AddTagWithUuidRequest) GetTagName() *string {
	return s.TagName
}

func (s *AddTagWithUuidRequest) GetUuidList() *string {
	return s.UuidList
}

func (s *AddTagWithUuidRequest) SetTagName(v string) *AddTagWithUuidRequest {
	s.TagName = &v
	return s
}

func (s *AddTagWithUuidRequest) SetUuidList(v string) *AddTagWithUuidRequest {
	s.UuidList = &v
	return s
}

func (s *AddTagWithUuidRequest) Validate() error {
	return dara.Validate(s)
}
