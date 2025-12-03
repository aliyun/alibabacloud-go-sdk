// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceNameResponseBody
	GetRequestId() *string
}

type ModifyInstanceNameResponseBody struct {
	// example:
	//
	// 959DA199-54E5-569D-AD46-92BED8515E62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
