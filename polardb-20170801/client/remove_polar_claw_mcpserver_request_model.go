// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawMCPServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemovePolarClawMCPServerRequest
	GetApplicationId() *string
	SetServerName(v string) *RemovePolarClawMCPServerRequest
	GetServerName() *string
}

type RemovePolarClawMCPServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-v1
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s RemovePolarClawMCPServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawMCPServerRequest) GoString() string {
	return s.String()
}

func (s *RemovePolarClawMCPServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemovePolarClawMCPServerRequest) GetServerName() *string {
	return s.ServerName
}

func (s *RemovePolarClawMCPServerRequest) SetApplicationId(v string) *RemovePolarClawMCPServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemovePolarClawMCPServerRequest) SetServerName(v string) *RemovePolarClawMCPServerRequest {
	s.ServerName = &v
	return s
}

func (s *RemovePolarClawMCPServerRequest) Validate() error {
	return dara.Validate(s)
}
