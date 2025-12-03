// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *CreateAccessGroupResponseBody
	GetAccessGroupId() *string
	SetRequestId(v string) *CreateAccessGroupResponseBody
	GetRequestId() *string
}

type CreateAccessGroupResponseBody struct {
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponseBody) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *CreateAccessGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessGroupResponseBody) SetAccessGroupId(v string) *CreateAccessGroupResponseBody {
	s.AccessGroupId = &v
	return s
}

func (s *CreateAccessGroupResponseBody) SetRequestId(v string) *CreateAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
