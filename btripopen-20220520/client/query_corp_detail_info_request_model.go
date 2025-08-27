// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCorpDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *QueryCorpDetailInfoRequest
	GetAccountId() *string
	SetTargetCorpId(v string) *QueryCorpDetailInfoRequest
	GetTargetCorpId() *string
	SetTargetThirdCorpId(v string) *QueryCorpDetailInfoRequest
	GetTargetThirdCorpId() *string
}

type QueryCorpDetailInfoRequest struct {
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// example:
	//
	// btripxxxxxx
	TargetCorpId      *string `json:"target_corp_id,omitempty" xml:"target_corp_id,omitempty"`
	TargetThirdCorpId *string `json:"target_third_corp_id,omitempty" xml:"target_third_corp_id,omitempty"`
}

func (s QueryCorpDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCorpDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpDetailInfoRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryCorpDetailInfoRequest) GetTargetCorpId() *string {
	return s.TargetCorpId
}

func (s *QueryCorpDetailInfoRequest) GetTargetThirdCorpId() *string {
	return s.TargetThirdCorpId
}

func (s *QueryCorpDetailInfoRequest) SetAccountId(v string) *QueryCorpDetailInfoRequest {
	s.AccountId = &v
	return s
}

func (s *QueryCorpDetailInfoRequest) SetTargetCorpId(v string) *QueryCorpDetailInfoRequest {
	s.TargetCorpId = &v
	return s
}

func (s *QueryCorpDetailInfoRequest) SetTargetThirdCorpId(v string) *QueryCorpDetailInfoRequest {
	s.TargetThirdCorpId = &v
	return s
}

func (s *QueryCorpDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
