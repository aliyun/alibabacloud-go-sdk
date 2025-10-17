// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDDLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableDDLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableDDLResponse
	GetStatusCode() *int32
	SetBody(v *GetTableDDLResponseBody) *GetTableDDLResponse
	GetBody() *GetTableDDLResponseBody
}

type GetTableDDLResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableDDLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableDDLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableDDLResponse) GoString() string {
	return s.String()
}

func (s *GetTableDDLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableDDLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableDDLResponse) GetBody() *GetTableDDLResponseBody {
	return s.Body
}

func (s *GetTableDDLResponse) SetHeaders(v map[string]*string) *GetTableDDLResponse {
	s.Headers = v
	return s
}

func (s *GetTableDDLResponse) SetStatusCode(v int32) *GetTableDDLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDDLResponse) SetBody(v *GetTableDDLResponseBody) *GetTableDDLResponse {
	s.Body = v
	return s
}

func (s *GetTableDDLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
