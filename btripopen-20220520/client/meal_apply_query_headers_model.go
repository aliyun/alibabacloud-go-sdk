// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealApplyQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type MealApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *MealApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealApplyQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *MealApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealApplyQueryHeaders) SetXAcsBtripCorpToken(v string) *MealApplyQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
