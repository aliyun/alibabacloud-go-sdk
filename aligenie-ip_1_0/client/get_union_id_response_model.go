// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnionIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUnionIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUnionIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUnionIdResponseBody) *GetUnionIdResponse
	GetBody() *GetUnionIdResponseBody
}

type GetUnionIdResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnionIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUnionIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnionIdResponse) GetBody() *GetUnionIdResponseBody {
	return s.Body
}

func (s *GetUnionIdResponse) SetHeaders(v map[string]*string) *GetUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUnionIdResponse) SetStatusCode(v int32) *GetUnionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnionIdResponse) SetBody(v *GetUnionIdResponseBody) *GetUnionIdResponse {
	s.Body = v
	return s
}

func (s *GetUnionIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
