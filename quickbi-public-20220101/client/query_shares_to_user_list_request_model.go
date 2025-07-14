// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySharesToUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QuerySharesToUserListRequest
	GetUserId() *string
}

type QuerySharesToUserListRequest struct {
	// The ID of the user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e53****5ba4b679ee22e2a2927****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySharesToUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySharesToUserListRequest) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySharesToUserListRequest) SetUserId(v string) *QuerySharesToUserListRequest {
	s.UserId = &v
	return s
}

func (s *QuerySharesToUserListRequest) Validate() error {
	return dara.Validate(s)
}
