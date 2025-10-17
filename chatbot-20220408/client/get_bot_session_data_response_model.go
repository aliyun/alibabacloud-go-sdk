// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBotSessionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBotSessionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBotSessionDataResponse
	GetStatusCode() *int32
	SetBody(v *GetBotSessionDataResponseBody) *GetBotSessionDataResponse
	GetBody() *GetBotSessionDataResponseBody
}

type GetBotSessionDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBotSessionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBotSessionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBotSessionDataResponse) GoString() string {
	return s.String()
}

func (s *GetBotSessionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBotSessionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBotSessionDataResponse) GetBody() *GetBotSessionDataResponseBody {
	return s.Body
}

func (s *GetBotSessionDataResponse) SetHeaders(v map[string]*string) *GetBotSessionDataResponse {
	s.Headers = v
	return s
}

func (s *GetBotSessionDataResponse) SetStatusCode(v int32) *GetBotSessionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBotSessionDataResponse) SetBody(v *GetBotSessionDataResponseBody) *GetBotSessionDataResponse {
	s.Body = v
	return s
}

func (s *GetBotSessionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
