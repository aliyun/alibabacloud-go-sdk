// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelServiceId(v string) *CreateModelServiceResponseBody
	GetModelServiceId() *string
	SetRequestId(v string) *CreateModelServiceResponseBody
	GetRequestId() *string
}

type CreateModelServiceResponseBody struct {
	// The model service ID.
	//
	// example:
	//
	// ms-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelServiceResponseBody) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *CreateModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelServiceResponseBody) SetModelServiceId(v string) *CreateModelServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *CreateModelServiceResponseBody) SetRequestId(v string) *CreateModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
