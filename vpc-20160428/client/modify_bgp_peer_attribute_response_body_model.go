// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpPeerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBgpPeerAttributeResponseBody
	GetRequestId() *string
}

type ModifyBgpPeerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D4B7649A-61BB-4C64-A586-1DFF1EDA6A42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBgpPeerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpPeerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBgpPeerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBgpPeerAttributeResponseBody) SetRequestId(v string) *ModifyBgpPeerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBgpPeerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
