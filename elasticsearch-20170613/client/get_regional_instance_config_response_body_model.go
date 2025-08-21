// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionalInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRegionalInstanceConfigResponseBody
	GetRequestId() *string
	SetResult(v *GetRegionalInstanceConfigResponseBodyResult) *GetRegionalInstanceConfigResponseBody
	GetResult() *GetRegionalInstanceConfigResponseBodyResult
}

type GetRegionalInstanceConfigResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetRegionalInstanceConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRegionalInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegionalInstanceConfigResponseBody) GetResult() *GetRegionalInstanceConfigResponseBodyResult {
	return s.Result
}

func (s *GetRegionalInstanceConfigResponseBody) SetRequestId(v string) *GetRegionalInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBody) SetResult(v *GetRegionalInstanceConfigResponseBodyResult) *GetRegionalInstanceConfigResponseBody {
	s.Result = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResult struct {
	ClientNodeAmountRange *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange `json:"clientNodeAmountRange,omitempty" xml:"clientNodeAmountRange,omitempty" type:"Struct"`
	ClientNodeDiskList    []*GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList  `json:"clientNodeDiskList,omitempty" xml:"clientNodeDiskList,omitempty" type:"Repeated"`
	ClientSpecs           []*string                                                         `json:"clientSpecs,omitempty" xml:"clientSpecs,omitempty" type:"Repeated"`
	DataNodeAmountRange   *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange   `json:"dataNodeAmountRange,omitempty" xml:"dataNodeAmountRange,omitempty" type:"Struct"`
	DataNodeDiskList      []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList    `json:"dataNodeDiskList,omitempty" xml:"dataNodeDiskList,omitempty" type:"Repeated"`
	DataNodeSpecs         []*string                                                         `json:"dataNodeSpecs,omitempty" xml:"dataNodeSpecs,omitempty" type:"Repeated"`
	KibanaSpecs           []*string                                                         `json:"kibanaSpecs,omitempty" xml:"kibanaSpecs,omitempty" type:"Repeated"`
	MasterAmountRange     []*string                                                         `json:"masterAmountRange,omitempty" xml:"masterAmountRange,omitempty" type:"Repeated"`
	MasterDiskList        []*GetRegionalInstanceConfigResponseBodyResultMasterDiskList      `json:"masterDiskList,omitempty" xml:"masterDiskList,omitempty" type:"Repeated"`
	MasterSpecs           []*string                                                         `json:"masterSpecs,omitempty" xml:"masterSpecs,omitempty" type:"Repeated"`
	SpecInfoMap           map[string]*ResultSpecInfoMapValue                                `json:"specInfoMap,omitempty" xml:"specInfoMap,omitempty"`
	Versions              []*string                                                         `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
	WarmNodeAmountRange   *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange   `json:"warmNodeAmountRange,omitempty" xml:"warmNodeAmountRange,omitempty" type:"Struct"`
	WarmNodeDiskList      []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList    `json:"warmNodeDiskList,omitempty" xml:"warmNodeDiskList,omitempty" type:"Repeated"`
	WarmNodeSpecs         []*string                                                         `json:"warmNodeSpecs,omitempty" xml:"warmNodeSpecs,omitempty" type:"Repeated"`
}

func (s GetRegionalInstanceConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetClientNodeAmountRange() *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange {
	return s.ClientNodeAmountRange
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetClientNodeDiskList() []*GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList {
	return s.ClientNodeDiskList
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetClientSpecs() []*string {
	return s.ClientSpecs
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetDataNodeAmountRange() *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange {
	return s.DataNodeAmountRange
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetDataNodeDiskList() []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	return s.DataNodeDiskList
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetDataNodeSpecs() []*string {
	return s.DataNodeSpecs
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetKibanaSpecs() []*string {
	return s.KibanaSpecs
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetMasterAmountRange() []*string {
	return s.MasterAmountRange
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetMasterDiskList() []*GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	return s.MasterDiskList
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetMasterSpecs() []*string {
	return s.MasterSpecs
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetSpecInfoMap() map[string]*ResultSpecInfoMapValue {
	return s.SpecInfoMap
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetVersions() []*string {
	return s.Versions
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetWarmNodeAmountRange() *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange {
	return s.WarmNodeAmountRange
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetWarmNodeDiskList() []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	return s.WarmNodeDiskList
}

func (s *GetRegionalInstanceConfigResponseBodyResult) GetWarmNodeSpecs() []*string {
	return s.WarmNodeSpecs
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetClientNodeAmountRange(v *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) *GetRegionalInstanceConfigResponseBodyResult {
	s.ClientNodeAmountRange = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetClientNodeDiskList(v []*GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) *GetRegionalInstanceConfigResponseBodyResult {
	s.ClientNodeDiskList = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetClientSpecs(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.ClientSpecs = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetDataNodeAmountRange(v *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) *GetRegionalInstanceConfigResponseBodyResult {
	s.DataNodeAmountRange = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetDataNodeDiskList(v []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) *GetRegionalInstanceConfigResponseBodyResult {
	s.DataNodeDiskList = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetDataNodeSpecs(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.DataNodeSpecs = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetKibanaSpecs(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.KibanaSpecs = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetMasterAmountRange(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.MasterAmountRange = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetMasterDiskList(v []*GetRegionalInstanceConfigResponseBodyResultMasterDiskList) *GetRegionalInstanceConfigResponseBodyResult {
	s.MasterDiskList = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetMasterSpecs(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.MasterSpecs = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetSpecInfoMap(v map[string]*ResultSpecInfoMapValue) *GetRegionalInstanceConfigResponseBodyResult {
	s.SpecInfoMap = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetVersions(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.Versions = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetWarmNodeAmountRange(v *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) *GetRegionalInstanceConfigResponseBodyResult {
	s.WarmNodeAmountRange = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetWarmNodeDiskList(v []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) *GetRegionalInstanceConfigResponseBodyResult {
	s.WarmNodeDiskList = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) SetWarmNodeSpecs(v []*string) *GetRegionalInstanceConfigResponseBodyResult {
	s.WarmNodeSpecs = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) SetMaxAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) SetMinAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList struct {
	DiskType   *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize    *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize    *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit *int32  `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) SetDiskType(v string) *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) SetScaleLimit(v int32) *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultClientNodeDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) SetMaxAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) SetMinAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList struct {
	DiskType                  *string                                                                                 `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize                   *int32                                                                                  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize                   *int32                                                                                  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit                *int32                                                                                  `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	SubClassificationConfines []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines `json:"subClassificationConfines,omitempty" xml:"subClassificationConfines,omitempty" type:"Repeated"`
	ValueLimitSet             []*int32                                                                                `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetSubClassificationConfines() []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines {
	return s.SubClassificationConfines
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) GetValueLimitSet() []*int32 {
	return s.ValueLimitSet
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetDiskType(v string) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetScaleLimit(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetSubClassificationConfines(v []*GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.SubClassificationConfines = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) SetValueLimitSet(v []*int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList {
	s.ValueLimitSet = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines struct {
	MaxSize          *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize          *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) SetPerformanceLevel(v string) *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines {
	s.PerformanceLevel = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultDataNodeDiskListSubClassificationConfines) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultMasterDiskList struct {
	DiskType                  *string                                                                               `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize                   *int32                                                                                `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize                   *int32                                                                                `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit                *int32                                                                                `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	SubClassificationConfines []*GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines `json:"subClassificationConfines,omitempty" xml:"subClassificationConfines,omitempty" type:"Repeated"`
}

func (s GetRegionalInstanceConfigResponseBodyResultMasterDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) GetSubClassificationConfines() []*GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines {
	return s.SubClassificationConfines
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) SetDiskType(v string) *GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) SetScaleLimit(v int32) *GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) SetSubClassificationConfines(v []*GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) *GetRegionalInstanceConfigResponseBodyResultMasterDiskList {
	s.SubClassificationConfines = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines struct {
	MaxSize          *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize          *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) SetPerformanceLevel(v string) *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines {
	s.PerformanceLevel = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultMasterDiskListSubClassificationConfines) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) SetMaxAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) SetMinAmount(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList struct {
	DiskType                  *string                                                                                 `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize                   *int32                                                                                  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize                   *int32                                                                                  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit                *int32                                                                                  `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	SubClassificationConfines []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines `json:"subClassificationConfines,omitempty" xml:"subClassificationConfines,omitempty" type:"Repeated"`
	ValueLimitSet             []*int32                                                                                `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetSubClassificationConfines() []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines {
	return s.SubClassificationConfines
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) GetValueLimitSet() []*int32 {
	return s.ValueLimitSet
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetDiskType(v string) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetScaleLimit(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetSubClassificationConfines(v []*GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.SubClassificationConfines = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) SetValueLimitSet(v []*int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList {
	s.ValueLimitSet = v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines struct {
	MaxSize          *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize          *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) SetMaxSize(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines {
	s.MaxSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) SetMinSize(v int32) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines {
	s.MinSize = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) SetPerformanceLevel(v string) *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines {
	s.PerformanceLevel = &v
	return s
}

func (s *GetRegionalInstanceConfigResponseBodyResultWarmNodeDiskListSubClassificationConfines) Validate() error {
	return dara.Validate(s)
}
