// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceEstimateOutput interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*PriceEstimateOutputApps) *PriceEstimateOutput
	GetApps() []*PriceEstimateOutputApps
	SetItems(v []*PriceEstimateOutputItems) *PriceEstimateOutput
	GetItems() []*PriceEstimateOutputItems
	SetPostPayItems(v []*PriceEstimateOutputPostPayItems) *PriceEstimateOutput
	GetPostPayItems() []*PriceEstimateOutputPostPayItems
	SetPostPayTotalPrice(v float32) *PriceEstimateOutput
	GetPostPayTotalPrice() *float32
	SetTotalPrice(v float32) *PriceEstimateOutput
	GetTotalPrice() *float32
}

type PriceEstimateOutput struct {
	Apps         []*PriceEstimateOutputApps         `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	Items        []*PriceEstimateOutputItems        `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PostPayItems []*PriceEstimateOutputPostPayItems `json:"PostPayItems,omitempty" xml:"PostPayItems,omitempty" type:"Repeated"`
	// example:
	//
	// 235.66
	PostPayTotalPrice *float32 `json:"PostPayTotalPrice,omitempty" xml:"PostPayTotalPrice,omitempty"`
	// example:
	//
	// 235.66
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s PriceEstimateOutput) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutput) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutput) GetApps() []*PriceEstimateOutputApps {
	return s.Apps
}

func (s *PriceEstimateOutput) GetItems() []*PriceEstimateOutputItems {
	return s.Items
}

func (s *PriceEstimateOutput) GetPostPayItems() []*PriceEstimateOutputPostPayItems {
	return s.PostPayItems
}

func (s *PriceEstimateOutput) GetPostPayTotalPrice() *float32 {
	return s.PostPayTotalPrice
}

func (s *PriceEstimateOutput) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *PriceEstimateOutput) SetApps(v []*PriceEstimateOutputApps) *PriceEstimateOutput {
	s.Apps = v
	return s
}

func (s *PriceEstimateOutput) SetItems(v []*PriceEstimateOutputItems) *PriceEstimateOutput {
	s.Items = v
	return s
}

func (s *PriceEstimateOutput) SetPostPayItems(v []*PriceEstimateOutputPostPayItems) *PriceEstimateOutput {
	s.PostPayItems = v
	return s
}

func (s *PriceEstimateOutput) SetPostPayTotalPrice(v float32) *PriceEstimateOutput {
	s.PostPayTotalPrice = &v
	return s
}

func (s *PriceEstimateOutput) SetTotalPrice(v float32) *PriceEstimateOutput {
	s.TotalPrice = &v
	return s
}

