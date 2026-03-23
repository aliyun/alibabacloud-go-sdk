// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendifyInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSendifyInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSendifyInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSendifyInfoResponseBody) *GetSendifyInfoResponse
	GetBody() *GetSendifyInfoResponseBody
}

type GetSendifyInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendifyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendifyInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSendifyInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSendifyInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSendifyInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSendifyInfoResponse) GetBody() *GetSendifyInfoResponseBody {
	return s.Body
}

func (s *GetSendifyInfoResponse) SetHeaders(v map[string]*string) *GetSendifyInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSendifyInfoResponse) SetStatusCode(v int32) *GetSendifyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendifyInfoResponse) SetBody(v *GetSendifyInfoResponseBody) *GetSendifyInfoResponse {
	s.Body = v
	return s
}

func (s *GetSendifyInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
