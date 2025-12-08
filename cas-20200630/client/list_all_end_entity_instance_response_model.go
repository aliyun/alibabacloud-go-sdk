// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllEndEntityInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllEndEntityInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllEndEntityInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListAllEndEntityInstanceResponseBody) *ListAllEndEntityInstanceResponse
	GetBody() *ListAllEndEntityInstanceResponseBody
}

type ListAllEndEntityInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllEndEntityInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllEndEntityInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllEndEntityInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListAllEndEntityInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllEndEntityInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllEndEntityInstanceResponse) GetBody() *ListAllEndEntityInstanceResponseBody {
	return s.Body
}

func (s *ListAllEndEntityInstanceResponse) SetHeaders(v map[string]*string) *ListAllEndEntityInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListAllEndEntityInstanceResponse) SetStatusCode(v int32) *ListAllEndEntityInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllEndEntityInstanceResponse) SetBody(v *ListAllEndEntityInstanceResponseBody) *ListAllEndEntityInstanceResponse {
	s.Body = v
	return s
}

func (s *ListAllEndEntityInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
