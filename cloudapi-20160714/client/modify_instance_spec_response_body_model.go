// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyInstanceSpecResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 36BBBAD4-1CFB-489F-841A-8CA52EEA787E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSpecResponseBody) SetRequestId(v string) *ModifyInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
