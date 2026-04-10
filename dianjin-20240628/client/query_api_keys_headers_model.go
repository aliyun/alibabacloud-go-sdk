// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApiKeysHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryApiKeysHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *QueryApiKeysHeaders
	GetXLoadTest() *bool
}

type QueryApiKeysHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// true
	XLoadTest *bool `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s QueryApiKeysHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryApiKeysHeaders) GoString() string {
	return s.String()
}

func (s *QueryApiKeysHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryApiKeysHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *QueryApiKeysHeaders) SetCommonHeaders(v map[string]*string) *QueryApiKeysHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryApiKeysHeaders) SetXLoadTest(v bool) *QueryApiKeysHeaders {
	s.XLoadTest = &v
	return s
}

func (s *QueryApiKeysHeaders) Validate() error {
	return dara.Validate(s)
}
