// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreatePublishTaskShrinkRequest
	GetAgentKey() *string
	SetBizType(v string) *CreatePublishTaskShrinkRequest
	GetBizType() *string
	SetDataIdListShrink(v string) *CreatePublishTaskShrinkRequest
	GetDataIdListShrink() *string
}

type CreatePublishTaskShrinkRequest struct {
	// The workspace key. If not specified, the default workspace is accessed. Obtain the key value from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The publish unit type. To publish a chatbot, use the CreateInstancePublishTask API.
	//
	// example:
	//
	// faq
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Additional publish information. Currently supported: if BizType is faq, this field specifies category IDs, indicating that only knowledge under these categories is published.
	//
	// example:
	//
	// ["8521"]
	DataIdListShrink *string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty"`
}

func (s CreatePublishTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreatePublishTaskShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreatePublishTaskShrinkRequest) GetDataIdListShrink() *string {
	return s.DataIdListShrink
}

func (s *CreatePublishTaskShrinkRequest) SetAgentKey(v string) *CreatePublishTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePublishTaskShrinkRequest) SetBizType(v string) *CreatePublishTaskShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreatePublishTaskShrinkRequest) SetDataIdListShrink(v string) *CreatePublishTaskShrinkRequest {
	s.DataIdListShrink = &v
	return s
}

func (s *CreatePublishTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
