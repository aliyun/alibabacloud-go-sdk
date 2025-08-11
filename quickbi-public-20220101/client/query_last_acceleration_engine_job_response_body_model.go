// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLastAccelerationEngineJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryLastAccelerationEngineJobResponseBody
	GetRequestId() *string
	SetResult(v *QueryLastAccelerationEngineJobResponseBodyResult) *QueryLastAccelerationEngineJobResponseBody
	GetResult() *QueryLastAccelerationEngineJobResponseBodyResult
	SetSuccess(v bool) *QueryLastAccelerationEngineJobResponseBody
	GetSuccess() *bool
}

type QueryLastAccelerationEngineJobResponseBody struct {
	// example:
	//
	// 46e53*********92704c8
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryLastAccelerationEngineJobResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLastAccelerationEngineJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLastAccelerationEngineJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLastAccelerationEngineJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLastAccelerationEngineJobResponseBody) GetResult() *QueryLastAccelerationEngineJobResponseBodyResult {
	return s.Result
}

func (s *QueryLastAccelerationEngineJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryLastAccelerationEngineJobResponseBody) SetRequestId(v string) *QueryLastAccelerationEngineJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBody) SetResult(v *QueryLastAccelerationEngineJobResponseBodyResult) *QueryLastAccelerationEngineJobResponseBody {
	s.Result = v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBody) SetSuccess(v bool) *QueryLastAccelerationEngineJobResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryLastAccelerationEngineJobResponseBodyResult struct {
	// example:
	//
	// 2025-06-18 17:07:43
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-18 17:08:26
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 46e53********5464564
	JobHistoryId *string `json:"JobHistoryId,omitempty" xml:"JobHistoryId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryLastAccelerationEngineJobResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryLastAccelerationEngineJobResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) GetJobHistoryId() *string {
	return s.JobHistoryId
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) SetGmtCreate(v string) *QueryLastAccelerationEngineJobResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) SetGmtModified(v string) *QueryLastAccelerationEngineJobResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) SetJobHistoryId(v string) *QueryLastAccelerationEngineJobResponseBodyResult {
	s.JobHistoryId = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) SetStatus(v string) *QueryLastAccelerationEngineJobResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
