// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDrdsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDrdsInstanceResponseBody
	GetSuccess() *bool
}

type RemoveDrdsInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDrdsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDrdsInstanceResponseBody) SetRequestId(v string) *RemoveDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsInstanceResponseBody) SetSuccess(v bool) *RemoveDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDrdsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