func (s *PriceEstimateOutput) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputApps struct {
	Feature *PriceEstimateFeature `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// example:
	//
	// 1
	Id     *int64                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Usages []*PriceEstimateOutputAppsUsages `json:"Usages,omitempty" xml:"Usages,omitempty" type:"Repeated"`
}

func (s PriceEstimateOutputApps) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputApps) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputApps) GetFeature() *PriceEstimateFeature {
	return s.Feature
}

func (s *PriceEstimateOutputApps) GetId() *int64 {
	return s.Id
}

func (s *PriceEstimateOutputApps) GetUsages() []*PriceEstimateOutputAppsUsages {
	return s.Usages
}

func (s *PriceEstimateOutputApps) SetFeature(v *PriceEstimateFeature) *PriceEstimateOutputApps {
	s.Feature = v
	return s
}

func (s *PriceEstimateOutputApps) SetId(v int64) *PriceEstimateOutputApps {
	s.Id = &v
	return s
}

func (s *PriceEstimateOutputApps) SetUsages(v []*PriceEstimateOutputAppsUsages) *PriceEstimateOutputApps {
	s.Usages = v
	return s
}

func (s *PriceEstimateOutputApps) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputAppsUsages struct {
	// example:
	//
	// 3600.00
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// Microservice_cpuUsage
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 核*秒
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s PriceEstimateOutputAppsUsages) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputAppsUsages) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputAppsUsages) GetAmount() *float32 {
	return s.Amount
}

func (s *PriceEstimateOutputAppsUsages) GetId() *string {
	return s.Id
}

func (s *PriceEstimateOutputAppsUsages) GetUnit() *string {
	return s.Unit
}

func (s *PriceEstimateOutputAppsUsages) SetAmount(v float32) *PriceEstimateOutputAppsUsages {
	s.Amount = &v
	return s
}

func (s *PriceEstimateOutputAppsUsages) SetId(v string) *PriceEstimateOutputAppsUsages {
	s.Id = &v
	return s
}

func (s *PriceEstimateOutputAppsUsages) SetUnit(v string) *PriceEstimateOutputAppsUsages {
	s.Unit = &v
	return s
}

func (s *PriceEstimateOutputAppsUsages) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputItems struct {
	// example:
	//
	// 3600.00
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// p_micro_service_cpu
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1.00
	Price *float32                         `json:"Price,omitempty" xml:"Price,omitempty"`
	Steps []*PriceEstimateOutputItemsSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// example:
	//
	// pack/post
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 核*秒
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s PriceEstimateOutputItems) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputItems) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputItems) GetAmount() *float32 {
	return s.Amount
}

func (s *PriceEstimateOutputItems) GetCount() *int64 {
	return s.Count
}

func (s *PriceEstimateOutputItems) GetId() *string {
	return s.Id
}

func (s *PriceEstimateOutputItems) GetPrice() *float32 {
	return s.Price
}

func (s *PriceEstimateOutputItems) GetSteps() []*PriceEstimateOutputItemsSteps {
	return s.Steps
}

func (s *PriceEstimateOutputItems) GetType() *string {
	return s.Type
}

func (s *PriceEstimateOutputItems) GetUnit() *string {
	return s.Unit
}

func (s *PriceEstimateOutputItems) SetAmount(v float32) *PriceEstimateOutputItems {
	s.Amount = &v
	return s
}

func (s *PriceEstimateOutputItems) SetCount(v int64) *PriceEstimateOutputItems {
	s.Count = &v
	return s
}

func (s *PriceEstimateOutputItems) SetId(v string) *PriceEstimateOutputItems {
	s.Id = &v
	return s
}

func (s *PriceEstimateOutputItems) SetPrice(v float32) *PriceEstimateOutputItems {
	s.Price = &v
	return s
}

func (s *PriceEstimateOutputItems) SetSteps(v []*PriceEstimateOutputItemsSteps) *PriceEstimateOutputItems {
	s.Steps = v
	return s
}

func (s *PriceEstimateOutputItems) SetType(v string) *PriceEstimateOutputItems {
	s.Type = &v
	return s
}

func (s *PriceEstimateOutputItems) SetUnit(v string) *PriceEstimateOutputItems {
	s.Unit = &v
	return s
}

func (s *PriceEstimateOutputItems) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputItemsSteps struct {
	// example:
	//
	// 0
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 10000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 0.0001
	Price     *float32  `json:"Price,omitempty" xml:"Price,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 核*秒
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s PriceEstimateOutputItemsSteps) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputItemsSteps) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputItemsSteps) GetBegin() *int64 {
	return s.Begin
}

func (s *PriceEstimateOutputItemsSteps) GetEnd() *int64 {
	return s.End
}

func (s *PriceEstimateOutputItemsSteps) GetPrice() *float32 {
	return s.Price
}

func (s *PriceEstimateOutputItemsSteps) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *PriceEstimateOutputItemsSteps) GetUnit() *string {
	return s.Unit
}

func (s *PriceEstimateOutputItemsSteps) SetBegin(v int64) *PriceEstimateOutputItemsSteps {
	s.Begin = &v
	return s
}

func (s *PriceEstimateOutputItemsSteps) SetEnd(v int64) *PriceEstimateOutputItemsSteps {
	s.End = &v
	return s
}

func (s *PriceEstimateOutputItemsSteps) SetPrice(v float32) *PriceEstimateOutputItemsSteps {
	s.Price = &v
	return s
}

func (s *PriceEstimateOutputItemsSteps) SetRegionIds(v []*string) *PriceEstimateOutputItemsSteps {
	s.RegionIds = v
	return s
}

