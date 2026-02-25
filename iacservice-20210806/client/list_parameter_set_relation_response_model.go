// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListParameterSetRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListParameterSetRelationResponse
	GetStatusCode() *int32
	SetBody(v *ListParameterSetRelationResponseBody) *ListParameterSetRelationResponse
	GetBody() *ListParameterSetRelationResponseBody
}

type ListParameterSetRelationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterSetRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterSetRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetRelationResponse) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListParameterSetRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListParameterSetRelationResponse) GetBody() *ListParameterSetRelationResponseBody {
	return s.Body
}

func (s *ListParameterSetRelationResponse) SetHeaders(v map[string]*string) *ListParameterSetRelationResponse {
	s.Headers = v
	return s
}

func (s *ListParameterSetRelationResponse) SetStatusCode(v int32) *ListParameterSetRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterSetRelationResponse) SetBody(v *ListParameterSetRelationResponseBody) *ListParameterSetRelationResponse {
	s.Body = v
	return s
}

func (s *ListParameterSetRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
