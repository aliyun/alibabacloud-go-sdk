// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessAssignmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessAssignmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessAssignmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessAssignmentResponseBody) *DeleteAccessAssignmentResponse
	GetBody() *DeleteAccessAssignmentResponseBody
}

type DeleteAccessAssignmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessAssignmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessAssignmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessAssignmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessAssignmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessAssignmentResponse) GetBody() *DeleteAccessAssignmentResponseBody {
	return s.Body
}

func (s *DeleteAccessAssignmentResponse) SetHeaders(v map[string]*string) *DeleteAccessAssignmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessAssignmentResponse) SetStatusCode(v int32) *DeleteAccessAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessAssignmentResponse) SetBody(v *DeleteAccessAssignmentResponseBody) *DeleteAccessAssignmentResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessAssignmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