func (s *PriceEstimateOutputItemsSteps) SetUnit(v string) *PriceEstimateOutputItemsSteps {
	s.Unit = &v
	return s
}

func (s *PriceEstimateOutputItemsSteps) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputPostPayItems struct {
	// example:
	//
	// 3600.00
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// p_micro_service_cpu
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1.00
	Price *float32                                `json:"Price,omitempty" xml:"Price,omitempty"`
	Steps []*PriceEstimateOutputPostPayItemsSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// example:
	//
	// pack/post
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 核*秒
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s PriceEstimateOutputPostPayItems) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputPostPayItems) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputPostPayItems) GetAmount() *float32 {
	return s.Amount
}

func (s *PriceEstimateOutputPostPayItems) GetCount() *int64 {
	return s.Count
}

func (s *PriceEstimateOutputPostPayItems) GetId() *string {
	return s.Id
}

func (s *PriceEstimateOutputPostPayItems) GetPrice() *float32 {
	return s.Price
}

func (s *PriceEstimateOutputPostPayItems) GetSteps() []*PriceEstimateOutputPostPayItemsSteps {
	return s.Steps
}

func (s *PriceEstimateOutputPostPayItems) GetType() *string {
	return s.Type
}

func (s *PriceEstimateOutputPostPayItems) GetUnit() *string {
	return s.Unit
}

func (s *PriceEstimateOutputPostPayItems) SetAmount(v float32) *PriceEstimateOutputPostPayItems {
	s.Amount = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetCount(v int64) *PriceEstimateOutputPostPayItems {
	s.Count = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetId(v string) *PriceEstimateOutputPostPayItems {
	s.Id = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetPrice(v float32) *PriceEstimateOutputPostPayItems {
	s.Price = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetSteps(v []*PriceEstimateOutputPostPayItemsSteps) *PriceEstimateOutputPostPayItems {
	s.Steps = v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetType(v string) *PriceEstimateOutputPostPayItems {
	s.Type = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) SetUnit(v string) *PriceEstimateOutputPostPayItems {
	s.Unit = &v
	return s
}

func (s *PriceEstimateOutputPostPayItems) Validate() error {
	return dara.Validate(s)
}

type PriceEstimateOutputPostPayItemsSteps struct {
	// example:
	//
	// 0
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 10000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 0.0001
	Price     *float32  `json:"Price,omitempty" xml:"Price,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 核*秒
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s PriceEstimateOutputPostPayItemsSteps) String() string {
	return dara.Prettify(s)
}

func (s PriceEstimateOutputPostPayItemsSteps) GoString() string {
	return s.String()
}

func (s *PriceEstimateOutputPostPayItemsSteps) GetBegin() *int64 {
	return s.Begin
}

func (s *PriceEstimateOutputPostPayItemsSteps) GetEnd() *int64 {
	return s.End
}

func (s *PriceEstimateOutputPostPayItemsSteps) GetPrice() *float32 {
	return s.Price
}

func (s *PriceEstimateOutputPostPayItemsSteps) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *PriceEstimateOutputPostPayItemsSteps) GetUnit() *string {
	return s.Unit
}

func (s *PriceEstimateOutputPostPayItemsSteps) SetBegin(v int64) *PriceEstimateOutputPostPayItemsSteps {
	s.Begin = &v
	return s
}

func (s *PriceEstimateOutputPostPayItemsSteps) SetEnd(v int64) *PriceEstimateOutputPostPayItemsSteps {
	s.End = &v
	return s
}

func (s *PriceEstimateOutputPostPayItemsSteps) SetPrice(v float32) *PriceEstimateOutputPostPayItemsSteps {
	s.Price = &v
	return s
}

func (s *PriceEstimateOutputPostPayItemsSteps) SetRegionIds(v []*string) *PriceEstimateOutputPostPayItemsSteps {
	s.RegionIds = v
	return s
}

func (s *PriceEstimateOutputPostPayItemsSteps) SetUnit(v string) *PriceEstimateOutputPostPayItemsSteps {
	s.Unit = &v
	return s
}

func (s *PriceEstimateOutputPostPayItemsSteps) Validate() error {
	return dara.Validate(s)
}
