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
