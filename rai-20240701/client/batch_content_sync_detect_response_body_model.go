// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentSyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchContentSyncDetectResponseBody
	GetCode() *string
	SetData(v *BatchContentSyncDetectResponseBodyData) *BatchContentSyncDetectResponseBody
	GetData() *BatchContentSyncDetectResponseBodyData
	SetHttpStatusCode(v int32) *BatchContentSyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchContentSyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchContentSyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchContentSyncDetectResponseBody
	GetSuccess() *bool
}

type BatchContentSyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchContentSyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchContentSyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchContentSyncDetectResponseBody) GetData() *BatchContentSyncDetectResponseBodyData {
	return s.Data
}

func (s *BatchContentSyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchContentSyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchContentSyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchContentSyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchContentSyncDetectResponseBody) SetCode(v string) *BatchContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetData(v *BatchContentSyncDetectResponseBodyData) *BatchContentSyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *BatchContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetMessage(v string) *BatchContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetRequestId(v string) *BatchContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetSuccess(v bool) *BatchContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchContentSyncDetectResponseBodyData struct {
	DetectResultList []*BatchContentSyncDetectResponseBodyDataDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
}

func (s BatchContentSyncDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBodyData) GetDetectResultList() []*BatchContentSyncDetectResponseBodyDataDetectResultList {
	return s.DetectResultList
}

func (s *BatchContentSyncDetectResponseBodyData) SetDetectResultList(v []*BatchContentSyncDetectResponseBodyDataDetectResultList) *BatchContentSyncDetectResponseBodyData {
	s.DetectResultList = v
	return s
}

func (s *BatchContentSyncDetectResponseBodyData) Validate() error {
	if s.DetectResultList != nil {
		for _, item := range s.DetectResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchContentSyncDetectResponseBodyDataDetectResultList struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s BatchContentSyncDetectResponseBodyDataDetectResultList) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectResponseBodyDataDetectResultList) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) GetRiskInfo() *string {
	return s.RiskInfo
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) SetRiskInfo(v string) *BatchContentSyncDetectResponseBodyDataDetectResultList {
	s.RiskInfo = &v
	return s
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) SetRiskResult(v int32) *BatchContentSyncDetectResponseBodyDataDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) Validate() error {
	return dara.Validate(s)
}
