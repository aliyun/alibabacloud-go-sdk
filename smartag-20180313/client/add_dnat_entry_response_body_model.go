// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnatEntryId(v string) *AddDnatEntryResponseBody
	GetDnatEntryId() *string
	SetRequestId(v string) *AddDnatEntryResponseBody
	GetRequestId() *string
}

type AddDnatEntryResponseBody struct {
	// The ID of the DNAT entry.
	//
	// example:
	//
	// fwd-kxe4fq3xuzczze****
	DnatEntryId *string `json:"DnatEntryId,omitempty" xml:"DnatEntryId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56BF6C79-C77D-41A0-86DD-A4B156E784EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnatEntryResponseBody) GetDnatEntryId() *string {
	return s.DnatEntryId
}

func (s *AddDnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnatEntryResponseBody) SetDnatEntryId(v string) *AddDnatEntryResponseBody {
	s.DnatEntryId = &v
	return s
}

func (s *AddDnatEntryResponseBody) SetRequestId(v string) *AddDnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
