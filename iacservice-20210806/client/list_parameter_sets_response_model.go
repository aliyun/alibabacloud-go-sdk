// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListParameterSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListParameterSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListParameterSetsResponseBody) *ListParameterSetsResponse
	GetBody() *ListParameterSetsResponseBody
}

type ListParameterSetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsResponse) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListParameterSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListParameterSetsResponse) GetBody() *ListParameterSetsResponseBody {
	return s.Body
}

func (s *ListParameterSetsResponse) SetHeaders(v map[string]*string) *ListParameterSetsResponse {
	s.Headers = v
	return s
}

func (s *ListParameterSetsResponse) SetStatusCode(v int32) *ListParameterSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterSetsResponse) SetBody(v *ListParameterSetsResponseBody) *ListParameterSetsResponse {
	s.Body = v
	return s
}

func (s *ListParameterSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
