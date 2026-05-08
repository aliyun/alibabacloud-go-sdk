// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextThemesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustry(v string) *ListTextThemesRequest
	GetIndustry() *string
}

type ListTextThemesRequest struct {
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s ListTextThemesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextThemesRequest) GoString() string {
	return s.String()
}

func (s *ListTextThemesRequest) GetIndustry() *string {
	return s.Industry
}

func (s *ListTextThemesRequest) SetIndustry(v string) *ListTextThemesRequest {
	s.Industry = &v
	return s
}

func (s *ListTextThemesRequest) Validate() error {
	return dara.Validate(s)
}
