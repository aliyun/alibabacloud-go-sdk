// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *ListTargetsRequest
	GetArn() *string
	SetEventBusName(v string) *ListTargetsRequest
	GetEventBusName() *string
	SetLimit(v int32) *ListTargetsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListTargetsRequest
	GetNextToken() *string
	SetRuleName(v string) *ListTargetsRequest
	GetRuleName() *string
}

type ListTargetsRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the event rule.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:118609547428****:services/fc-connector.a1/functions/event
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of returned entries in a call.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the event rule.
	//
	// example:
	//
	// tf-testacc-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListTargetsRequest) GetArn() *string {
	return s.Arn
}

func (s *ListTargetsRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListTargetsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListTargetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTargetsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListTargetsRequest) SetArn(v string) *ListTargetsRequest {
	s.Arn = &v
	return s
}

func (s *ListTargetsRequest) SetEventBusName(v string) *ListTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *ListTargetsRequest) SetLimit(v int32) *ListTargetsRequest {
	s.Limit = &v
	return s
}

func (s *ListTargetsRequest) SetNextToken(v string) *ListTargetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTargetsRequest) SetRuleName(v string) *ListTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *ListTargetsRequest) Validate() error {
	return dara.Validate(s)
}
