// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderChangeConfirmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainOrderChangeConfirmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *TrainOrderChangeConfirmHeaders
	GetXAcsBtripCorpToken() *string
}

type TrainOrderChangeConfirmHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s TrainOrderChangeConfirmHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderChangeConfirmHeaders) GoString() string {
	return s.String()
}

func (s *TrainOrderChangeConfirmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainOrderChangeConfirmHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *TrainOrderChangeConfirmHeaders) SetCommonHeaders(v map[string]*string) *TrainOrderChangeConfirmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainOrderChangeConfirmHeaders) SetXAcsBtripCorpToken(v string) *TrainOrderChangeConfirmHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *TrainOrderChangeConfirmHeaders) Validate() error {
	return dara.Validate(s)
}
