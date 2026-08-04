// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnumConfigByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *QueryEnumConfigByTypeRequest
	GetType() *string
}

type QueryEnumConfigByTypeRequest struct {
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryEnumConfigByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEnumConfigByTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryEnumConfigByTypeRequest) GetType() *string {
	return s.Type
}

func (s *QueryEnumConfigByTypeRequest) SetType(v string) *QueryEnumConfigByTypeRequest {
	s.Type = &v
	return s
}

func (s *QueryEnumConfigByTypeRequest) Validate() error {
	return dara.Validate(s)
}
