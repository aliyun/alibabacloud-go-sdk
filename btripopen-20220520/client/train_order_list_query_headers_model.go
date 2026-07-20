// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainOrderListQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderListQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
