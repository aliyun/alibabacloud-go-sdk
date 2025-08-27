// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCitySearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *CitySearchRequest
	GetKeyword() *string
}

type CitySearchRequest struct {
	// This parameter is required.
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s CitySearchRequest) String() string {
	return dara.Prettify(s)
}

func (s CitySearchRequest) GoString() string {
	return s.String()
}

func (s *CitySearchRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *CitySearchRequest) SetKeyword(v string) *CitySearchRequest {
	s.Keyword = &v
	return s
}

func (s *CitySearchRequest) Validate() error {
	return dara.Validate(s)
}
