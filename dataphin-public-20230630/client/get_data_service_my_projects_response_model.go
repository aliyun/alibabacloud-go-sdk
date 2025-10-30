// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceMyProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceMyProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceMyProjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceMyProjectsResponseBody) *GetDataServiceMyProjectsResponse
	GetBody() *GetDataServiceMyProjectsResponseBody
}

type GetDataServiceMyProjectsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceMyProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceMyProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceMyProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceMyProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceMyProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceMyProjectsResponse) GetBody() *GetDataServiceMyProjectsResponseBody {
	return s.Body
}

func (s *GetDataServiceMyProjectsResponse) SetHeaders(v map[string]*string) *GetDataServiceMyProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceMyProjectsResponse) SetStatusCode(v int32) *GetDataServiceMyProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceMyProjectsResponse) SetBody(v *GetDataServiceMyProjectsResponseBody) *GetDataServiceMyProjectsResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceMyProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
