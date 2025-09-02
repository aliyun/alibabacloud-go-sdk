// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataServiceApiTestResponseBodyData) *GetDataServiceApiTestResponseBody
	GetData() *GetDataServiceApiTestResponseBodyData
	SetRequestId(v string) *GetDataServiceApiTestResponseBody
	GetRequestId() *string
}

type GetDataServiceApiTestResponseBody struct {
	// Return object
	Data *GetDataServiceApiTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// adssdsewe
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataServiceApiTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiTestResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiTestResponseBody) GetData() *GetDataServiceApiTestResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApiTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiTestResponseBody) SetData(v *GetDataServiceApiTestResponseBodyData) *GetDataServiceApiTestResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApiTestResponseBody) SetRequestId(v string) *GetDataServiceApiTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiTestResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiTestResponseBodyData struct {
	// Test API Id
	//
	// example:
	//
	// 12343
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// Time consuming
	//
	// example:
	//
	// 10
	CostTime *string `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// Debug information
	//
	// example:
	//
	// [<span style=\\"color: #92D581;\\">INFO</span>] [16:15:13.240] resource group is 498774069027041[<span style=\\"color: #92D581;\\">INFO</span>]
	DebugInfo *string `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// Node Debug information
	//
	// example:
	//
	// [<span style=\\"color: #92D581;\\">INFO</span>] [16:15:13.240] resource group is 498774069027041[<span style=\\"color: #92D581;\\">INFO</span>]
	NodesDebugInfo *string `json:"NodesDebugInfo,omitempty" xml:"NodesDebugInfo,omitempty"`
	// Test API request parameters
	//
	// example:
	//
	// {"name":"test"}
	ParamMap *string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// The test API returns the code. If it is not completed, the data is empty.
	//
	// example:
	//
	// 0
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// Return data
	//
	// example:
	//
	// {"id":2}
	RetResult *string `json:"RetResult,omitempty" xml:"RetResult,omitempty"`
	// Whether the task has been completed, including: RUNNING,FINISHED
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataServiceApiTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiTestResponseBodyData) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetDataServiceApiTestResponseBodyData) GetCostTime() *string {
	return s.CostTime
}

func (s *GetDataServiceApiTestResponseBodyData) GetDebugInfo() *string {
	return s.DebugInfo
}

func (s *GetDataServiceApiTestResponseBodyData) GetNodesDebugInfo() *string {
	return s.NodesDebugInfo
}

func (s *GetDataServiceApiTestResponseBodyData) GetParamMap() *string {
	return s.ParamMap
}

func (s *GetDataServiceApiTestResponseBodyData) GetRetCode() *int64 {
	return s.RetCode
}

func (s *GetDataServiceApiTestResponseBodyData) GetRetResult() *string {
	return s.RetResult
}

func (s *GetDataServiceApiTestResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDataServiceApiTestResponseBodyData) SetApiId(v int64) *GetDataServiceApiTestResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetCostTime(v string) *GetDataServiceApiTestResponseBodyData {
	s.CostTime = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetDebugInfo(v string) *GetDataServiceApiTestResponseBodyData {
	s.DebugInfo = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetNodesDebugInfo(v string) *GetDataServiceApiTestResponseBodyData {
	s.NodesDebugInfo = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetParamMap(v string) *GetDataServiceApiTestResponseBodyData {
	s.ParamMap = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetRetCode(v int64) *GetDataServiceApiTestResponseBodyData {
	s.RetCode = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetRetResult(v string) *GetDataServiceApiTestResponseBodyData {
	s.RetResult = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) SetStatus(v string) *GetDataServiceApiTestResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDataServiceApiTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
