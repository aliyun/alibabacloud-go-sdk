// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskAttributeResponseBody) *GetTaskAttributeResponse
	GetBody() *GetTaskAttributeResponseBody
}

type GetTaskAttributeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskAttributeResponse) GetBody() *GetTaskAttributeResponseBody {
	return s.Body
}

func (s *GetTaskAttributeResponse) SetHeaders(v map[string]*string) *GetTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetTaskAttributeResponse) SetStatusCode(v int32) *GetTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskAttributeResponse) SetBody(v *GetTaskAttributeResponseBody) *GetTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *GetTaskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
