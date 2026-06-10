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
	// The key of the business space. If you omit this parameter, the default business space is used. You can obtain the key from the Business Management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The type of content to publish. To publish a bot, use the `CreateInstancePublishTask` API.
	//
	// example:
	//
	// faq
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies data to publish by ID. This parameter is used only when `BizType` is set to `faq`. In this case, provide one or more category IDs to publish knowledge exclusively from the specified categories.
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
