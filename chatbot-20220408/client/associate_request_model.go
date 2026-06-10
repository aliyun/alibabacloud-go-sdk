// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *AssociateRequest
	GetAgentKey() *string
	SetInstanceId(v string) *AssociateRequest
	GetInstanceId() *string
	SetPerspective(v []*string) *AssociateRequest
	GetPerspective() []*string
	SetRecommendNum(v int64) *AssociateRequest
	GetRecommendNum() *int64
	SetSessionId(v string) *AssociateRequest
	GetSessionId() *string
	SetUtterance(v string) *AssociateRequest
	GetUtterance() *string
}

type AssociateRequest struct {
	// The workspace key. If this parameter is not specified, the service uses the default workspace. You can obtain the key from the workspace management page of your Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The code for a perspective. Use this to retrieve answers from different perspectives that share the same knowledge title. When constructing the request, pass the value in the "Perspective={perspective_code}" format, for example, \\&Perspective=["FZJBY3raWr"]. If you use an SDK, pass the parameter as defined by the SDK. You can specify only one perspective in each request.
	//
	// example:
	//
	// ["qyzzVfyFfa"]
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The number of recommended FAQs to return. The value must be an integer from 1 to 10. This parameter takes effect only when recommendations are generated. The number of FAQs returned is less than or equal to the value you specify.
	//
	// example:
	//
	// 8
	RecommendNum *int64 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The session ID that identifies a conversation and maintains context. For the first request from a new user, you can omit this parameter. The service automatically starts a session and returns a session ID in the response. To continue the conversation, you must include the session ID in subsequent requests.
	//
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The user\\"s utterance.
	//
	// example:
	//
	// 公积金提取
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s AssociateRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRequest) GoString() string {
	return s.String()
}

func (s *AssociateRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *AssociateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateRequest) GetPerspective() []*string {
	return s.Perspective
}

func (s *AssociateRequest) GetRecommendNum() *int64 {
	return s.RecommendNum
}

func (s *AssociateRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AssociateRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *AssociateRequest) SetAgentKey(v string) *AssociateRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateRequest) SetInstanceId(v string) *AssociateRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateRequest) SetPerspective(v []*string) *AssociateRequest {
	s.Perspective = v
	return s
}

func (s *AssociateRequest) SetRecommendNum(v int64) *AssociateRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateRequest) SetSessionId(v string) *AssociateRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateRequest) SetUtterance(v string) *AssociateRequest {
	s.Utterance = &v
	return s
}

func (s *AssociateRequest) Validate() error {
	return dara.Validate(s)
}
