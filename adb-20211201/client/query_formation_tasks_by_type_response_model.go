// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTasksByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFormationTasksByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFormationTasksByTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryFormationTasksByTypeResponseBody) *QueryFormationTasksByTypeResponse
	GetBody() *QueryFormationTasksByTypeResponseBody
}

type QueryFormationTasksByTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFormationTasksByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFormationTasksByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTasksByTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryFormationTasksByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFormationTasksByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFormationTasksByTypeResponse) GetBody() *QueryFormationTasksByTypeResponseBody {
	return s.Body
}

func (s *QueryFormationTasksByTypeResponse) SetHeaders(v map[string]*string) *QueryFormationTasksByTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryFormationTasksByTypeResponse) SetStatusCode(v int32) *QueryFormationTasksByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFormationTasksByTypeResponse) SetBody(v *QueryFormationTasksByTypeResponseBody) *QueryFormationTasksByTypeResponse {
	s.Body = v
	return s
}

func (s *QueryFormationTasksByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
