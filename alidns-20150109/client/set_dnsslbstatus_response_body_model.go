// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDNSSLBStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpen(v bool) *SetDNSSLBStatusResponseBody
	GetOpen() *bool
	SetRecordCount(v int64) *SetDNSSLBStatusResponseBody
	GetRecordCount() *int64
	SetRequestId(v string) *SetDNSSLBStatusResponseBody
	GetRequestId() *string
}

type SetDNSSLBStatusResponseBody struct {
	// Indicates whether weighted round-robin is enabled for the subdomain name.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The number of A records that are matched.
	//
	// example:
	//
	// 8
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDNSSLBStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDNSSLBStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusResponseBody) GetOpen() *bool {
	return s.Open
}

func (s *SetDNSSLBStatusResponseBody) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *SetDNSSLBStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDNSSLBStatusResponseBody) SetOpen(v bool) *SetDNSSLBStatusResponseBody {
	s.Open = &v
	return s
}

func (s *SetDNSSLBStatusResponseBody) SetRecordCount(v int64) *SetDNSSLBStatusResponseBody {
	s.RecordCount = &v
	return s
}

func (s *SetDNSSLBStatusResponseBody) SetRequestId(v string) *SetDNSSLBStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDNSSLBStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
