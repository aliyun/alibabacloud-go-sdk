// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionLOAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhysicalConnectionLOAType(v *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) *DescribePhysicalConnectionLOAResponseBody
	GetPhysicalConnectionLOAType() *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType
	SetRequestId(v string) *DescribePhysicalConnectionLOAResponseBody
	GetRequestId() *string
}

type DescribePhysicalConnectionLOAResponseBody struct {
	// The LOA information about the Express Connect circuit.
	PhysicalConnectionLOAType *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType `json:"PhysicalConnectionLOAType,omitempty" xml:"PhysicalConnectionLOAType,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 318BB676-0A2B-43A0-9AD8-F1D34E93750F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhysicalConnectionLOAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOAResponseBody) GetPhysicalConnectionLOAType() *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	return s.PhysicalConnectionLOAType
}

func (s *DescribePhysicalConnectionLOAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhysicalConnectionLOAResponseBody) SetPhysicalConnectionLOAType(v *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) *DescribePhysicalConnectionLOAResponseBody {
	s.PhysicalConnectionLOAType = v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBody) SetRequestId(v string) *DescribePhysicalConnectionLOAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBody) Validate() error {
	if s.PhysicalConnectionLOAType != nil {
		if err := s.PhysicalConnectionLOAType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType struct {
	// The name of the construction company.
	//
	// example:
	//
	// company
	CompanyLocalizedName *string `json:"CompanyLocalizedName,omitempty" xml:"CompanyLocalizedName,omitempty"`
	// The name of the organization that requires the Express Connect circuit.
	//
	// example:
	//
	// test1234
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The time when construction starts.
	//
	// example:
	//
	// 2019-02-26T08:00:00Z
	ConstructionTime *string `json:"ConstructionTime,omitempty" xml:"ConstructionTime,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-bp1ca4wca27****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The circuit code provided by the connectivity provider.
	//
	// example:
	//
	// aaa111
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The label of the cable in the data center.
	//
	// example:
	//
	// bbb222
	LineLabel *string `json:"LineLabel,omitempty" xml:"LineLabel,omitempty"`
	// The contact information about line O\\&M.
	//
	// example:
	//
	// 1388888****
	LineSPContactInfo *string `json:"LineSPContactInfo,omitempty" xml:"LineSPContactInfo,omitempty"`
	// The ISP. Valid values:
	//
	// 	- **China Telecom**
	//
	// 	- **China Unicom**
	//
	// 	- **China Mobile**
	//
	// 	- **Other ISPs in China**
	//
	// example:
	//
	// Other ISPs in China
	LineServiceProvider *string `json:"LineServiceProvider,omitempty" xml:"LineServiceProvider,omitempty"`
	// The type of the Express Connect circuit. Valid values:
	//
	// 	- **MSTP**
	//
	// 	- **MPLSVPN**
	//
	// 	- **FIBRE**
	//
	// 	- **Other**
	//
	// example:
	//
	// FIBRE
	LineType *string `json:"LineType,omitempty" xml:"LineType,omitempty"`
	// The download URL of the LOA file.
	//
	// example:
	//
	// http://******
	LoaUrl *string `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	// The information about the construction workers.
	PMInfo *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo `json:"PMInfo,omitempty" xml:"PMInfo,omitempty" type:"Struct"`
	// The on-site construction company.
	//
	// example:
	//
	// ctcu
	SI *string `json:"SI,omitempty" xml:"SI,omitempty"`
	// The status of the LOA. Valid values:
	//
	// 	- **Applying**: The LOA is pending for approval.
	//
	// 	- **Accept**: The LOA is approved.
	//
	// 	- **Available**: The LOA is available.
	//
	// 	- **Rejected**: The LOA is rejected.
	//
	// 	- **Completing**: The Express Connect circuit is under construction.
	//
	// 	- **Complete**: The Express Connect circuit is installed.
	//
	// 	- **Deleted**: The LOA is deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetCompanyLocalizedName() *string {
	return s.CompanyLocalizedName
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetCompanyName() *string {
	return s.CompanyName
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetConstructionTime() *string {
	return s.ConstructionTime
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLineLabel() *string {
	return s.LineLabel
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLineSPContactInfo() *string {
	return s.LineSPContactInfo
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLineServiceProvider() *string {
	return s.LineServiceProvider
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLineType() *string {
	return s.LineType
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetLoaUrl() *string {
	return s.LoaUrl
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetPMInfo() *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo {
	return s.PMInfo
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetSI() *string {
	return s.SI
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) GetStatus() *string {
	return s.Status
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetCompanyLocalizedName(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.CompanyLocalizedName = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetCompanyName(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.CompanyName = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetConstructionTime(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.ConstructionTime = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetInstanceId(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.InstanceId = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLineCode(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LineCode = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLineLabel(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LineLabel = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLineSPContactInfo(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LineSPContactInfo = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLineServiceProvider(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LineServiceProvider = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLineType(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LineType = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetLoaUrl(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.LoaUrl = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetPMInfo(v *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.PMInfo = v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetSI(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.SI = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) SetStatus(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType {
	s.Status = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOAType) Validate() error {
	if s.PMInfo != nil {
		if err := s.PMInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo struct {
	PMInfo []*DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo `json:"PMInfo,omitempty" xml:"PMInfo,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) GetPMInfo() []*DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	return s.PMInfo
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) SetPMInfo(v []*DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo {
	s.PMInfo = v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfo) Validate() error {
	if s.PMInfo != nil {
		for _, item := range s.PMInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo struct {
	// The identity document number of the construction worker.
	//
	// example:
	//
	// 12345671****
	PMCertificateNo *string `json:"PMCertificateNo,omitempty" xml:"PMCertificateNo,omitempty"`
	// The identity document type of the construction worker. Valid values:
	//
	// 	- **IDCard**
	//
	// 	- **Passport**
	//
	// 	- **Other**
	//
	// example:
	//
	// Other
	PMCertificateType *string `json:"PMCertificateType,omitempty" xml:"PMCertificateType,omitempty"`
	// The phone number of the construction worker.
	//
	// example:
	//
	// 18910010****
	PMContactInfo *string `json:"PMContactInfo,omitempty" xml:"PMContactInfo,omitempty"`
	// The gender of the construction worker. Valid values:
	//
	// 	- **Male**
	//
	// 	- **Female**
	//
	// example:
	//
	// Male
	PMGender *string `json:"PMGender,omitempty" xml:"PMGender,omitempty"`
	// The name of the construction worker.
	//
	// example:
	//
	// name
	PMName *string `json:"PMName,omitempty" xml:"PMName,omitempty"`
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GetPMCertificateNo() *string {
	return s.PMCertificateNo
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GetPMCertificateType() *string {
	return s.PMCertificateType
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GetPMContactInfo() *string {
	return s.PMContactInfo
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GetPMGender() *string {
	return s.PMGender
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) GetPMName() *string {
	return s.PMName
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) SetPMCertificateNo(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	s.PMCertificateNo = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) SetPMCertificateType(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	s.PMCertificateType = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) SetPMContactInfo(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	s.PMContactInfo = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) SetPMGender(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	s.PMGender = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) SetPMName(v string) *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo {
	s.PMName = &v
	return s
}

func (s *DescribePhysicalConnectionLOAResponseBodyPhysicalConnectionLOATypePMInfoPMInfo) Validate() error {
	return dara.Validate(s)
}
