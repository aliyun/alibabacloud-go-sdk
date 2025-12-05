// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesByKeyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliasesByKeyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliasesByKeyIdResponse
	GetStatusCode() *int32
	SetBody(v *ListAliasesByKeyIdResponseBody) *ListAliasesByKeyIdResponse
	GetBody() *ListAliasesByKeyIdResponseBody
}

type ListAliasesByKeyIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliasesByKeyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliasesByKeyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesByKeyIdResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliasesByKeyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliasesByKeyIdResponse) GetBody() *ListAliasesByKeyIdResponseBody {
	return s.Body
}

func (s *ListAliasesByKeyIdResponse) SetHeaders(v map[string]*string) *ListAliasesByKeyIdResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetStatusCode(v int32) *ListAliasesByKeyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetBody(v *ListAliasesByKeyIdResponseBody) *ListAliasesByKeyIdResponse {
	s.Body = v
	return s
}

func (s *ListAliasesByKeyIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
