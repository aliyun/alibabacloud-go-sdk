// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceTypeAutoEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourceTypeAutoEnableResponseBody
	GetRequestId() *string
}

type ModifyResourceTypeAutoEnableResponseBody struct {
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceTypeAutoEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceTypeAutoEnableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceTypeAutoEnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceTypeAutoEnableResponseBody) SetRequestId(v string) *ModifyResourceTypeAutoEnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceTypeAutoEnableResponseBody) Validate() error {
	return dara.Validate(s)
}
