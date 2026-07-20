// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealApplyModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealApplyModifyHeaders
	GetXAcsBtripCorpToken() *string
}

type MealApplyModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealApplyModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyHeaders) GoString() string {
	return s.String()
}

func (s *MealApplyModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealApplyModifyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealApplyModifyHeaders) SetCommonHeaders(v map[string]*string) *MealApplyModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealApplyModifyHeaders) SetXAcsBtripCorpToken(v string) *MealApplyModifyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealApplyModifyHeaders) Validate() error {
	return dara.Validate(s)
}
