// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApsOptimizationStrategyResponseBody
	GetCode() *string
	SetData(v []*ListApsOptimizationStrategyResponseBodyData) *ListApsOptimizationStrategyResponseBody
	GetData() []*ListApsOptimizationStrategyResponseBodyData
	SetHttpStatusCode(v int32) *ListApsOptimizationStrategyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApsOptimizationStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApsOptimizationStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApsOptimizationStrategyResponseBody
	GetSuccess() *bool
}

type ListApsOptimizationStrategyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// -
	Data []*ListApsOptimizationStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApsOptimizationStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApsOptimizationStrategyResponseBody) GetData() []*ListApsOptimizationStrategyResponseBodyData {
	return s.Data
}

func (s *ListApsOptimizationStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApsOptimizationStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApsOptimizationStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApsOptimizationStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApsOptimizationStrategyResponseBody) SetCode(v string) *ListApsOptimizationStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) SetData(v []*ListApsOptimizationStrategyResponseBodyData) *ListApsOptimizationStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) SetHttpStatusCode(v int32) *ListApsOptimizationStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) SetMessage(v string) *ListApsOptimizationStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) SetRequestId(v string) *ListApsOptimizationStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) SetSuccess(v bool) *ListApsOptimizationStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBody) Validate() error {
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

type ListApsOptimizationStrategyResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// amv-23xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The status of the lifecycle management policy. Valid values:
	//
	// 1.  on: enabled.
	//
	// 2.  off: disabled.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the lifecycle management policy.
	//
	// example:
	//
	// test
	StrategyDesc *string `json:"StrategyDesc,omitempty" xml:"StrategyDesc,omitempty"`
	// The name of the lifecycle management policy.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The type of the lifecycle management policy. Only StrategyValue is returned.
	//
	// example:
	//
	// StrategyValue
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s ListApsOptimizationStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationStrategyResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsOptimizationStrategyResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListApsOptimizationStrategyResponseBodyData) GetStrategyDesc() *string {
	return s.StrategyDesc
}

func (s *ListApsOptimizationStrategyResponseBodyData) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListApsOptimizationStrategyResponseBodyData) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ListApsOptimizationStrategyResponseBodyData) SetDBClusterId(v string) *ListApsOptimizationStrategyResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBodyData) SetStatus(v string) *ListApsOptimizationStrategyResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBodyData) SetStrategyDesc(v string) *ListApsOptimizationStrategyResponseBodyData {
	s.StrategyDesc = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBodyData) SetStrategyName(v string) *ListApsOptimizationStrategyResponseBodyData {
	s.StrategyName = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBodyData) SetStrategyType(v string) *ListApsOptimizationStrategyResponseBodyData {
	s.StrategyType = &v
	return s
}

func (s *ListApsOptimizationStrategyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
