// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceBelongIdentifiersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifierIds(v []*string) *AddDeviceBelongIdentifiersResponseBody
	GetIdentifierIds() []*string
	SetRequestId(v string) *AddDeviceBelongIdentifiersResponseBody
	GetRequestId() *string
}

type AddDeviceBelongIdentifiersResponseBody struct {
	IdentifierIds []*string `json:"IdentifierIds,omitempty" xml:"IdentifierIds,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceBelongIdentifiersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceBelongIdentifiersResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceBelongIdentifiersResponseBody) GetIdentifierIds() []*string {
	return s.IdentifierIds
}

func (s *AddDeviceBelongIdentifiersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDeviceBelongIdentifiersResponseBody) SetIdentifierIds(v []*string) *AddDeviceBelongIdentifiersResponseBody {
	s.IdentifierIds = v
	return s
}

func (s *AddDeviceBelongIdentifiersResponseBody) SetRequestId(v string) *AddDeviceBelongIdentifiersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceBelongIdentifiersResponseBody) Validate() error {
	return dara.Validate(s)
}
