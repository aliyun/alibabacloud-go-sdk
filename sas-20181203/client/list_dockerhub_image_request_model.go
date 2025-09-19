// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDockerhubImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *ListDockerhubImageRequest
	GetQuery() *string
}

type ListDockerhubImageRequest struct {
	// The query condition for images. You can query images in the `[namespace/]repoName[:version]` format. Conditions in `[]` are optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// python:3.9
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListDockerhubImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDockerhubImageRequest) GoString() string {
	return s.String()
}

func (s *ListDockerhubImageRequest) GetQuery() *string {
	return s.Query
}

func (s *ListDockerhubImageRequest) SetQuery(v string) *ListDockerhubImageRequest {
	s.Query = &v
	return s
}

func (s *ListDockerhubImageRequest) Validate() error {
	return dara.Validate(s)
}
