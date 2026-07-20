// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateRefundHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainFeeCalculateRefundHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainFeeCalculateRefundHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainFeeCalculateRefundHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainFeeCalculateRefundHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundHeaders) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainFeeCalculateRefundHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainFeeCalculateRefundHeaders) SetCommonHeaders(v map[string]*string) *TrainFeeCalculateRefundHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainFeeCalculateRefundHeaders) SetXAcsBtripCorpToken(v string) *TrainFeeCalculateRefundHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainFeeCalculateRefundHeaders) Validate() error {
	return dara.Validate(s)
}
