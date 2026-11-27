// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumSymbolFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRumSymbolFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRumSymbolFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRumSymbolFilesResponseBody) *DeleteRumSymbolFilesResponse
	GetBody() *DeleteRumSymbolFilesResponseBody
}

type DeleteRumSymbolFilesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRumSymbolFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRumSymbolFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumSymbolFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRumSymbolFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRumSymbolFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRumSymbolFilesResponse) GetBody() *DeleteRumSymbolFilesResponseBody {
	return s.Body
}

func (s *DeleteRumSymbolFilesResponse) SetHeaders(v map[string]*string) *DeleteRumSymbolFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRumSymbolFilesResponse) SetStatusCode(v int32) *DeleteRumSymbolFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRumSymbolFilesResponse) SetBody(v *DeleteRumSymbolFilesResponseBody) *DeleteRumSymbolFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteRumSymbolFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
