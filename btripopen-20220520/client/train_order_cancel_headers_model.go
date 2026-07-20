// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainOrderCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCancelHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderCancelHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderCancelHeaders) SetXAcsBtripCorpToken(v string) *TrainOrderCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderCancelHeaders) Validate() error {
	return dara.Validate(s)
}
