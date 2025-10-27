// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormInstanceResponseBody) *GetLindormInstanceResponse
	GetBody() *GetLindormInstanceResponseBody
}

type GetLindormInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormInstanceResponse) GetBody() *GetLindormInstanceResponseBody {
	return s.Body
}

func (s *GetLindormInstanceResponse) SetHeaders(v map[string]*string) *GetLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceResponse) SetStatusCode(v int32) *GetLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceResponse) SetBody(v *GetLindormInstanceResponseBody) *GetLindormInstanceResponse {
	s.Body = v
	return s
}

func (s *GetLindormInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
