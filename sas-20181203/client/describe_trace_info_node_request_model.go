// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *DescribeTraceInfoNodeRequest
	GetEventName() *string
	SetFrom(v string) *DescribeTraceInfoNodeRequest
	GetFrom() *string
	SetIncidentTime(v int64) *DescribeTraceInfoNodeRequest
	GetIncidentTime() *int64
	SetLang(v string) *DescribeTraceInfoNodeRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeTraceInfoNodeRequest
	GetSourceIp() *string
	SetType(v string) *DescribeTraceInfoNodeRequest
	GetType() *string
	SetUuid(v string) *DescribeTraceInfoNodeRequest
	GetUuid() *string
	SetVertexId(v string) *DescribeTraceInfoNodeRequest
	GetVertexId() *string
}

type DescribeTraceInfoNodeRequest struct {
	// The event name.
	//
	// >For more information, call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain this parameter.
	//
	// example:
	//
	// WEBSHELL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The source identifier of the request. Set the value to sas.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The time when the event was first detected.
	//
	// example:
	//
	// 1635978934000
	IncidentTime *int64 `json:"IncidentTime,omitempty" xml:"IncidentTime,omitempty"`
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains the value.
	//
	// example:
	//
	// 127.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The vertex type. You can call the [DescribeTraceInfoDetail](~~DescribeTraceInfoDetail~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// SAS_ASSET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server to query. You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6f346617-eef9-45e6-b6d1-946xxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The vertex ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03da4e2350a3eb50cd25a18cexxxxxxx
	VertexId *string `json:"VertexId,omitempty" xml:"VertexId,omitempty"`
}

func (s DescribeTraceInfoNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoNodeRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeTraceInfoNodeRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeTraceInfoNodeRequest) GetIncidentTime() *int64 {
	return s.IncidentTime
}

func (s *DescribeTraceInfoNodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTraceInfoNodeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeTraceInfoNodeRequest) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoNodeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeTraceInfoNodeRequest) GetVertexId() *string {
	return s.VertexId
}

func (s *DescribeTraceInfoNodeRequest) SetEventName(v string) *DescribeTraceInfoNodeRequest {
	s.EventName = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetFrom(v string) *DescribeTraceInfoNodeRequest {
	s.From = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetIncidentTime(v int64) *DescribeTraceInfoNodeRequest {
	s.IncidentTime = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetLang(v string) *DescribeTraceInfoNodeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetSourceIp(v string) *DescribeTraceInfoNodeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetType(v string) *DescribeTraceInfoNodeRequest {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetUuid(v string) *DescribeTraceInfoNodeRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) SetVertexId(v string) *DescribeTraceInfoNodeRequest {
	s.VertexId = &v
	return s
}

func (s *DescribeTraceInfoNodeRequest) Validate() error {
	return dara.Validate(s)
}
