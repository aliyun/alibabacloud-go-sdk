// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPartialBuy(v int32) *GetAuthSummaryResponseBody
	GetAllowPartialBuy() *int32
	SetAllowUpgradePartialBuy(v int32) *GetAuthSummaryResponseBody
	GetAllowUpgradePartialBuy() *int32
	SetAllowUserUnbind(v int32) *GetAuthSummaryResponseBody
	GetAllowUserUnbind() *int32
	SetAutoBind(v int32) *GetAuthSummaryResponseBody
	GetAutoBind() *int32
	SetClusterNodeCheck(v int32) *GetAuthSummaryResponseBody
	GetClusterNodeCheck() *int32
	SetDefaultAuthToAll(v int32) *GetAuthSummaryResponseBody
	GetDefaultAuthToAll() *int32
	SetHasPreBindSetting(v bool) *GetAuthSummaryResponseBody
	GetHasPreBindSetting() *bool
	SetHighestVersion(v int32) *GetAuthSummaryResponseBody
	GetHighestVersion() *int32
	SetInvalidBindStatus(v string) *GetAuthSummaryResponseBody
	GetInvalidBindStatus() *string
	SetIsMultiVersion(v int32) *GetAuthSummaryResponseBody
	GetIsMultiVersion() *int32
	SetMachine(v *GetAuthSummaryResponseBodyMachine) *GetAuthSummaryResponseBody
	GetMachine() *GetAuthSummaryResponseBodyMachine
	SetPostPaidHighestVersion(v string) *GetAuthSummaryResponseBody
	GetPostPaidHighestVersion() *string
	SetPostPaidHostAutoBind(v string) *GetAuthSummaryResponseBody
	GetPostPaidHostAutoBind() *string
	SetPostPaidHostAutoBindVersion(v string) *GetAuthSummaryResponseBody
	GetPostPaidHostAutoBindVersion() *string
	SetPostPaidVersionSummary(v []*GetAuthSummaryResponseBodyPostPaidVersionSummary) *GetAuthSummaryResponseBody
	GetPostPaidVersionSummary() []*GetAuthSummaryResponseBodyPostPaidVersionSummary
	SetRequestId(v string) *GetAuthSummaryResponseBody
	GetRequestId() *string
	SetVersionSummary(v []*GetAuthSummaryResponseBodyVersionSummary) *GetAuthSummaryResponseBody
	GetVersionSummary() []*GetAuthSummaryResponseBodyVersionSummary
}

