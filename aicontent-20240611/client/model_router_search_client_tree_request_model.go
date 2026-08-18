// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSearchClientTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelRouterSearchClientTreeRequest
	GetKeyword() *string
}

type ModelRouterSearchClientTreeRequest struct {
	// The search keyword, which can be a department name, username, or phone number.
	//
	// example:
	//
	// R&D Department
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s ModelRouterSearchClientTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSearchClientTreeRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterSearchClientTreeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterSearchClientTreeRequest) SetKeyword(v string) *ModelRouterSearchClientTreeRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterSearchClientTreeRequest) Validate() error {
	return dara.Validate(s)
}
