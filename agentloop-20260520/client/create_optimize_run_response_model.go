// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOptimizeRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOptimizeRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOptimizeRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateOptimizeRunResponseBody) *CreateOptimizeRunResponse
	GetBody() *CreateOptimizeRunResponseBody
}

type CreateOptimizeRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOptimizeRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOptimizeRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOptimizeRunResponse) GoString() string {
	return s.String()
}

func (s *CreateOptimizeRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOptimizeRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOptimizeRunResponse) GetBody() *CreateOptimizeRunResponseBody {
	return s.Body
}

func (s *CreateOptimizeRunResponse) SetHeaders(v map[string]*string) *CreateOptimizeRunResponse {
	s.Headers = v
	return s
}

func (s *CreateOptimizeRunResponse) SetStatusCode(v int32) *CreateOptimizeRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOptimizeRunResponse) SetBody(v *CreateOptimizeRunResponseBody) *CreateOptimizeRunResponse {
	s.Body = v
	return s
}

func (s *CreateOptimizeRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
