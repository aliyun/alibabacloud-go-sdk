// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbnormalyEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAbnormalyEventsResponseBody
	GetCode() *string
	SetData(v []*ListAbnormalyEventsResponseBodyData) *ListAbnormalyEventsResponseBody
	GetData() []*ListAbnormalyEventsResponseBodyData
	SetMessage(v string) *ListAbnormalyEventsResponseBody
	GetMessage() *string
	SetTotal(v int32) *ListAbnormalyEventsResponseBody
	GetTotal() *int32
}

type ListAbnormalyEventsResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization was successful.
	//
	// - Other status codes indicate that the authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data []*ListAbnormalyEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code description. This value is empty if no error occurred.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 4
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAbnormalyEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAbnormalyEventsResponseBody) GetData() []*ListAbnormalyEventsResponseBodyData {
	return s.Data
}

func (s *ListAbnormalyEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAbnormalyEventsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAbnormalyEventsResponseBody) SetCode(v string) *ListAbnormalyEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAbnormalyEventsResponseBody) SetData(v []*ListAbnormalyEventsResponseBodyData) *ListAbnormalyEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListAbnormalyEventsResponseBody) SetMessage(v string) *ListAbnormalyEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAbnormalyEventsResponseBody) SetTotal(v int32) *ListAbnormalyEventsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAbnormalyEventsResponseBody) Validate() error {
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

type ListAbnormalyEventsResponseBodyData struct {
	// The creation time.
	//
	// example:
	//
	// 1725801090000
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The detailed description of the anomaly item.
	//
	// example:
	//
	// 节点发生OOM, 可查看OOM发生原因
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The diagnostic status. Valid values:
	//
	// - 1: diagnosis ready.
	//
	// - 2: diagnosis in progress.
	//
	// - 3: diagnosis completed.
	//
	// - 4: not diagnosable.
	//
	// - 5: diagnosis failed.
	//
	// example:
	//
	// 3
	DiagStatus *int32 `json:"diag_status,omitempty" xml:"diag_status,omitempty"`
	// The end time of the anomaly event.
	//
	// example:
	//
	// 1725797727754
	EndAt *int64 `json:"end_at,omitempty" xml:"end_at,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The name of the anomaly item.
	//
	// example:
	//
	// 节点CPU使用率检测
	Item *string `json:"item,omitempty" xml:"item,omitempty"`
	// The level of the anomaly item.
	//
	// example:
	//
	// potential
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The namespace of the pod.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The list of operations for the anomaly item.
	Opts []*ListAbnormalyEventsResponseBodyDataOpts `json:"opts,omitempty" xml:"opts,omitempty" type:"Repeated"`
	// The pod name.
	//
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// The raw metrics.
	RawMetrics *ListAbnormalyEventsResponseBodyDataRawMetrics `json:"raw_metrics,omitempty" xml:"raw_metrics,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The type of the anomaly item.
	//
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The UUID of the anomaly event.
	//
	// example:
	//
	// 43f05b46-1034-42e8-a528-6e5ca1108277
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyData) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListAbnormalyEventsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAbnormalyEventsResponseBodyData) GetDiagStatus() *int32 {
	return s.DiagStatus
}

func (s *ListAbnormalyEventsResponseBodyData) GetEndAt() *int64 {
	return s.EndAt
}

func (s *ListAbnormalyEventsResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *ListAbnormalyEventsResponseBodyData) GetItem() *string {
	return s.Item
}

func (s *ListAbnormalyEventsResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *ListAbnormalyEventsResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAbnormalyEventsResponseBodyData) GetOpts() []*ListAbnormalyEventsResponseBodyDataOpts {
	return s.Opts
}

func (s *ListAbnormalyEventsResponseBodyData) GetPod() *string {
	return s.Pod
}

func (s *ListAbnormalyEventsResponseBodyData) GetRawMetrics() *ListAbnormalyEventsResponseBodyDataRawMetrics {
	return s.RawMetrics
}

