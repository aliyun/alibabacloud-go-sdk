// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMiningTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMiningTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMiningTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMiningTaskResponseBody) *CreateMiningTaskResponse
	GetBody() *CreateMiningTaskResponseBody
}

type CreateMiningTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMiningTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMiningTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMiningTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMiningTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMiningTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMiningTaskResponse) GetBody() *CreateMiningTaskResponseBody {
	return s.Body
}

func (s *CreateMiningTaskResponse) SetHeaders(v map[string]*string) *CreateMiningTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMiningTaskResponse) SetStatusCode(v int32) *CreateMiningTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMiningTaskResponse) SetBody(v *CreateMiningTaskResponseBody) *CreateMiningTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMiningTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
