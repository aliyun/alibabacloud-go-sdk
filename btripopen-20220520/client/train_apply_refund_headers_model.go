// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyRefundHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainApplyRefundHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainApplyRefundHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainApplyRefundHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainApplyRefundHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundHeaders) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainApplyRefundHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainApplyRefundHeaders) SetCommonHeaders(v map[string]*string) *TrainApplyRefundHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainApplyRefundHeaders) SetXAcsBtripCorpToken(v string) *TrainApplyRefundHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainApplyRefundHeaders) Validate() error {
	return dara.Validate(s)
}
