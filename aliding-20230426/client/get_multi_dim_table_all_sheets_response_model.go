// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiDimTableAllSheetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiDimTableAllSheetsResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiDimTableAllSheetsResponseBody) *GetMultiDimTableAllSheetsResponse
	GetBody() *GetMultiDimTableAllSheetsResponseBody
}

type GetMultiDimTableAllSheetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiDimTableAllSheetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiDimTableAllSheetsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiDimTableAllSheetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiDimTableAllSheetsResponse) GetBody() *GetMultiDimTableAllSheetsResponseBody {
	return s.Body
}

func (s *GetMultiDimTableAllSheetsResponse) SetHeaders(v map[string]*string) *GetMultiDimTableAllSheetsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiDimTableAllSheetsResponse) SetStatusCode(v int32) *GetMultiDimTableAllSheetsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponse) SetBody(v *GetMultiDimTableAllSheetsResponseBody) *GetMultiDimTableAllSheetsResponse {
	s.Body = v
	return s
}

func (s *GetMultiDimTableAllSheetsResponse) Validate() error {
	return dara.Validate(s)
}
