// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupDepartSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GroupDepartSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *GroupDepartSaveHeaders
	GetXAcsBtripCorpToken() *string
}

type GroupDepartSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s GroupDepartSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s GroupDepartSaveHeaders) GoString() string {
	return s.String()
}

func (s *GroupDepartSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GroupDepartSaveHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *GroupDepartSaveHeaders) SetCommonHeaders(v map[string]*string) *GroupDepartSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupDepartSaveHeaders) SetXAcsBtripCorpToken(v string) *GroupDepartSaveHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *GroupDepartSaveHeaders) Validate() error {
	return dara.Validate(s)
}
