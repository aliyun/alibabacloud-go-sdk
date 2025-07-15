// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullNatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFullNatEntryId(v string) *CreateFullNatEntryResponseBody
	GetFullNatEntryId() *string
	SetRequestId(v string) *CreateFullNatEntryResponseBody
	GetRequestId() *string
}

type CreateFullNatEntryResponseBody struct {
	// The FULLNAT entry ID.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatEntryId *string `json:"FullNatEntryId,omitempty" xml:"FullNatEntryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFullNatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFullNatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFullNatEntryResponseBody) GetFullNatEntryId() *string {
	return s.FullNatEntryId
}

func (s *CreateFullNatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFullNatEntryResponseBody) SetFullNatEntryId(v string) *CreateFullNatEntryResponseBody {
	s.FullNatEntryId = &v
	return s
}

func (s *CreateFullNatEntryResponseBody) SetRequestId(v string) *CreateFullNatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFullNatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
