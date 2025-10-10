// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyProductResponseBody
	GetRequestId() *string
}

type ModifyProductResponseBody struct {
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyProductResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyProductResponseBody) SetRequestId(v string) *ModifyProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyProductResponseBody) Validate() error {
	return dara.Validate(s)
}
