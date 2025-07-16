// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v map[string]interface{}) *InvokeSkillRequest
	GetParams() map[string]interface{}
	SetSkillId(v string) *InvokeSkillRequest
	GetSkillId() *string
	SetStream(v bool) *InvokeSkillRequest
	GetStream() *bool
}

type InvokeSkillRequest struct {
	// example:
	//
	// {}
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a1d033dd-xxxx-49cf-b49b-2068081bb551
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	Stream  *bool   `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s InvokeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillRequest) GoString() string {
	return s.String()
}

func (s *InvokeSkillRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *InvokeSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *InvokeSkillRequest) GetStream() *bool {
	return s.Stream
}

func (s *InvokeSkillRequest) SetParams(v map[string]interface{}) *InvokeSkillRequest {
	s.Params = v
	return s
}

func (s *InvokeSkillRequest) SetSkillId(v string) *InvokeSkillRequest {
	s.SkillId = &v
	return s
}

func (s *InvokeSkillRequest) SetStream(v bool) *InvokeSkillRequest {
	s.Stream = &v
	return s
}

func (s *InvokeSkillRequest) Validate() error {
	return dara.Validate(s)
}
