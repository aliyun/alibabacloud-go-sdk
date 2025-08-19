// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAncestorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildId(v string) *ListAncestorsRequest
	GetChildId() *string
}

type ListAncestorsRequest struct {
	// The ID of the subfolder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-i1c9nr****
	ChildId *string `json:"ChildId,omitempty" xml:"ChildId,omitempty"`
}

func (s ListAncestorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAncestorsRequest) GoString() string {
	return s.String()
}

func (s *ListAncestorsRequest) GetChildId() *string {
	return s.ChildId
}

func (s *ListAncestorsRequest) SetChildId(v string) *ListAncestorsRequest {
	s.ChildId = &v
	return s
}

func (s *ListAncestorsRequest) Validate() error {
	return dara.Validate(s)
}
