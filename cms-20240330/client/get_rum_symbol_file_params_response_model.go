// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFileParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumSymbolFileParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumSymbolFileParamsResponse
	GetStatusCode() *int32
	SetBody(v *GetRumSymbolFileParamsResponseBody) *GetRumSymbolFileParamsResponse
	GetBody() *GetRumSymbolFileParamsResponseBody
}

type GetRumSymbolFileParamsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumSymbolFileParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumSymbolFileParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFileParamsResponse) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFileParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumSymbolFileParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumSymbolFileParamsResponse) GetBody() *GetRumSymbolFileParamsResponseBody {
	return s.Body
}

func (s *GetRumSymbolFileParamsResponse) SetHeaders(v map[string]*string) *GetRumSymbolFileParamsResponse {
	s.Headers = v
	return s
}

func (s *GetRumSymbolFileParamsResponse) SetStatusCode(v int32) *GetRumSymbolFileParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumSymbolFileParamsResponse) SetBody(v *GetRumSymbolFileParamsResponseBody) *GetRumSymbolFileParamsResponse {
	s.Body = v
	return s
}

func (s *GetRumSymbolFileParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
