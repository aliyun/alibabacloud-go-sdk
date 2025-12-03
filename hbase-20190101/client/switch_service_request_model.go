// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *SwitchServiceRequest
	GetClusterId() *string
	SetOperate(v string) *SwitchServiceRequest
	GetOperate() *string
	SetServiceName(v string) *SwitchServiceRequest
	GetServiceName() *string
}

type SwitchServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-uf6r2hn2zrxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HBaseProxy
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s SwitchServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchServiceRequest) GoString() string {
	return s.String()
}

func (s *SwitchServiceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SwitchServiceRequest) GetOperate() *string {
	return s.Operate
}

func (s *SwitchServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SwitchServiceRequest) SetClusterId(v string) *SwitchServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *SwitchServiceRequest) SetOperate(v string) *SwitchServiceRequest {
	s.Operate = &v
	return s
}

func (s *SwitchServiceRequest) SetServiceName(v string) *SwitchServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *SwitchServiceRequest) Validate() error {
	return dara.Validate(s)
}
