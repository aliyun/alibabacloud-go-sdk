// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyApproveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealApplyApproveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealApplyApproveHeaders
	GetXAcsBtripCorpToken() *string
}

type MealApplyApproveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealApplyApproveHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealApplyApproveHeaders) GoString() string {
	return s.String()
}

func (s *MealApplyApproveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealApplyApproveHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealApplyApproveHeaders) SetCommonHeaders(v map[string]*string) *MealApplyApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealApplyApproveHeaders) SetXAcsBtripCorpToken(v string) *MealApplyApproveHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealApplyApproveHeaders) Validate() error {
	return dara.Validate(s)
}
