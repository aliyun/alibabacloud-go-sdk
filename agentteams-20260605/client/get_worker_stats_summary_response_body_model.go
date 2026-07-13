// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerStatsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkerStatsSummaryResponseBody
	GetCode() *string
	SetData(v *GetWorkerStatsSummaryResponseBodyData) *GetWorkerStatsSummaryResponseBody
	GetData() *GetWorkerStatsSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkerStatsSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkerStatsSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkerStatsSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkerStatsSummaryResponseBody
	GetSuccess() *bool
}

type GetWorkerStatsSummaryResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetWorkerStatsSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerStatsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerStatsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerStatsSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkerStatsSummaryResponseBody) GetData() *GetWorkerStatsSummaryResponseBodyData {
	return s.Data
}

func (s *GetWorkerStatsSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkerStatsSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkerStatsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkerStatsSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkerStatsSummaryResponseBody) SetCode(v string) *GetWorkerStatsSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) SetData(v *GetWorkerStatsSummaryResponseBodyData) *GetWorkerStatsSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) SetHttpStatusCode(v int32) *GetWorkerStatsSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) SetMessage(v string) *GetWorkerStatsSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) SetRequestId(v string) *GetWorkerStatsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) SetSuccess(v bool) *GetWorkerStatsSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkerStatsSummaryResponseBodyData struct {
	OtherWorkers   *int32 `json:"OtherWorkers,omitempty" xml:"OtherWorkers,omitempty"`
	RunningWorkers *int32 `json:"RunningWorkers,omitempty" xml:"RunningWorkers,omitempty"`
	StoppedWorkers *int32 `json:"StoppedWorkers,omitempty" xml:"StoppedWorkers,omitempty"`
	TotalWorkers   *int32 `json:"TotalWorkers,omitempty" xml:"TotalWorkers,omitempty"`
}

func (s GetWorkerStatsSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerStatsSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerStatsSummaryResponseBodyData) GetOtherWorkers() *int32 {
	return s.OtherWorkers
}

func (s *GetWorkerStatsSummaryResponseBodyData) GetRunningWorkers() *int32 {
	return s.RunningWorkers
}

func (s *GetWorkerStatsSummaryResponseBodyData) GetStoppedWorkers() *int32 {
	return s.StoppedWorkers
}

func (s *GetWorkerStatsSummaryResponseBodyData) GetTotalWorkers() *int32 {
	return s.TotalWorkers
}

func (s *GetWorkerStatsSummaryResponseBodyData) SetOtherWorkers(v int32) *GetWorkerStatsSummaryResponseBodyData {
	s.OtherWorkers = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBodyData) SetRunningWorkers(v int32) *GetWorkerStatsSummaryResponseBodyData {
	s.RunningWorkers = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBodyData) SetStoppedWorkers(v int32) *GetWorkerStatsSummaryResponseBodyData {
	s.StoppedWorkers = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBodyData) SetTotalWorkers(v int32) *GetWorkerStatsSummaryResponseBodyData {
	s.TotalWorkers = &v
	return s
}

func (s *GetWorkerStatsSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
