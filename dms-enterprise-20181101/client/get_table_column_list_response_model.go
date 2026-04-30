// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableColumnListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableColumnListResponse
	GetStatusCode() *int32
	SetBody(v *GetTableColumnListResponseBody) *GetTableColumnListResponse
	GetBody() *GetTableColumnListResponseBody
}

type GetTableColumnListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableColumnListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableColumnListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnListResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableColumnListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableColumnListResponse) GetBody() *GetTableColumnListResponseBody {
	return s.Body
}

func (s *GetTableColumnListResponse) SetHeaders(v map[string]*string) *GetTableColumnListResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnListResponse) SetStatusCode(v int32) *GetTableColumnListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnListResponse) SetBody(v *GetTableColumnListResponseBody) *GetTableColumnListResponse {
	s.Body = v
	return s
}

func (s *GetTableColumnListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
