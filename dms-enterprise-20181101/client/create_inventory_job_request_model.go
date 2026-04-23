// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInventoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParam(v string) *CreateInventoryJobRequest
	GetParam() *string
}

type CreateInventoryJobRequest struct {
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s CreateInventoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInventoryJobRequest) GoString() string {
	return s.String()
}

func (s *CreateInventoryJobRequest) GetParam() *string {
	return s.Param
}

func (s *CreateInventoryJobRequest) SetParam(v string) *CreateInventoryJobRequest {
	s.Param = &v
	return s
}

func (s *CreateInventoryJobRequest) Validate() error {
	return dara.Validate(s)
}
