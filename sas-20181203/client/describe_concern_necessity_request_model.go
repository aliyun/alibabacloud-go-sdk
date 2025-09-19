// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConcernNecessityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeConcernNecessityRequest
	GetLang() *string
}

type DescribeConcernNecessityRequest struct {
	// The language of the content within the request and response. Valid values:
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

func (s DescribeConcernNecessityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConcernNecessityRequest) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeConcernNecessityRequest) SetLang(v string) *DescribeConcernNecessityRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConcernNecessityRequest) Validate() error {
	return dara.Validate(s)
}
