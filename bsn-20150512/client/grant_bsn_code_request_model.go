// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantBsnCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GrantBsnCodeRequest
	GetAliUid() *int64
	SetGrantToAliuid(v int64) *GrantBsnCodeRequest
	GetGrantToAliuid() *int64
	SetNotes(v string) *GrantBsnCodeRequest
	GetNotes() *string
	SetSn(v string) *GrantBsnCodeRequest
	GetSn() *string
}

type GrantBsnCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 456*******163
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 124*******345
	GrantToAliuid *int64  `json:"GrantToAliuid,omitempty" xml:"GrantToAliuid,omitempty"`
	Notes         *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1131-*******-233
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GrantBsnCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantBsnCodeRequest) GoString() string {
	return s.String()
}

func (s *GrantBsnCodeRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GrantBsnCodeRequest) GetGrantToAliuid() *int64 {
	return s.GrantToAliuid
}

func (s *GrantBsnCodeRequest) GetNotes() *string {
	return s.Notes
}

func (s *GrantBsnCodeRequest) GetSn() *string {
	return s.Sn
}

func (s *GrantBsnCodeRequest) SetAliUid(v int64) *GrantBsnCodeRequest {
	s.AliUid = &v
	return s
}

func (s *GrantBsnCodeRequest) SetGrantToAliuid(v int64) *GrantBsnCodeRequest {
	s.GrantToAliuid = &v
	return s
}

func (s *GrantBsnCodeRequest) SetNotes(v string) *GrantBsnCodeRequest {
	s.Notes = &v
	return s
}

func (s *GrantBsnCodeRequest) SetSn(v string) *GrantBsnCodeRequest {
	s.Sn = &v
	return s
}

func (s *GrantBsnCodeRequest) Validate() error {
	return dara.Validate(s)
}
