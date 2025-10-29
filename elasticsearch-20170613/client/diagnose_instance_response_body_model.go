// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DiagnoseInstanceResponseBody
	GetRequestId() *string
	SetResult(v *DiagnoseInstanceResponseBodyResult) *DiagnoseInstanceResponseBody
	GetResult() *DiagnoseInstanceResponseBodyResult
}

type DiagnoseInstanceResponseBody struct {
	// The ID of the report.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The diagnosis status. Valid values: Supported: SUCCESS, FAILED, and RUNNING.
	Result *DiagnoseInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DiagnoseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiagnoseInstanceResponseBody) GetResult() *DiagnoseInstanceResponseBodyResult {
	return s.Result
}

func (s *DiagnoseInstanceResponseBody) SetRequestId(v string) *DiagnoseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseInstanceResponseBody) SetResult(v *DiagnoseInstanceResponseBodyResult) *DiagnoseInstanceResponseBody {
	s.Result = v
	return s
}

func (s *DiagnoseInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiagnoseInstanceResponseBodyResult struct {
	// The ID of the diagnostic instance.
	//
	// example:
	//
	// 1535745731000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// trigger__2020-08-17T17:09:02
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DiagnoseInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DiagnoseInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DiagnoseInstanceResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *DiagnoseInstanceResponseBodyResult) GetState() *string {
	return s.State
}

func (s *DiagnoseInstanceResponseBodyResult) SetCreateTime(v int64) *DiagnoseInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetInstanceId(v string) *DiagnoseInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetReportId(v string) *DiagnoseInstanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetState(v string) *DiagnoseInstanceResponseBodyResult {
	s.State = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
