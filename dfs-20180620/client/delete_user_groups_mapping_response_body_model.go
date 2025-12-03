// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupsMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserGroupsMappingResponseBody
	GetRequestId() *string
}

type DeleteUserGroupsMappingResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserGroupsMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupsMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupsMappingResponseBody) SetRequestId(v string) *DeleteUserGroupsMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupsMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
