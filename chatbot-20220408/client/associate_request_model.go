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
	// The business space key. If not specified, the default business space is accessed. Obtain the key value from the business management page of the primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique identifier of the robot.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The perspective code, which is used to call answers from different perspectives under the same knowledge title. When assembling request parameters, pass the parameter in the format Perspective={perspective code}. For example: &Perspective=["FZJBY3raWr"]. When using the SDK, use the parameters defined in the SDK. Currently, only one perspective answer can be called.
	//
	// example:
	//
	// ["qyzzVfyFfa"]
	Perspective []*string `json:"Perspective,omitempty" xml:"Perspective,omitempty" type:"Repeated"`
	// The number of recommended questions, ranging from 1 to 10. This parameter takes effect only when recommendations are returned. The number of returned recommended questions is less than or equal to RecommendNum.
	//
	// example:
	//
	// 8
	RecommendNum *int64 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// The session ID, which is used to identify a visitor\\"s session and maintain context information. For a new visitor, you do not need to pass this field when calling the Chat operation for the first time. The robot starts a session and returns the SessionId of the session in the response of the Chat operation. For subsequent conversation turns of the visitor, pass the SessionId of the current session when calling the Chat operation so that the robot can continue the conversation based on the SessionId.
	//
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The user utterance.
	//
	// example:
	//
	// Housing fund withdrawal
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
