// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishFlowVersionRequest
	GetDescription() *string
	SetFlowName(v string) *PublishFlowVersionRequest
	GetFlowName() *string
}

type PublishFlowVersionRequest struct {
	// example:
	//
	// example flow description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s PublishFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishFlowVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishFlowVersionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *PublishFlowVersionRequest) SetDescription(v string) *PublishFlowVersionRequest {
	s.Description = &v
	return s
}

func (s *PublishFlowVersionRequest) SetFlowName(v string) *PublishFlowVersionRequest {
	s.FlowName = &v
	return s
}

func (s *PublishFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
