// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaTableResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaTableResponseBody) *GetLumaTableResponse
	GetBody() *GetLumaTableResponseBody
}

type GetLumaTableResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaTableResponse) GoString() string {
	return s.String()
}

func (s *GetLumaTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaTableResponse) GetBody() *GetLumaTableResponseBody {
	return s.Body
}

func (s *GetLumaTableResponse) SetHeaders(v map[string]*string) *GetLumaTableResponse {
	s.Headers = v
	return s
}

func (s *GetLumaTableResponse) SetStatusCode(v int32) *GetLumaTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaTableResponse) SetBody(v *GetLumaTableResponseBody) *GetLumaTableResponse {
	s.Body = v
	return s
}

func (s *GetLumaTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
