// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUnionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUnionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateUnionTaskResponseBody) *CreateUnionTaskResponse
	GetBody() *CreateUnionTaskResponseBody
}

type CreateUnionTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUnionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUnionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUnionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUnionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUnionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUnionTaskResponse) GetBody() *CreateUnionTaskResponseBody {
	return s.Body
}

func (s *CreateUnionTaskResponse) SetHeaders(v map[string]*string) *CreateUnionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUnionTaskResponse) SetStatusCode(v int32) *CreateUnionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUnionTaskResponse) SetBody(v *CreateUnionTaskResponseBody) *CreateUnionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateUnionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
