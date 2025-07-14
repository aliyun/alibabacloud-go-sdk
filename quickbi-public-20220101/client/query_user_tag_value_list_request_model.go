// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserTagValueListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryUserTagValueListRequest
	GetUserId() *string
}

type QueryUserTagValueListRequest struct {
	// This UserID refers to the Quick BI UserID, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserTagValueListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagValueListRequest) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserTagValueListRequest) SetUserId(v string) *QueryUserTagValueListRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserTagValueListRequest) Validate() error {
	return dara.Validate(s)
}
