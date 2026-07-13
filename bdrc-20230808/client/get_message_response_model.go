// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageResponseBody) *GetMessageResponse
	GetBody() *GetMessageResponseBody
}

type GetMessageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageResponse) GoString() string {
	return s.String()
}

func (s *GetMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageResponse) GetBody() *GetMessageResponseBody {
	return s.Body
}

func (s *GetMessageResponse) SetHeaders(v map[string]*string) *GetMessageResponse {
	s.Headers = v
	return s
}

func (s *GetMessageResponse) SetStatusCode(v int32) *GetMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageResponse) SetBody(v *GetMessageResponseBody) *GetMessageResponse {
	s.Body = v
	return s
}

func (s *GetMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
