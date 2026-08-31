// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSecurityOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceSecurityOptionsResponseBody
	GetRequestId() *string
}

type ModifyInstanceSecurityOptionsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceSecurityOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSecurityOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSecurityOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSecurityOptionsResponseBody) SetRequestId(v string) *ModifyInstanceSecurityOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSecurityOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
