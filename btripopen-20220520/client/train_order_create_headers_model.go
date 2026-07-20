// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCreateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderCreateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderCreateHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainOrderCreateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderCreateHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderCreateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderCreateHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderCreateHeaders) SetXAcsBtripCorpToken(v string) *TrainOrderCreateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderCreateHeaders) Validate() error {
	return dara.Validate(s)
}
