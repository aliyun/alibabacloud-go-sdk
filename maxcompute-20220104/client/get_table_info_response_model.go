// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTableInfoResponseBody) *GetTableInfoResponse
	GetBody() *GetTableInfoResponseBody
}

type GetTableInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableInfoResponse) GetBody() *GetTableInfoResponseBody {
	return s.Body
}

func (s *GetTableInfoResponse) SetHeaders(v map[string]*string) *GetTableInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTableInfoResponse) SetStatusCode(v int32) *GetTableInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableInfoResponse) SetBody(v *GetTableInfoResponseBody) *GetTableInfoResponse {
	s.Body = v
	return s
}

func (s *GetTableInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
