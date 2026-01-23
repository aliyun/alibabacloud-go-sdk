// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableColumnLineagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableColumnLineagesResponse
	GetStatusCode() *int32
	SetBody(v *GetTableColumnLineagesResponseBody) *GetTableColumnLineagesResponse
	GetBody() *GetTableColumnLineagesResponseBody
}

type GetTableColumnLineagesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableColumnLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableColumnLineagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableColumnLineagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableColumnLineagesResponse) GetBody() *GetTableColumnLineagesResponseBody {
	return s.Body
}

func (s *GetTableColumnLineagesResponse) SetHeaders(v map[string]*string) *GetTableColumnLineagesResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnLineagesResponse) SetStatusCode(v int32) *GetTableColumnLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnLineagesResponse) SetBody(v *GetTableColumnLineagesResponseBody) *GetTableColumnLineagesResponse {
	s.Body = v
	return s
}

func (s *GetTableColumnLineagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
