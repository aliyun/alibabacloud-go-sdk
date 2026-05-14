// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceModelResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceModelResponseBody) *ListInstanceModelResponse
	GetBody() *ListInstanceModelResponseBody
}

type ListInstanceModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceModelResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceModelResponse) GetBody() *ListInstanceModelResponseBody {
	return s.Body
}

func (s *ListInstanceModelResponse) SetHeaders(v map[string]*string) *ListInstanceModelResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceModelResponse) SetStatusCode(v int32) *ListInstanceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceModelResponse) SetBody(v *ListInstanceModelResponseBody) *ListInstanceModelResponse {
	s.Body = v
	return s
}

func (s *ListInstanceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
