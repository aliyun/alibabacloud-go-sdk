// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSupplementDagrunResponseBody
	GetCode() *string
	SetDagrunList(v []*GetSupplementDagrunResponseBodyDagrunList) *GetSupplementDagrunResponseBody
	GetDagrunList() []*GetSupplementDagrunResponseBodyDagrunList
	SetHttpStatusCode(v int32) *GetSupplementDagrunResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSupplementDagrunResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSupplementDagrunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSupplementDagrunResponseBody
	GetSuccess() *bool
}

type GetSupplementDagrunResponseBody struct {
	// example:
	//
	// OK
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DagrunList []*GetSupplementDagrunResponseBodyDagrunList `json:"DagrunList,omitempty" xml:"DagrunList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSupplementDagrunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSupplementDagrunResponseBody) GetDagrunList() []*GetSupplementDagrunResponseBodyDagrunList {
	return s.DagrunList
}

func (s *GetSupplementDagrunResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSupplementDagrunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSupplementDagrunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupplementDagrunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSupplementDagrunResponseBody) SetCode(v string) *GetSupplementDagrunResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetDagrunList(v []*GetSupplementDagrunResponseBodyDagrunList) *GetSupplementDagrunResponseBody {
	s.DagrunList = v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetHttpStatusCode(v int32) *GetSupplementDagrunResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetMessage(v string) *GetSupplementDagrunResponseBody {
	s.Message = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetRequestId(v string) *GetSupplementDagrunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetSuccess(v bool) *GetSupplementDagrunResponseBody {
	s.Success = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunResponseBodyDagrunList struct {
	// example:
	//
	// 2024-04-01
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1717081789000
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// Dagrun ID
	//
	// example:
	//
	// dr_2242792_14542
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717081729000
	StartExecuteTime *int64 `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// f_8241792_20201202_2099680
	SupplementId *string `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s GetSupplementDagrunResponseBodyDagrunList) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunResponseBodyDagrunList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetBizDate() *string {
	return s.BizDate
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetDuration() *string {
	return s.Duration
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetStartExecuteTime() *int64 {
	return s.StartExecuteTime
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetStatus() *string {
	return s.Status
}

func (s *GetSupplementDagrunResponseBodyDagrunList) GetSupplementId() *string {
	return s.SupplementId
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetBizDate(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.BizDate = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetDuration(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Duration = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetEndExecuteTime(v int64) *GetSupplementDagrunResponseBodyDagrunList {
	s.EndExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetId(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetStartExecuteTime(v int64) *GetSupplementDagrunResponseBodyDagrunList {
	s.StartExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetStatus(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Status = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetSupplementId(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.SupplementId = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) Validate() error {
	return dara.Validate(s)
}
