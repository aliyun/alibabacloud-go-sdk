// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAScripts(v []*ListAScriptsResponseBodyAScripts) *ListAScriptsResponseBody
	GetAScripts() []*ListAScriptsResponseBodyAScripts
	SetMaxResults(v int32) *ListAScriptsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAScriptsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAScriptsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAScriptsResponseBody
	GetTotalCount() *int32
}

type ListAScriptsResponseBody struct {
	// The AScript rules.
	AScripts []*ListAScriptsResponseBodyAScripts `json:"AScripts,omitempty" xml:"AScripts,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// This parameter is required.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2CA81429-F160-593A-8AB5-A2A9617845B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// > This parameter is optional. By default, this parameter is not returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponseBody) GetAScripts() []*ListAScriptsResponseBodyAScripts {
	return s.AScripts
}

func (s *ListAScriptsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAScriptsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAScriptsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAScriptsResponseBody) SetAScripts(v []*ListAScriptsResponseBodyAScripts) *ListAScriptsResponseBody {
	s.AScripts = v
	return s
}

func (s *ListAScriptsResponseBody) SetMaxResults(v int32) *ListAScriptsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAScriptsResponseBody) SetNextToken(v string) *ListAScriptsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAScriptsResponseBody) SetRequestId(v string) *ListAScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAScriptsResponseBody) SetTotalCount(v int32) *ListAScriptsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAScriptsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAScriptsResponseBodyAScripts struct {
	// The AScript rule ID.
	//
	// example:
	//
	// as-aznwocxofkakf7****
	AScriptId *string `json:"AScriptId,omitempty" xml:"AScriptId,omitempty"`
	// The name of the AScript rule.
	//
	// example:
	//
	// test
	AScriptName *string `json:"AScriptName,omitempty" xml:"AScriptName,omitempty"`
	// The status of the AScript rule. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **Configuring**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Available
	AScriptStatus *string `json:"AScriptStatus,omitempty" xml:"AScriptStatus,omitempty"`
	// Indicates whether the AScript rule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsn-t0w1m9r6suiwmc****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The Application Load Balancer (ALB) instance ID.
	//
	// example:
	//
	// alb-vv9rg2ub31tyec****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The content of the AScript rule.
	//
	// example:
	//
	// {test}
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
}

func (s ListAScriptsResponseBodyAScripts) String() string {
	return dara.Prettify(s)
}

func (s ListAScriptsResponseBodyAScripts) GoString() string {
	return s.String()
}

func (s *ListAScriptsResponseBodyAScripts) GetAScriptId() *string {
	return s.AScriptId
}

func (s *ListAScriptsResponseBodyAScripts) GetAScriptName() *string {
	return s.AScriptName
}

func (s *ListAScriptsResponseBodyAScripts) GetAScriptStatus() *string {
	return s.AScriptStatus
}

func (s *ListAScriptsResponseBodyAScripts) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAScriptsResponseBodyAScripts) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListAScriptsResponseBodyAScripts) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListAScriptsResponseBodyAScripts) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptId(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptName(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptName = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetAScriptStatus(v string) *ListAScriptsResponseBodyAScripts {
	s.AScriptStatus = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetEnabled(v bool) *ListAScriptsResponseBodyAScripts {
	s.Enabled = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetListenerId(v string) *ListAScriptsResponseBodyAScripts {
	s.ListenerId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetLoadBalancerId(v string) *ListAScriptsResponseBodyAScripts {
	s.LoadBalancerId = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) SetScriptContent(v string) *ListAScriptsResponseBodyAScripts {
	s.ScriptContent = &v
	return s
}

func (s *ListAScriptsResponseBodyAScripts) Validate() error {
	return dara.Validate(s)
}
