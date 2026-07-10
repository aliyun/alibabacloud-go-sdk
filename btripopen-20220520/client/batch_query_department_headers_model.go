// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryDepartmentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchQueryDepartmentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *BatchQueryDepartmentHeaders
	GetXAcsBtripCorpToken() *string
}

type BatchQueryDepartmentHeaders struct {
	CommonHeaders      map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsBtripCorpToken *string            `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s BatchQueryDepartmentHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchQueryDepartmentHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *BatchQueryDepartmentHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryDepartmentHeaders) SetXAcsBtripCorpToken(v string) *BatchQueryDepartmentHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *BatchQueryDepartmentHeaders) Validate() error {
	return dara.Validate(s)
}
