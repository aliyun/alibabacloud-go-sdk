// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllSheetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllSheetsResponse
	GetStatusCode() *int32
	SetBody(v *GetAllSheetsResponseBody) *GetAllSheetsResponse
	GetBody() *GetAllSheetsResponseBody
}

type GetAllSheetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllSheetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllSheetsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsResponse) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllSheetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllSheetsResponse) GetBody() *GetAllSheetsResponseBody {
	return s.Body
}

func (s *GetAllSheetsResponse) SetHeaders(v map[string]*string) *GetAllSheetsResponse {
	s.Headers = v
	return s
}

func (s *GetAllSheetsResponse) SetStatusCode(v int32) *GetAllSheetsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllSheetsResponse) SetBody(v *GetAllSheetsResponseBody) *GetAllSheetsResponse {
	s.Body = v
	return s
}

func (s *GetAllSheetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
