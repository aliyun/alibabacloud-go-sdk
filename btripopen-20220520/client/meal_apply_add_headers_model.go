// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealApplyAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealApplyAddHeaders
	GetXAcsBtripCorpToken() *string
}

type MealApplyAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealApplyAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddHeaders) GoString() string {
	return s.String()
}

func (s *MealApplyAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealApplyAddHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealApplyAddHeaders) SetCommonHeaders(v map[string]*string) *MealApplyAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealApplyAddHeaders) SetXAcsBtripCorpToken(v string) *MealApplyAddHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealApplyAddHeaders) Validate() error {
	return dara.Validate(s)
}
