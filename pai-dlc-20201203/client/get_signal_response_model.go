// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSignalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSignalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSignalResponse
	GetStatusCode() *int32
	SetBody(v *GetSignalResponseBody) *GetSignalResponse
	GetBody() *GetSignalResponseBody
}

type GetSignalResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSignalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSignalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSignalResponse) GoString() string {
	return s.String()
}

func (s *GetSignalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSignalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSignalResponse) GetBody() *GetSignalResponseBody {
	return s.Body
}

func (s *GetSignalResponse) SetHeaders(v map[string]*string) *GetSignalResponse {
	s.Headers = v
	return s
}

func (s *GetSignalResponse) SetStatusCode(v int32) *GetSignalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignalResponse) SetBody(v *GetSignalResponseBody) *GetSignalResponse {
	s.Body = v
	return s
}

func (s *GetSignalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
