// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRatioStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstallRatio(v *GetClientRatioStatisticResponseBodyClientInstallRatio) *GetClientRatioStatisticResponseBody
	GetClientInstallRatio() *GetClientRatioStatisticResponseBodyClientInstallRatio
	SetClientOnlineRatio(v *GetClientRatioStatisticResponseBodyClientOnlineRatio) *GetClientRatioStatisticResponseBody
	GetClientOnlineRatio() *GetClientRatioStatisticResponseBodyClientOnlineRatio
	SetDates(v []*int64) *GetClientRatioStatisticResponseBody
	GetDates() []*int64
	SetRequestId(v string) *GetClientRatioStatisticResponseBody
	GetRequestId() *string
}

type GetClientRatioStatisticResponseBody struct {
	// The statistics on the client installation rate.
	ClientInstallRatio *GetClientRatioStatisticResponseBodyClientInstallRatio `json:"ClientInstallRatio,omitempty" xml:"ClientInstallRatio,omitempty" type:"Struct"`
	// The statistics on the client online rate.
	ClientOnlineRatio *GetClientRatioStatisticResponseBodyClientOnlineRatio `json:"ClientOnlineRatio,omitempty" xml:"ClientOnlineRatio,omitempty" type:"Struct"`
	// The list of time when statistics were collected.
	Dates []*int64 `json:"Dates,omitempty" xml:"Dates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientRatioStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBody) GetClientInstallRatio() *GetClientRatioStatisticResponseBodyClientInstallRatio {
	return s.ClientInstallRatio
}

func (s *GetClientRatioStatisticResponseBody) GetClientOnlineRatio() *GetClientRatioStatisticResponseBodyClientOnlineRatio {
	return s.ClientOnlineRatio
}

func (s *GetClientRatioStatisticResponseBody) GetDates() []*int64 {
	return s.Dates
}

func (s *GetClientRatioStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientRatioStatisticResponseBody) SetClientInstallRatio(v *GetClientRatioStatisticResponseBodyClientInstallRatio) *GetClientRatioStatisticResponseBody {
	s.ClientInstallRatio = v
	return s
}

func (s *GetClientRatioStatisticResponseBody) SetClientOnlineRatio(v *GetClientRatioStatisticResponseBodyClientOnlineRatio) *GetClientRatioStatisticResponseBody {
	s.ClientOnlineRatio = v
	return s
}

func (s *GetClientRatioStatisticResponseBody) SetDates(v []*int64) *GetClientRatioStatisticResponseBody {
	s.Dates = v
	return s
}

func (s *GetClientRatioStatisticResponseBody) SetRequestId(v string) *GetClientRatioStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientRatioStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientInstallRatio struct {
	// The list of current statistics on the installation rate of the client.
	CurrentItems []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems `json:"CurrentItems,omitempty" xml:"CurrentItems,omitempty" type:"Repeated"`
	// The list of historical statistics on the installation rate of the client.
	HistoryItems []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatio) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatio) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatio) GetCurrentItems() []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems {
	return s.CurrentItems
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatio) GetHistoryItems() []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems {
	return s.HistoryItems
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatio) SetCurrentItems(v []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) *GetClientRatioStatisticResponseBodyClientInstallRatio {
	s.CurrentItems = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatio) SetHistoryItems(v []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) *GetClientRatioStatisticResponseBodyClientInstallRatio {
	s.HistoryItems = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatio) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems struct {
	// The list of the statistics on the installation rate of the client by vendor.
	Items []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 0
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) GetItems() []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems {
	return s.Items
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) GetVendor() *int64 {
	return s.Vendor
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) SetItems(v []*GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems {
	s.Items = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) SetVendor(v int64) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems {
	s.Vendor = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems struct {
	// The total number of assets.
	//
	// example:
	//
	// 100
	AssetTotalCount *int32 `json:"AssetTotalCount,omitempty" xml:"AssetTotalCount,omitempty"`
	// The timestamp of the calculation. Unit: milliseconds.
	//
	// example:
	//
	// 1687759630045
	CalculateTime *int64 `json:"CalculateTime,omitempty" xml:"CalculateTime,omitempty"`
	// The installation rate. Unit: %.
	//
	// example:
	//
	// 70.00
	InstallRatio *float64 `json:"InstallRatio,omitempty" xml:"InstallRatio,omitempty"`
	// The number of assets on which the client is installed.
	//
	// example:
	//
	// 70
	InstalledAssetCount *int32 `json:"InstalledAssetCount,omitempty" xml:"InstalledAssetCount,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) GetAssetTotalCount() *int32 {
	return s.AssetTotalCount
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) GetCalculateTime() *int64 {
	return s.CalculateTime
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) GetInstallRatio() *float64 {
	return s.InstallRatio
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) GetInstalledAssetCount() *int32 {
	return s.InstalledAssetCount
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) SetAssetTotalCount(v int32) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems {
	s.AssetTotalCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) SetCalculateTime(v int64) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems {
	s.CalculateTime = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) SetInstallRatio(v float64) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems {
	s.InstallRatio = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) SetInstalledAssetCount(v int32) *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems {
	s.InstalledAssetCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioCurrentItemsItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems struct {
	// The list of statistics on the client installation rate.
	Items []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a simple application server
	//
	// example:
	//
	// 3
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) GetItems() []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems {
	return s.Items
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) GetVendor() *int64 {
	return s.Vendor
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) SetItems(v []*GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems {
	s.Items = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) SetVendor(v int64) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems {
	s.Vendor = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems struct {
	// The total number of assets.
	//
	// example:
	//
	// 100
	AssetTotalCount *int32 `json:"AssetTotalCount,omitempty" xml:"AssetTotalCount,omitempty"`
	// The timestamp of the calculation. Unit: milliseconds.
	//
	// example:
	//
	// 1687759630045
	CalculateTime *int64 `json:"CalculateTime,omitempty" xml:"CalculateTime,omitempty"`
	// The installation rate. Unit: %.
	//
	// example:
	//
	// 80.00
	InstallRatio *float64 `json:"InstallRatio,omitempty" xml:"InstallRatio,omitempty"`
	// The number of assets on which the client is installed.
	//
	// example:
	//
	// 80
	InstalledAssetCount *int32 `json:"InstalledAssetCount,omitempty" xml:"InstalledAssetCount,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) GetAssetTotalCount() *int32 {
	return s.AssetTotalCount
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) GetCalculateTime() *int64 {
	return s.CalculateTime
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) GetInstallRatio() *float64 {
	return s.InstallRatio
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) GetInstalledAssetCount() *int32 {
	return s.InstalledAssetCount
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) SetAssetTotalCount(v int32) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems {
	s.AssetTotalCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) SetCalculateTime(v int64) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems {
	s.CalculateTime = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) SetInstallRatio(v float64) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems {
	s.InstallRatio = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) SetInstalledAssetCount(v int32) *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems {
	s.InstalledAssetCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientInstallRatioHistoryItemsItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientOnlineRatio struct {
	// The list of current statistics on the online rate of the client.
	CurrentItems []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems `json:"CurrentItems,omitempty" xml:"CurrentItems,omitempty" type:"Repeated"`
	// The list of historical statistics on the online rate of the client.
	HistoryItems []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatio) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatio) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatio) GetCurrentItems() []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems {
	return s.CurrentItems
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatio) GetHistoryItems() []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems {
	return s.HistoryItems
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatio) SetCurrentItems(v []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) *GetClientRatioStatisticResponseBodyClientOnlineRatio {
	s.CurrentItems = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatio) SetHistoryItems(v []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) *GetClientRatioStatisticResponseBodyClientOnlineRatio {
	s.HistoryItems = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatio) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems struct {
	// The list of current statistics on the online rate of the client by vendor.
	Items []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a simple application server
	//
	// example:
	//
	// 3
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) GetItems() []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems {
	return s.Items
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) GetVendor() *int64 {
	return s.Vendor
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) SetItems(v []*GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems {
	s.Items = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) SetVendor(v int64) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems {
	s.Vendor = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems struct {
	// The number of assets on which the client is installed.
	//
	// example:
	//
	// 50
	AssetInstallCount *int32 `json:"AssetInstallCount,omitempty" xml:"AssetInstallCount,omitempty"`
	// The timestamp of the calculation. Unit: milliseconds.
	//
	// example:
	//
	// 1687759630045
	CalculateTime *int64 `json:"CalculateTime,omitempty" xml:"CalculateTime,omitempty"`
	// The number of online assets.
	//
	// example:
	//
	// 10
	OnlineAssetCount *int32 `json:"OnlineAssetCount,omitempty" xml:"OnlineAssetCount,omitempty"`
	// The online rate. Unit: %.
	//
	// example:
	//
	// 20.00
	OnlineRatio *float64 `json:"OnlineRatio,omitempty" xml:"OnlineRatio,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) GetAssetInstallCount() *int32 {
	return s.AssetInstallCount
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) GetCalculateTime() *int64 {
	return s.CalculateTime
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) GetOnlineAssetCount() *int32 {
	return s.OnlineAssetCount
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) GetOnlineRatio() *float64 {
	return s.OnlineRatio
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) SetAssetInstallCount(v int32) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems {
	s.AssetInstallCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) SetCalculateTime(v int64) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems {
	s.CalculateTime = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) SetOnlineAssetCount(v int32) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems {
	s.OnlineAssetCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) SetOnlineRatio(v float64) *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems {
	s.OnlineRatio = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioCurrentItemsItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems struct {
	// The list of historical statistics on the online rate of the client by vendor.
	Items []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 7
	Vendor *int64 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) GetItems() []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems {
	return s.Items
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) GetVendor() *int64 {
	return s.Vendor
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) SetItems(v []*GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems {
	s.Items = v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) SetVendor(v int64) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems {
	s.Vendor = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItems) Validate() error {
	return dara.Validate(s)
}

type GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems struct {
	// The number of assets on which the client is installed.
	//
	// example:
	//
	// 50
	AssetInstallCount *int32 `json:"AssetInstallCount,omitempty" xml:"AssetInstallCount,omitempty"`
	// The timestamp of the calculation. Unit: milliseconds.
	//
	// example:
	//
	// 1687759630045
	CalculateTime *int64 `json:"CalculateTime,omitempty" xml:"CalculateTime,omitempty"`
	// The number of online assets.
	//
	// example:
	//
	// 20
	OnlineAssetCount *int32 `json:"OnlineAssetCount,omitempty" xml:"OnlineAssetCount,omitempty"`
	// The online rate. Unit: %.
	//
	// example:
	//
	// 40.00
	OnlineRatio *float64 `json:"OnlineRatio,omitempty" xml:"OnlineRatio,omitempty"`
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) String() string {
	return dara.Prettify(s)
}

func (s GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) GoString() string {
	return s.String()
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) GetAssetInstallCount() *int32 {
	return s.AssetInstallCount
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) GetCalculateTime() *int64 {
	return s.CalculateTime
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) GetOnlineAssetCount() *int32 {
	return s.OnlineAssetCount
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) GetOnlineRatio() *float64 {
	return s.OnlineRatio
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) SetAssetInstallCount(v int32) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems {
	s.AssetInstallCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) SetCalculateTime(v int64) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems {
	s.CalculateTime = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) SetOnlineAssetCount(v int32) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems {
	s.OnlineAssetCount = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) SetOnlineRatio(v float64) *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems {
	s.OnlineRatio = &v
	return s
}

func (s *GetClientRatioStatisticResponseBodyClientOnlineRatioHistoryItemsItems) Validate() error {
	return dara.Validate(s)
}
