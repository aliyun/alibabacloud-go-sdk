// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawMCPServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawMCPServersRequest
	GetApplicationId() *string
	SetServerList(v []*string) *DescribePolarClawMCPServersRequest
	GetServerList() []*string
}

type DescribePolarClawMCPServersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ServerList    []*string `json:"ServerList,omitempty" xml:"ServerList,omitempty" type:"Repeated"`
}

func (s DescribePolarClawMCPServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawMCPServersRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawMCPServersRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawMCPServersRequest) GetServerList() []*string {
	return s.ServerList
}

func (s *DescribePolarClawMCPServersRequest) SetApplicationId(v string) *DescribePolarClawMCPServersRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawMCPServersRequest) SetServerList(v []*string) *DescribePolarClawMCPServersRequest {
	s.ServerList = v
	return s
}

func (s *DescribePolarClawMCPServersRequest) Validate() error {
	return dara.Validate(s)
}
