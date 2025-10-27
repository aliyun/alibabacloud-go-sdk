// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *PublishFlowVersionResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *PublishFlowVersionResponseBody
	GetDescription() *string
	SetFlowName(v string) *PublishFlowVersionResponseBody
	GetFlowName() *string
	SetRequestId(v string) *PublishFlowVersionResponseBody
	GetRequestId() *string
	SetVersion(v string) *PublishFlowVersionResponseBody
	GetVersion() *string
}

type PublishFlowVersionResponseBody struct {
	// example:
	//
	// 2025-10-24T14:11:26+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// my flow description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// my-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 294D68C1-5108-5971-853A-1A9CC87B4816
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PublishFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishFlowVersionResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *PublishFlowVersionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *PublishFlowVersionResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *PublishFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishFlowVersionResponseBody) GetVersion() *string {
	return s.Version
}

func (s *PublishFlowVersionResponseBody) SetCreatedTime(v string) *PublishFlowVersionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *PublishFlowVersionResponseBody) SetDescription(v string) *PublishFlowVersionResponseBody {
	s.Description = &v
	return s
}

func (s *PublishFlowVersionResponseBody) SetFlowName(v string) *PublishFlowVersionResponseBody {
	s.FlowName = &v
	return s
}

func (s *PublishFlowVersionResponseBody) SetRequestId(v string) *PublishFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishFlowVersionResponseBody) SetVersion(v string) *PublishFlowVersionResponseBody {
	s.Version = &v
	return s
}

func (s *PublishFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
