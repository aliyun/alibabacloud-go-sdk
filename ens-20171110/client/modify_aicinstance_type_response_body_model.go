// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAICInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAICInstanceTypeResponseBody
	GetRequestId() *string
}

type ModifyAICInstanceTypeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F590DD6C-B78B-571B-8FE9-B6ECEED9BE3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAICInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAICInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAICInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAICInstanceTypeResponseBody) SetRequestId(v string) *ModifyAICInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAICInstanceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
