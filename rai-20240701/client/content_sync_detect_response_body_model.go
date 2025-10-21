// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentSyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ContentSyncDetectResponseBody
	GetCode() *string
	SetData(v *ContentSyncDetectResponseBodyData) *ContentSyncDetectResponseBody
	GetData() *ContentSyncDetectResponseBodyData
	SetHttpStatusCode(v int32) *ContentSyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ContentSyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ContentSyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ContentSyncDetectResponseBody
	GetSuccess() *bool
}

type ContentSyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ContentSyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ContentSyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ContentSyncDetectResponseBody) GetData() *ContentSyncDetectResponseBodyData {
	return s.Data
}

func (s *ContentSyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ContentSyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ContentSyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContentSyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ContentSyncDetectResponseBody) SetCode(v string) *ContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetData(v *ContentSyncDetectResponseBodyData) *ContentSyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *ContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetMessage(v string) *ContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetRequestId(v string) *ContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetSuccess(v bool) *ContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ContentSyncDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContentSyncDetectResponseBodyData struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ContentSyncDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ContentSyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponseBodyData) GetRiskInfo() *string {
	return s.RiskInfo
}

func (s *ContentSyncDetectResponseBodyData) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *ContentSyncDetectResponseBodyData) SetRiskInfo(v string) *ContentSyncDetectResponseBodyData {
	s.RiskInfo = &v
	return s
}

func (s *ContentSyncDetectResponseBodyData) SetRiskResult(v int32) *ContentSyncDetectResponseBodyData {
	s.RiskResult = &v
	return s
}

func (s *ContentSyncDetectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
