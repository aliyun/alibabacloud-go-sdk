// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteCustomLineRequest
	GetLang() *string
	SetLineId(v string) *DeleteCustomLineRequest
	GetLineId() *string
}

type DeleteCustomLineRequest struct {
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique ID of the custom line.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
}

func (s DeleteCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteCustomLineRequest) GetLineId() *string {
	return s.LineId
}

func (s *DeleteCustomLineRequest) SetLang(v string) *DeleteCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCustomLineRequest) SetLineId(v string) *DeleteCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *DeleteCustomLineRequest) Validate() error {
	return dara.Validate(s)
}
