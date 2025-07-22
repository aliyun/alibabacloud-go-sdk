// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSpnSummaryInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIdsShrink(v string) *DescribeUserSpnSummaryInfoShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *DescribeUserSpnSummaryInfoShrinkRequest
	GetNbid() *string
}

type DescribeUserSpnSummaryInfoShrinkRequest struct {
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DescribeUserSpnSummaryInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeUserSpnSummaryInfoShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) SetNbid(v string) *DescribeUserSpnSummaryInfoShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
