// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsLogDispatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlsLogDispatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlsLogDispatchResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlsLogDispatchResponseBody) *CreateSlsLogDispatchResponse
	GetBody() *CreateSlsLogDispatchResponseBody
}

type CreateSlsLogDispatchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlsLogDispatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlsLogDispatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsLogDispatchResponse) GoString() string {
	return s.String()
}

func (s *CreateSlsLogDispatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlsLogDispatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlsLogDispatchResponse) GetBody() *CreateSlsLogDispatchResponseBody {
	return s.Body
}

func (s *CreateSlsLogDispatchResponse) SetHeaders(v map[string]*string) *CreateSlsLogDispatchResponse {
	s.Headers = v
	return s
}

func (s *CreateSlsLogDispatchResponse) SetStatusCode(v int32) *CreateSlsLogDispatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlsLogDispatchResponse) SetBody(v *CreateSlsLogDispatchResponseBody) *CreateSlsLogDispatchResponse {
	s.Body = v
	return s
}

func (s *CreateSlsLogDispatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
