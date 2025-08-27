// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainOrderQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainOrderQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainOrderQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainOrderQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainOrderQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainOrderQueryHeaders) Validate() error {
	return dara.Validate(s)
}
