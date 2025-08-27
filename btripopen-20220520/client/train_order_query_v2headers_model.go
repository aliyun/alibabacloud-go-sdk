// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderQueryV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderQueryV2Headers
	GetXAcsBtripCorpToken() *string
}

type TrainOrderQueryV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 112dcasca
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderQueryV2Headers) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2Headers) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderQueryV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderQueryV2Headers) SetCommonHeaders(v map[string]*string) *TrainOrderQueryV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderQueryV2Headers) SetXAcsBtripCorpToken(v string) *TrainOrderQueryV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderQueryV2Headers) Validate() error {
	return dara.Validate(s)
}
