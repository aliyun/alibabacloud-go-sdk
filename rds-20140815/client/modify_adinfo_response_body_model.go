// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyADInfoResponseBody
	GetRequestId() *string
}

type ModifyADInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyADInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyADInfoResponseBody) SetRequestId(v string) *ModifyADInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyADInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
