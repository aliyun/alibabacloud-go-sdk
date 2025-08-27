// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommonApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommonApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *CommonApplyQueryResponseBody) *CommonApplyQueryResponse
	GetBody() *CommonApplyQueryResponseBody
}

type CommonApplyQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommonApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CommonApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommonApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommonApplyQueryResponse) GetBody() *CommonApplyQueryResponseBody {
	return s.Body
}

func (s *CommonApplyQueryResponse) SetHeaders(v map[string]*string) *CommonApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *CommonApplyQueryResponse) SetStatusCode(v int32) *CommonApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonApplyQueryResponse) SetBody(v *CommonApplyQueryResponseBody) *CommonApplyQueryResponse {
	s.Body = v
	return s
}

func (s *CommonApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
