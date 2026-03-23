// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendifyAutoLoginURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSendifyAutoLoginURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSendifyAutoLoginURLResponse
	GetStatusCode() *int32
	SetBody(v *GetSendifyAutoLoginURLResponseBody) *GetSendifyAutoLoginURLResponse
	GetBody() *GetSendifyAutoLoginURLResponseBody
}

type GetSendifyAutoLoginURLResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendifyAutoLoginURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendifyAutoLoginURLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSendifyAutoLoginURLResponse) GoString() string {
	return s.String()
}

func (s *GetSendifyAutoLoginURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSendifyAutoLoginURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSendifyAutoLoginURLResponse) GetBody() *GetSendifyAutoLoginURLResponseBody {
	return s.Body
}

func (s *GetSendifyAutoLoginURLResponse) SetHeaders(v map[string]*string) *GetSendifyAutoLoginURLResponse {
	s.Headers = v
	return s
}

func (s *GetSendifyAutoLoginURLResponse) SetStatusCode(v int32) *GetSendifyAutoLoginURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendifyAutoLoginURLResponse) SetBody(v *GetSendifyAutoLoginURLResponseBody) *GetSendifyAutoLoginURLResponse {
	s.Body = v
	return s
}

func (s *GetSendifyAutoLoginURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
