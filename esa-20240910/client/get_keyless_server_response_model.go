// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeylessServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKeylessServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKeylessServerResponse
	GetStatusCode() *int32
	SetBody(v *GetKeylessServerResponseBody) *GetKeylessServerResponse
	GetBody() *GetKeylessServerResponseBody
}

type GetKeylessServerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeylessServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeylessServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKeylessServerResponse) GoString() string {
	return s.String()
}

func (s *GetKeylessServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKeylessServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKeylessServerResponse) GetBody() *GetKeylessServerResponseBody {
	return s.Body
}

func (s *GetKeylessServerResponse) SetHeaders(v map[string]*string) *GetKeylessServerResponse {
	s.Headers = v
	return s
}

func (s *GetKeylessServerResponse) SetStatusCode(v int32) *GetKeylessServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeylessServerResponse) SetBody(v *GetKeylessServerResponseBody) *GetKeylessServerResponse {
	s.Body = v
	return s
}

func (s *GetKeylessServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
