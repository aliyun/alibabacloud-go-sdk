// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRequestVodTestToolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *HttpRequestVodTestToolResponseBody
	GetBody() *string
	SetHeader(v string) *HttpRequestVodTestToolResponseBody
	GetHeader() *string
	SetRequestId(v string) *HttpRequestVodTestToolResponseBody
	GetRequestId() *string
	SetStatusCode(v int32) *HttpRequestVodTestToolResponseBody
	GetStatusCode() *int32
}

type HttpRequestVodTestToolResponseBody struct {
	Body       *string `json:"Body,omitempty" xml:"Body,omitempty"`
	Header     *string `json:"Header,omitempty" xml:"Header,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s HttpRequestVodTestToolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HttpRequestVodTestToolResponseBody) GoString() string {
	return s.String()
}

func (s *HttpRequestVodTestToolResponseBody) GetBody() *string {
	return s.Body
}

func (s *HttpRequestVodTestToolResponseBody) GetHeader() *string {
	return s.Header
}

func (s *HttpRequestVodTestToolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HttpRequestVodTestToolResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HttpRequestVodTestToolResponseBody) SetBody(v string) *HttpRequestVodTestToolResponseBody {
	s.Body = &v
	return s
}

func (s *HttpRequestVodTestToolResponseBody) SetHeader(v string) *HttpRequestVodTestToolResponseBody {
	s.Header = &v
	return s
}

func (s *HttpRequestVodTestToolResponseBody) SetRequestId(v string) *HttpRequestVodTestToolResponseBody {
	s.RequestId = &v
	return s
}

func (s *HttpRequestVodTestToolResponseBody) SetStatusCode(v int32) *HttpRequestVodTestToolResponseBody {
	s.StatusCode = &v
	return s
}

func (s *HttpRequestVodTestToolResponseBody) Validate() error {
	return dara.Validate(s)
}
