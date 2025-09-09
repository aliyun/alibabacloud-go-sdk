// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEventStatusResponseBody
	GetRequestId() *string
}

type ModifyEventStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8491DBFD-48C0-4E11-B6FC-6F38921244A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEventStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEventStatusResponseBody) SetRequestId(v string) *ModifyEventStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEventStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
