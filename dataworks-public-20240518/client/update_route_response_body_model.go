// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRouteResponseBody
	GetSuccess() *bool
}

type UpdateRouteResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
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

func (s UpdateRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRouteResponseBody) SetRequestId(v string) *UpdateRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteResponseBody) SetSuccess(v bool) *UpdateRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
