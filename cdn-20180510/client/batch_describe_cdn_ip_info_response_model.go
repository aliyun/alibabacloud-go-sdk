// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDescribeCdnIpInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDescribeCdnIpInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDescribeCdnIpInfoResponse
	GetStatusCode() *int32
	SetBody(v *BatchDescribeCdnIpInfoResponseBody) *BatchDescribeCdnIpInfoResponse
	GetBody() *BatchDescribeCdnIpInfoResponseBody
}

type BatchDescribeCdnIpInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDescribeCdnIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDescribeCdnIpInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDescribeCdnIpInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchDescribeCdnIpInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDescribeCdnIpInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDescribeCdnIpInfoResponse) GetBody() *BatchDescribeCdnIpInfoResponseBody {
	return s.Body
}

func (s *BatchDescribeCdnIpInfoResponse) SetHeaders(v map[string]*string) *BatchDescribeCdnIpInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchDescribeCdnIpInfoResponse) SetStatusCode(v int32) *BatchDescribeCdnIpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponse) SetBody(v *BatchDescribeCdnIpInfoResponseBody) *BatchDescribeCdnIpInfoResponse {
	s.Body = v
	return s
}

func (s *BatchDescribeCdnIpInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
