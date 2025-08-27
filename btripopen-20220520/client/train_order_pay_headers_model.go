// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderPayHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainOrderPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderPayHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderPayHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderPayHeaders) SetXAcsBtripCorpToken(v string) *TrainOrderPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderPayHeaders) Validate() error {
	return dara.Validate(s)
}
