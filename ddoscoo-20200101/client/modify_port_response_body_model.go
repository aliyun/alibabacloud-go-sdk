// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPortResponseBody
	GetRequestId() *string
}

type ModifyPortResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPortResponseBody) SetRequestId(v string) *ModifyPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPortResponseBody) Validate() error {
	return dara.Validate(s)
}
