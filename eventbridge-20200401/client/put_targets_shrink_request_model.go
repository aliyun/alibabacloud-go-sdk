// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTargetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *PutTargetsShrinkRequest
	GetEventBusName() *string
	SetRuleName(v string) *PutTargetsShrinkRequest
	GetRuleName() *string
	SetTargetsShrink(v string) *PutTargetsShrinkRequest
	GetTargetsShrink() *string
}

type PutTargetsShrinkRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// eventTest
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssr-send-to-vendor-test01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The event targets to be created or updated. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s PutTargetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutTargetsShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *PutTargetsShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutTargetsShrinkRequest) GetTargetsShrink() *string {
	return s.TargetsShrink
}

func (s *PutTargetsShrinkRequest) SetEventBusName(v string) *PutTargetsShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *PutTargetsShrinkRequest) SetRuleName(v string) *PutTargetsShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *PutTargetsShrinkRequest) SetTargetsShrink(v string) *PutTargetsShrinkRequest {
	s.TargetsShrink = &v
	return s
}

func (s *PutTargetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
