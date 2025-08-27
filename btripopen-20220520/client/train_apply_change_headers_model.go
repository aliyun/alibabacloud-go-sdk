// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyChangeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainApplyChangeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainApplyChangeHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainApplyChangeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainApplyChangeHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeHeaders) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainApplyChangeHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainApplyChangeHeaders) SetCommonHeaders(v map[string]*string) *TrainApplyChangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainApplyChangeHeaders) SetXAcsBtripCorpToken(v string) *TrainApplyChangeHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainApplyChangeHeaders) Validate() error {
	return dara.Validate(s)
}
