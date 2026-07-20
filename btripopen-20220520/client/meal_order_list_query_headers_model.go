// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealOrderListQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type MealOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *MealOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealOrderListQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *MealOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealOrderListQueryHeaders) SetXAcsBtripCorpToken(v string) *MealOrderListQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
