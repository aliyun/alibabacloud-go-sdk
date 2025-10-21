// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContentDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContentDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *GetContentDetectResultResponseBody) *GetContentDetectResultResponse
	GetBody() *GetContentDetectResultResponseBody
}

type GetContentDetectResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContentDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContentDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContentDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContentDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContentDetectResultResponse) GetBody() *GetContentDetectResultResponseBody {
	return s.Body
}

func (s *GetContentDetectResultResponse) SetHeaders(v map[string]*string) *GetContentDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetContentDetectResultResponse) SetStatusCode(v int32) *GetContentDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContentDetectResultResponse) SetBody(v *GetContentDetectResultResponseBody) *GetContentDetectResultResponse {
	s.Body = v
	return s
}

func (s *GetContentDetectResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
