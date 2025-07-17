// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateResourceGroupResponseBody
	GetSuccess() *bool
}

type UpdateResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetSuccess(v bool) *UpdateResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
