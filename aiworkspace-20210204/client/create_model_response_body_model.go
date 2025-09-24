// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *CreateModelResponseBody
	GetModelId() *string
	SetRequestId(v string) *CreateModelResponseBody
	GetRequestId() *string
}

type CreateModelResponseBody struct {
	// The model ID.
	//
	// example:
	//
	// model-rbvg5wzljz****ks92
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DAD3112-AE22-5563-9A02-5C7E8****E35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) GetModelId() *string {
	return s.ModelId
}

func (s *CreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelResponseBody) SetModelId(v string) *CreateModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
