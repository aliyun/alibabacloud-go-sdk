// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteModelResourceResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteModelResourceResponseBody
	GetRequestId() *string
}

type DeleteModelResourceResponseBody struct {
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

func (s DeleteModelResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteModelResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelResourceResponseBody) SetData(v bool) *DeleteModelResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteModelResourceResponseBody) SetRequestId(v string) *DeleteModelResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
