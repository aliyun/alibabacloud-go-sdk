// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSoftDelete(v bool) *DeleteSourceRequest
	GetSoftDelete() *bool
	SetSourceLocationName(v string) *DeleteSourceRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *DeleteSourceRequest
	GetSourceName() *string
	SetSourceType(v string) *DeleteSourceRequest
	GetSourceType() *string
}

type DeleteSourceRequest struct {
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
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyVodSource
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The source type. Valid values: vodSource and liveSource.
	//
	// This parameter is required.
	//
	// example:
	//
	// vodSource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DeleteSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceRequest) GetSoftDelete() *bool {
	return s.SoftDelete
}

func (s *DeleteSourceRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *DeleteSourceRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *DeleteSourceRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteSourceRequest) SetSoftDelete(v bool) *DeleteSourceRequest {
	s.SoftDelete = &v
	return s
}

func (s *DeleteSourceRequest) SetSourceLocationName(v string) *DeleteSourceRequest {
	s.SourceLocationName = &v
	return s
}

func (s *DeleteSourceRequest) SetSourceName(v string) *DeleteSourceRequest {
	s.SourceName = &v
	return s
}

func (s *DeleteSourceRequest) SetSourceType(v string) *DeleteSourceRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteSourceRequest) Validate() error {
	return dara.Validate(s)
}
