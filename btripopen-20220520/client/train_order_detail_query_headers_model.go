// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderDetailQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderDetailQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderDetailQueryHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainOrderDetailQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderDetailQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderDetailQueryHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderDetailQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderDetailQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderDetailQueryHeaders) SetXAcsBtripCorpToken(v string) *TrainOrderDetailQueryHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderDetailQueryHeaders) Validate() error {
	return dara.Validate(s)
}
