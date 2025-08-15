// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchCode(v string) *DescribeScenesRequest
	GetSearchCode() *string
}

type DescribeScenesRequest struct {
	// example:
	//
	// ak
	SearchCode *string `json:"SearchCode,omitempty" xml:"SearchCode,omitempty"`
}

func (s DescribeScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScenesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScenesRequest) GetSearchCode() *string {
	return s.SearchCode
}

func (s *DescribeScenesRequest) SetSearchCode(v string) *DescribeScenesRequest {
	s.SearchCode = &v
	return s
}

func (s *DescribeScenesRequest) Validate() error {
	return dara.Validate(s)
}
