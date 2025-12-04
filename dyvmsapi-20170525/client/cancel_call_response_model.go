// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCallResponse
	GetStatusCode() *int32
	SetBody(v *CancelCallResponseBody) *CancelCallResponse
	GetBody() *CancelCallResponseBody
}

type CancelCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCallResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCallResponse) GoString() string {
	return s.String()
}

func (s *CancelCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCallResponse) GetBody() *CancelCallResponseBody {
	return s.Body
}

func (s *CancelCallResponse) SetHeaders(v map[string]*string) *CancelCallResponse {
	s.Headers = v
	return s
}

func (s *CancelCallResponse) SetStatusCode(v int32) *CancelCallResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCallResponse) SetBody(v *CancelCallResponseBody) *CancelCallResponse {
	s.Body = v
	return s
}

func (s *CancelCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
