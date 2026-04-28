// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeEndpoint(v string) *CreateModelApiResponseBody
	GetInvokeEndpoint() *string
	SetModelApiId(v string) *CreateModelApiResponseBody
	GetModelApiId() *string
	SetRequestId(v string) *CreateModelApiResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateModelApiResponseBody
	GetStatus() *string
}

type CreateModelApiResponseBody struct {
	// example:
	//
	// xxx
	InvokeEndpoint *string `json:"InvokeEndpoint,omitempty" xml:"InvokeEndpoint,omitempty"`
	// example:
	//
	// mi-xxxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateModelApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelApiResponseBody) GetInvokeEndpoint() *string {
	return s.InvokeEndpoint
}

func (s *CreateModelApiResponseBody) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *CreateModelApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelApiResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateModelApiResponseBody) SetInvokeEndpoint(v string) *CreateModelApiResponseBody {
	s.InvokeEndpoint = &v
	return s
}

func (s *CreateModelApiResponseBody) SetModelApiId(v string) *CreateModelApiResponseBody {
	s.ModelApiId = &v
	return s
}

func (s *CreateModelApiResponseBody) SetRequestId(v string) *CreateModelApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelApiResponseBody) SetStatus(v string) *CreateModelApiResponseBody {
	s.Status = &v
	return s
}

func (s *CreateModelApiResponseBody) Validate() error {
	return dara.Validate(s)
}
