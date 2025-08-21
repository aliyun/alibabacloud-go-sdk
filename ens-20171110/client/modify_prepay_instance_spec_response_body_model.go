// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyPrepayInstanceSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FD94C8E8-128E-525C-A0C3-60E063B70330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPrepayInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
