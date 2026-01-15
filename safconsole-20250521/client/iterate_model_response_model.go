// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIterateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IterateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IterateModelResponse
	GetStatusCode() *int32
	SetBody(v *IterateModelResponseBody) *IterateModelResponse
	GetBody() *IterateModelResponseBody
}

type IterateModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IterateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IterateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s IterateModelResponse) GoString() string {
	return s.String()
}

func (s *IterateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IterateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IterateModelResponse) GetBody() *IterateModelResponseBody {
	return s.Body
}

func (s *IterateModelResponse) SetHeaders(v map[string]*string) *IterateModelResponse {
	s.Headers = v
	return s
}

func (s *IterateModelResponse) SetStatusCode(v int32) *IterateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *IterateModelResponse) SetBody(v *IterateModelResponseBody) *IterateModelResponse {
	s.Body = v
	return s
}

func (s *IterateModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
