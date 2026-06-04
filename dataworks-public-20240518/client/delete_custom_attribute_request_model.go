// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteCustomAttributeRequest
	GetId() *string
}

type DeleteCustomAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// custom-attribute:biz_owner
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCustomAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAttributeRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAttributeRequest) GetId() *string {
	return s.Id
}

func (s *DeleteCustomAttributeRequest) SetId(v string) *DeleteCustomAttributeRequest {
	s.Id = &v
	return s
}

func (s *DeleteCustomAttributeRequest) Validate() error {
	return dara.Validate(s)
}
