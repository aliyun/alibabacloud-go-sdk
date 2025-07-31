// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizUnitInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizUnitInfo(v *GetBizUnitInfoResponseBodyBizUnitInfo) *GetBizUnitInfoResponseBody
	GetBizUnitInfo() *GetBizUnitInfoResponseBodyBizUnitInfo
	SetCode(v string) *GetBizUnitInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBizUnitInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBizUnitInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBizUnitInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBizUnitInfoResponseBody
	GetSuccess() *bool
}

type GetBizUnitInfoResponseBody struct {
	BizUnitInfo *GetBizUnitInfoResponseBodyBizUnitInfo `json:"BizUnitInfo,omitempty" xml:"BizUnitInfo,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBizUnitInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBody) GetBizUnitInfo() *GetBizUnitInfoResponseBodyBizUnitInfo {
	return s.BizUnitInfo
}

func (s *GetBizUnitInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBizUnitInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBizUnitInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBizUnitInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBizUnitInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBizUnitInfoResponseBody) SetBizUnitInfo(v *GetBizUnitInfoResponseBodyBizUnitInfo) *GetBizUnitInfoResponseBody {
	s.BizUnitInfo = v
	return s
}

func (s *GetBizUnitInfoResponseBody) SetCode(v string) *GetBizUnitInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetBizUnitInfoResponseBody) SetHttpStatusCode(v int32) *GetBizUnitInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBizUnitInfoResponseBody) SetMessage(v string) *GetBizUnitInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetBizUnitInfoResponseBody) SetRequestId(v string) *GetBizUnitInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBizUnitInfoResponseBody) SetSuccess(v bool) *GetBizUnitInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBizUnitInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBizUnitInfoResponseBodyBizUnitInfo struct {
	AccountList []*GetBizUnitInfoResponseBodyBizUnitInfoAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BizObjectCount *int32 `json:"BizObjectCount,omitempty" xml:"BizObjectCount,omitempty"`
	// example:
	//
	// 1
	BizProcessCount    *int32                                                     `json:"BizProcessCount,omitempty" xml:"BizProcessCount,omitempty"`
	BusinessLeaderList []*GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList `json:"BusinessLeaderList,omitempty" xml:"BusinessLeaderList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DataDomainCount *int32                                                 `json:"DataDomainCount,omitempty" xml:"DataDomainCount,omitempty"`
	DataLeaderList  []*GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList `json:"DataLeaderList,omitempty" xml:"DataLeaderList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 测试
	DisplayName *string                                         `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvList     []*GetBizUnitInfoResponseBodyBizUnitInfoEnvList `json:"EnvList,omitempty" xml:"EnvList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// icon-e-commerce
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// DEV_PROD
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s GetBizUnitInfoResponseBodyBizUnitInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBodyBizUnitInfo) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetAccountList() []*GetBizUnitInfoResponseBodyBizUnitInfoAccountList {
	return s.AccountList
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetBizObjectCount() *int32 {
	return s.BizObjectCount
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetBizProcessCount() *int32 {
	return s.BizProcessCount
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetBusinessLeaderList() []*GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList {
	return s.BusinessLeaderList
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetDataDomainCount() *int32 {
	return s.DataDomainCount
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetDataLeaderList() []*GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList {
	return s.DataLeaderList
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetDescription() *string {
	return s.Description
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetEnvList() []*GetBizUnitInfoResponseBodyBizUnitInfoEnvList {
	return s.EnvList
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetIcon() *string {
	return s.Icon
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetId() *int64 {
	return s.Id
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetMode() *string {
	return s.Mode
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetName() *string {
	return s.Name
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetAccountList(v []*GetBizUnitInfoResponseBodyBizUnitInfoAccountList) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.AccountList = v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetBizObjectCount(v int32) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.BizObjectCount = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetBizProcessCount(v int32) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.BizProcessCount = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetBusinessLeaderList(v []*GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.BusinessLeaderList = v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetDataDomainCount(v int32) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.DataDomainCount = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetDataLeaderList(v []*GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.DataLeaderList = v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetDescription(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.Description = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetDisplayName(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.DisplayName = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetEnvList(v []*GetBizUnitInfoResponseBodyBizUnitInfoEnvList) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.EnvList = v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetGmtCreate(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetGmtModified(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.GmtModified = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetIcon(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.Icon = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetId(v int64) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.Id = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetLastModifier(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.LastModifier = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetLastModifierName(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.LastModifierName = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetMode(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.Mode = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetName(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.Name = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetOwnerName(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.OwnerName = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) SetOwnerUserId(v string) *GetBizUnitInfoResponseBodyBizUnitInfo {
	s.OwnerUserId = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfo) Validate() error {
	return dara.Validate(s)
}

type GetBizUnitInfoResponseBodyBizUnitInfoAccountList struct {
	// example:
	//
	// 20001201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoAccountList) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoAccountList) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoAccountList) GetId() *string {
	return s.Id
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoAccountList) SetId(v string) *GetBizUnitInfoResponseBodyBizUnitInfoAccountList {
	s.Id = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoAccountList) Validate() error {
	return dara.Validate(s)
}

type GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList struct {
	// example:
	//
	// 20001201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) GetId() *string {
	return s.Id
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) SetId(v string) *GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList {
	s.Id = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoBusinessLeaderList) Validate() error {
	return dara.Validate(s)
}

type GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList struct {
	// example:
	//
	// 20001201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) GetId() *string {
	return s.Id
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) SetId(v string) *GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList {
	s.Id = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoDataLeaderList) Validate() error {
	return dara.Validate(s)
}

type GetBizUnitInfoResponseBodyBizUnitInfoEnvList struct {
	// example:
	//
	// 测试数据板块001_开发
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// example:
	//
	// LD_test001_dev
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoEnvList) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoResponseBodyBizUnitInfoEnvList) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) GetEnvName() *string {
	return s.EnvName
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) GetName() *string {
	return s.Name
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) SetDisplayName(v string) *GetBizUnitInfoResponseBodyBizUnitInfoEnvList {
	s.DisplayName = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) SetEnvName(v string) *GetBizUnitInfoResponseBodyBizUnitInfoEnvList {
	s.EnvName = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) SetName(v string) *GetBizUnitInfoResponseBodyBizUnitInfoEnvList {
	s.Name = &v
	return s
}

func (s *GetBizUnitInfoResponseBodyBizUnitInfoEnvList) Validate() error {
	return dara.Validate(s)
}
