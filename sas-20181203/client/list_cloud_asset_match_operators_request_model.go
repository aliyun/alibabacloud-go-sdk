// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetMatchOperatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListCloudAssetMatchOperatorsRequest
	GetLang() *string
}

type ListCloudAssetMatchOperatorsRequest struct {
	// Sets the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListCloudAssetMatchOperatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetMatchOperatorsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAssetMatchOperatorsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCloudAssetMatchOperatorsRequest) SetLang(v string) *ListCloudAssetMatchOperatorsRequest {
	s.Lang = &v
	return s
}

func (s *ListCloudAssetMatchOperatorsRequest) Validate() error {
	return dara.Validate(s)
}
