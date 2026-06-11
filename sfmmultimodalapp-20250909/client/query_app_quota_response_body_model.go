// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveLicenseCount(v int32) *QueryAppQuotaResponseBody
	GetActiveLicenseCount() *int32
	SetAppId(v string) *QueryAppQuotaResponseBody
	GetAppId() *string
	SetLicenseCount(v int32) *QueryAppQuotaResponseBody
	GetLicenseCount() *int32
	SetRequestId(v string) *QueryAppQuotaResponseBody
	GetRequestId() *string
	SetUsagePercent(v float64) *QueryAppQuotaResponseBody
	GetUsagePercent() *float64
	SetWorkspaceId(v string) *QueryAppQuotaResponseBody
	GetWorkspaceId() *string
}

type QueryAppQuotaResponseBody struct {
	ActiveLicenseCount *int32   `json:"ActiveLicenseCount,omitempty" xml:"ActiveLicenseCount,omitempty"`
	AppId              *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LicenseCount       *int32   `json:"LicenseCount,omitempty" xml:"LicenseCount,omitempty"`
	RequestId          *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsagePercent       *float64 `json:"UsagePercent,omitempty" xml:"UsagePercent,omitempty"`
	WorkspaceId        *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAppQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppQuotaResponseBody) GetActiveLicenseCount() *int32 {
	return s.ActiveLicenseCount
}

func (s *QueryAppQuotaResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *QueryAppQuotaResponseBody) GetLicenseCount() *int32 {
	return s.LicenseCount
}

func (s *QueryAppQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppQuotaResponseBody) GetUsagePercent() *float64 {
	return s.UsagePercent
}

func (s *QueryAppQuotaResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryAppQuotaResponseBody) SetActiveLicenseCount(v int32) *QueryAppQuotaResponseBody {
	s.ActiveLicenseCount = &v
	return s
}

func (s *QueryAppQuotaResponseBody) SetAppId(v string) *QueryAppQuotaResponseBody {
	s.AppId = &v
	return s
}

func (s *QueryAppQuotaResponseBody) SetLicenseCount(v int32) *QueryAppQuotaResponseBody {
	s.LicenseCount = &v
	return s
}

func (s *QueryAppQuotaResponseBody) SetRequestId(v string) *QueryAppQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppQuotaResponseBody) SetUsagePercent(v float64) *QueryAppQuotaResponseBody {
	s.UsagePercent = &v
	return s
}

func (s *QueryAppQuotaResponseBody) SetWorkspaceId(v string) *QueryAppQuotaResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *QueryAppQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
