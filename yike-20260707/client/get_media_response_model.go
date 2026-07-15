// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaResponseBody) *GetMediaResponse
	GetBody() *GetMediaResponseBody
}

type GetMediaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponse) GoString() string {
	return s.String()
}

func (s *GetMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaResponse) GetBody() *GetMediaResponseBody {
	return s.Body
}

func (s *GetMediaResponse) SetHeaders(v map[string]*string) *GetMediaResponse {
	s.Headers = v
	return s
}

func (s *GetMediaResponse) SetStatusCode(v int32) *GetMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaResponse) SetBody(v *GetMediaResponseBody) *GetMediaResponse {
	s.Body = v
	return s
}

func (s *GetMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
