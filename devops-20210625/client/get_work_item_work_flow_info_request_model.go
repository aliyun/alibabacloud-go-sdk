// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemWorkFlowInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationId(v string) *GetWorkItemWorkFlowInfoRequest
	GetConfigurationId() *string
}

type GetWorkItemWorkFlowInfoRequest struct {
	// example:
	//
	// 711d33c738b9171c45fa......
	ConfigurationId *string `json:"configurationId,omitempty" xml:"configurationId,omitempty"`
}

func (s GetWorkItemWorkFlowInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoRequest) GetConfigurationId() *string {
	return s.ConfigurationId
}

func (s *GetWorkItemWorkFlowInfoRequest) SetConfigurationId(v string) *GetWorkItemWorkFlowInfoRequest {
	s.ConfigurationId = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoRequest) Validate() error {
	return dara.Validate(s)
}
