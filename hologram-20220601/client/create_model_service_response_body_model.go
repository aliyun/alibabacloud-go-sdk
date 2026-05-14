// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateModelServiceResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateModelServiceResponseBody
	GetRequestId() *string
}

type CreateModelServiceResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelServiceResponseBody) SetData(v bool) *CreateModelServiceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateModelServiceResponseBody) SetRequestId(v string) *CreateModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
