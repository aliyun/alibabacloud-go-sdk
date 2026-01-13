// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetTableMetaResponseBody) *GetTableMetaResponse
	GetBody() *GetTableMetaResponseBody
}

type GetTableMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableMetaResponse) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableMetaResponse) GetBody() *GetTableMetaResponseBody {
	return s.Body
}

func (s *GetTableMetaResponse) SetHeaders(v map[string]*string) *GetTableMetaResponse {
	s.Headers = v
	return s
}

func (s *GetTableMetaResponse) SetStatusCode(v int32) *GetTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableMetaResponse) SetBody(v *GetTableMetaResponseBody) *GetTableMetaResponse {
	s.Body = v
	return s
}

func (s *GetTableMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
