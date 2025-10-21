// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListModelInstanceResponseBody) *ListModelInstanceResponse
	GetBody() *ListModelInstanceResponseBody
}

type ListModelInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListModelInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelInstanceResponse) GetBody() *ListModelInstanceResponseBody {
	return s.Body
}

func (s *ListModelInstanceResponse) SetHeaders(v map[string]*string) *ListModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListModelInstanceResponse) SetStatusCode(v int32) *ListModelInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelInstanceResponse) SetBody(v *ListModelInstanceResponseBody) *ListModelInstanceResponse {
	s.Body = v
	return s
}

func (s *ListModelInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
