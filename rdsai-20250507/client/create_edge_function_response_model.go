// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeFunctionResponseBody) *CreateEdgeFunctionResponse
	GetBody() *CreateEdgeFunctionResponseBody
}

type CreateEdgeFunctionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeFunctionResponse) GetBody() *CreateEdgeFunctionResponseBody {
	return s.Body
}

func (s *CreateEdgeFunctionResponse) SetHeaders(v map[string]*string) *CreateEdgeFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeFunctionResponse) SetStatusCode(v int32) *CreateEdgeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeFunctionResponse) SetBody(v *CreateEdgeFunctionResponseBody) *CreateEdgeFunctionResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
