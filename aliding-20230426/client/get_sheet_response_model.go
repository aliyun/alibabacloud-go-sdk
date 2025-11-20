// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSheetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSheetResponse
	GetStatusCode() *int32
	SetBody(v *GetSheetResponseBody) *GetSheetResponse
	GetBody() *GetSheetResponseBody
}

type GetSheetResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSheetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSheetResponse) GoString() string {
	return s.String()
}

func (s *GetSheetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSheetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSheetResponse) GetBody() *GetSheetResponseBody {
	return s.Body
}

func (s *GetSheetResponse) SetHeaders(v map[string]*string) *GetSheetResponse {
	s.Headers = v
	return s
}

func (s *GetSheetResponse) SetStatusCode(v int32) *GetSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSheetResponse) SetBody(v *GetSheetResponseBody) *GetSheetResponse {
	s.Body = v
	return s
}

func (s *GetSheetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
