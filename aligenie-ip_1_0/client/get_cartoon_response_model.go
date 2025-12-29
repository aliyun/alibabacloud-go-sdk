// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCartoonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCartoonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCartoonResponse
	GetStatusCode() *int32
	SetBody(v *GetCartoonResponseBody) *GetCartoonResponse
	GetBody() *GetCartoonResponseBody
}

type GetCartoonResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCartoonResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCartoonResponse) GoString() string {
	return s.String()
}

func (s *GetCartoonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCartoonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCartoonResponse) GetBody() *GetCartoonResponseBody {
	return s.Body
}

func (s *GetCartoonResponse) SetHeaders(v map[string]*string) *GetCartoonResponse {
	s.Headers = v
	return s
}

func (s *GetCartoonResponse) SetStatusCode(v int32) *GetCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCartoonResponse) SetBody(v *GetCartoonResponseBody) *GetCartoonResponse {
	s.Body = v
	return s
}

func (s *GetCartoonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
