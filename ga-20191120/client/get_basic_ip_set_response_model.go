// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicIpSetResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicIpSetResponseBody) *GetBasicIpSetResponse
	GetBody() *GetBasicIpSetResponseBody
}

type GetBasicIpSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicIpSetResponse) GetBody() *GetBasicIpSetResponseBody {
	return s.Body
}

func (s *GetBasicIpSetResponse) SetHeaders(v map[string]*string) *GetBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *GetBasicIpSetResponse) SetStatusCode(v int32) *GetBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicIpSetResponse) SetBody(v *GetBasicIpSetResponseBody) *GetBasicIpSetResponse {
	s.Body = v
	return s
}

func (s *GetBasicIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
