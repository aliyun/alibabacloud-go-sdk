// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTaskByIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFormationTaskByIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFormationTaskByIDResponse
	GetStatusCode() *int32
	SetBody(v *QueryFormationTaskByIDResponseBody) *QueryFormationTaskByIDResponse
	GetBody() *QueryFormationTaskByIDResponseBody
}

type QueryFormationTaskByIDResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFormationTaskByIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFormationTaskByIDResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTaskByIDResponse) GoString() string {
	return s.String()
}

func (s *QueryFormationTaskByIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFormationTaskByIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFormationTaskByIDResponse) GetBody() *QueryFormationTaskByIDResponseBody {
	return s.Body
}

func (s *QueryFormationTaskByIDResponse) SetHeaders(v map[string]*string) *QueryFormationTaskByIDResponse {
	s.Headers = v
	return s
}

func (s *QueryFormationTaskByIDResponse) SetStatusCode(v int32) *QueryFormationTaskByIDResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFormationTaskByIDResponse) SetBody(v *QueryFormationTaskByIDResponseBody) *QueryFormationTaskByIDResponse {
	s.Body = v
	return s
}

func (s *QueryFormationTaskByIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
