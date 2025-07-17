// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateNetworkResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateNetworkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNetworkResponseBody
	GetSuccess() *bool
}

type CreateNetworkResponseBody struct {
	// The network ID.
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

func (s CreateNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNetworkResponseBody) SetId(v int64) *CreateNetworkResponseBody {
	s.Id = &v
	return s
}

func (s *CreateNetworkResponseBody) SetRequestId(v string) *CreateNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkResponseBody) SetSuccess(v bool) *CreateNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
