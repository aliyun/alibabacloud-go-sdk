// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedemptionRecordsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRedemptionRecordsHeaders
	GetCommonHeaders() map[string]*string
	SetXLoadTest(v bool) *QueryRedemptionRecordsHeaders
	GetXLoadTest() *bool
}

type QueryRedemptionRecordsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// false
	XLoadTest *bool `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s QueryRedemptionRecordsHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsHeaders) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRedemptionRecordsHeaders) GetXLoadTest() *bool {
	return s.XLoadTest
}

func (s *QueryRedemptionRecordsHeaders) SetCommonHeaders(v map[string]*string) *QueryRedemptionRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRedemptionRecordsHeaders) SetXLoadTest(v bool) *QueryRedemptionRecordsHeaders {
	s.XLoadTest = &v
	return s
}

func (s *QueryRedemptionRecordsHeaders) Validate() error {
	return dara.Validate(s)
}
