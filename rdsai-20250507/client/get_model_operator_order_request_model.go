// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOperatorOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetModelOperatorOrderRequest
	GetRegion() *string
}

type GetModelOperatorOrderRequest struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetModelOperatorOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderRequest) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderRequest) GetRegion() *string {
	return s.Region
}

func (s *GetModelOperatorOrderRequest) SetRegion(v string) *GetModelOperatorOrderRequest {
	s.Region = &v
	return s
}

func (s *GetModelOperatorOrderRequest) Validate() error {
	return dara.Validate(s)
}
