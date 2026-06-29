// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokeSeatsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemsShrink(v string) *BatchRevokeSeatsShrinkRequest
	GetItemsShrink() *string
	SetLocale(v string) *BatchRevokeSeatsShrinkRequest
	GetLocale() *string
}

type BatchRevokeSeatsShrinkRequest struct {
	// The list of revocation items. This parameter is required.
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// The language. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s BatchRevokeSeatsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokeSeatsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchRevokeSeatsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *BatchRevokeSeatsShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *BatchRevokeSeatsShrinkRequest) SetItemsShrink(v string) *BatchRevokeSeatsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *BatchRevokeSeatsShrinkRequest) SetLocale(v string) *BatchRevokeSeatsShrinkRequest {
	s.Locale = &v
	return s
}

func (s *BatchRevokeSeatsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
