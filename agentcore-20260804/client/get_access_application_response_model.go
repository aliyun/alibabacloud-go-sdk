// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessApplicationResponseBody) *GetAccessApplicationResponse
	GetBody() *GetAccessApplicationResponseBody
}

type GetAccessApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessApplicationResponse) GetBody() *GetAccessApplicationResponseBody {
	return s.Body
}

func (s *GetAccessApplicationResponse) SetHeaders(v map[string]*string) *GetAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetAccessApplicationResponse) SetStatusCode(v int32) *GetAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessApplicationResponse) SetBody(v *GetAccessApplicationResponseBody) *GetAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *GetAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
