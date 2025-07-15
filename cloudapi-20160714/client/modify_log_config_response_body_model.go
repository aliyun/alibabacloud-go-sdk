// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLogConfigResponseBody
	GetRequestId() *string
}

type ModifyLogConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 75A4ADCB-AA26-51FB-94D4-AB3240040974
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLogConfigResponseBody) SetRequestId(v string) *ModifyLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
