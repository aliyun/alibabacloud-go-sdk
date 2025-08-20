// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateScratchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateScratchResponseBody
	GetRequestId() *string
	SetTemplateScratchId(v string) *UpdateTemplateScratchResponseBody
	GetTemplateScratchId() *string
}

type UpdateTemplateScratchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 221DA822-B8CF-50DF-A9D2-BA197BF97BD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s UpdateTemplateScratchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateScratchResponseBody) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *UpdateTemplateScratchResponseBody) SetRequestId(v string) *UpdateTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateScratchResponseBody) SetTemplateScratchId(v string) *UpdateTemplateScratchResponseBody {
	s.TemplateScratchId = &v
	return s
}

func (s *UpdateTemplateScratchResponseBody) Validate() error {
	return dara.Validate(s)
}
