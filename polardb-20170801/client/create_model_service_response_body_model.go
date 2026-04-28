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
	SetStatus(v string) *CreateModelServiceResponseBody
	GetStatus() *string
}

type CreateModelServiceResponseBody struct {
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *CreateModelServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateModelServiceResponseBody) SetModelServiceId(v string) *CreateModelServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *CreateModelServiceResponseBody) SetRequestId(v string) *CreateModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelServiceResponseBody) SetStatus(v string) *CreateModelServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