type GetAuthSummaryResponseBody struct {
	// Indicates whether you can purchase protection quota on demand when you purchase Security Center. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllowPartialBuy *int32 `json:"AllowPartialBuy,omitempty" xml:"AllowPartialBuy,omitempty"`
	// Indicates whether you can purchase protection quota on demand after an upgrade. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllowUpgradePartialBuy *int32 `json:"AllowUpgradePartialBuy,omitempty" xml:"AllowUpgradePartialBuy,omitempty"`
	// Indicates whether all bound assets can be immediately unbound. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllowUserUnbind *int32 `json:"AllowUserUnbind,omitempty" xml:"AllowUserUnbind,omitempty"`
	// Indicates whether automatic binding is enabled. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// Cluster node need to check the machine version,  Value:
	//
	// - **0*	- : Not required
	//
	// - **1*	- : Required
	//
	// example:
	//
	// Required
	ClusterNodeCheck *int32 `json:"ClusterNodeCheck,omitempty" xml:"ClusterNodeCheck,omitempty"`
	// Indicates whether the protection quota is supported for all assets. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	DefaultAuthToAll *int32 `json:"DefaultAuthToAll,omitempty" xml:"DefaultAuthToAll,omitempty"`
	// Indicates whether pre-bound assets exist. If you select assets to bind when you purchase Security Center, pre-bound assets exist. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	HasPreBindSetting *bool `json:"HasPreBindSetting,omitempty" xml:"HasPreBindSetting,omitempty"`
	// The most advanced edition that is used. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// >  If you purchase Security Center Multi-edition, the value indicates the most advanced edition that is used. Otherwise, the value indicates the specific edition that is purchased.
	//
	// example:
	//
	// 1
	HighestVersion *int32 `json:"HighestVersion,omitempty" xml:"HighestVersion,omitempty"`
	// Binding effective status, value:
	//
	// - **NORMAL*	- : Effective
	//
	// - **INVALID_NODE_VERSION**: Invalid
	//
	// example:
	//
	// Effective
	InvalidBindStatus *string `json:"InvalidBindStatus,omitempty" xml:"InvalidBindStatus,omitempty"`
	// Indicates whether Security Center Multi-edition is purchased. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	IsMultiVersion *int32 `json:"IsMultiVersion,omitempty" xml:"IsMultiVersion,omitempty"`
	// The statistics of the protection quota for assets.
	Machine *GetAuthSummaryResponseBodyMachine `json:"Machine,omitempty" xml:"Machine,omitempty" type:"Struct"`
	// Activate the pay-as-you-go service protection version for hosts and container security, which is the highest protection version among all bound hosts. Values:   - **1**: Free Edition  - **3**: Enterprise Edition - **5**: Advanced Edition - **6**: Antivirus Edition     - **7**: Flagship Edition
	//
	// example:
	//
	// 7
	PostPaidHighestVersion *string `json:"PostPaidHighestVersion,omitempty" xml:"PostPaidHighestVersion,omitempty"`
	// The pay-as-you-go service for host and container security adds an automatic binding identifier for new hosts, with values: - **0**: Off - **1**: On
	//
	// example:
	//
	// 1
	PostPaidHostAutoBind *string `json:"PostPaidHostAutoBind,omitempty" xml:"PostPaidHostAutoBind,omitempty"`
	// The version for the pay-as-you-go service of host and container security to automatically bind new assets, with values: - **1**: Free Edition - **3**: Enterprise Edition - **5**: Advanced Edition - **6**: Antivirus Edition - **7**: Flagship Edition
	//
	// example:
	//
	// 7
	PostPaidHostAutoBindVersion *string `json:"PostPaidHostAutoBindVersion,omitempty" xml:"PostPaidHostAutoBindVersion,omitempty"`
	// Statistics on pay-as-you-go service authorization for host and container security.
	PostPaidVersionSummary []*GetAuthSummaryResponseBodyPostPaidVersionSummary `json:"PostPaidVersionSummary,omitempty" xml:"PostPaidVersionSummary,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-***-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The quota consumption statistics.
	VersionSummary []*GetAuthSummaryResponseBodyVersionSummary `json:"VersionSummary,omitempty" xml:"VersionSummary,omitempty" type:"Repeated"`
}

func (s GetAuthSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthSummaryResponseBody) GetAllowPartialBuy() *int32 {
	return s.AllowPartialBuy
}

func (s *GetAuthSummaryResponseBody) GetAllowUpgradePartialBuy() *int32 {
	return s.AllowUpgradePartialBuy
}

func (s *GetAuthSummaryResponseBody) GetAllowUserUnbind() *int32 {
	return s.AllowUserUnbind
}

func (s *GetAuthSummaryResponseBody) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *GetAuthSummaryResponseBody) GetClusterNodeCheck() *int32 {
	return s.ClusterNodeCheck
}

func (s *GetAuthSummaryResponseBody) GetDefaultAuthToAll() *int32 {
	return s.DefaultAuthToAll
}

func (s *GetAuthSummaryResponseBody) GetHasPreBindSetting() *bool {
	return s.HasPreBindSetting
}

func (s *GetAuthSummaryResponseBody) GetHighestVersion() *int32 {
	return s.HighestVersion
}

func (s *GetAuthSummaryResponseBody) GetInvalidBindStatus() *string {
	return s.InvalidBindStatus
}

func (s *GetAuthSummaryResponseBody) GetIsMultiVersion() *int32 {
	return s.IsMultiVersion
}

func (s *GetAuthSummaryResponseBody) GetMachine() *GetAuthSummaryResponseBodyMachine {
	return s.Machine
}

func (s *GetAuthSummaryResponseBody) GetPostPaidHighestVersion() *string {
	return s.PostPaidHighestVersion
}

func (s *GetAuthSummaryResponseBody) GetPostPaidHostAutoBind() *string {
	return s.PostPaidHostAutoBind
}

func (s *GetAuthSummaryResponseBody) GetPostPaidHostAutoBindVersion() *string {
	return s.PostPaidHostAutoBindVersion
}

func (s *GetAuthSummaryResponseBody) GetPostPaidVersionSummary() []*GetAuthSummaryResponseBodyPostPaidVersionSummary {
	return s.PostPaidVersionSummary
}

func (s *GetAuthSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthSummaryResponseBody) GetVersionSummary() []*GetAuthSummaryResponseBodyVersionSummary {
	return s.VersionSummary
}

func (s *GetAuthSummaryResponseBody) SetAllowPartialBuy(v int32) *GetAuthSummaryResponseBody {
	s.AllowPartialBuy = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetAllowUpgradePartialBuy(v int32) *GetAuthSummaryResponseBody {
	s.AllowUpgradePartialBuy = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetAllowUserUnbind(v int32) *GetAuthSummaryResponseBody {
	s.AllowUserUnbind = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetAutoBind(v int32) *GetAuthSummaryResponseBody {
	s.AutoBind = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetClusterNodeCheck(v int32) *GetAuthSummaryResponseBody {
	s.ClusterNodeCheck = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetDefaultAuthToAll(v int32) *GetAuthSummaryResponseBody {
	s.DefaultAuthToAll = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetHasPreBindSetting(v bool) *GetAuthSummaryResponseBody {
	s.HasPreBindSetting = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetHighestVersion(v int32) *GetAuthSummaryResponseBody {
	s.HighestVersion = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetInvalidBindStatus(v string) *GetAuthSummaryResponseBody {
	s.InvalidBindStatus = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetIsMultiVersion(v int32) *GetAuthSummaryResponseBody {
	s.IsMultiVersion = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetMachine(v *GetAuthSummaryResponseBodyMachine) *GetAuthSummaryResponseBody {
	s.Machine = v
	return s
}

func (s *GetAuthSummaryResponseBody) SetPostPaidHighestVersion(v string) *GetAuthSummaryResponseBody {
	s.PostPaidHighestVersion = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetPostPaidHostAutoBind(v string) *GetAuthSummaryResponseBody {
	s.PostPaidHostAutoBind = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetPostPaidHostAutoBindVersion(v string) *GetAuthSummaryResponseBody {
	s.PostPaidHostAutoBindVersion = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetPostPaidVersionSummary(v []*GetAuthSummaryResponseBodyPostPaidVersionSummary) *GetAuthSummaryResponseBody {
	s.PostPaidVersionSummary = v
	return s
}

func (s *GetAuthSummaryResponseBody) SetRequestId(v string) *GetAuthSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthSummaryResponseBody) SetVersionSummary(v []*GetAuthSummaryResponseBodyVersionSummary) *GetAuthSummaryResponseBody {
	s.VersionSummary = v
	return s
}

func (s *GetAuthSummaryResponseBody) Validate() error {
	if s.Machine != nil {
		if err := s.Machine.Validate(); err != nil {
			return err
		}
	}
	if s.PostPaidVersionSummary != nil {
		for _, item := range s.PostPaidVersionSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VersionSummary != nil {
		for _, item := range s.VersionSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAuthSummaryResponseBodyMachine struct {
	// The number of cores of the assets that are bound to Security Center.
	//
	// example:
	//
	// 10
	BindCoreCount *int32 `json:"BindCoreCount,omitempty" xml:"BindCoreCount,omitempty"`
	// The number of the assets that are bound to Security Center.
	//
	// example:
	//
	// 10
	BindEcsCount *int32 `json:"BindEcsCount,omitempty" xml:"BindEcsCount,omitempty"`
	// Bind the number of cores for postpaid authorization assets.
	//
	// example:
	//
	// 10
	PostPaidBindCoreCount *int32 `json:"PostPaidBindCoreCount,omitempty" xml:"PostPaidBindCoreCount,omitempty"`
	// The number of assets bound to the postpaid authorization.
	//
	// example:
	//
	// 10
	PostPaidBindEcsCount *int32 `json:"PostPaidBindEcsCount,omitempty" xml:"PostPaidBindEcsCount,omitempty"`
	// The number of cores of the assets that are at risk.
	//
	// example:
	//
	// 10
	RiskCoreCount *int32 `json:"RiskCoreCount,omitempty" xml:"RiskCoreCount,omitempty"`
	// The number of the assets that are at risk.
	//
	// example:
	//
	// 10
	RiskEcsCount *int32 `json:"RiskEcsCount,omitempty" xml:"RiskEcsCount,omitempty"`
	// The total number of asset cores.
	//
	// example:
	//
	// 10
	TotalCoreCount *int32 `json:"TotalCoreCount,omitempty" xml:"TotalCoreCount,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 10
	TotalEcsCount *int32 `json:"TotalEcsCount,omitempty" xml:"TotalEcsCount,omitempty"`
	// The number of cores of unbound assets.
	//
	// example:
	//
	// 10
	UnBindCoreCount *int32 `json:"UnBindCoreCount,omitempty" xml:"UnBindCoreCount,omitempty"`
	// The number of unbound assets.
	//
	// example:
	//
	// 10
	UnBindEcsCount *int32 `json:"UnBindEcsCount,omitempty" xml:"UnBindEcsCount,omitempty"`
}

