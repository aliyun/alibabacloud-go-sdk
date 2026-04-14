// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackResponse
	GetStatusCode() *int32
	SetBody(v *GetStackResponseBody) *GetStackResponse
	GetBody() *GetStackResponseBody
}

type GetStackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponse) GoString() string {
	return s.String()
}

func (s *GetStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackResponse) GetBody() *GetStackResponseBody {
	return s.Body
}

func (s *GetStackResponse) SetHeaders(v map[string]*string) *GetStackResponse {
	s.Headers = v
	return s
}

func (s *GetStackResponse) SetStatusCode(v int32) *GetStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResponse) SetBody(v *GetStackResponseBody) *GetStackResponse {
	s.Body = v
	return s
}

func (s *GetStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
