// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullNatEntryAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFullNatEntryAttributeResponseBody
	GetRequestId() *string
}

type ModifyFullNatEntryAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFullNatEntryAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullNatEntryAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFullNatEntryAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFullNatEntryAttributeResponseBody) SetRequestId(v string) *ModifyFullNatEntryAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFullNatEntryAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
