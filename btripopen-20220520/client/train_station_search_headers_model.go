// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStationSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainStationSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainStationSearchHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainStationSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainStationSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchHeaders) GoString() string {
	return s.String()
}

func (s *TrainStationSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainStationSearchHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainStationSearchHeaders) SetCommonHeaders(v map[string]*string) *TrainStationSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainStationSearchHeaders) SetXAcsBtripSoCorpToken(v string) *TrainStationSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainStationSearchHeaders) Validate() error {
	return dara.Validate(s)
}
