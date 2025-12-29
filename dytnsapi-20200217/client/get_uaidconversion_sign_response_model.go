// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDConversionSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUAIDConversionSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUAIDConversionSignResponse
	GetStatusCode() *int32
	SetBody(v *GetUAIDConversionSignResponseBody) *GetUAIDConversionSignResponse
	GetBody() *GetUAIDConversionSignResponseBody
}

type GetUAIDConversionSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUAIDConversionSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUAIDConversionSignResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDConversionSignResponse) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUAIDConversionSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUAIDConversionSignResponse) GetBody() *GetUAIDConversionSignResponseBody {
	return s.Body
}

func (s *GetUAIDConversionSignResponse) SetHeaders(v map[string]*string) *GetUAIDConversionSignResponse {
	s.Headers = v
	return s
}

func (s *GetUAIDConversionSignResponse) SetStatusCode(v int32) *GetUAIDConversionSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUAIDConversionSignResponse) SetBody(v *GetUAIDConversionSignResponseBody) *GetUAIDConversionSignResponse {
	s.Body = v
	return s
}

func (s *GetUAIDConversionSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
