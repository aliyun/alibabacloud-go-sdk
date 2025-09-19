// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGraph4InvestigationOnlineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnomalyId(v string) *DescribeGraph4InvestigationOnlineRequest
	GetAnomalyId() *string
	SetAnomalyUuid(v string) *DescribeGraph4InvestigationOnlineRequest
	GetAnomalyUuid() *string
	SetLang(v string) *DescribeGraph4InvestigationOnlineRequest
	GetLang() *string
	SetNamespace(v string) *DescribeGraph4InvestigationOnlineRequest
	GetNamespace() *string
	SetVertexId(v string) *DescribeGraph4InvestigationOnlineRequest
	GetVertexId() *string
}

type DescribeGraph4InvestigationOnlineRequest struct {
	// The ID of the alert event. You can call [DescribeSuspEvents](~~DescribeSuspEvents~~) to obtain the alert event ID, with the value path being: data.SuspEvents[index].UniqueInfo.
	//
	// example:
	//
	// 786fc80896b25422b5324cb6e57bxxxx
	AnomalyId *string `json:"AnomalyId,omitempty" xml:"AnomalyId,omitempty"`
	// The UUID of the alert event asset. You can call [DescribeSuspEvents](~~DescribeSuspEvents~~) to obtain the asset UUID, with the value path being: data.SuspEvents[index].Uuid.
	//
	// example:
	//
	// 3502e4b0-f362-4059-84a2-f47fa2b5exxx
	AnomalyUuid *string `json:"AnomalyUuid,omitempty" xml:"AnomalyUuid,omitempty"`
	// Sets the language type for the request and response messages. The default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The namespace of the graph, which is fixed as: hundun_dc_online.
	//
	// This parameter is required.
	//
	// example:
	//
	// hundun_dc_online
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Vertex ID. This does not need to be proactively provided.
	//
	// example:
	//
	// 29872354f741b1b044b8a9b4e2ab****
	VertexId *string `json:"VertexId,omitempty" xml:"VertexId,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineRequest) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineRequest) GetAnomalyId() *string {
	return s.AnomalyId
}

func (s *DescribeGraph4InvestigationOnlineRequest) GetAnomalyUuid() *string {
	return s.AnomalyUuid
}

func (s *DescribeGraph4InvestigationOnlineRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGraph4InvestigationOnlineRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeGraph4InvestigationOnlineRequest) GetVertexId() *string {
	return s.VertexId
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetAnomalyId(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.AnomalyId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetAnomalyUuid(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.AnomalyUuid = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetLang(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetNamespace(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetVertexId(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.VertexId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) Validate() error {
	return dara.Validate(s)
}