func (s GetAuthSummaryResponseBodyMachine) String() string {
	return dara.Prettify(s)
}

func (s GetAuthSummaryResponseBodyMachine) GoString() string {
	return s.String()
}

func (s *GetAuthSummaryResponseBodyMachine) GetBindCoreCount() *int32 {
	return s.BindCoreCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetBindEcsCount() *int32 {
	return s.BindEcsCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetPostPaidBindCoreCount() *int32 {
	return s.PostPaidBindCoreCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetPostPaidBindEcsCount() *int32 {
	return s.PostPaidBindEcsCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetRiskCoreCount() *int32 {
	return s.RiskCoreCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetRiskEcsCount() *int32 {
	return s.RiskEcsCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetTotalCoreCount() *int32 {
	return s.TotalCoreCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetTotalEcsCount() *int32 {
	return s.TotalEcsCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetUnBindCoreCount() *int32 {
	return s.UnBindCoreCount
}

func (s *GetAuthSummaryResponseBodyMachine) GetUnBindEcsCount() *int32 {
	return s.UnBindEcsCount
}

func (s *GetAuthSummaryResponseBodyMachine) SetBindCoreCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.BindCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetBindEcsCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.BindEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetPostPaidBindCoreCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.PostPaidBindCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetPostPaidBindEcsCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.PostPaidBindEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetRiskCoreCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.RiskCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetRiskEcsCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.RiskEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetTotalCoreCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.TotalCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetTotalEcsCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.TotalEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetUnBindCoreCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.UnBindCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) SetUnBindEcsCount(v int32) *GetAuthSummaryResponseBodyMachine {
	s.UnBindEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyMachine) Validate() error {
	return dara.Validate(s)
}

type GetAuthSummaryResponseBodyPostPaidVersionSummary struct {
	// The type of authorization consumed during binding, with values: - **ASSET**: Consumes the number of authorized devices - **CORE**: Consumes the number of authorized cores - **ASSET_AND_CORE**: Consumes both the number of authorized devices and cores.
	//
	// example:
	//
	// ASSET
	AuthBindType *string `json:"AuthBindType,omitempty" xml:"AuthBindType,omitempty"`
	// Current version index, the higher the number, the newer the version, used for sorting. Values: - **1**: Free Edition - **2**: Anti-virus Edition - **3**: Advanced Edition - **4**: Enterprise Edition - **5**: Flagship Edition
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Number of authorized cores used. > This parameter is valid when AuthBindType is set to CORE or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UsedCoreCount *int64 `json:"UsedCoreCount,omitempty" xml:"UsedCoreCount,omitempty"`
	// Number of authorized devices used. > This parameter is valid when AuthBindType is ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UsedEcsCount *int64 `json:"UsedEcsCount,omitempty" xml:"UsedEcsCount,omitempty"`
	// Bound host assets with postpaid versions, values:   - **1**: Free version  - **3**: Enterprise version - **5**: Advanced version - **6**: Anti-virus version     - **7**: Flagship version
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAuthSummaryResponseBodyPostPaidVersionSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAuthSummaryResponseBodyPostPaidVersionSummary) GoString() string {
	return s.String()
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) GetAuthBindType() *string {
	return s.AuthBindType
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) GetIndex() *int32 {
	return s.Index
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) GetUsedCoreCount() *int64 {
	return s.UsedCoreCount
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) GetUsedEcsCount() *int64 {
	return s.UsedEcsCount
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) GetVersion() *int32 {
	return s.Version
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) SetAuthBindType(v string) *GetAuthSummaryResponseBodyPostPaidVersionSummary {
	s.AuthBindType = &v
	return s
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) SetIndex(v int32) *GetAuthSummaryResponseBodyPostPaidVersionSummary {
	s.Index = &v
	return s
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) SetUsedCoreCount(v int64) *GetAuthSummaryResponseBodyPostPaidVersionSummary {
	s.UsedCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) SetUsedEcsCount(v int64) *GetAuthSummaryResponseBodyPostPaidVersionSummary {
	s.UsedEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) SetVersion(v int32) *GetAuthSummaryResponseBodyPostPaidVersionSummary {
	s.Version = &v
	return s
}

func (s *GetAuthSummaryResponseBodyPostPaidVersionSummary) Validate() error {
	return dara.Validate(s)
}

type GetAuthSummaryResponseBodyVersionSummary struct {
	// The type of the quota that is consumed. Valid values:
	//
	// 	- ASSET: quota of servers.
	//
	// 	- CORE: quota of server cores.
	//
	// 	- ASSET_AND_CORE: both.
	//
	// example:
	//
	// ASSET
	AuthBindType *string `json:"AuthBindType,omitempty" xml:"AuthBindType,omitempty"`
	// The index of the current edition. The smaller the value, the higher the edition. The index is used for sorting.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The total quota of server cores.
	//
	// >  This parameter takes effect only if AuthBindType is set to CORE or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	TotalCoreAuthCount *int32 `json:"TotalCoreAuthCount,omitempty" xml:"TotalCoreAuthCount,omitempty"`
	// The total quota of servers in the current edition.
	//
	// >  This parameter takes effect only if AuthBindType is set to ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total quota of servers.
	//
	// >  This parameter takes effect only if AuthBindType is set to ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	TotalEcsAuthCount *int32 `json:"TotalEcsAuthCount,omitempty" xml:"TotalEcsAuthCount,omitempty"`
	// The remaining quota of servers.
	//
	// >  This parameter takes effect only if AuthBindType is set to ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UnUsedCount *int32 `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	// The remaining quota of server cores.
	//
	// >  This parameter takes effect only if AuthBindType is set to CORE or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UnusedCoreAuthCount *int32 `json:"UnusedCoreAuthCount,omitempty" xml:"UnusedCoreAuthCount,omitempty"`
	// The remaining quota of servers.
	//
	// >  This parameter takes effect only if AuthBindType is set to ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UnusedEcsAuthCount *int32 `json:"UnusedEcsAuthCount,omitempty" xml:"UnusedEcsAuthCount,omitempty"`
	// The consumed quota of server cores.
	//
	// >  This parameter takes effect only if AuthBindType is set to CORE or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UsedCoreCount *int32 `json:"UsedCoreCount,omitempty" xml:"UsedCoreCount,omitempty"`
	// The used quota of servers.
	//
	// >  This parameter takes effect only if AuthBindType is set to ASSET or ASSET_AND_CORE.
	//
	// example:
	//
	// 10
	UsedEcsCount *int32 `json:"UsedEcsCount,omitempty" xml:"UsedEcsCount,omitempty"`
	// The edition of purchased Security Center. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **8**: Multi-edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAuthSummaryResponseBodyVersionSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAuthSummaryResponseBodyVersionSummary) GoString() string {
	return s.String()
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetAuthBindType() *string {
	return s.AuthBindType
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetIndex() *int32 {
	return s.Index
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetTotalCoreAuthCount() *int32 {
	return s.TotalCoreAuthCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetTotalEcsAuthCount() *int32 {
	return s.TotalEcsAuthCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetUnUsedCount() *int32 {
	return s.UnUsedCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetUnusedCoreAuthCount() *int32 {
	return s.UnusedCoreAuthCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetUnusedEcsAuthCount() *int32 {
	return s.UnusedEcsAuthCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetUsedCoreCount() *int32 {
	return s.UsedCoreCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetUsedEcsCount() *int32 {
	return s.UsedEcsCount
}

func (s *GetAuthSummaryResponseBodyVersionSummary) GetVersion() *int32 {
	return s.Version
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetAuthBindType(v string) *GetAuthSummaryResponseBodyVersionSummary {
	s.AuthBindType = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetIndex(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.Index = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetTotalCoreAuthCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.TotalCoreAuthCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetTotalCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.TotalCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetTotalEcsAuthCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.TotalEcsAuthCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetUnUsedCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.UnUsedCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetUnusedCoreAuthCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.UnusedCoreAuthCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetUnusedEcsAuthCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.UnusedEcsAuthCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetUsedCoreCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.UsedCoreCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetUsedEcsCount(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.UsedEcsCount = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) SetVersion(v int32) *GetAuthSummaryResponseBodyVersionSummary {
	s.Version = &v
	return s
}

func (s *GetAuthSummaryResponseBodyVersionSummary) Validate() error {
	return dara.Validate(s)
}
