// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPodLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPodLogsResponse
	GetStatusCode() *int32
	SetBody(v *GetPodLogsResponseBody) *GetPodLogsResponse
	GetBody() *GetPodLogsResponseBody
}

type GetPodLogsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPodLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPodLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPodLogsResponse) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPodLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPodLogsResponse) GetBody() *GetPodLogsResponseBody {
	return s.Body
}

func (s *GetPodLogsResponse) SetHeaders(v map[string]*string) *GetPodLogsResponse {
	s.Headers = v
	return s
}

func (s *GetPodLogsResponse) SetStatusCode(v int32) *GetPodLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodLogsResponse) SetBody(v *GetPodLogsResponseBody) *GetPodLogsResponse {
	s.Body = v
	return s
}

func (s *GetPodLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
