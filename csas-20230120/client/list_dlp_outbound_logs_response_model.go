// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDlpOutboundLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDlpOutboundLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDlpOutboundLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListDlpOutboundLogsResponseBody) *ListDlpOutboundLogsResponse
	GetBody() *ListDlpOutboundLogsResponseBody
}

type ListDlpOutboundLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDlpOutboundLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDlpOutboundLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponse) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDlpOutboundLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDlpOutboundLogsResponse) GetBody() *ListDlpOutboundLogsResponseBody {
	return s.Body
}

func (s *ListDlpOutboundLogsResponse) SetHeaders(v map[string]*string) *ListDlpOutboundLogsResponse {
	s.Headers = v
	return s
}

func (s *ListDlpOutboundLogsResponse) SetStatusCode(v int32) *ListDlpOutboundLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDlpOutboundLogsResponse) SetBody(v *ListDlpOutboundLogsResponseBody) *ListDlpOutboundLogsResponse {
	s.Body = v
	return s
}

func (s *ListDlpOutboundLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
