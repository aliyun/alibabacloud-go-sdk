// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamJobManagerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamJobManagerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamJobManagerLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamJobManagerLogsResponseBody) *ListStreamJobManagerLogsResponse
	GetBody() *ListStreamJobManagerLogsResponseBody
}

type ListStreamJobManagerLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamJobManagerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamJobManagerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamJobManagerLogsResponse) GoString() string {
	return s.String()
}

func (s *ListStreamJobManagerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamJobManagerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamJobManagerLogsResponse) GetBody() *ListStreamJobManagerLogsResponseBody {
	return s.Body
}

func (s *ListStreamJobManagerLogsResponse) SetHeaders(v map[string]*string) *ListStreamJobManagerLogsResponse {
	s.Headers = v
	return s
}

func (s *ListStreamJobManagerLogsResponse) SetStatusCode(v int32) *ListStreamJobManagerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamJobManagerLogsResponse) SetBody(v *ListStreamJobManagerLogsResponseBody) *ListStreamJobManagerLogsResponse {
	s.Body = v
	return s
}

func (s *ListStreamJobManagerLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
