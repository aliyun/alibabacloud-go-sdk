// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSingleUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SyncSingleUserHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *SyncSingleUserHeaders
	GetXAcsBtripSoCorpToken() *string
}

type SyncSingleUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s SyncSingleUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s SyncSingleUserHeaders) GoString() string {
	return s.String()
}

func (s *SyncSingleUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SyncSingleUserHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *SyncSingleUserHeaders) SetCommonHeaders(v map[string]*string) *SyncSingleUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncSingleUserHeaders) SetXAcsBtripSoCorpToken(v string) *SyncSingleUserHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *SyncSingleUserHeaders) Validate() error {
	return dara.Validate(s)
}
