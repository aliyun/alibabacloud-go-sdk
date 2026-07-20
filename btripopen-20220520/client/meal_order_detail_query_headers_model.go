// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealOrderDetailQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MealOrderDetailQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *MealOrderDetailQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type MealOrderDetailQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s MealOrderDetailQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s MealOrderDetailQueryHeaders) GoString() string {
	return s.String()
}

func (s *MealOrderDetailQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MealOrderDetailQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *MealOrderDetailQueryHeaders) SetCommonHeaders(v map[string]*string) *MealOrderDetailQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MealOrderDetailQueryHeaders) SetXAcsBtripCorpToken(v string) *MealOrderDetailQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *MealOrderDetailQueryHeaders) Validate() error {
	return dara.Validate(s)
}
