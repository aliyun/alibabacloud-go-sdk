// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableFieldsResponse
	GetStatusCode() *int32
	SetBody(v *GetTableFieldsResponseBody) *GetTableFieldsResponse
	GetBody() *GetTableFieldsResponseBody
}

type GetTableFieldsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableFieldsResponse) GoString() string {
	return s.String()
}

func (s *GetTableFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableFieldsResponse) GetBody() *GetTableFieldsResponseBody {
	return s.Body
}

func (s *GetTableFieldsResponse) SetHeaders(v map[string]*string) *GetTableFieldsResponse {
	s.Headers = v
	return s
}

func (s *GetTableFieldsResponse) SetStatusCode(v int32) *GetTableFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableFieldsResponse) SetBody(v *GetTableFieldsResponseBody) *GetTableFieldsResponse {
	s.Body = v
	return s
}

func (s *GetTableFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
