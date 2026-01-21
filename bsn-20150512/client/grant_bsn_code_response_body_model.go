// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantBsnCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GrantBsnCodeResponseBody
	GetAliUid() *int64
	SetGrantToAliuid(v int64) *GrantBsnCodeResponseBody
	GetGrantToAliuid() *int64
	SetNotes(v string) *GrantBsnCodeResponseBody
	GetNotes() *string
	SetRequestId(v string) *GrantBsnCodeResponseBody
	GetRequestId() *string
	SetSn(v string) *GrantBsnCodeResponseBody
	GetSn() *string
	SetSuccess(v bool) *GrantBsnCodeResponseBody
	GetSuccess() *bool
}

type GrantBsnCodeResponseBody struct {
	// example:
	//
	// 456*******163
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 124*******345
	GrantToAliuid *int64  `json:"GrantToAliuid,omitempty" xml:"GrantToAliuid,omitempty"`
	Notes         *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// example:
	//
	// FA30A2AF-9542-5927-9B73-2030F1A68DBF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1131-*******-233
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantBsnCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantBsnCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GrantBsnCodeResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GrantBsnCodeResponseBody) GetGrantToAliuid() *int64 {
	return s.GrantToAliuid
}

func (s *GrantBsnCodeResponseBody) GetNotes() *string {
	return s.Notes
}

func (s *GrantBsnCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantBsnCodeResponseBody) GetSn() *string {
	return s.Sn
}

func (s *GrantBsnCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantBsnCodeResponseBody) SetAliUid(v int64) *GrantBsnCodeResponseBody {
	s.AliUid = &v
	return s
}

func (s *GrantBsnCodeResponseBody) SetGrantToAliuid(v int64) *GrantBsnCodeResponseBody {
	s.GrantToAliuid = &v
	return s
}

func (s *GrantBsnCodeResponseBody) SetNotes(v string) *GrantBsnCodeResponseBody {
	s.Notes = &v
	return s
}

func (s *GrantBsnCodeResponseBody) SetRequestId(v string) *GrantBsnCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantBsnCodeResponseBody) SetSn(v string) *GrantBsnCodeResponseBody {
	s.Sn = &v
	return s
}

func (s *GrantBsnCodeResponseBody) SetSuccess(v bool) *GrantBsnCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GrantBsnCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
