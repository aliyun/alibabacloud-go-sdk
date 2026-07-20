// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainExceedApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TrainExceedApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *TrainExceedApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type TrainExceedApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s TrainExceedApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TrainExceedApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *TrainExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *TrainExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TrainExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *TrainExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *TrainExceedApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
