// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorAccountAKListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudVendorAccountAKs(v []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) *DescribeCloudVendorAccountAKListResponseBody
	GetCloudVendorAccountAKs() []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs
	SetPageInfo(v *DescribeCloudVendorAccountAKListResponseBodyPageInfo) *DescribeCloudVendorAccountAKListResponseBody
	GetPageInfo() *DescribeCloudVendorAccountAKListResponseBodyPageInfo
	SetRequestId(v string) *DescribeCloudVendorAccountAKListResponseBody
	GetRequestId() *string
}

type DescribeCloudVendorAccountAKListResponseBody struct {
	// The information about the AccessKey pairs.
	CloudVendorAccountAKs []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs `json:"CloudVendorAccountAKs,omitempty" xml:"CloudVendorAccountAKs,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeCloudVendorAccountAKListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudVendorAccountAKListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListResponseBody) GetCloudVendorAccountAKs() []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	return s.CloudVendorAccountAKs
}

func (s *DescribeCloudVendorAccountAKListResponseBody) GetPageInfo() *DescribeCloudVendorAccountAKListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCloudVendorAccountAKListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudVendorAccountAKListResponseBody) SetCloudVendorAccountAKs(v []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) *DescribeCloudVendorAccountAKListResponseBody {
	s.CloudVendorAccountAKs = v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBody) SetPageInfo(v *DescribeCloudVendorAccountAKListResponseBodyPageInfo) *DescribeCloudVendorAccountAKListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBody) SetRequestId(v string) *DescribeCloudVendorAccountAKListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBody) Validate() error {
	if s.CloudVendorAccountAKs != nil {
		for _, item := range s.CloudVendorAccountAKs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs struct {
	// The type of the account to which the AccessKey pair belongs. Valid values:
	//
	// 	- **primary**: a primary account
	//
	// 	- **sub**: a sub-account
	//
	// example:
	//
	// primary
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The unique ID of the AccessKey pair.
	//
	// example:
	//
	// 2345
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The modules that are associated with the AccessKey pair.
	AuthModules []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// The Account ID.
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The extended information of the module.
	//
	// example:
	//
	// {\\"product\\":\\"webFirewall\\",\\"remark\\":\\"remark\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The error message of the AccessKey pair.
	//
	// example:
	//
	// The IAM user is forbidden in the currently selected region
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The AccessKey ID.
	//
	// example:
	//
	// S3D6c4O***
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The service status of the AccessKey pair. Valid values:
	//
	// 	- **0**: being used
	//
	// 	- **1**: exception occurred
	//
	// 	- **2**: being validated
	//
	// 	- **3**: validation timed out
	//
	// example:
	//
	// 0
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- **0**: enabled
	//
	// 	- **1**: disabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud service provider. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud
	//
	// 	- **Azure**: Microsoft Azure
	//
	// 	- **AWS**: Amazon Web Services (AWS)
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the AccessKey pair.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
	// Account ID of the multi-cloud instance.
	//
	// example:
	//
	// 123
	VendorUid *string `json:"VendorUid,omitempty" xml:"VendorUid,omitempty"`
	// Account name of the multi-cloud instance.
	//
	// example:
	//
	// VendorUserName
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
}

func (s DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetAkType() *string {
	return s.AkType
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetAuthId() *int64 {
	return s.AuthId
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetAuthModules() []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	return s.AuthModules
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetSecretId() *string {
	return s.SecretId
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetVendorUid() *string {
	return s.VendorUid
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetAkType(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.AkType = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetAuthId(v int64) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.AuthId = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetAuthModules(v []*DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.AuthModules = v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetCtdrCloudUserId(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.CtdrCloudUserId = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetExtendInfo(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.ExtendInfo = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetMessage(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.Message = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetSecretId(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.SecretId = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetServiceStatus(v int32) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetStatus(v int32) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.Status = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetVendor(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetVendorAuthAlias(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.VendorAuthAlias = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetVendorUid(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.VendorUid = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) SetVendorUserName(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs {
	s.VendorUserName = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKs) Validate() error {
	if s.AuthModules != nil {
		for _, item := range s.AuthModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules struct {
	// The error message of the module.
	//
	// example:
	//
	// ak_domain_error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The code of the module. Valid values:
	//
	// 	- **HOST**: host
	//
	// 	- **CSPM**: configuration assessment
	//
	// 	- **SIEM**: CloudSiem
	//
	// 	- **TRIAL**: log audit
	//
	// example:
	//
	// HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The cloud asset that is associated with the module.
	//
	// example:
	//
	// Cloud server or virtual machine
	ModuleAssetType *string `json:"ModuleAssetType,omitempty" xml:"ModuleAssetType,omitempty"`
	// The display name of the module.
	//
	// example:
	//
	// Host Assets
	ModuleDisp *string `json:"ModuleDisp,omitempty" xml:"ModuleDisp,omitempty"`
	// The service status of the module. Valid values:
	//
	// 	- **0**: being used
	//
	// 	- **1**: exception occurred
	//
	// 	- **2**: being validated
	//
	// 	- **3**: validation timed out
	//
	// example:
	//
	// 0
	ModuleServiceStatus *int32 `json:"ModuleServiceStatus,omitempty" xml:"ModuleServiceStatus,omitempty"`
	// The permission description of the module.
	//
	// example:
	//
	// Read permission of the cloud server or virtual machine
	ModuleStatement *string `json:"ModuleStatement,omitempty" xml:"ModuleStatement,omitempty"`
	// The error message of the log audit service.
	//
	// example:
	//
	// timeout
	TrailMessage *string `json:"TrailMessage,omitempty" xml:"TrailMessage,omitempty"`
	// The status of the log audit service. Valid values:
	//
	// 	- **init**: being initialized
	//
	// 	- **verify**: being validated
	//
	// 	- **enable**: enabled
	//
	// 	- **disable**: disabled
	//
	// 	- **error**: exception occurred
	//
	// 	- **timeout**: validation timed out
	//
	// example:
	//
	// enable
	TrailStatus *string `json:"TrailStatus,omitempty" xml:"TrailStatus,omitempty"`
}

func (s DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetModule() *string {
	return s.Module
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetModuleAssetType() *string {
	return s.ModuleAssetType
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetModuleDisp() *string {
	return s.ModuleDisp
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetModuleServiceStatus() *int32 {
	return s.ModuleServiceStatus
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetModuleStatement() *string {
	return s.ModuleStatement
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetTrailMessage() *string {
	return s.TrailMessage
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) GetTrailStatus() *string {
	return s.TrailStatus
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetMessage(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.Message = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetModule(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.Module = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetModuleAssetType(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.ModuleAssetType = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetModuleDisp(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.ModuleDisp = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetModuleServiceStatus(v int32) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.ModuleServiceStatus = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetModuleStatement(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.ModuleStatement = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetTrailMessage(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.TrailMessage = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) SetTrailStatus(v string) *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules {
	s.TrailStatus = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyCloudVendorAccountAKsAuthModules) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudVendorAccountAKListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudVendorAccountAKListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) SetCount(v int32) *DescribeCloudVendorAccountAKListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCloudVendorAccountAKListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) SetPageSize(v int32) *DescribeCloudVendorAccountAKListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCloudVendorAccountAKListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
