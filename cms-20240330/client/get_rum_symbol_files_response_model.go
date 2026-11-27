// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumSymbolFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumSymbolFilesResponse
	GetStatusCode() *int32
	SetBody(v *GetRumSymbolFilesResponseBody) *GetRumSymbolFilesResponse
	GetBody() *GetRumSymbolFilesResponseBody
}

type GetRumSymbolFilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumSymbolFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumSymbolFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFilesResponse) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumSymbolFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumSymbolFilesResponse) GetBody() *GetRumSymbolFilesResponseBody {
	return s.Body
}

func (s *GetRumSymbolFilesResponse) SetHeaders(v map[string]*string) *GetRumSymbolFilesResponse {
	s.Headers = v
	return s
}

func (s *GetRumSymbolFilesResponse) SetStatusCode(v int32) *GetRumSymbolFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumSymbolFilesResponse) SetBody(v *GetRumSymbolFilesResponseBody) *GetRumSymbolFilesResponse {
	s.Body = v
	return s
}

func (s *GetRumSymbolFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
