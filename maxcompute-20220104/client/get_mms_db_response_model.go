// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsDbResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsDbResponseBody) *GetMmsDbResponse
	GetBody() *GetMmsDbResponseBody
}

type GetMmsDbResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsDbResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDbResponse) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsDbResponse) GetBody() *GetMmsDbResponseBody {
	return s.Body
}

func (s *GetMmsDbResponse) SetHeaders(v map[string]*string) *GetMmsDbResponse {
	s.Headers = v
	return s
}

func (s *GetMmsDbResponse) SetStatusCode(v int32) *GetMmsDbResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsDbResponse) SetBody(v *GetMmsDbResponseBody) *GetMmsDbResponse {
	s.Body = v
	return s
}

func (s *GetMmsDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
