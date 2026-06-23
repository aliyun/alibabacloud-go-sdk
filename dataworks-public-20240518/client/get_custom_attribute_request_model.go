// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetCustomAttributeRequest
	GetId() *string
}

type GetCustomAttributeRequest struct {
	// The custom attribute ID. The ID must match the regular expression `^custom-attribute:[A-Za-z][A-Za-z0-9_]{0,98}$`. The part after \\"custom-attribute:\\" must be less than 100 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom-attribute:biz_owner
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCustomAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetCustomAttributeRequest) GetId() *string {
	return s.Id
}

func (s *GetCustomAttributeRequest) SetId(v string) *GetCustomAttributeRequest {
	s.Id = &v
	return s
}

func (s *GetCustomAttributeRequest) Validate() error {
	return dara.Validate(s)
}
