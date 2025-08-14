// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchLibName(v string) *QuerySearchLibRequest
	GetSearchLibName() *string
}

type QuerySearchLibRequest struct {
	// The name of the search library.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s QuerySearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchLibRequest) GoString() string {
	return s.String()
}

func (s *QuerySearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *QuerySearchLibRequest) SetSearchLibName(v string) *QuerySearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *QuerySearchLibRequest) Validate() error {
	return dara.Validate(s)
}
