// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceLocationName(v string) *GetSourceRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *GetSourceRequest
	GetSourceName() *string
	SetSourceType(v string) *GetSourceRequest
	GetSourceType() *string
}

type GetSourceRequest struct {
	// The source location.
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

func (s GetSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceRequest) GoString() string {
	return s.String()
}

func (s *GetSourceRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *GetSourceRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *GetSourceRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetSourceRequest) SetSourceLocationName(v string) *GetSourceRequest {
	s.SourceLocationName = &v
	return s
}

func (s *GetSourceRequest) SetSourceName(v string) *GetSourceRequest {
	s.SourceName = &v
	return s
}

func (s *GetSourceRequest) SetSourceType(v string) *GetSourceRequest {
	s.SourceType = &v
	return s
}

func (s *GetSourceRequest) Validate() error {
	return dara.Validate(s)
}
