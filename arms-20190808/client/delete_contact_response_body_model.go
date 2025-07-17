// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteContactResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteContactResponseBody
	GetRequestId() *string
}

type DeleteContactResponseBody struct {
	// Indicates whether the alert contact is deleted. Valid values:
	//
	// 	- `true`: The alert contact is deleted.
	//
	// 	- `false`: The alert contact is not deleted.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactResponseBody) SetIsSuccess(v bool) *DeleteContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactResponseBody) Validate() error {
	return dara.Validate(s)
}
