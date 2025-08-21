// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNCInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeNCInformationResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeNCInformationResponseBodyData) *DescribeNCInformationResponseBody
	GetData() []*DescribeNCInformationResponseBodyData
	SetDesc(v string) *DescribeNCInformationResponseBody
	GetDesc() *string
	SetMsg(v string) *DescribeNCInformationResponseBody
	GetMsg() *string
	SetPager(v *DescribeNCInformationResponseBodyPager) *DescribeNCInformationResponseBody
	GetPager() *DescribeNCInformationResponseBodyPager
	SetRequestId(v string) *DescribeNCInformationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNCInformationResponseBody
	GetTotalCount() *int32
}

type DescribeNCInformationResponseBody struct {
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*DescribeNCInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Desc        *string                                  `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Msg         *string                                  `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Pager       *DescribeNCInformationResponseBodyPager  `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNCInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNCInformationResponseBody) GetData() []*DescribeNCInformationResponseBodyData {
	return s.Data
}

func (s *DescribeNCInformationResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *DescribeNCInformationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeNCInformationResponseBody) GetPager() *DescribeNCInformationResponseBodyPager {
	return s.Pager
}

func (s *DescribeNCInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNCInformationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNCInformationResponseBody) SetCurrentPage(v int32) *DescribeNCInformationResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNCInformationResponseBody) SetData(v []*DescribeNCInformationResponseBodyData) *DescribeNCInformationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNCInformationResponseBody) SetDesc(v string) *DescribeNCInformationResponseBody {
	s.Desc = &v
	return s
}

func (s *DescribeNCInformationResponseBody) SetMsg(v string) *DescribeNCInformationResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeNCInformationResponseBody) SetPager(v *DescribeNCInformationResponseBodyPager) *DescribeNCInformationResponseBody {
	s.Pager = v
	return s
}

func (s *DescribeNCInformationResponseBody) SetRequestId(v string) *DescribeNCInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNCInformationResponseBody) SetTotalCount(v int32) *DescribeNCInformationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNCInformationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyData struct {
	Cpu     *DescribeNCInformationResponseBodyDataCpu    `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	Gpu     *DescribeNCInformationResponseBodyDataGpu    `json:"Gpu,omitempty" xml:"Gpu,omitempty" type:"Struct"`
	Hdd     *DescribeNCInformationResponseBodyDataHdd    `json:"Hdd,omitempty" xml:"Hdd,omitempty" type:"Struct"`
	Info    *DescribeNCInformationResponseBodyDataInfo   `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Memory  *DescribeNCInformationResponseBodyDataMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Nvme    *DescribeNCInformationResponseBodyDataNvme   `json:"Nvme,omitempty" xml:"Nvme,omitempty" type:"Struct"`
	Online  *bool                                        `json:"Online,omitempty" xml:"Online,omitempty"`
	Region  *string                                      `json:"Region,omitempty" xml:"Region,omitempty"`
	Ssd     *DescribeNCInformationResponseBodyDataSsd    `json:"Ssd,omitempty" xml:"Ssd,omitempty" type:"Struct"`
	Virtual *string                                      `json:"Virtual,omitempty" xml:"Virtual,omitempty"`
}

func (s DescribeNCInformationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyData) GetCpu() *DescribeNCInformationResponseBodyDataCpu {
	return s.Cpu
}

func (s *DescribeNCInformationResponseBodyData) GetGpu() *DescribeNCInformationResponseBodyDataGpu {
	return s.Gpu
}

func (s *DescribeNCInformationResponseBodyData) GetHdd() *DescribeNCInformationResponseBodyDataHdd {
	return s.Hdd
}

func (s *DescribeNCInformationResponseBodyData) GetInfo() *DescribeNCInformationResponseBodyDataInfo {
	return s.Info
}

func (s *DescribeNCInformationResponseBodyData) GetMemory() *DescribeNCInformationResponseBodyDataMemory {
	return s.Memory
}

func (s *DescribeNCInformationResponseBodyData) GetNvme() *DescribeNCInformationResponseBodyDataNvme {
	return s.Nvme
}

func (s *DescribeNCInformationResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *DescribeNCInformationResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *DescribeNCInformationResponseBodyData) GetSsd() *DescribeNCInformationResponseBodyDataSsd {
	return s.Ssd
}

func (s *DescribeNCInformationResponseBodyData) GetVirtual() *string {
	return s.Virtual
}

func (s *DescribeNCInformationResponseBodyData) SetCpu(v *DescribeNCInformationResponseBodyDataCpu) *DescribeNCInformationResponseBodyData {
	s.Cpu = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetGpu(v *DescribeNCInformationResponseBodyDataGpu) *DescribeNCInformationResponseBodyData {
	s.Gpu = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetHdd(v *DescribeNCInformationResponseBodyDataHdd) *DescribeNCInformationResponseBodyData {
	s.Hdd = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetInfo(v *DescribeNCInformationResponseBodyDataInfo) *DescribeNCInformationResponseBodyData {
	s.Info = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetMemory(v *DescribeNCInformationResponseBodyDataMemory) *DescribeNCInformationResponseBodyData {
	s.Memory = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetNvme(v *DescribeNCInformationResponseBodyDataNvme) *DescribeNCInformationResponseBodyData {
	s.Nvme = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetOnline(v bool) *DescribeNCInformationResponseBodyData {
	s.Online = &v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetRegion(v string) *DescribeNCInformationResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetSsd(v *DescribeNCInformationResponseBodyDataSsd) *DescribeNCInformationResponseBodyData {
	s.Ssd = v
	return s
}

func (s *DescribeNCInformationResponseBodyData) SetVirtual(v string) *DescribeNCInformationResponseBodyData {
	s.Virtual = &v
	return s
}

func (s *DescribeNCInformationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataCpu struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataCpu) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataCpu) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataCpu) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataCpu {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetRemain(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataCpu {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetReserved(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataCpu {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetTotal(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetType(v string) *DescribeNCInformationResponseBodyDataCpu {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetUsed(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataCpu {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataCpu) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataGpu struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataGpu) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataGpu) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataGpu) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataGpu {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetRemain(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataGpu {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetReserved(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataGpu {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetTotal(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetType(v string) *DescribeNCInformationResponseBodyDataGpu {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetUsed(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataGpu {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataGpu) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataHdd struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataHdd) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataHdd) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataHdd) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataHdd {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetRemain(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataHdd {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetReserved(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataHdd {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetTotal(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetType(v string) *DescribeNCInformationResponseBodyDataHdd {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetUsed(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataHdd {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataHdd) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataInfo struct {
	Ip   *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag  []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Uuid *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataInfo) GetIp() *string {
	return s.Ip
}

func (s *DescribeNCInformationResponseBodyDataInfo) GetName() *string {
	return s.Name
}

func (s *DescribeNCInformationResponseBodyDataInfo) GetTag() []*string {
	return s.Tag
}

func (s *DescribeNCInformationResponseBodyDataInfo) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeNCInformationResponseBodyDataInfo) SetIp(v string) *DescribeNCInformationResponseBodyDataInfo {
	s.Ip = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataInfo) SetName(v string) *DescribeNCInformationResponseBodyDataInfo {
	s.Name = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataInfo) SetTag(v []*string) *DescribeNCInformationResponseBodyDataInfo {
	s.Tag = v
	return s
}

func (s *DescribeNCInformationResponseBodyDataInfo) SetUuid(v string) *DescribeNCInformationResponseBodyDataInfo {
	s.Uuid = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataMemory struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataMemory) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataMemory) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataMemory) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataMemory {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetRemain(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataMemory {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetReserved(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataMemory {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetTotal(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetType(v string) *DescribeNCInformationResponseBodyDataMemory {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetUsed(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataMemory {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataMemory) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataNvme struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataNvme) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataNvme) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataNvme) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataNvme {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetRemain(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataNvme {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetReserved(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataNvme {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetTotal(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetType(v string) *DescribeNCInformationResponseBodyDataNvme {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetUsed(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataNvme {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataNvme) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyDataSsd struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeNCInformationResponseBodyDataSsd) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyDataSsd) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetType() *string {
	return s.Type
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeNCInformationResponseBodyDataSsd) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetDisplay(v bool) *DescribeNCInformationResponseBodyDataSsd {
	s.Display = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetOversellRatio(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.OversellRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetRemain(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.Remain = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetReserveDisable(v bool) *DescribeNCInformationResponseBodyDataSsd {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetReserveDisableTotal(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetReserved(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.Reserved = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetStatusDisable(v bool) *DescribeNCInformationResponseBodyDataSsd {
	s.StatusDisable = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetStatusDisableTotal(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetTotal(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetType(v string) *DescribeNCInformationResponseBodyDataSsd {
	s.Type = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetUsed(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.Used = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) SetUsedRatio(v int64) *DescribeNCInformationResponseBodyDataSsd {
	s.UsedRatio = &v
	return s
}

func (s *DescribeNCInformationResponseBodyDataSsd) Validate() error {
	return dara.Validate(s)
}

type DescribeNCInformationResponseBodyPager struct {
	Page  *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	Size  *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeNCInformationResponseBodyPager) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationResponseBodyPager) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationResponseBodyPager) GetPage() *int64 {
	return s.Page
}

func (s *DescribeNCInformationResponseBodyPager) GetSize() *int64 {
	return s.Size
}

func (s *DescribeNCInformationResponseBodyPager) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeNCInformationResponseBodyPager) SetPage(v int64) *DescribeNCInformationResponseBodyPager {
	s.Page = &v
	return s
}

func (s *DescribeNCInformationResponseBodyPager) SetSize(v int64) *DescribeNCInformationResponseBodyPager {
	s.Size = &v
	return s
}

func (s *DescribeNCInformationResponseBodyPager) SetTotal(v int64) *DescribeNCInformationResponseBodyPager {
	s.Total = &v
	return s
}

func (s *DescribeNCInformationResponseBodyPager) Validate() error {
	return dara.Validate(s)
}
