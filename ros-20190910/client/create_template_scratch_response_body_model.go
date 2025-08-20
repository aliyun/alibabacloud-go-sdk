// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateScratchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTemplateScratchResponseBody
	GetRequestId() *string
	SetTemplateScratchId(v string) *CreateTemplateScratchResponseBody
	GetTemplateScratchId() *string
}

type CreateTemplateScratchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 84980977-22F0-5421-B30D-B201311D5DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s CreateTemplateScratchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateScratchResponseBody) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *CreateTemplateScratchResponseBody) SetRequestId(v string) *CreateTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateScratchResponseBody) SetTemplateScratchId(v string) *CreateTemplateScratchResponseBody {
	s.TemplateScratchId = &v
	return s
}

func (s *CreateTemplateScratchResponseBody) Validate() error {
	return dara.Validate(s)
}
