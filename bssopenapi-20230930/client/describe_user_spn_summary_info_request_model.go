// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSpnSummaryInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds) *DescribeUserSpnSummaryInfoRequest
	GetEcIdAccountIds() []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds
	SetNbid(v string) *DescribeUserSpnSummaryInfoRequest
	GetNbid() *string
}

type DescribeUserSpnSummaryInfoRequest struct {
	EcIdAccountIds []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                            `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DescribeUserSpnSummaryInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoRequest) GetEcIdAccountIds() []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DescribeUserSpnSummaryInfoRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeUserSpnSummaryInfoRequest) SetEcIdAccountIds(v []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds) *DescribeUserSpnSummaryInfoRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequest) SetNbid(v string) *DescribeUserSpnSummaryInfoRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeUserSpnSummaryInfoRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeUserSpnSummaryInfoRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeUserSpnSummaryInfoRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) SetEcId(v string) *DescribeUserSpnSummaryInfoRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
