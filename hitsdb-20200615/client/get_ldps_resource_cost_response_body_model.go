// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsResourceCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetLdpsResourceCostResponseBody
	GetEndTime() *int64
	SetInstanceId(v string) *GetLdpsResourceCostResponseBody
	GetInstanceId() *string
	SetJobId(v string) *GetLdpsResourceCostResponseBody
	GetJobId() *string
	SetRequestId(v string) *GetLdpsResourceCostResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *GetLdpsResourceCostResponseBody
	GetStartTime() *int64
	SetTotalResource(v int64) *GetLdpsResourceCostResponseBody
	GetTotalResource() *int64
}

type GetLdpsResourceCostResponseBody struct {
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalResource *int64  `json:"TotalResource,omitempty" xml:"TotalResource,omitempty"`
}

func (s GetLdpsResourceCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsResourceCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetLdpsResourceCostResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLdpsResourceCostResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetLdpsResourceCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLdpsResourceCostResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetLdpsResourceCostResponseBody) GetTotalResource() *int64 {
	return s.TotalResource
}

func (s *GetLdpsResourceCostResponseBody) SetEndTime(v int64) *GetLdpsResourceCostResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetInstanceId(v string) *GetLdpsResourceCostResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetJobId(v string) *GetLdpsResourceCostResponseBody {
	s.JobId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetRequestId(v string) *GetLdpsResourceCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetStartTime(v int64) *GetLdpsResourceCostResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetTotalResource(v int64) *GetLdpsResourceCostResponseBody {
	s.TotalResource = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) Validate() error {
	return dara.Validate(s)
}
