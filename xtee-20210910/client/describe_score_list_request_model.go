// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeScoreListRequest
	GetId() *string
	SetLang(v string) *DescribeScoreListRequest
	GetLang() *string
}

type DescribeScoreListRequest struct {
	// Primary key ID.
	//
	// example:
	//
	// 300126
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Language type of the returned message. Values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeScoreListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScoreListRequest) GetId() *string {
	return s.Id
}

func (s *DescribeScoreListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeScoreListRequest) SetId(v string) *DescribeScoreListRequest {
	s.Id = &v
	return s
}

func (s *DescribeScoreListRequest) SetLang(v string) *DescribeScoreListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScoreListRequest) Validate() error {
	return dara.Validate(s)
}
