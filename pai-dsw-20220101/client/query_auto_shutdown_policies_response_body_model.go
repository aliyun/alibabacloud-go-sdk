// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutoShutdownPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoShutdownPolicies(v []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) *QueryAutoShutdownPoliciesResponseBody
	GetAutoShutdownPolicies() []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies
	SetRequestId(v string) *QueryAutoShutdownPoliciesResponseBody
	GetRequestId() *string
}

type QueryAutoShutdownPoliciesResponseBody struct {
	AutoShutdownPolicies []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies `json:"AutoShutdownPolicies,omitempty" xml:"AutoShutdownPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAutoShutdownPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAutoShutdownPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAutoShutdownPoliciesResponseBody) GetAutoShutdownPolicies() []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	return s.AutoShutdownPolicies
}

func (s *QueryAutoShutdownPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAutoShutdownPoliciesResponseBody) SetAutoShutdownPolicies(v []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) *QueryAutoShutdownPoliciesResponseBody {
	s.AutoShutdownPolicies = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBody) SetRequestId(v string) *QueryAutoShutdownPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBody) Validate() error {
	if s.AutoShutdownPolicies != nil {
		for _, item := range s.AutoShutdownPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies struct {
	Conditions []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Deprecated
	Context map[string]interface{} `json:"Context,omitempty" xml:"Context,omitempty"`
	// Deprecated
	//
	// example:
	//
	// IdleDuration > 1440 min AND CpuUtilization <= 1%
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// dsw-05cefd0be2e5a278
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// Workspace
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) String() string {
	return dara.Prettify(s)
}

func (s QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GoString() string {
	return s.String()
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GetConditions() []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions {
	return s.Conditions
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GetContext() map[string]interface{} {
	return s.Context
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GetExpression() *string {
	return s.Expression
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) SetConditions(v []*QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	s.Conditions = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) SetContext(v map[string]interface{}) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	s.Context = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) SetExpression(v string) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	s.Expression = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) SetInstanceId(v string) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	s.InstanceId = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) SetSourceType(v string) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies {
	s.SourceType = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPolicies) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions struct {
	Context map[string]interface{} `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// IdleDuration > 1440 min AND CpuUtilization <= 1%
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Workspace
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) String() string {
	return dara.Prettify(s)
}

func (s QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) GoString() string {
	return s.String()
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) GetContext() map[string]interface{} {
	return s.Context
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) GetExpression() *string {
	return s.Expression
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) SetContext(v map[string]interface{}) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions {
	s.Context = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) SetExpression(v string) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions {
	s.Expression = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) SetSourceType(v string) *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions {
	s.SourceType = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponseBodyAutoShutdownPoliciesConditions) Validate() error {
	return dara.Validate(s)
}
