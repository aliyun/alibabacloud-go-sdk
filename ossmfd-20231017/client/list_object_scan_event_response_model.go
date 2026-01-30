// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectScanEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListObjectScanEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListObjectScanEventResponse
	GetStatusCode() *int32
	SetBody(v *ListObjectScanEventResponseBody) *ListObjectScanEventResponse
	GetBody() *ListObjectScanEventResponseBody
}

type ListObjectScanEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListObjectScanEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListObjectScanEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventResponse) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListObjectScanEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListObjectScanEventResponse) GetBody() *ListObjectScanEventResponseBody {
	return s.Body
}

func (s *ListObjectScanEventResponse) SetHeaders(v map[string]*string) *ListObjectScanEventResponse {
	s.Headers = v
	return s
}

func (s *ListObjectScanEventResponse) SetStatusCode(v int32) *ListObjectScanEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListObjectScanEventResponse) SetBody(v *ListObjectScanEventResponseBody) *ListObjectScanEventResponse {
	s.Body = v
	return s
}

func (s *ListObjectScanEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
