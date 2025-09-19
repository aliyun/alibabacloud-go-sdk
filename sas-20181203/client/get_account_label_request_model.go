// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelList(v []*string) *GetAccountLabelRequest
	GetLabelList() []*string
	SetLang(v string) *GetAccountLabelRequest
	GetLang() *string
}

type GetAccountLabelRequest struct {
	// The tags.
	//
	// This parameter is required.
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetAccountLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountLabelRequest) GoString() string {
	return s.String()
}

func (s *GetAccountLabelRequest) GetLabelList() []*string {
	return s.LabelList
}

func (s *GetAccountLabelRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAccountLabelRequest) SetLabelList(v []*string) *GetAccountLabelRequest {
	s.LabelList = v
	return s
}

func (s *GetAccountLabelRequest) SetLang(v string) *GetAccountLabelRequest {
	s.Lang = &v
	return s
}

func (s *GetAccountLabelRequest) Validate() error {
	return dara.Validate(s)
}
