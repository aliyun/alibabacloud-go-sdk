// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceDependenciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceDependenciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceDependenciesResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceDependenciesResponseBody) *GetDataSourceDependenciesResponse
	GetBody() *GetDataSourceDependenciesResponseBody
}

type GetDataSourceDependenciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceDependenciesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDependenciesResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceDependenciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceDependenciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceDependenciesResponse) GetBody() *GetDataSourceDependenciesResponseBody {
	return s.Body
}

func (s *GetDataSourceDependenciesResponse) SetHeaders(v map[string]*string) *GetDataSourceDependenciesResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceDependenciesResponse) SetStatusCode(v int32) *GetDataSourceDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceDependenciesResponse) SetBody(v *GetDataSourceDependenciesResponseBody) *GetDataSourceDependenciesResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceDependenciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
