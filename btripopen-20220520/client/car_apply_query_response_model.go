// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *CarApplyQueryResponseBody) *CarApplyQueryResponse
	GetBody() *CarApplyQueryResponseBody
}

type CarApplyQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarApplyQueryResponse) GetBody() *CarApplyQueryResponseBody {
	return s.Body
}

func (s *CarApplyQueryResponse) SetHeaders(v map[string]*string) *CarApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *CarApplyQueryResponse) SetStatusCode(v int32) *CarApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarApplyQueryResponse) SetBody(v *CarApplyQueryResponseBody) *CarApplyQueryResponse {
	s.Body = v
	return s
}

func (s *CarApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
