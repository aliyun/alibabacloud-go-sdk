// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiChannelRecordingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiChannelRecordingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiChannelRecordingsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiChannelRecordingsResponseBody) *ListMultiChannelRecordingsResponse
	GetBody() *ListMultiChannelRecordingsResponseBody
}

type ListMultiChannelRecordingsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiChannelRecordingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiChannelRecordingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiChannelRecordingsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiChannelRecordingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiChannelRecordingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiChannelRecordingsResponse) GetBody() *ListMultiChannelRecordingsResponseBody {
	return s.Body
}

func (s *ListMultiChannelRecordingsResponse) SetHeaders(v map[string]*string) *ListMultiChannelRecordingsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiChannelRecordingsResponse) SetStatusCode(v int32) *ListMultiChannelRecordingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiChannelRecordingsResponse) SetBody(v *ListMultiChannelRecordingsResponseBody) *ListMultiChannelRecordingsResponse {
	s.Body = v
	return s
}

func (s *ListMultiChannelRecordingsResponse) Validate() error {
	return dara.Validate(s)
}
