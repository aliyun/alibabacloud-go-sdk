// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DuplicateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DuplicateModelResponse
	GetStatusCode() *int32
	SetBody(v *DuplicateModelResponseBody) *DuplicateModelResponse
	GetBody() *DuplicateModelResponseBody
}

type DuplicateModelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DuplicateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DuplicateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DuplicateModelResponse) GoString() string {
	return s.String()
}

func (s *DuplicateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DuplicateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DuplicateModelResponse) GetBody() *DuplicateModelResponseBody {
	return s.Body
}

func (s *DuplicateModelResponse) SetHeaders(v map[string]*string) *DuplicateModelResponse {
	s.Headers = v
	return s
}

func (s *DuplicateModelResponse) SetStatusCode(v int32) *DuplicateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DuplicateModelResponse) SetBody(v *DuplicateModelResponseBody) *DuplicateModelResponse {
	s.Body = v
	return s
}

func (s *DuplicateModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
