// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDrdsDbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDrdsDbResponseBody
	GetSuccess() *bool
}

type RemoveDrdsDbResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B12FC174-D5CE-4A6E-83C1-0F8F86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDrdsDbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDrdsDbResponseBody) SetRequestId(v string) *RemoveDrdsDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbResponseBody) SetSuccess(v bool) *RemoveDrdsDbResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDrdsDbResponseBody) Validate() error {
	return dara.Validate(s)
}
