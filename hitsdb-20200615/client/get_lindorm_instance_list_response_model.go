// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormInstanceListResponseBody) *GetLindormInstanceListResponse
	GetBody() *GetLindormInstanceListResponseBody
}

type GetLindormInstanceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormInstanceListResponse) GetBody() *GetLindormInstanceListResponseBody {
	return s.Body
}

func (s *GetLindormInstanceListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceListResponse) SetStatusCode(v int32) *GetLindormInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceListResponse) SetBody(v *GetLindormInstanceListResponseBody) *GetLindormInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetLindormInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
