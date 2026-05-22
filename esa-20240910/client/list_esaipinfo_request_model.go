// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListESAIPInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVipInfo(v string) *ListESAIPInfoRequest
	GetVipInfo() *string
}

type ListESAIPInfoRequest struct {
	// This parameter is required.
	VipInfo *string `json:"VipInfo,omitempty" xml:"VipInfo,omitempty"`
}

func (s ListESAIPInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListESAIPInfoRequest) GoString() string {
	return s.String()
}

func (s *ListESAIPInfoRequest) GetVipInfo() *string {
	return s.VipInfo
}

func (s *ListESAIPInfoRequest) SetVipInfo(v string) *ListESAIPInfoRequest {
	s.VipInfo = &v
	return s
}

func (s *ListESAIPInfoRequest) Validate() error {
	return dara.Validate(s)
}
