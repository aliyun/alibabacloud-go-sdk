// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReservedNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteReservedNodePoolRequest
	GetClusterId() *string
	SetId(v string) *DeleteReservedNodePoolRequest
	GetId() *string
}

type DeleteReservedNodePoolRequest struct {
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// rnp-cdx****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteReservedNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReservedNodePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteReservedNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteReservedNodePoolRequest) GetId() *string {
	return s.Id
}

func (s *DeleteReservedNodePoolRequest) SetClusterId(v string) *DeleteReservedNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteReservedNodePoolRequest) SetId(v string) *DeleteReservedNodePoolRequest {
	s.Id = &v
	return s
}

func (s *DeleteReservedNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
