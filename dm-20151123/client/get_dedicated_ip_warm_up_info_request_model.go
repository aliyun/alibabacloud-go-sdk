// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedIp(v string) *GetDedicatedIpWarmUpInfoRequest
	GetDedicatedIp() *string
}

type GetDedicatedIpWarmUpInfoRequest struct {
	DedicatedIp *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
}

func (s GetDedicatedIpWarmUpInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpInfoRequest) GetDedicatedIp() *string {
	return s.DedicatedIp
}

func (s *GetDedicatedIpWarmUpInfoRequest) SetDedicatedIp(v string) *GetDedicatedIpWarmUpInfoRequest {
	s.DedicatedIp = &v
	return s
}

func (s *GetDedicatedIpWarmUpInfoRequest) Validate() error {
	return dara.Validate(s)
}
