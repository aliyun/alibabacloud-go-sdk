// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest
	GetDataDisk() []*CreateInstanceRequestDataDisk
	SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest
	GetSystemDisk() *CreateInstanceRequestSystemDisk
	SetAutoRenew(v string) *CreateInstanceRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v string) *CreateInstanceRequest
	GetAutoRenewPeriod() *string
	SetEnsRegionId(v string) *CreateInstanceRequest
	GetEnsRegionId() *string
	SetHostName(v string) *CreateInstanceRequest
	GetHostName() *string
	SetImageId(v string) *CreateInstanceRequest
	GetImageId() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateInstanceRequest
	GetInternetChargeType() *string
	SetIpType(v string) *CreateInstanceRequest
	GetIpType() *string
	SetKeyPairName(v string) *CreateInstanceRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateInstanceRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateInstanceRequest
	GetPasswordInherit() *bool
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetPeriod(v string) *CreateInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CreateInstanceRequest
	GetPrivateIpAddress() *string
	SetPublicIpIdentification(v bool) *CreateInstanceRequest
	GetPublicIpIdentification() *bool
	SetQuantity(v string) *CreateInstanceRequest
	GetQuantity() *string
	SetUniqueSuffix(v bool) *CreateInstanceRequest
	GetUniqueSuffix() *bool
	SetUserData(v string) *CreateInstanceRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
}

type CreateInstanceRequest struct {
	DataDisk   []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SystemDisk *CreateInstanceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to enable auto-renewal. Valid values: **True*	- | **False*	- (default).
	//
	// example:
	//
	// True
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required when **AutoRenew*	- is set to **True**. Valid values: **1*	- to **12**. Unit: months.
	//
	// example:
	//
	// 12
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The hostname of the instance.
	//
	// General naming conventions: Periods (.) and hyphens (-) cannot be used as the first or last character and cannot be used consecutively.
	//
	// Specific naming conventions:
	//
	// - Windows instances: The hostname must be **2*	- to **15*	- characters in length. Periods (.) are not supported. The hostname cannot consist entirely of digits. Uppercase letters, lowercase letters, digits, and hyphens (-) are allowed.
	//
	// - Other instances (such as Linux): The hostname must be **2*	- to **64*	- characters in length. Multiple periods (.) are supported. Each segment separated by a period can contain uppercase letters, lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// test-HostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image that is used to start the instance.
	//
	// example:
	//
	// yourImage ID
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the instance. Naming conventions:
	//
	// - The name must be **2*	- to **128*	- characters in length and can contain letters and Chinese characters.
	//
	// - The name must start with an uppercase letter, a lowercase letter, or a Chinese character. It cannot start with http:// or https://.
	//
	// - The name can contain digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// If this parameter is not specified, the default value is the InstanceId of the instance.
	//
	// example:
	//
	// test：Instance_Name.1-2
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// For details about instance types, see [Limits](https://help.aliyun.com/document_detail/66124.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ens.se1.tiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required when you create an instance for the first time. If you already have instance resources, the existing billing method is used by default. Valid values:
	//
	// - **BandwidthByDay**: daily peak bandwidth.
	//
	// - **95BandwidthByMonth**: monthly 95th percentile bandwidth.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP type. Valid values:
	//
	// - **ipv4*	- (default): IPv4.
	//
	// - **ipv6**: IPv6.
	//
	// - **ipv4Andipv6**: IPv4 and IPv6.
	//
	// example:
	//
	// ipv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The name of the key pair. Only a single name is supported.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance.
	//
	// The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: ```()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/```
	//
	// example:
	//
	// yourPassword:1
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. Valid values:
	//
	// - **true**: The preset password of the image is used. The **Password*	- parameter must be empty, and you must make sure that the image has a preset password.
	//
	// - **false**: The preset password of the image is not used.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The payment method. Currently, this is a fixed value of Subscription (subscription instance). Modification is not supported.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The subscription duration. Valid values: **1*	- to **9*	- and **12**. Unit: months.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The private IP address. If you specify a private IP address, you must also specify the VSwitchId parameter. The vSwitch must be created in advance. Otherwise, an error is returned.
	//
	// example:
	//
	// 10.10.10.10
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Specifies whether a public IP address can be assigned to the specified instance. Valid values:
	//
	// - **true*	- (default): A public IP address can be assigned.
	//
	// - **false**: A public IP address cannot be assigned.
	//
	// example:
	//
	// true
	PublicIpIdentification *bool `json:"PublicIpIdentification,omitempty" xml:"PublicIpIdentification,omitempty"`
	// The number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Specifies whether to add sequential suffixes to **HostName*	- and **InstanceName**. The sequential suffixes start from **001*	- and cannot exceed **999**.
	//
	// Example: **LocalHost001**, **LocalHost002**, **MyInstance001**, and **MyInstance002**.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The custom data. The data must start with `#!`. The maximum length is 256 characters and the maximum size is 16 KB. Only custom data scripts are supported. InstanceMetaData rendering is not supported.
	//
	//
	// You can pass in custom data that is Base64-encoded. When you call the API to initiate a request, the custom data you specify is not encrypted. Do not pass in confidential information such as passwords and private keys in plaintext. If you need to pass in passwords or private keys, encrypt the content first and then Base64-encode it before passing it in. Decrypt the content in the same way inside the instance.
	//
	// For more information, see [Data format](https://cloudinit.readthedocs.io/en/latest/topics/format.html).
	//
	// example:
	//
	// #!/bin/sh  echo "Hello World.  The time is now $(date -R)!" | tee /home/output.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// YourVSwitchId
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetDataDisk() []*CreateInstanceRequestDataDisk {
	return s.DataDisk
}

func (s *CreateInstanceRequest) GetSystemDisk() *CreateInstanceRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateInstanceRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateInstanceRequest) GetIpType() *string {
	return s.IpType
}

func (s *CreateInstanceRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateInstanceRequest) GetPublicIpIdentification() *bool {
	return s.PublicIpIdentification
}

func (s *CreateInstanceRequest) GetQuantity() *string {
	return s.Quantity
}

func (s *CreateInstanceRequest) GetUniqueSuffix() *bool {
	return s.UniqueSuffix
}

func (s *CreateInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v string) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetEnsRegionId(v string) *CreateInstanceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetIpType(v string) *CreateInstanceRequest {
	s.IpType = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPasswordInherit(v bool) *CreateInstanceRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v string) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetPublicIpIdentification(v bool) *CreateInstanceRequest {
	s.PublicIpIdentification = &v
	return s
}

func (s *CreateInstanceRequest) SetQuantity(v string) *CreateInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateInstanceRequest) SetUniqueSuffix(v bool) *CreateInstanceRequest {
	s.UniqueSuffix = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceRequestDataDisk struct {
	// The size of the first data cloud disk. Unit: GiB. The size must be at least 20 GiB and a multiple of 10.
	//
	// example:
	//
	// 50
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) GetSize() *string {
	return s.Size
}

func (s *CreateInstanceRequestDataDisk) SetSize(v string) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestSystemDisk struct {
	// The size of the system cloud disk. Unit: GiB. The value must be a multiple of 10 and cannot be smaller than the image size. For the valid range, refer to the SystemDiskMinSize and SystemDiskMaxSize values returned by DescribeAvailableResourceInfo. 20 GiB and 40 GiB are common configurations and do not indicate that only these two values are supported.
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) GetSize() *string {
	return s.Size
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v string) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}
