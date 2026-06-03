// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogListResponse
	GetStatusCode() *int32
	SetBody(v *GetLogListResponseBody) *GetLogListResponse
	GetBody() *GetLogListResponseBody
}

type GetLogListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogListResponse) GoString() string {
	return s.String()
}

func (s *GetLogListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogListResponse) GetBody() *GetLogListResponseBody {
	return s.Body
}

func (s *GetLogListResponse) SetHeaders(v map[string]*string) *GetLogListResponse {
	s.Headers = v
	return s
}

func (s *GetLogListResponse) SetStatusCode(v int32) *GetLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogListResponse) SetBody(v *GetLogListResponseBody) *GetLogListResponse {
	s.Body = v
	return s
}

func (s *GetLogListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
