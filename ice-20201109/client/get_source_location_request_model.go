// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceLocationName(v string) *GetSourceLocationRequest
	GetSourceLocationName() *string
}

type GetSourceLocationRequest struct {
	// The name of the source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
}

func (s GetSourceLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceLocationRequest) GoString() string {
	return s.String()
}

func (s *GetSourceLocationRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *GetSourceLocationRequest) SetSourceLocationName(v string) *GetSourceLocationRequest {
	s.SourceLocationName = &v
	return s
}

func (s *GetSourceLocationRequest) Validate() error {
	return dara.Validate(s)
}
