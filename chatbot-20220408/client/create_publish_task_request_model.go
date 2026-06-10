// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublishTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreatePublishTaskRequest
	GetAgentKey() *string
	SetBizType(v string) *CreatePublishTaskRequest
	GetBizType() *string
	SetDataIdList(v []*string) *CreatePublishTaskRequest
	GetDataIdList() []*string
}

type CreatePublishTaskRequest struct {
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
	DataIdList []*string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty" type:"Repeated"`
}

func (s CreatePublishTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreatePublishTaskRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreatePublishTaskRequest) GetDataIdList() []*string {
	return s.DataIdList
}

func (s *CreatePublishTaskRequest) SetAgentKey(v string) *CreatePublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CreatePublishTaskRequest) SetBizType(v string) *CreatePublishTaskRequest {
	s.BizType = &v
	return s
}

func (s *CreatePublishTaskRequest) SetDataIdList(v []*string) *CreatePublishTaskRequest {
	s.DataIdList = v
	return s
}

func (s *CreatePublishTaskRequest) Validate() error {
	return dara.Validate(s)
}
