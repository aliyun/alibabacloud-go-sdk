// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiInterpretation(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetAiInterpretation() *string
	SetAttackSources(v []*DescribeAiEngineEvidenceRequestResponseBodyAttackSources) *DescribeAiEngineEvidenceRequestResponseBody
	GetAttackSources() []*DescribeAiEngineEvidenceRequestResponseBodyAttackSources
	SetEvidenceId(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetEvidenceId() *string
	SetHost(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetHost() *string
	SetHttpStatus(v int32) *DescribeAiEngineEvidenceRequestResponseBody
	GetHttpStatus() *int32
	SetImpactType(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetImpactType() *string
	SetMethod(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetMethod() *string
	SetNodeId(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetNodeId() *string
	SetRequest(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetRequest() *string
	SetRequestId(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetRequestId() *string
	SetRequestPath(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetRequestPath() *string
	SetRequestTime(v int64) *DescribeAiEngineEvidenceRequestResponseBody
	GetRequestTime() *int64
	SetResource(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetResource() *string
	SetResponse(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetResponse() *string
	SetSuspectedBypass(v int32) *DescribeAiEngineEvidenceRequestResponseBody
	GetSuspectedBypass() *int32
	SetTraceId(v string) *DescribeAiEngineEvidenceRequestResponseBody
	GetTraceId() *string
}

type DescribeAiEngineEvidenceRequestResponseBody struct {
	// The AI interpretation result in the specified language, in Markdown format.
	//
	// example:
	//
	// This request has SQL injection risk.
	AiInterpretation *string `json:"AiInterpretation,omitempty" xml:"AiInterpretation,omitempty"`
	// The list of attack source groups for the request.
	AttackSources []*DescribeAiEngineEvidenceRequestResponseBodyAttackSources `json:"AttackSources,omitempty" xml:"AttackSources,omitempty" type:"Repeated"`
	// The evidence ID.
	//
	// example:
	//
	// EVD-20241025-001
	EvidenceId *string `json:"EvidenceId,omitempty" xml:"EvidenceId,omitempty"`
	// The host of the request.
	//
	// example:
	//
	// api.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The HTTP response status code.
	//
	// example:
	//
	// 500
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The impact tag of the evidence. Valid values:
	//
	// - **block**: The request is blocked.
	//
	// - **challenge**: The request triggers a challenge.
	//
	// - **monitor**: The request is monitored.
	//
	// - **bypass**: The request is suspected to be a bypass.
	//
	// - **attempt**: The request is a reconnaissance probe.
	//
	// example:
	//
	// bypass
	ImpactType *string `json:"ImpactType,omitempty" xml:"ImpactType,omitempty"`
	// The HTTP method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The ID of the associated node, which corresponds to the NodeId in the node list.
	//
	// example:
	//
	// node-109-04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The request content.
	//
	// example:
	//
	// GET /api/orders/query?id=1%27 HTTP/1.1
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A***2-1219-019A-8B44-255*****2C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The masked request path.
	//
	// example:
	//
	// /api/orders/query
	RequestPath *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	// The time when the original request was made. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1792973400000
	RequestTime *int64 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The ID of the protected object to which the request belongs.
	//
	// example:
	//
	// www.example.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The response content.
	//
	// example:
	//
	// HTTP/1.1 500 Internal Server Error
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// Specifies whether the request is suspected to be a bypass. Valid values:
	//
	// - **1**: The request is suspected to be a bypass.
	//
	// - **0**: The request is not a bypass.
	//
	// example:
	//
	// 1
	SuspectedBypass *int32 `json:"SuspectedBypass,omitempty" xml:"SuspectedBypass,omitempty"`
	// The trace ID of the original request, which can be up to 128 ASCII characters in length.
	//
	// example:
	//
	// trace-20241025-109-04
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetAiInterpretation() *string {
	return s.AiInterpretation
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetAttackSources() []*DescribeAiEngineEvidenceRequestResponseBodyAttackSources {
	return s.AttackSources
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetEvidenceId() *string {
	return s.EvidenceId
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetImpactType() *string {
	return s.ImpactType
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetMethod() *string {
	return s.Method
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetRequest() *string {
	return s.Request
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetResource() *string {
	return s.Resource
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetResponse() *string {
	return s.Response
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetSuspectedBypass() *int32 {
	return s.SuspectedBypass
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetAiInterpretation(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.AiInterpretation = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetAttackSources(v []*DescribeAiEngineEvidenceRequestResponseBodyAttackSources) *DescribeAiEngineEvidenceRequestResponseBody {
	s.AttackSources = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetEvidenceId(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.EvidenceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetHost(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetHttpStatus(v int32) *DescribeAiEngineEvidenceRequestResponseBody {
	s.HttpStatus = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetImpactType(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.ImpactType = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetMethod(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.Method = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetNodeId(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetRequest(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.Request = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetRequestId(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetRequestPath(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.RequestPath = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetRequestTime(v int64) *DescribeAiEngineEvidenceRequestResponseBody {
	s.RequestTime = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetResource(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.Resource = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetResponse(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.Response = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetSuspectedBypass(v int32) *DescribeAiEngineEvidenceRequestResponseBody {
	s.SuspectedBypass = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) SetTraceId(v string) *DescribeAiEngineEvidenceRequestResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBody) Validate() error {
	if s.AttackSources != nil {
		for _, item := range s.AttackSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEvidenceRequestResponseBodyAttackSources struct {
	// The list of sources of this type.
	Sources []*DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The type of the attack source.
	//
	// example:
	//
	// ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestResponseBodyAttackSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestResponseBodyAttackSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSources) GetSources() []*DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources {
	return s.Sources
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSources) GetType() *string {
	return s.Type
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSources) SetSources(v []*DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) *DescribeAiEngineEvidenceRequestResponseBodyAttackSources {
	s.Sources = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSources) SetType(v string) *DescribeAiEngineEvidenceRequestResponseBodyAttackSources {
	s.Type = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSources) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources struct {
	// The display attribute of the source, such as the region code.
	//
	// example:
	//
	// US-CA
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The value of the source, such as an IP address.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) GetValue() *string {
	return s.Value
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) SetAttribute(v string) *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources {
	s.Attribute = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) SetValue(v string) *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources {
	s.Value = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponseBodyAttackSourcesSources) Validate() error {
	return dara.Validate(s)
}
