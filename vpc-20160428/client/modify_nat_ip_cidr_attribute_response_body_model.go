// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpCidrAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatIpCidrAttributeResponseBody
	GetRequestId() *string
}

type ModifyNatIpCidrAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatIpCidrAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpCidrAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatIpCidrAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatIpCidrAttributeResponseBody) SetRequestId(v string) *ModifyNatIpCidrAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatIpCidrAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
