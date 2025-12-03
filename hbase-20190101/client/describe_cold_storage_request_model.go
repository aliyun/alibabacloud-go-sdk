// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeColdStorageRequest
	GetClusterId() *string
}

type DescribeColdStorageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeColdStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeColdStorageRequest) SetClusterId(v string) *DescribeColdStorageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageRequest) Validate() error {
	return dara.Validate(s)
}
