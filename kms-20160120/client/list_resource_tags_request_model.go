// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *ListResourceTagsRequest
	GetKeyId() *string
}

type ListResourceTagsRequest struct {
	// The globally unique ID of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s ListResourceTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTagsRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ListResourceTagsRequest) SetKeyId(v string) *ListResourceTagsRequest {
	s.KeyId = &v
	return s
}

func (s *ListResourceTagsRequest) Validate() error {
	return dara.Validate(s)
}
