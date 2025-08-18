// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v *GetWafQuotaResponseBodyQuota) *GetWafQuotaResponseBody
	GetQuota() *GetWafQuotaResponseBodyQuota
	SetRequestId(v string) *GetWafQuotaResponseBody
	GetRequestId() *string
}

type GetWafQuotaResponseBody struct {
	// Returned quota information.
	Quota *GetWafQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBody) GetQuota() *GetWafQuotaResponseBodyQuota {
	return s.Quota
}

func (s *GetWafQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWafQuotaResponseBody) SetQuota(v *GetWafQuotaResponseBodyQuota) *GetWafQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetWafQuotaResponseBody) SetRequestId(v string) *GetWafQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWafQuotaResponseBodyQuota struct {
	// Quota information related to custom lists.
	List *GetWafQuotaResponseBodyQuotaList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// Quota information related to the WAF managed rules group.
	ManagedRulesGroup *GetWafQuotaResponseBodyQuotaManagedRulesGroup `json:"ManagedRulesGroup,omitempty" xml:"ManagedRulesGroup,omitempty" type:"Struct"`
	// Quota information related to custom response pages.
	Page *GetWafQuotaResponseBodyQuotaPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// Quota information related to scene protection.
	ScenePolicy *GetWafQuotaResponseBodyQuotaScenePolicy `json:"ScenePolicy,omitempty" xml:"ScenePolicy,omitempty" type:"Struct"`
}

func (s GetWafQuotaResponseBodyQuota) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuota) GetList() *GetWafQuotaResponseBodyQuotaList {
	return s.List
}

func (s *GetWafQuotaResponseBodyQuota) GetManagedRulesGroup() *GetWafQuotaResponseBodyQuotaManagedRulesGroup {
	return s.ManagedRulesGroup
}

func (s *GetWafQuotaResponseBodyQuota) GetPage() *GetWafQuotaResponseBodyQuotaPage {
	return s.Page
}

func (s *GetWafQuotaResponseBodyQuota) GetScenePolicy() *GetWafQuotaResponseBodyQuotaScenePolicy {
	return s.ScenePolicy
}

func (s *GetWafQuotaResponseBodyQuota) SetList(v *GetWafQuotaResponseBodyQuotaList) *GetWafQuotaResponseBodyQuota {
	s.List = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetManagedRulesGroup(v *GetWafQuotaResponseBodyQuotaManagedRulesGroup) *GetWafQuotaResponseBodyQuota {
	s.ManagedRulesGroup = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetPage(v *GetWafQuotaResponseBodyQuotaPage) *GetWafQuotaResponseBodyQuota {
	s.Page = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) SetScenePolicy(v *GetWafQuotaResponseBodyQuotaScenePolicy) *GetWafQuotaResponseBodyQuota {
	s.ScenePolicy = v
	return s
}

func (s *GetWafQuotaResponseBodyQuota) Validate() error {
	return dara.Validate(s)
}

type GetWafQuotaResponseBodyQuotaList struct {
	// Indicates whether the custom list is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// An object containing quota information for each type of item in the custom list.
	Items map[string]*QuotaListItemsValue `json:"Items,omitempty" xml:"Items,omitempty"`
	// The number quota allowed per custom list.
	NumberItemsPerList *WafQuotaInteger `json:"NumberItemsPerList,omitempty" xml:"NumberItemsPerList,omitempty"`
	// The total number quota allowed for items in all custom lists.
	NumberItemsTotal *WafQuotaInteger `json:"NumberItemsTotal,omitempty" xml:"NumberItemsTotal,omitempty"`
	// The total number quota allowed for custom lists.
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaList) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaList) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaList) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafQuotaResponseBodyQuotaList) GetItems() map[string]*QuotaListItemsValue {
	return s.Items
}

func (s *GetWafQuotaResponseBodyQuotaList) GetNumberItemsPerList() *WafQuotaInteger {
	return s.NumberItemsPerList
}

func (s *GetWafQuotaResponseBodyQuotaList) GetNumberItemsTotal() *WafQuotaInteger {
	return s.NumberItemsTotal
}

func (s *GetWafQuotaResponseBodyQuotaList) GetNumberTotal() *WafQuotaInteger {
	return s.NumberTotal
}

func (s *GetWafQuotaResponseBodyQuotaList) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaList {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetItems(v map[string]*QuotaListItemsValue) *GetWafQuotaResponseBodyQuotaList {
	s.Items = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberItemsPerList(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberItemsPerList = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberItemsTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberItemsTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaList {
	s.NumberTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaList) Validate() error {
	return dara.Validate(s)
}

type GetWafQuotaResponseBodyQuotaManagedRulesGroup struct {
	// Indicates whether the WAF managed rules group is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The total number quota allowed for the WAF managed rules group.
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaManagedRulesGroup) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaManagedRulesGroup) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) GetNumberTotal() *WafQuotaInteger {
	return s.NumberTotal
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaManagedRulesGroup {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaManagedRulesGroup {
	s.NumberTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaManagedRulesGroup) Validate() error {
	return dara.Validate(s)
}

type GetWafQuotaResponseBodyQuotaPage struct {
	// An object containing quota information for each Content-Type in custom response pages.
	ContentTypes map[string]*QuotaPageContentTypesValue `json:"ContentTypes,omitempty" xml:"ContentTypes,omitempty"`
	// Indicates whether the custom response page is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The total number quota allowed for custom response pages.
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaPage) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaPage) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaPage) GetContentTypes() map[string]*QuotaPageContentTypesValue {
	return s.ContentTypes
}

func (s *GetWafQuotaResponseBodyQuotaPage) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafQuotaResponseBodyQuotaPage) GetNumberTotal() *WafQuotaInteger {
	return s.NumberTotal
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetContentTypes(v map[string]*QuotaPageContentTypesValue) *GetWafQuotaResponseBodyQuotaPage {
	s.ContentTypes = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaPage {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaPage) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaPage {
	s.NumberTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaPage) Validate() error {
	return dara.Validate(s)
}

type GetWafQuotaResponseBodyQuotaScenePolicy struct {
	// Indicates whether the scene protection feature is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The total number quota for scene protection rules.
	NumberTotal *WafQuotaInteger `json:"NumberTotal,omitempty" xml:"NumberTotal,omitempty"`
}

func (s GetWafQuotaResponseBodyQuotaScenePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponseBodyQuotaScenePolicy) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) GetNumberTotal() *WafQuotaInteger {
	return s.NumberTotal
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) SetEnable(v bool) *GetWafQuotaResponseBodyQuotaScenePolicy {
	s.Enable = &v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) SetNumberTotal(v *WafQuotaInteger) *GetWafQuotaResponseBodyQuotaScenePolicy {
	s.NumberTotal = v
	return s
}

func (s *GetWafQuotaResponseBodyQuotaScenePolicy) Validate() error {
	return dara.Validate(s)
}
