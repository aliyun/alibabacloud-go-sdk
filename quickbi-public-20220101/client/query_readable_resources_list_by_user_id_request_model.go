// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryReadableResourcesListByUserIdRequest
	GetUserId() *string
}

type QueryReadableResourcesListByUserIdRequest struct {
	// Quick BI the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryReadableResourcesListByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryReadableResourcesListByUserIdRequest) SetUserId(v string) *QueryReadableResourcesListByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
