// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncThirdUserMappingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SyncThirdUserMappingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *SyncThirdUserMappingHeaders
	GetXAcsBtripCorpToken() *string
}

type SyncThirdUserMappingHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s SyncThirdUserMappingHeaders) String() string {
	return dara.Prettify(s)
}

func (s SyncThirdUserMappingHeaders) GoString() string {
	return s.String()
}

func (s *SyncThirdUserMappingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SyncThirdUserMappingHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *SyncThirdUserMappingHeaders) SetCommonHeaders(v map[string]*string) *SyncThirdUserMappingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncThirdUserMappingHeaders) SetXAcsBtripCorpToken(v string) *SyncThirdUserMappingHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *SyncThirdUserMappingHeaders) Validate() error {
	return dara.Validate(s)
}
