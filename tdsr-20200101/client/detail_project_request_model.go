// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DetailProjectRequest
	GetId() *string
}

type DetailProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailProjectRequest) GoString() string {
	return s.String()
}

func (s *DetailProjectRequest) GetId() *string {
	return s.Id
}

func (s *DetailProjectRequest) SetId(v string) *DetailProjectRequest {
	s.Id = &v
	return s
}

func (s *DetailProjectRequest) Validate() error {
	return dara.Validate(s)
}
