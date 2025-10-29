// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpApiResponseBody) *GetHttpApiResponse
	GetBody() *GetHttpApiResponseBody
}

type GetHttpApiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpApiResponse) GetBody() *GetHttpApiResponseBody {
	return s.Body
}

func (s *GetHttpApiResponse) SetHeaders(v map[string]*string) *GetHttpApiResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiResponse) SetStatusCode(v int32) *GetHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiResponse) SetBody(v *GetHttpApiResponseBody) *GetHttpApiResponse {
	s.Body = v
	return s
}

func (s *GetHttpApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
