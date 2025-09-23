// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceShareId(v string) *DeleteResourceShareRequest
	GetResourceShareId() *string
}

type DeleteResourceShareRequest struct {
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-qSkW1HBY****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
}

func (s DeleteResourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceShareRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *DeleteResourceShareRequest) SetResourceShareId(v string) *DeleteResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *DeleteResourceShareRequest) Validate() error {
	return dara.Validate(s)
}
