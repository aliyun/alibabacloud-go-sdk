// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageLogResponse
	GetStatusCode() *int32
	SetBody(v *GetPageLogResponseBody) *GetPageLogResponse
	GetBody() *GetPageLogResponseBody
}

type GetPageLogResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageLogResponse) GoString() string {
	return s.String()
}

func (s *GetPageLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageLogResponse) GetBody() *GetPageLogResponseBody {
	return s.Body
}

func (s *GetPageLogResponse) SetHeaders(v map[string]*string) *GetPageLogResponse {
	s.Headers = v
	return s
}

func (s *GetPageLogResponse) SetStatusCode(v int32) *GetPageLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageLogResponse) SetBody(v *GetPageLogResponseBody) *GetPageLogResponse {
	s.Body = v
	return s
}

func (s *GetPageLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
