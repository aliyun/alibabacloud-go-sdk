// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetUserCanAnalysisVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurPage(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetCurPage() *int32
	SetPageSize(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetRequestId() *string
	SetTotalPage(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetTotalPage() *int32
	SetTotalSize(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetTotalSize() *int32
	SetUserCanAnalysisVpcsPopResults(v *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) *DescribeIntranetUserCanAnalysisVpcsResponseBody
	GetUserCanAnalysisVpcsPopResults() *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults
}

type DescribeIntranetUserCanAnalysisVpcsResponseBody struct {
	// example:
	//
	// 11
	CurPage *int32 `json:"CurPage,omitempty" xml:"CurPage,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 35
	TotalSize                     *int32                                                                        `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	UserCanAnalysisVpcsPopResults *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults `json:"UserCanAnalysisVpcsPopResults,omitempty" xml:"UserCanAnalysisVpcsPopResults,omitempty" type:"Struct"`
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetCurPage() *int32 {
	return s.CurPage
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) GetUserCanAnalysisVpcsPopResults() *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults {
	return s.UserCanAnalysisVpcsPopResults
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetCurPage(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.CurPage = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetPageSize(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetRequestId(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetTotalPage(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetTotalSize(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) SetUserCanAnalysisVpcsPopResults(v *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	s.UserCanAnalysisVpcsPopResults = v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBody) Validate() error {
	if s.UserCanAnalysisVpcsPopResults != nil {
		if err := s.UserCanAnalysisVpcsPopResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults struct {
	UserCanAnalysisVpcsPopResult []*DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult `json:"UserCanAnalysisVpcsPopResult,omitempty" xml:"UserCanAnalysisVpcsPopResult,omitempty" type:"Repeated"`
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) GoString() string {
	return s.String()
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) GetUserCanAnalysisVpcsPopResult() []*DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	return s.UserCanAnalysisVpcsPopResult
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) SetUserCanAnalysisVpcsPopResult(v []*DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults {
	s.UserCanAnalysisVpcsPopResult = v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResults) Validate() error {
	if s.UserCanAnalysisVpcsPopResult != nil {
		for _, item := range s.UserCanAnalysisVpcsPopResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult struct {
	AuthType         *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TreeLevel        *int32  `json:"TreeLevel,omitempty" xml:"TreeLevel,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType          *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GoString() string {
	return s.String()
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetTreeLevel() *int32 {
	return s.TreeLevel
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetAuthType(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.AuthType = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetAuthorizedUserId(v int64) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetNetworkType(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.NetworkType = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetRegionId(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.RegionId = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetTreeLevel(v int32) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.TreeLevel = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetVpcId(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.VpcId = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) SetVpcType(v string) *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult {
	s.VpcType = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponseBodyUserCanAnalysisVpcsPopResultsUserCanAnalysisVpcsPopResult) Validate() error {
	return dara.Validate(s)
}
