// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoInfoSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainNoInfoSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainNoInfoSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainNoInfoSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainNoInfoSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchHeaders) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainNoInfoSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainNoInfoSearchHeaders) SetCommonHeaders(v map[string]*string) *TrainNoInfoSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainNoInfoSearchHeaders) SetXAcsBtripCorpToken(v string) *TrainNoInfoSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainNoInfoSearchHeaders) Validate() error {
	return dara.Validate(s)
}
