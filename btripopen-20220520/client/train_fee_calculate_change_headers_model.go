// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateChangeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainFeeCalculateChangeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainFeeCalculateChangeHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainFeeCalculateChangeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainFeeCalculateChangeHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeHeaders) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainFeeCalculateChangeHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainFeeCalculateChangeHeaders) SetCommonHeaders(v map[string]*string) *TrainFeeCalculateChangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainFeeCalculateChangeHeaders) SetXAcsBtripCorpToken(v string) *TrainFeeCalculateChangeHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainFeeCalculateChangeHeaders) Validate() error {
	return dara.Validate(s)
}
