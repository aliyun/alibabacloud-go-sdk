// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReConfigApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ReConfigApplicationRequest
	GetAppId() *string
	SetVariables(v string) *ReConfigApplicationRequest
	GetVariables() *string
}

type ReConfigApplicationRequest struct {
	// example:
	//
	// Q2P4O9YSOKCT35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {"${instance_type}":"ecs.c6.3xlarge"}
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ReConfigApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ReConfigApplicationRequest) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ReConfigApplicationRequest) GetVariables() *string {
	return s.Variables
}

func (s *ReConfigApplicationRequest) SetAppId(v string) *ReConfigApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ReConfigApplicationRequest) SetVariables(v string) *ReConfigApplicationRequest {
	s.Variables = &v
	return s
}

func (s *ReConfigApplicationRequest) Validate() error {
	return dara.Validate(s)
}
