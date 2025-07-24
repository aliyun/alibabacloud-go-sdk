// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCredits interface {
	dara.Model
	String() string
	GoString() string
	SetGenericTextSearch(v int32) *SearchCredits
	GetGenericTextSearch() *int32
}

type SearchCredits struct {
	GenericTextSearch *int32 `json:"genericTextSearch,omitempty" xml:"genericTextSearch,omitempty"`
}

func (s SearchCredits) String() string {
	return dara.Prettify(s)
}

func (s SearchCredits) GoString() string {
	return s.String()
}

func (s *SearchCredits) GetGenericTextSearch() *int32 {
	return s.GenericTextSearch
}

func (s *SearchCredits) SetGenericTextSearch(v int32) *SearchCredits {
	s.GenericTextSearch = &v
	return s
}

func (s *SearchCredits) Validate() error {
	return dara.Validate(s)
}
