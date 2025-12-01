// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetDetailByIdResponseBody) *GetDetailByIdResponse
	GetBody() *GetDetailByIdResponseBody
}

type GetDetailByIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetailByIdResponse) GetBody() *GetDetailByIdResponseBody {
	return s.Body
}

func (s *GetDetailByIdResponse) SetHeaders(v map[string]*string) *GetDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetDetailByIdResponse) SetStatusCode(v int32) *GetDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetailByIdResponse) SetBody(v *GetDetailByIdResponseBody) *GetDetailByIdResponse {
	s.Body = v
	return s
}

func (s *GetDetailByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
