// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDcdnWafGroupResponseBody
	GetRequestId() *string
}

type ModifyDcdnWafGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 79B78B62-9006-5D6A-9DAB-303E134CD7AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDcdnWafGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDcdnWafGroupResponseBody) SetRequestId(v string) *ModifyDcdnWafGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDcdnWafGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
