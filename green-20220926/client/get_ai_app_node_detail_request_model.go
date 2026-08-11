// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppNodeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppNodeDetailRequest
	GetAppId() *string
	SetEndTime(v string) *GetAiAppNodeDetailRequest
	GetEndTime() *string
	SetNodeId(v string) *GetAiAppNodeDetailRequest
	GetNodeId() *string
	SetNodeName(v string) *GetAiAppNodeDetailRequest
	GetNodeName() *string
	SetNodeType(v string) *GetAiAppNodeDetailRequest
	GetNodeType() *string
	SetRegionId(v string) *GetAiAppNodeDetailRequest
	GetRegionId() *string
	SetStartTime(v string) *GetAiAppNodeDetailRequest
	GetStartTime() *string
}

type GetAiAppNodeDetailRequest struct {
	// The application ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2026-01-02 16:08:38
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// node-xxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node type. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// TOOL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2026-01-01 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAiAppNodeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppNodeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppNodeDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppNodeDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppNodeDetailRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAiAppNodeDetailRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAiAppNodeDetailRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *GetAiAppNodeDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppNodeDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppNodeDetailRequest) SetAppId(v string) *GetAiAppNodeDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetEndTime(v string) *GetAiAppNodeDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetNodeId(v string) *GetAiAppNodeDetailRequest {
	s.NodeId = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetNodeName(v string) *GetAiAppNodeDetailRequest {
	s.NodeName = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetNodeType(v string) *GetAiAppNodeDetailRequest {
	s.NodeType = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetRegionId(v string) *GetAiAppNodeDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) SetStartTime(v string) *GetAiAppNodeDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetAiAppNodeDetailRequest) Validate() error {
	return dara.Validate(s)
}
