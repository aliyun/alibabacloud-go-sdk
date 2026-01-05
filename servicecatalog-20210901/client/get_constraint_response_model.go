// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConstraintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConstraintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConstraintResponse
	GetStatusCode() *int32
	SetBody(v *GetConstraintResponseBody) *GetConstraintResponse
	GetBody() *GetConstraintResponseBody
}

type GetConstraintResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConstraintResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConstraintResponse) GoString() string {
	return s.String()
}

func (s *GetConstraintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConstraintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConstraintResponse) GetBody() *GetConstraintResponseBody {
	return s.Body
}

func (s *GetConstraintResponse) SetHeaders(v map[string]*string) *GetConstraintResponse {
	s.Headers = v
	return s
}

func (s *GetConstraintResponse) SetStatusCode(v int32) *GetConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConstraintResponse) SetBody(v *GetConstraintResponseBody) *GetConstraintResponse {
	s.Body = v
	return s
}

func (s *GetConstraintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
