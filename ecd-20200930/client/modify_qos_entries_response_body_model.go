// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQosEntriesResponseBody
	GetRequestId() *string
}

type ModifyQosEntriesResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQosEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQosEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQosEntriesResponseBody) SetRequestId(v string) *ModifyQosEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQosEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
