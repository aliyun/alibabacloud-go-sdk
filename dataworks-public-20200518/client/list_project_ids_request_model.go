// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListProjectIdsRequest
	GetUserId() *string
}

type ListProjectIdsRequest struct {
	// The ID of the desired Alibaba Cloud account.
	//
	// You can log on to the [DataWorks](https://workbench.data.aliyun.com/console) console and move the pointer over the profile picture in the upper-right corner to view the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 171111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListProjectIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectIdsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectIdsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListProjectIdsRequest) SetUserId(v string) *ListProjectIdsRequest {
	s.UserId = &v
	return s
}

func (s *ListProjectIdsRequest) Validate() error {
	return dara.Validate(s)
}
