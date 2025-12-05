// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllRegionsResponse
	GetStatusCode() *int32
	SetBody(v *GetAllRegionsResponseBody) *GetAllRegionsResponse
	GetBody() *GetAllRegionsResponseBody
}

type GetAllRegionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllRegionsResponse) GoString() string {
	return s.String()
}

func (s *GetAllRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllRegionsResponse) GetBody() *GetAllRegionsResponseBody {
	return s.Body
}

func (s *GetAllRegionsResponse) SetHeaders(v map[string]*string) *GetAllRegionsResponse {
	s.Headers = v
	return s
}

func (s *GetAllRegionsResponse) SetStatusCode(v int32) *GetAllRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllRegionsResponse) SetBody(v *GetAllRegionsResponseBody) *GetAllRegionsResponse {
	s.Body = v
	return s
}

func (s *GetAllRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
