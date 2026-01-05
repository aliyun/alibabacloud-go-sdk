// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConstraintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConstraintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConstraintResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConstraintResponseBody) *DeleteConstraintResponse
	GetBody() *DeleteConstraintResponseBody
}

type DeleteConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConstraintResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConstraintResponse) GoString() string {
	return s.String()
}

func (s *DeleteConstraintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConstraintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConstraintResponse) GetBody() *DeleteConstraintResponseBody {
	return s.Body
}

func (s *DeleteConstraintResponse) SetHeaders(v map[string]*string) *DeleteConstraintResponse {
	s.Headers = v
	return s
}

func (s *DeleteConstraintResponse) SetStatusCode(v int32) *DeleteConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConstraintResponse) SetBody(v *DeleteConstraintResponseBody) *DeleteConstraintResponse {
	s.Body = v
	return s
}

func (s *DeleteConstraintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
