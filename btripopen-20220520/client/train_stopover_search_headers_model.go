// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStopoverSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainStopoverSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainStopoverSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainStopoverSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainStopoverSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainStopoverSearchHeaders) GoString() string {
	return s.String()
}

func (s *TrainStopoverSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainStopoverSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainStopoverSearchHeaders) SetCommonHeaders(v map[string]*string) *TrainStopoverSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainStopoverSearchHeaders) SetXAcsBtripCorpToken(v string) *TrainStopoverSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainStopoverSearchHeaders) Validate() error {
	return dara.Validate(s)
}
