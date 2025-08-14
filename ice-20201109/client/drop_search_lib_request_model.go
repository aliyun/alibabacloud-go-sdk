// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchLibName(v string) *DropSearchLibRequest
	GetSearchLibName() *string
}

type DropSearchLibRequest struct {
	// The name of the search library.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s DropSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DropSearchLibRequest) GoString() string {
	return s.String()
}

func (s *DropSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *DropSearchLibRequest) SetSearchLibName(v string) *DropSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *DropSearchLibRequest) Validate() error {
	return dara.Validate(s)
}
