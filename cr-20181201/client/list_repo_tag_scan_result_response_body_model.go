// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagScanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepoTagScanResultResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoTagScanResultResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoTagScanResultResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoTagScanResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoTagScanResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRepoTagScanResultResponseBody
	GetTotalCount() *int32
	SetVulnerabilities(v []*ListRepoTagScanResultResponseBodyVulnerabilities) *ListRepoTagScanResultResponseBody
	GetVulnerabilities() []*ListRepoTagScanResultResponseBodyVulnerabilities
}

type ListRepoTagScanResultResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: Success
	//
	// - `false`: Failure
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the scan result list.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page in the scan result list.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 56B5C92F-F5D9-46E0-823F-EC71D1892DAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of scanned vulnerabilities.
	//
	// example:
	//
	// 196
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The array of scanned vulnerabilities.
	Vulnerabilities []*ListRepoTagScanResultResponseBodyVulnerabilities `json:"Vulnerabilities,omitempty" xml:"Vulnerabilities,omitempty" type:"Repeated"`
}

func (s ListRepoTagScanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoTagScanResultResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoTagScanResultResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoTagScanResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoTagScanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoTagScanResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRepoTagScanResultResponseBody) GetVulnerabilities() []*ListRepoTagScanResultResponseBodyVulnerabilities {
	return s.Vulnerabilities
}

func (s *ListRepoTagScanResultResponseBody) SetCode(v string) *ListRepoTagScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetIsSuccess(v bool) *ListRepoTagScanResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetPageNo(v int32) *ListRepoTagScanResultResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetPageSize(v int32) *ListRepoTagScanResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetRequestId(v string) *ListRepoTagScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetTotalCount(v int32) *ListRepoTagScanResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetVulnerabilities(v []*ListRepoTagScanResultResponseBodyVulnerabilities) *ListRepoTagScanResultResponseBody {
	s.Vulnerabilities = v
	return s
}

func (s *ListRepoTagScanResultResponseBody) Validate() error {
	if s.Vulnerabilities != nil {
		for _, item := range s.Vulnerabilities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepoTagScanResultResponseBodyVulnerabilities struct {
	// The image layer that introduced the vulnerability.
	//
	// example:
	//
	// sha256:123456717b8e40b6480979b739010d8d549989602bcdd07922119aec6f9dbe57
	AddedBy *string `json:"AddedBy,omitempty" xml:"AddedBy,omitempty"`
	// Deprecated
	//
	// The vulnerability name.
	//
	// example:
	//
	// Vulnerability
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The vulnerability link.
	//
	// example:
	//
	// https://security-tracker.debian.org/tracker/CVE-2009-5155
	CveLink *string `json:"CveLink,omitempty" xml:"CveLink,omitempty"`
	// The vulnerability location.
	//
	// example:
	//
	// /test.txt
	CveLocation *string `json:"CveLocation,omitempty" xml:"CveLocation,omitempty"`
	// The vulnerability name.
	//
	// example:
	//
	// CVE-2009-5155
	CveName *string `json:"CveName,omitempty" xml:"CveName,omitempty"`
	// The vulnerability description.
	//
	// example:
	//
	// description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The feature that introduced the vulnerability.
	//
	// example:
	//
	// eglibc
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The vulnerability fix command.
	//
	// example:
	//
	// yum install -y xxx
	FixCmd *string `json:"FixCmd,omitempty" xml:"FixCmd,omitempty"`
	// The vulnerability type. Valid values:
	//
	// - `cve`: System vulnerability
	//
	// - `sca`: Application vulnerability
	//
	// example:
	//
	// cve
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity level of the scanned vulnerability. Valid values:
	//
	// - `High`: High
	//
	// - `Medium`: Medium
	//
	// - `Low`: Low
	//
	// - `Unknown`: Unknown
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The vulnerability version.
	//
	// example:
	//
	// 2.19-6.9
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version that fixes the vulnerability.
	//
	// example:
	//
	// 2.19-18+deb8u5
	VersionFixed *string `json:"VersionFixed,omitempty" xml:"VersionFixed,omitempty"`
	// The version format.
	//
	// example:
	//
	// dpkg
	VersionFormat *string `json:"VersionFormat,omitempty" xml:"VersionFormat,omitempty"`
}

func (s ListRepoTagScanResultResponseBodyVulnerabilities) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagScanResultResponseBodyVulnerabilities) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetAddedBy() *string {
	return s.AddedBy
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetAliasName() *string {
	return s.AliasName
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetCveLink() *string {
	return s.CveLink
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetCveLocation() *string {
	return s.CveLocation
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetCveName() *string {
	return s.CveName
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetDescription() *string {
	return s.Description
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetFeature() *string {
	return s.Feature
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetFixCmd() *string {
	return s.FixCmd
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetScanType() *string {
	return s.ScanType
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetSeverity() *string {
	return s.Severity
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetVersion() *string {
	return s.Version
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetVersionFixed() *string {
	return s.VersionFixed
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) GetVersionFormat() *string {
	return s.VersionFormat
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetAddedBy(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.AddedBy = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetAliasName(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.AliasName = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveLink(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveLink = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveLocation(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveLocation = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveName(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveName = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetDescription(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Description = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetFeature(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Feature = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetFixCmd(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.FixCmd = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetScanType(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.ScanType = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetSeverity(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Severity = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersion(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Version = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersionFixed(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.VersionFixed = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersionFormat(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.VersionFormat = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) Validate() error {
	return dara.Validate(s)
}
