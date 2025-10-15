// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicense(v *GetInstanceLicenseResponseBodyLicense) *GetInstanceLicenseResponseBody
	GetLicense() *GetInstanceLicenseResponseBodyLicense
	SetRequestId(v string) *GetInstanceLicenseResponseBody
	GetRequestId() *string
}

type GetInstanceLicenseResponseBody struct {
	// Returned result.
	License *GetInstanceLicenseResponseBodyLicense `json:"License,omitempty" xml:"License,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponseBody) GetLicense() *GetInstanceLicenseResponseBodyLicense {
	return s.License
}

func (s *GetInstanceLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLicenseResponseBody) SetLicense(v *GetInstanceLicenseResponseBodyLicense) *GetInstanceLicenseResponseBody {
	s.License = v
	return s
}

func (s *GetInstanceLicenseResponseBody) SetRequestId(v string) *GetInstanceLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLicenseResponseBody) Validate() error {
	if s.License != nil {
		if err := s.License.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceLicenseResponseBodyLicense struct {
	// Edition of the License
	//
	// example:
	//
	// free
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// End date of the validity period of the License, timestamp
	//
	// example:
	//
	// 1723996800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Payment type of the License
	//
	// example:
	//
	// prepay
	LicenseChargeType *string `json:"LicenseChargeType,omitempty" xml:"LicenseChargeType,omitempty"`
	// Detailed configuration JSON string of the License
	//
	// example:
	//
	// {"modules":[{"features":[{"name":"urn:alibaba:idaas:license:module:ud:customField","status":"enabled"}]……{"name":"urn:alibaba:idaas:license:tag:enterprise","status":"enabled"}],"version":"1.0"}
	LicenseConfigJson *string `json:"LicenseConfigJson,omitempty" xml:"LicenseConfigJson,omitempty"`
	// Creation time of the License, timestamp
	//
	// example:
	//
	// 1720509699000
	LicenseCreateTime *int64 `json:"LicenseCreateTime,omitempty" xml:"LicenseCreateTime,omitempty"`
	// Unique identifier of the License
	//
	// example:
	//
	// license_1234xxxx
	LicenseId *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	// Status of the License
	//
	// example:
	//
	// valid
	LicenseStatus *string `json:"LicenseStatus,omitempty" xml:"LicenseStatus,omitempty"`
	// Purchase channel of the License
	//
	// example:
	//
	// alibaba_cloud
	PurchaseChannel *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	// Unique external product identifier corresponding to the License
	//
	// example:
	//
	// eiam-cn-xxxxx
	PurchaseInstanceId *string `json:"PurchaseInstanceId,omitempty" xml:"PurchaseInstanceId,omitempty"`
	// Start date of the validity period of the License, timestamp
	//
	// example:
	//
	// 1720509699000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// User quota of the License
	//
	// example:
	//
	// 100
	UserQuota *int64 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetInstanceLicenseResponseBodyLicense) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponseBodyLicense) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponseBodyLicense) GetEdition() *string {
	return s.Edition
}

func (s *GetInstanceLicenseResponseBodyLicense) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseChargeType() *string {
	return s.LicenseChargeType
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseConfigJson() *string {
	return s.LicenseConfigJson
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseCreateTime() *int64 {
	return s.LicenseCreateTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseId() *string {
	return s.LicenseId
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseStatus() *string {
	return s.LicenseStatus
}

func (s *GetInstanceLicenseResponseBodyLicense) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *GetInstanceLicenseResponseBodyLicense) GetPurchaseInstanceId() *string {
	return s.PurchaseInstanceId
}

func (s *GetInstanceLicenseResponseBodyLicense) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetUserQuota() *int64 {
	return s.UserQuota
}

func (s *GetInstanceLicenseResponseBodyLicense) SetEdition(v string) *GetInstanceLicenseResponseBodyLicense {
	s.Edition = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetEndTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.EndTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseChargeType(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseChargeType = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseConfigJson(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseConfigJson = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseCreateTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseCreateTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseId(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseId = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseStatus(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseStatus = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetPurchaseChannel(v string) *GetInstanceLicenseResponseBodyLicense {
	s.PurchaseChannel = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetPurchaseInstanceId(v string) *GetInstanceLicenseResponseBodyLicense {
	s.PurchaseInstanceId = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetStartTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.StartTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetUserQuota(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.UserQuota = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) Validate() error {
	return dara.Validate(s)
}
