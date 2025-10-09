// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeTraceInfoDetailRequest
	GetFrom() *string
	SetIncidentTime(v int64) *DescribeTraceInfoDetailRequest
	GetIncidentTime() *int64
	SetLang(v string) *DescribeTraceInfoDetailRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeTraceInfoDetailRequest
	GetSourceIp() *string
	SetType(v string) *DescribeTraceInfoDetailRequest
	GetType() *string
	SetUuid(v string) *DescribeTraceInfoDetailRequest
	GetUuid() *string
	SetVertexId(v string) *DescribeTraceInfoDetailRequest
	GetVertexId() *string
}

type DescribeTraceInfoDetailRequest struct {
	// The ID of the request source. Set the value to **sas**.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The timestamp of the detection. Unit: milliseconds.
	//
	// example:
	//
	// 1670555392000
	IncidentTime *int64 `json:"IncidentTime,omitempty" xml:"IncidentTime,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. The value of this parameter is specified by the system.
	//
	// example:
	//
	// 127.0.0.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the vertex. Set the value to **SAS_INCIDENT**.
	//
	// This parameter is required.
	//
	// example:
	//
	// SAS_INCIDENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627f2d7-aaa2-4ed1-b07a-xxxxxxxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Vertex ID, which can be obtained from the AlarmUniqueInfo in the response of the [DescribeSuspEvents](~~DescribeSuspEvents~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// cce1d28dxxxxxxxxxxxxxxxx
	VertexId *string `json:"VertexId,omitempty" xml:"VertexId,omitempty"`
}

func (s DescribeTraceInfoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeTraceInfoDetailRequest) GetIncidentTime() *int64 {
	return s.IncidentTime
}

func (s *DescribeTraceInfoDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTraceInfoDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeTraceInfoDetailRequest) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeTraceInfoDetailRequest) GetVertexId() *string {
	return s.VertexId
}

func (s *DescribeTraceInfoDetailRequest) SetFrom(v string) *DescribeTraceInfoDetailRequest {
	s.From = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetIncidentTime(v int64) *DescribeTraceInfoDetailRequest {
	s.IncidentTime = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetLang(v string) *DescribeTraceInfoDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetSourceIp(v string) *DescribeTraceInfoDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetType(v string) *DescribeTraceInfoDetailRequest {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetUuid(v string) *DescribeTraceInfoDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) SetVertexId(v string) *DescribeTraceInfoDetailRequest {
	s.VertexId = &v
	return s
}

func (s *DescribeTraceInfoDetailRequest) Validate() error {
	return dara.Validate(s)
}
