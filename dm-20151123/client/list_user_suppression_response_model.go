// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSuppressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserSuppressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserSuppressionResponse
	GetStatusCode() *int32
	SetBody(v *ListUserSuppressionResponseBody) *ListUserSuppressionResponse
	GetBody() *ListUserSuppressionResponseBody
}

type ListUserSuppressionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserSuppressionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserSuppressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserSuppressionResponse) GetBody() *ListUserSuppressionResponseBody {
	return s.Body
}

func (s *ListUserSuppressionResponse) SetHeaders(v map[string]*string) *ListUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *ListUserSuppressionResponse) SetStatusCode(v int32) *ListUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSuppressionResponse) SetBody(v *ListUserSuppressionResponseBody) *ListUserSuppressionResponse {
	s.Body = v
	return s
}

func (s *ListUserSuppressionResponse) Validate() error {
	return dara.Validate(s)
}
