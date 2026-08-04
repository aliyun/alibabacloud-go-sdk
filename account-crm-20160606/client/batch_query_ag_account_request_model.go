// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *BatchQueryAgAccountRequest
	GetMpk() *string
	SetPkList(v string) *BatchQueryAgAccountRequest
	GetPkList() *string
}

type BatchQueryAgAccountRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	PkList *string `json:"PkList,omitempty" xml:"PkList,omitempty"`
}

func (s BatchQueryAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryAgAccountRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *BatchQueryAgAccountRequest) GetPkList() *string {
	return s.PkList
}

func (s *BatchQueryAgAccountRequest) SetMpk(v string) *BatchQueryAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *BatchQueryAgAccountRequest) SetPkList(v string) *BatchQueryAgAccountRequest {
	s.PkList = &v
	return s
}

func (s *BatchQueryAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
