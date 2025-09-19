// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPartialBuyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OpenPartialBuyRequest
	GetInstanceId() *string
}

type OpenPartialBuyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OpenPartialBuyRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenPartialBuyRequest) GoString() string {
	return s.String()
}

func (s *OpenPartialBuyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenPartialBuyRequest) SetInstanceId(v string) *OpenPartialBuyRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenPartialBuyRequest) Validate() error {
	return dara.Validate(s)
}
