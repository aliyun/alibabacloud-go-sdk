// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetImageInfosResponseBody) *GetImageInfosResponse
	GetBody() *GetImageInfosResponseBody
}

type GetImageInfosResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfosResponse) GoString() string {
	return s.String()
}

func (s *GetImageInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageInfosResponse) GetBody() *GetImageInfosResponseBody {
	return s.Body
}

func (s *GetImageInfosResponse) SetHeaders(v map[string]*string) *GetImageInfosResponse {
	s.Headers = v
	return s
}

func (s *GetImageInfosResponse) SetStatusCode(v int32) *GetImageInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageInfosResponse) SetBody(v *GetImageInfosResponseBody) *GetImageInfosResponse {
	s.Body = v
	return s
}

func (s *GetImageInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
