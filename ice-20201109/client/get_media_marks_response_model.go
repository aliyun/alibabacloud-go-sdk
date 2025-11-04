// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaMarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaMarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaMarksResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaMarksResponseBody) *GetMediaMarksResponse
	GetBody() *GetMediaMarksResponseBody
}

type GetMediaMarksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaMarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaMarksResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaMarksResponse) GoString() string {
	return s.String()
}

func (s *GetMediaMarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaMarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaMarksResponse) GetBody() *GetMediaMarksResponseBody {
	return s.Body
}

func (s *GetMediaMarksResponse) SetHeaders(v map[string]*string) *GetMediaMarksResponse {
	s.Headers = v
	return s
}

func (s *GetMediaMarksResponse) SetStatusCode(v int32) *GetMediaMarksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaMarksResponse) SetBody(v *GetMediaMarksResponseBody) *GetMediaMarksResponse {
	s.Body = v
	return s
}

func (s *GetMediaMarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
