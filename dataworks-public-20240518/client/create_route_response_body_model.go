// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateRouteResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRouteResponseBody
	GetSuccess() *bool
}

type CreateRouteResponseBody struct {
	// The route ID.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s CreateRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRouteResponseBody) SetId(v int64) *CreateRouteResponseBody {
	s.Id = &v
	return s
}

func (s *CreateRouteResponseBody) SetRequestId(v string) *CreateRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteResponseBody) SetSuccess(v bool) *CreateRouteResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
