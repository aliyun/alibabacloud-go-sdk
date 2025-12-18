// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourceInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateResourceInventoryResponseBody
	GetRequestId() *string
}

type GenerateResourceInventoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E1225EEA-B5F8-538F-8E37-A943986B6290
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateResourceInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourceInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateResourceInventoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateResourceInventoryResponseBody) SetRequestId(v string) *GenerateResourceInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateResourceInventoryResponseBody) Validate() error {
	return dara.Validate(s)
}
