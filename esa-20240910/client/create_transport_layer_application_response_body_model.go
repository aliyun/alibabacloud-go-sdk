// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransportLayerApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *CreateTransportLayerApplicationResponseBody
	GetApplicationId() *int64
	SetRequestId(v string) *CreateTransportLayerApplicationResponseBody
	GetRequestId() *string
}

type CreateTransportLayerApplicationResponseBody struct {
	// Layer 4 application ID
	//
	// example:
	//
	// 165503967****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 9e5448c7-edaf-49aa-9887-0fcd0832306c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransportLayerApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransportLayerApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransportLayerApplicationResponseBody) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *CreateTransportLayerApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransportLayerApplicationResponseBody) SetApplicationId(v int64) *CreateTransportLayerApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateTransportLayerApplicationResponseBody) SetRequestId(v string) *CreateTransportLayerApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransportLayerApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
