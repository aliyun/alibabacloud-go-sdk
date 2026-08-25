// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetServerIdeInstanceResponseBody) *GetServerIdeInstanceResponse
	GetBody() *GetServerIdeInstanceResponseBody
}

type GetServerIdeInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServerIdeInstanceResponse) GetBody() *GetServerIdeInstanceResponseBody {
	return s.Body
}

func (s *GetServerIdeInstanceResponse) SetHeaders(v map[string]*string) *GetServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServerIdeInstanceResponse) SetStatusCode(v int32) *GetServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServerIdeInstanceResponse) SetBody(v *GetServerIdeInstanceResponseBody) *GetServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *GetServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
