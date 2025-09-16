// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewMode(v bool) *ListTablesRequest
	GetNewMode() *bool
}

type ListTablesRequest struct {
	// Specifies whether the OpenSearch Vector Search Edition instance is of the new version.
	//
	// example:
	//
	// true
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetNewMode() *bool {
	return s.NewMode
}

func (s *ListTablesRequest) SetNewMode(v bool) *ListTablesRequest {
	s.NewMode = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
