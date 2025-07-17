// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteContactGroupResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteContactGroupResponseBody
	GetRequestId() *string
}

type DeleteContactGroupResponseBody struct {
	// Indicates whether the alert contact group was deleted. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactGroupResponseBody) SetIsSuccess(v bool) *DeleteContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetRequestId(v string) *DeleteContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
