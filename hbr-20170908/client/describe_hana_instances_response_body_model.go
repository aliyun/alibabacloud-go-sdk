// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaInstancesResponseBody
	GetCode() *string
	SetHanas(v *DescribeHanaInstancesResponseBodyHanas) *DescribeHanaInstancesResponseBody
	GetHanas() *DescribeHanaInstancesResponseBodyHanas
	SetMessage(v string) *DescribeHanaInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHanaInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHanaInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeHanaInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeHanaInstancesResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the SAP HANA instances.
	Hanas *DescribeHanaInstancesResponseBodyHanas `json:"Hanas,omitempty" xml:"Hanas,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4003DD68-3C3C-5071-B4FC-631A6C1BAC1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaInstancesResponseBody) GetHanas() *DescribeHanaInstancesResponseBodyHanas {
	return s.Hanas
}

func (s *DescribeHanaInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHanaInstancesResponseBody) SetCode(v string) *DescribeHanaInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetHanas(v *DescribeHanaInstancesResponseBodyHanas) *DescribeHanaInstancesResponseBody {
	s.Hanas = v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetMessage(v string) *DescribeHanaInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetPageNumber(v int32) *DescribeHanaInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetPageSize(v int32) *DescribeHanaInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetRequestId(v string) *DescribeHanaInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetSuccess(v bool) *DescribeHanaInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetTotalCount(v int32) *DescribeHanaInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) Validate() error {
	if s.Hanas != nil {
		if err := s.Hanas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaInstancesResponseBodyHanas struct {
	Hana []*DescribeHanaInstancesResponseBodyHanasHana `json:"Hana,omitempty" xml:"Hana,omitempty" type:"Repeated"`
}

func (s DescribeHanaInstancesResponseBodyHanas) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanas) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanas) GetHana() []*DescribeHanaInstancesResponseBodyHanasHana {
	return s.Hana
}

func (s *DescribeHanaInstancesResponseBodyHanas) SetHana(v []*DescribeHanaInstancesResponseBodyHanasHana) *DescribeHanaInstancesResponseBodyHanas {
	s.Hana = v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanas) Validate() error {
	if s.Hana != nil {
		for _, item := range s.Hana {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHanaInstancesResponseBodyHanasHana struct {
	// The alert settings. Valid value: INHERITED, which indicates that the Cloud Backup client sends alert notifications by using the same method configured for the backup vault.
	//
	// example:
	//
	// INHERITED
	AlertSetting *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0004cf6g6******0yd7y
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- **SELF_ACCOUNT**: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- **CROSS_ACCOUNT**: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The name of the SAP HANA instance.
	//
	// example:
	//
	// HANA-DEV
	HanaName *string `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	// The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
	//
	// example:
	//
	// 47.100.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The instance number of the SAP HANA system.
	//
	// example:
	//
	// 00
	InstanceNumber *int32 `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmvnf22m7itha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the SAP HANA instance. Valid values:
	//
	// 	- INITIALIZING: The instance is being initialized.
	//
	// 	- INITIALIZED: The instance is registered.
	//
	// 	- INVALID_HANA_NODE: The instance is invalid.
	//
	// 	- INITIALIZE_FAILED: The client fails to be installed on the instance.
	//
	// example:
	//
	// INITIALIZED
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information.
	//
	// example:
	//
	// INSTALL_CLIENT_FAILED
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The tags of the SAP HANA instance.
	Tags *DescribeHanaInstancesResponseBodyHanasHanaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// Indicates whether the SAP HANA instance is connected over Secure Sockets Layer (SSL). Valid values:
	//
	// 	- true: The SAP HANA instance is connected over SSL.
	//
	// 	- false: The SAP HANA instance is not connected over SSL.
	//
	// example:
	//
	// true
	UseSsl *bool `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The username of the SYSTEMDB database.
	//
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Indicates whether the SSL certificate of the SAP HANA instance is verified.
	//
	// example:
	//
	// false
	ValidateCertificate *bool `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0000s974******1hl
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaInstancesResponseBodyHanasHana) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHana) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetAlertSetting() *string {
	return s.AlertSetting
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetHanaName() *string {
	return s.HanaName
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetHost() *string {
	return s.Host
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetInstanceNumber() *int32 {
	return s.InstanceNumber
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetTags() *DescribeHanaInstancesResponseBodyHanasHanaTags {
	return s.Tags
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetUseSsl() *bool {
	return s.UseSsl
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetUserName() *string {
	return s.UserName
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetValidateCertificate() *bool {
	return s.ValidateCertificate
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetAlertSetting(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.AlertSetting = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetClusterId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetCrossAccountRoleName(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetCrossAccountType(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetCrossAccountUserId(v int64) *DescribeHanaInstancesResponseBodyHanasHana {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetHanaName(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.HanaName = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetHost(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Host = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetInstanceNumber(v int32) *DescribeHanaInstancesResponseBodyHanasHana {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetResourceGroupId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetStatus(v int64) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Status = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetStatusMessage(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.StatusMessage = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetTags(v *DescribeHanaInstancesResponseBodyHanasHanaTags) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Tags = v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetUseSsl(v bool) *DescribeHanaInstancesResponseBodyHanasHana {
	s.UseSsl = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetUserName(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.UserName = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetValidateCertificate(v bool) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ValidateCertificate = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetVaultId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaInstancesResponseBodyHanasHanaTags struct {
	Tag []*DescribeHanaInstancesResponseBodyHanasHanaTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTags) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTags) GetTag() []*DescribeHanaInstancesResponseBodyHanasHanaTagsTag {
	return s.Tag
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTags) SetTag(v []*DescribeHanaInstancesResponseBodyHanasHanaTagsTag) *DescribeHanaInstancesResponseBodyHanasHanaTags {
	s.Tag = v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHanaInstancesResponseBodyHanasHanaTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// ace:rm:rgld
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// rg-acfmwutpyat2kwy
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) SetKey(v string) *DescribeHanaInstancesResponseBodyHanasHanaTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) SetValue(v string) *DescribeHanaInstancesResponseBodyHanasHanaTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) Validate() error {
	return dara.Validate(s)
}
