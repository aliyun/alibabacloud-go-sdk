// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbMessengerPagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFbMessengerPagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFbMessengerPagesResponse
	GetStatusCode() *int32
	SetBody(v *GetFbMessengerPagesResponseBody) *GetFbMessengerPagesResponse
	GetBody() *GetFbMessengerPagesResponseBody
}

type GetFbMessengerPagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFbMessengerPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFbMessengerPagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFbMessengerPagesResponse) GoString() string {
	return s.String()
}

func (s *GetFbMessengerPagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFbMessengerPagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFbMessengerPagesResponse) GetBody() *GetFbMessengerPagesResponseBody {
	return s.Body
}

func (s *GetFbMessengerPagesResponse) SetHeaders(v map[string]*string) *GetFbMessengerPagesResponse {
	s.Headers = v
	return s
}

func (s *GetFbMessengerPagesResponse) SetStatusCode(v int32) *GetFbMessengerPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFbMessengerPagesResponse) SetBody(v *GetFbMessengerPagesResponseBody) *GetFbMessengerPagesResponse {
	s.Body = v
	return s
}

func (s *GetFbMessengerPagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
