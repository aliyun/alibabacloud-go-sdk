// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorSyncPayStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CooperatorSyncPayStatusHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *CooperatorSyncPayStatusHeaders
	GetXAcsBtripCorpToken() *string
}

type CooperatorSyncPayStatusHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s CooperatorSyncPayStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s CooperatorSyncPayStatusHeaders) GoString() string {
	return s.String()
}

func (s *CooperatorSyncPayStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CooperatorSyncPayStatusHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *CooperatorSyncPayStatusHeaders) SetCommonHeaders(v map[string]*string) *CooperatorSyncPayStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CooperatorSyncPayStatusHeaders) SetXAcsBtripCorpToken(v string) *CooperatorSyncPayStatusHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *CooperatorSyncPayStatusHeaders) Validate() error {
	return dara.Validate(s)
}
