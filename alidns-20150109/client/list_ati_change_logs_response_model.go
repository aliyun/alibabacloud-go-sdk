// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiChangeLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAtiChangeLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAtiChangeLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListAtiChangeLogsResponseBody) *ListAtiChangeLogsResponse
	GetBody() *ListAtiChangeLogsResponseBody
}

type ListAtiChangeLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAtiChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAtiChangeLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAtiChangeLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAtiChangeLogsResponse) GetBody() *ListAtiChangeLogsResponseBody {
	return s.Body
}

func (s *ListAtiChangeLogsResponse) SetHeaders(v map[string]*string) *ListAtiChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAtiChangeLogsResponse) SetStatusCode(v int32) *ListAtiChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAtiChangeLogsResponse) SetBody(v *ListAtiChangeLogsResponseBody) *ListAtiChangeLogsResponse {
	s.Body = v
	return s
}

func (s *ListAtiChangeLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
