// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulNumStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCnt(v int32) *DescribeVulNumStatisticsResponseBody
	GetAppCnt() *int32
	SetAppNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetAppNum() *int32
	SetCmsDealedTotalNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetCmsDealedTotalNum() *int32
	SetCmsNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetCmsNum() *int32
	SetCveNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetCveNum() *int32
	SetEmgNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetEmgNum() *int32
	SetRequestId(v string) *DescribeVulNumStatisticsResponseBody
	GetRequestId() *string
	SetScaNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetScaNum() *int32
	SetSysNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetSysNum() *int32
	SetVulAsapSum(v int32) *DescribeVulNumStatisticsResponseBody
	GetVulAsapSum() *int32
	SetVulDealedTotalNum(v int32) *DescribeVulNumStatisticsResponseBody
	GetVulDealedTotalNum() *int32
	SetVulLaterSum(v int32) *DescribeVulNumStatisticsResponseBody
	GetVulLaterSum() *int32
	SetVulNntfSum(v int32) *DescribeVulNumStatisticsResponseBody
	GetVulNntfSum() *int32
}

type DescribeVulNumStatisticsResponseBody struct {
	// The number of application vulnerabilities that are detected on the asset by using the web scanner.
	//
	// example:
	//
	// 0
	AppCnt *int32 `json:"AppCnt,omitempty" xml:"AppCnt,omitempty"`
	// The number of application vulnerabilities that are detected on the asset by using the web scanner.
	//
	// example:
	//
	// 0
	AppNum *int32 `json:"AppNum,omitempty" xml:"AppNum,omitempty"`
	// The number of Web-CMS vulnerabilities that are handled.
	//
	// example:
	//
	// 0
	CmsDealedTotalNum *int32 `json:"CmsDealedTotalNum,omitempty" xml:"CmsDealedTotalNum,omitempty"`
	// The number of Web-CMS vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 0
	CmsNum *int32 `json:"CmsNum,omitempty" xml:"CmsNum,omitempty"`
	// The number of Linux software vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 0
	CveNum *int32 `json:"CveNum,omitempty" xml:"CveNum,omitempty"`
	// The number of urgent vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 0
	EmgNum *int32 `json:"EmgNum,omitempty" xml:"EmgNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E22C89D2-FE13-5800-8746-9D0EF1827A59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of middleware vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 0
	ScaNum *int32 `json:"ScaNum,omitempty" xml:"ScaNum,omitempty"`
	// The number of Windows system vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 0
	SysNum *int32 `json:"SysNum,omitempty" xml:"SysNum,omitempty"`
	// The number of vulnerabilities that have the high priority.
	//
	// example:
	//
	// 0
	VulAsapSum *int32 `json:"VulAsapSum,omitempty" xml:"VulAsapSum,omitempty"`
	// The number of vulnerabilities that are handled.
	//
	// example:
	//
	// 0
	VulDealedTotalNum *int32 `json:"VulDealedTotalNum,omitempty" xml:"VulDealedTotalNum,omitempty"`
	// The number of vulnerabilities that have the medium priority.
	//
	// example:
	//
	// 0
	VulLaterSum *int32 `json:"VulLaterSum,omitempty" xml:"VulLaterSum,omitempty"`
	// The number of vulnerabilities that have the low priority.
	//
	// example:
	//
	// 0
	VulNntfSum *int32 `json:"VulNntfSum,omitempty" xml:"VulNntfSum,omitempty"`
}

func (s DescribeVulNumStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulNumStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulNumStatisticsResponseBody) GetAppCnt() *int32 {
	return s.AppCnt
}

func (s *DescribeVulNumStatisticsResponseBody) GetAppNum() *int32 {
	return s.AppNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetCmsDealedTotalNum() *int32 {
	return s.CmsDealedTotalNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetCmsNum() *int32 {
	return s.CmsNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetCveNum() *int32 {
	return s.CveNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetEmgNum() *int32 {
	return s.EmgNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulNumStatisticsResponseBody) GetScaNum() *int32 {
	return s.ScaNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetSysNum() *int32 {
	return s.SysNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetVulAsapSum() *int32 {
	return s.VulAsapSum
}

func (s *DescribeVulNumStatisticsResponseBody) GetVulDealedTotalNum() *int32 {
	return s.VulDealedTotalNum
}

func (s *DescribeVulNumStatisticsResponseBody) GetVulLaterSum() *int32 {
	return s.VulLaterSum
}

func (s *DescribeVulNumStatisticsResponseBody) GetVulNntfSum() *int32 {
	return s.VulNntfSum
}

func (s *DescribeVulNumStatisticsResponseBody) SetAppCnt(v int32) *DescribeVulNumStatisticsResponseBody {
	s.AppCnt = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetAppNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.AppNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetCmsDealedTotalNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.CmsDealedTotalNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetCmsNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.CmsNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetCveNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.CveNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetEmgNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.EmgNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetRequestId(v string) *DescribeVulNumStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetScaNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.ScaNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetSysNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.SysNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetVulAsapSum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.VulAsapSum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetVulDealedTotalNum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.VulDealedTotalNum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetVulLaterSum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.VulLaterSum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) SetVulNntfSum(v int32) *DescribeVulNumStatisticsResponseBody {
	s.VulNntfSum = &v
	return s
}

func (s *DescribeVulNumStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
