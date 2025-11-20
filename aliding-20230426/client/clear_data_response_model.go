// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearDataResponse
	GetStatusCode() *int32
	SetBody(v *ClearDataResponseBody) *ClearDataResponse
	GetBody() *ClearDataResponseBody
}

type ClearDataResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearDataResponse) GoString() string {
	return s.String()
}

func (s *ClearDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearDataResponse) GetBody() *ClearDataResponseBody {
	return s.Body
}

func (s *ClearDataResponse) SetHeaders(v map[string]*string) *ClearDataResponse {
	s.Headers = v
	return s
}

func (s *ClearDataResponse) SetStatusCode(v int32) *ClearDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearDataResponse) SetBody(v *ClearDataResponseBody) *ClearDataResponse {
	s.Body = v
	return s
}

func (s *ClearDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
