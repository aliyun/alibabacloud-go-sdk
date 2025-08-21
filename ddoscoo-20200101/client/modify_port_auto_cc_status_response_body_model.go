// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortAutoCcStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPortAutoCcStatusResponseBody
	GetRequestId() *string
}

type ModifyPortAutoCcStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPortAutoCcStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortAutoCcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPortAutoCcStatusResponseBody) SetRequestId(v string) *ModifyPortAutoCcStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPortAutoCcStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