func (s *ListAbnormalyEventsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAbnormalyEventsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListAbnormalyEventsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListAbnormalyEventsResponseBodyData) SetCreatedAt(v int64) *ListAbnormalyEventsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetDescription(v string) *ListAbnormalyEventsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetDiagStatus(v int32) *ListAbnormalyEventsResponseBodyData {
	s.DiagStatus = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetEndAt(v int64) *ListAbnormalyEventsResponseBodyData {
	s.EndAt = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetInstance(v string) *ListAbnormalyEventsResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetItem(v string) *ListAbnormalyEventsResponseBodyData {
	s.Item = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetLevel(v string) *ListAbnormalyEventsResponseBodyData {
	s.Level = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetNamespace(v string) *ListAbnormalyEventsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetOpts(v []*ListAbnormalyEventsResponseBodyDataOpts) *ListAbnormalyEventsResponseBodyData {
	s.Opts = v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetPod(v string) *ListAbnormalyEventsResponseBodyData {
	s.Pod = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetRawMetrics(v *ListAbnormalyEventsResponseBodyDataRawMetrics) *ListAbnormalyEventsResponseBodyData {
	s.RawMetrics = v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetRegionId(v string) *ListAbnormalyEventsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetType(v string) *ListAbnormalyEventsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetUuid(v string) *ListAbnormalyEventsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) Validate() error {
	if s.Opts != nil {
		for _, item := range s.Opts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RawMetrics != nil {
		if err := s.RawMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAbnormalyEventsResponseBodyDataOpts struct {
	// The operation type.
	//
	// example:
	//
	// diagnose
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The diagnostic result of the anomaly item.
	Result *ListAbnormalyEventsResponseBodyDataOptsResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// The diagnostic type.
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyDataOpts) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyDataOpts) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) GetLabel() *string {
	return s.Label
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) GetResult() *ListAbnormalyEventsResponseBodyDataOptsResult {
	return s.Result
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) GetType() *string {
	return s.Type
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetLabel(v string) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Label = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetResult(v *ListAbnormalyEventsResponseBodyDataOptsResult) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Result = v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetType(v string) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Type = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAbnormalyEventsResponseBodyDataOptsResult struct {
	// The diagnostic status.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The URL of the diagnostic details.
	//
	// example:
	//
	// /diagnose/result/PhfFg456
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyDataOptsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyDataOptsResult) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyDataOptsResult) GetStatus() *string {
	return s.Status
}

func (s *ListAbnormalyEventsResponseBodyDataOptsResult) GetUrl() *string {
	return s.Url
}

func (s *ListAbnormalyEventsResponseBodyDataOptsResult) SetStatus(v string) *ListAbnormalyEventsResponseBodyDataOptsResult {
	s.Status = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOptsResult) SetUrl(v string) *ListAbnormalyEventsResponseBodyDataOptsResult {
	s.Url = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOptsResult) Validate() error {
	return dara.Validate(s)
}

type ListAbnormalyEventsResponseBodyDataRawMetrics struct {
	// The end time.
	//
	// example:
	//
	// 1761814928
	EndTime *float32 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// The list of metric values.
	Metrics []*string `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The start time.
	//
	// example:
	//
	// 1761814928
	StartTime *float32 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyDataRawMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyDataRawMetrics) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) GetEndTime() *float32 {
	return s.EndTime
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) GetMetrics() []*string {
	return s.Metrics
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) GetStartTime() *float32 {
	return s.StartTime
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) SetEndTime(v float32) *ListAbnormalyEventsResponseBodyDataRawMetrics {
	s.EndTime = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) SetMetrics(v []*string) *ListAbnormalyEventsResponseBodyDataRawMetrics {
	s.Metrics = v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) SetStartTime(v float32) *ListAbnormalyEventsResponseBodyDataRawMetrics {
	s.StartTime = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataRawMetrics) Validate() error {
	return dara.Validate(s)
}
