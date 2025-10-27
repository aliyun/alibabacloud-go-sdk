// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourcePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBResourcePoolResponseBody
	GetRequestId() *string
}

type ModifyDBResourcePoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourcePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBResourcePoolResponseBody) SetRequestId(v string) *ModifyDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBResourcePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
