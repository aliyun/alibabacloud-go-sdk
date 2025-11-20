// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveFormDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveFormDataResponse
	GetStatusCode() *int32
	SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse
	GetBody() *SaveFormDataResponseBody
}

type SaveFormDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveFormDataResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveFormDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveFormDataResponse) GetBody() *SaveFormDataResponseBody {
	return s.Body
}

func (s *SaveFormDataResponse) SetHeaders(v map[string]*string) *SaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *SaveFormDataResponse) SetStatusCode(v int32) *SaveFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

func (s *SaveFormDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
