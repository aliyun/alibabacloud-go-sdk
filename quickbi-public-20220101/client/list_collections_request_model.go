// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *ListCollectionsRequest
	GetUserId() *string
}

type ListCollectionsRequest struct {
	// The ID of the user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListCollectionsRequest) SetUserId(v string) *ListCollectionsRequest {
	s.UserId = &v
	return s
}

func (s *ListCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
