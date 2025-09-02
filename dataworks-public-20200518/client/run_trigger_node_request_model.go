// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTriggerNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *RunTriggerNodeRequest
	GetAppId() *int64
	SetBizDate(v int64) *RunTriggerNodeRequest
	GetBizDate() *int64
	SetCycleTime(v int64) *RunTriggerNodeRequest
	GetCycleTime() *int64
	SetNodeId(v int64) *RunTriggerNodeRequest
	GetNodeId() *int64
}

type RunTriggerNodeRequest struct {
	// The ID of the DataWorks workspace to which the manually triggered node belongs. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data timestamp of the instance that is generated for the manually triggered node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606200230105
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The scheduling time to run the manually triggered node. Set the value to a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606200230105
	CycleTime *int64 `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The ID of the manually triggered node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000011
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s RunTriggerNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTriggerNodeRequest) GoString() string {
	return s.String()
}

func (s *RunTriggerNodeRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *RunTriggerNodeRequest) GetBizDate() *int64 {
	return s.BizDate
}

func (s *RunTriggerNodeRequest) GetCycleTime() *int64 {
	return s.CycleTime
}

func (s *RunTriggerNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *RunTriggerNodeRequest) SetAppId(v int64) *RunTriggerNodeRequest {
	s.AppId = &v
	return s
}

func (s *RunTriggerNodeRequest) SetBizDate(v int64) *RunTriggerNodeRequest {
	s.BizDate = &v
	return s
}

func (s *RunTriggerNodeRequest) SetCycleTime(v int64) *RunTriggerNodeRequest {
	s.CycleTime = &v
	return s
}

func (s *RunTriggerNodeRequest) SetNodeId(v int64) *RunTriggerNodeRequest {
	s.NodeId = &v
	return s
}

func (s *RunTriggerNodeRequest) Validate() error {
	return dara.Validate(s)
}
