// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComponentsVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CheckComponentsVersionRequest
	GetClusterId() *string
	SetComponents(v string) *CheckComponentsVersionRequest
	GetComponents() *string
}

type CheckComponentsVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HBASE,HADOOP
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s CheckComponentsVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckComponentsVersionRequest) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CheckComponentsVersionRequest) GetComponents() *string {
	return s.Components
}

func (s *CheckComponentsVersionRequest) SetClusterId(v string) *CheckComponentsVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckComponentsVersionRequest) SetComponents(v string) *CheckComponentsVersionRequest {
	s.Components = &v
	return s
}

func (s *CheckComponentsVersionRequest) Validate() error {
	return dara.Validate(s)
}
