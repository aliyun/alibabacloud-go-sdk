// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7UsKeepaliveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeL7UsKeepaliveResponseBody
	GetRequestId() *string
	SetRsKeepalive(v *DescribeL7UsKeepaliveResponseBodyRsKeepalive) *DescribeL7UsKeepaliveResponseBody
	GetRsKeepalive() *DescribeL7UsKeepaliveResponseBodyRsKeepalive
}

type DescribeL7UsKeepaliveResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of the Back-to-origin Persistent Connections parameter.
	RsKeepalive *DescribeL7UsKeepaliveResponseBodyRsKeepalive `json:"RsKeepalive,omitempty" xml:"RsKeepalive,omitempty" type:"Struct"`
}

func (s DescribeL7UsKeepaliveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7UsKeepaliveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL7UsKeepaliveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeL7UsKeepaliveResponseBody) GetRsKeepalive() *DescribeL7UsKeepaliveResponseBodyRsKeepalive {
	return s.RsKeepalive
}

func (s *DescribeL7UsKeepaliveResponseBody) SetRequestId(v string) *DescribeL7UsKeepaliveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBody) SetRsKeepalive(v *DescribeL7UsKeepaliveResponseBodyRsKeepalive) *DescribeL7UsKeepaliveResponseBody {
	s.RsKeepalive = v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBody) Validate() error {
	if s.RsKeepalive != nil {
		if err := s.RsKeepalive.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeL7UsKeepaliveResponseBodyRsKeepalive struct {
	DsKeepaliveTimeout *int64 `json:"DsKeepaliveTimeout,omitempty" xml:"DsKeepaliveTimeout,omitempty"`
	// Indicates whether Back-to-origin Persistent Connections is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The number of requests that reuse persistent connections.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int64 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of idle persistent connections.
	//
	// example:
	//
	// 30
	KeepaliveTimeout *int64 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
}

func (s DescribeL7UsKeepaliveResponseBodyRsKeepalive) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7UsKeepaliveResponseBodyRsKeepalive) GoString() string {
	return s.String()
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) GetDsKeepaliveTimeout() *int64 {
	return s.DsKeepaliveTimeout
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) GetKeepaliveRequests() *int64 {
	return s.KeepaliveRequests
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) GetKeepaliveTimeout() *int64 {
	return s.KeepaliveTimeout
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) SetDsKeepaliveTimeout(v int64) *DescribeL7UsKeepaliveResponseBodyRsKeepalive {
	s.DsKeepaliveTimeout = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) SetEnabled(v bool) *DescribeL7UsKeepaliveResponseBodyRsKeepalive {
	s.Enabled = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) SetKeepaliveRequests(v int64) *DescribeL7UsKeepaliveResponseBodyRsKeepalive {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) SetKeepaliveTimeout(v int64) *DescribeL7UsKeepaliveResponseBodyRsKeepalive {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeL7UsKeepaliveResponseBodyRsKeepalive) Validate() error {
	return dara.Validate(s)
}
