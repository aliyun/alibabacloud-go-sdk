// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNetworkResponseBody
	GetSuccess() *bool
}

type DeleteNetworkResponseBody struct {
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

func (s DeleteNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNetworkResponseBody) SetRequestId(v string) *DeleteNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkResponseBody) SetSuccess(v bool) *DeleteNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
