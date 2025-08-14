// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSoftDelete(v bool) *DeleteSourceLocationRequest
	GetSoftDelete() *bool
	SetSourceLocationName(v string) *DeleteSourceLocationRequest
	GetSourceLocationName() *string
}

type DeleteSourceLocationRequest struct {
	// Specifies whether to use delete markers.
	//
	// example:
	//
	// true
	SoftDelete *bool `json:"SoftDelete,omitempty" xml:"SoftDelete,omitempty"`
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
}

func (s DeleteSourceLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceLocationRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceLocationRequest) GetSoftDelete() *bool {
	return s.SoftDelete
}

func (s *DeleteSourceLocationRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *DeleteSourceLocationRequest) SetSoftDelete(v bool) *DeleteSourceLocationRequest {
	s.SoftDelete = &v
	return s
}

func (s *DeleteSourceLocationRequest) SetSourceLocationName(v string) *DeleteSourceLocationRequest {
	s.SourceLocationName = &v
	return s
}

func (s *DeleteSourceLocationRequest) Validate() error {
	return dara.Validate(s)
}
