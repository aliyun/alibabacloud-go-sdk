// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoListSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainNoListSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainNoListSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainNoListSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainNoListSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchHeaders) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainNoListSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainNoListSearchHeaders) SetCommonHeaders(v map[string]*string) *TrainNoListSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainNoListSearchHeaders) SetXAcsBtripCorpToken(v string) *TrainNoListSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainNoListSearchHeaders) Validate() error {
	return dara.Validate(s)
}
