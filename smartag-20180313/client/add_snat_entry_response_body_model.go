// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddSnatEntryResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *AddSnatEntryResponseBody
	GetRequestId() *string
}

type AddSnatEntryResponseBody struct {
	// The ID of the SNAT instance.
	//
	// example:
	//
	// snat-m2obgkt5ya1puz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56BF6C79-C77D-41A0-86DD-A4B156E784EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddSnatEntryResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSnatEntryResponseBody) SetInstanceId(v string) *AddSnatEntryResponseBody {
	s.InstanceId = &v
	return s
}

func (s *AddSnatEntryResponseBody) SetRequestId(v string) *AddSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
