// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *AssociateShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *AssociateShrinkRequest
	GetInstanceId() *string
	SetPerspectiveShrink(v string) *AssociateShrinkRequest
	GetPerspectiveShrink() *string
	SetRecommendNum(v int64) *AssociateShrinkRequest
	GetRecommendNum() *int64
	SetSessionId(v string) *AssociateShrinkRequest
	GetSessionId() *string
	SetUtterance(v string) *AssociateShrinkRequest
	GetUtterance() *string
}

type AssociateShrinkRequest struct {
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
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
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

func (s AssociateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssociateShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *AssociateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateShrinkRequest) GetPerspectiveShrink() *string {
	return s.PerspectiveShrink
}

func (s *AssociateShrinkRequest) GetRecommendNum() *int64 {
	return s.RecommendNum
}

func (s *AssociateShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AssociateShrinkRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *AssociateShrinkRequest) SetAgentKey(v string) *AssociateShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateShrinkRequest) SetInstanceId(v string) *AssociateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateShrinkRequest) SetPerspectiveShrink(v string) *AssociateShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *AssociateShrinkRequest) SetRecommendNum(v int64) *AssociateShrinkRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateShrinkRequest) SetSessionId(v string) *AssociateShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateShrinkRequest) SetUtterance(v string) *AssociateShrinkRequest {
	s.Utterance = &v
	return s
}

func (s *AssociateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
