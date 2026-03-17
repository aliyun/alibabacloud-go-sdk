// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagCipherResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearSagCipherResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearSagCipherResponse
	GetStatusCode() *int32
	SetBody(v *ClearSagCipherResponseBody) *ClearSagCipherResponse
	GetBody() *ClearSagCipherResponseBody
}

type ClearSagCipherResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearSagCipherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearSagCipherResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearSagCipherResponse) GoString() string {
	return s.String()
}

func (s *ClearSagCipherResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearSagCipherResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearSagCipherResponse) GetBody() *ClearSagCipherResponseBody {
	return s.Body
}

func (s *ClearSagCipherResponse) SetHeaders(v map[string]*string) *ClearSagCipherResponse {
	s.Headers = v
	return s
}

func (s *ClearSagCipherResponse) SetStatusCode(v int32) *ClearSagCipherResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearSagCipherResponse) SetBody(v *ClearSagCipherResponseBody) *ClearSagCipherResponse {
	s.Body = v
	return s
}

func (s *ClearSagCipherResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
