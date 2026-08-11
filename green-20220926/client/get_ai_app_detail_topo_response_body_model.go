// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailTopoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetAiAppDetailTopoResponseBodyData) *GetAiAppDetailTopoResponseBody
	GetData() []*GetAiAppDetailTopoResponseBodyData
	SetRequestId(v string) *GetAiAppDetailTopoResponseBody
	GetRequestId() *string
}

type GetAiAppDetailTopoResponseBody struct {
	// The returned data.
	Data []*GetAiAppDetailTopoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiAppDetailTopoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoResponseBody) GetData() []*GetAiAppDetailTopoResponseBodyData {
	return s.Data
}

func (s *GetAiAppDetailTopoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppDetailTopoResponseBody) SetData(v []*GetAiAppDetailTopoResponseBodyData) *GetAiAppDetailTopoResponseBody {
	s.Data = v
	return s
}

func (s *GetAiAppDetailTopoResponseBody) SetRequestId(v string) *GetAiAppDetailTopoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppDetailTopoResponseBodyData struct {
	// The node category.
	//
	// - LLM
	//
	// - Knowledge
	//
	// - Tools
	//
	// - Others
	//
	// example:
	//
	// LLM
	NodeCategory *string `json:"NodeCategory,omitempty" xml:"NodeCategory,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// idxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// namexxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node type. Valid values:
	//
	// - **APP**: end-to-end agent.
	//
	// - **MODEL**: large language model.
	//
	// - **TOOL**: tool.
	//
	// example:
	//
	// TOOL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The request count.
	//
	// example:
	//
	// 100
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The number of alerts.
	//
	// example:
	//
	// 20
	WarningCount *int32 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s GetAiAppDetailTopoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoResponseBodyData) GetNodeCategory() *string {
	return s.NodeCategory
}

func (s *GetAiAppDetailTopoResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAiAppDetailTopoResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAiAppDetailTopoResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *GetAiAppDetailTopoResponseBodyData) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *GetAiAppDetailTopoResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAiAppDetailTopoResponseBodyData) GetWarningCount() *int32 {
	return s.WarningCount
}

func (s *GetAiAppDetailTopoResponseBodyData) SetNodeCategory(v string) *GetAiAppDetailTopoResponseBodyData {
	s.NodeCategory = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetNodeId(v string) *GetAiAppDetailTopoResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetNodeName(v string) *GetAiAppDetailTopoResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetNodeType(v string) *GetAiAppDetailTopoResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetRequestCount(v int64) *GetAiAppDetailTopoResponseBodyData {
	s.RequestCount = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetRiskLevel(v string) *GetAiAppDetailTopoResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) SetWarningCount(v int32) *GetAiAppDetailTopoResponseBodyData {
	s.WarningCount = &v
	return s
}

func (s *GetAiAppDetailTopoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
