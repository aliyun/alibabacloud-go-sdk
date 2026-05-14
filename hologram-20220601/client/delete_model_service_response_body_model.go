// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteModelServiceResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteModelServiceResponseBody
	GetRequestId() *string
}

type DeleteModelServiceResponseBody struct {
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

func (s DeleteModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelServiceResponseBody) SetData(v bool) *DeleteModelServiceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteModelServiceResponseBody) SetRequestId(v string) *DeleteModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
