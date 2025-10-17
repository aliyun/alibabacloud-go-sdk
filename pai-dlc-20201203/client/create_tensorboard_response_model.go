// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *CreateTensorboardResponseBody) *CreateTensorboardResponse
	GetBody() *CreateTensorboardResponseBody
}

type CreateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTensorboardResponse) GetBody() *CreateTensorboardResponseBody {
	return s.Body
}

func (s *CreateTensorboardResponse) SetHeaders(v map[string]*string) *CreateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *CreateTensorboardResponse) SetStatusCode(v int32) *CreateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTensorboardResponse) SetBody(v *CreateTensorboardResponseBody) *CreateTensorboardResponse {
	s.Body = v
	return s
}

func (s *CreateTensorboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
