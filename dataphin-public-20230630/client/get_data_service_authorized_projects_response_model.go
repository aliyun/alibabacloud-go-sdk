// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAuthorizedProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAuthorizedProjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAuthorizedProjectsResponseBody) *GetDataServiceAuthorizedProjectsResponse
	GetBody() *GetDataServiceAuthorizedProjectsResponseBody
}

type GetDataServiceAuthorizedProjectsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAuthorizedProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAuthorizedProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAuthorizedProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAuthorizedProjectsResponse) GetBody() *GetDataServiceAuthorizedProjectsResponseBody {
	return s.Body
}

func (s *GetDataServiceAuthorizedProjectsResponse) SetHeaders(v map[string]*string) *GetDataServiceAuthorizedProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponse) SetStatusCode(v int32) *GetDataServiceAuthorizedProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponse) SetBody(v *GetDataServiceAuthorizedProjectsResponseBody) *GetDataServiceAuthorizedProjectsResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
