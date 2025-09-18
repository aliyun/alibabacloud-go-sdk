// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsInternal(v bool) *GetTableTokenRequest
	GetIsInternal() *bool
}

type GetTableTokenRequest struct {
	// example:
	//
	// true
	IsInternal *bool `json:"isInternal,omitempty" xml:"isInternal,omitempty"`
}

func (s GetTableTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTableTokenRequest) GetIsInternal() *bool {
	return s.IsInternal
}

func (s *GetTableTokenRequest) SetIsInternal(v bool) *GetTableTokenRequest {
	s.IsInternal = &v
	return s
}

func (s *GetTableTokenRequest) Validate() error {
	return dara.Validate(s)
}
