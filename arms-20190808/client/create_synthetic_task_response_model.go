// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSyntheticTaskResponseBody) *CreateSyntheticTaskResponse
	GetBody() *CreateSyntheticTaskResponseBody
}

type CreateSyntheticTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSyntheticTaskResponse) GetBody() *CreateSyntheticTaskResponseBody {
	return s.Body
}

func (s *CreateSyntheticTaskResponse) SetHeaders(v map[string]*string) *CreateSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSyntheticTaskResponse) SetStatusCode(v int32) *CreateSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSyntheticTaskResponse) SetBody(v *CreateSyntheticTaskResponseBody) *CreateSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSyntheticTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
