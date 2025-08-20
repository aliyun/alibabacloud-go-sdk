// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachUserENIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DetachUserENIRequest
	GetDBClusterId() *string
}

type DetachUserENIRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DetachUserENIRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachUserENIRequest) GoString() string {
	return s.String()
}

func (s *DetachUserENIRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DetachUserENIRequest) SetDBClusterId(v string) *DetachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *DetachUserENIRequest) Validate() error {
	return dara.Validate(s)
}
