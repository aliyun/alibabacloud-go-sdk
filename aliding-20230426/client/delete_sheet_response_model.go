// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSheetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSheetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSheetResponseBody) *DeleteSheetResponse
	GetBody() *DeleteSheetResponseBody
}

type DeleteSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSheetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSheetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSheetResponse) GetBody() *DeleteSheetResponseBody {
	return s.Body
}

func (s *DeleteSheetResponse) SetHeaders(v map[string]*string) *DeleteSheetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSheetResponse) SetStatusCode(v int32) *DeleteSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSheetResponse) SetBody(v *DeleteSheetResponseBody) *DeleteSheetResponse {
	s.Body = v
	return s
}

func (s *DeleteSheetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
