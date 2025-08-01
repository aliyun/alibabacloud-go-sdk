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
	// example:
	//
	// Success
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*ListAbnormalyEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Total   *int32                                 `json:"total,omitempty" xml:"total,omitempty"`
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
	return dara.Validate(s)
}

type ListAbnormalyEventsResponseBodyData struct {
	// example:
	//
	// 1725801090000
	CreatedAt   *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DiagStatus  *int32  `json:"diag_status,omitempty" xml:"diag_status,omitempty"`
	EndAt       *int64  `json:"end_at,omitempty" xml:"end_at,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance  *string                                    `json:"instance,omitempty" xml:"instance,omitempty"`
	Item      *string                                    `json:"item,omitempty" xml:"item,omitempty"`
	Level     *string                                    `json:"level,omitempty" xml:"level,omitempty"`
	Namespace *string                                    `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Opts      []*ListAbnormalyEventsResponseBodyDataOpts `json:"opts,omitempty" xml:"opts,omitempty" type:"Repeated"`
	Pod       *string                                    `json:"pod,omitempty" xml:"pod,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	return dara.Validate(s)
}

type ListAbnormalyEventsResponseBodyDataOpts struct {
	Label  *string                                        `json:"label,omitempty" xml:"label,omitempty"`
	Result *ListAbnormalyEventsResponseBodyDataOptsResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Type   *string                                        `json:"type,omitempty" xml:"type,omitempty"`
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
	return dara.Validate(s)
}

type ListAbnormalyEventsResponseBodyDataOptsResult struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
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
