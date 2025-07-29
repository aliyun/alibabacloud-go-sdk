// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixNodePoolVulsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FixNodePoolVulsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FixNodePoolVulsResponse
	GetStatusCode() *int32
	SetBody(v *FixNodePoolVulsResponseBody) *FixNodePoolVulsResponse
	GetBody() *FixNodePoolVulsResponseBody
}

type FixNodePoolVulsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FixNodePoolVulsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FixNodePoolVulsResponse) String() string {
	return dara.Prettify(s)
}

func (s FixNodePoolVulsResponse) GoString() string {
	return s.String()
}

func (s *FixNodePoolVulsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FixNodePoolVulsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FixNodePoolVulsResponse) GetBody() *FixNodePoolVulsResponseBody {
	return s.Body
}

func (s *FixNodePoolVulsResponse) SetHeaders(v map[string]*string) *FixNodePoolVulsResponse {
	s.Headers = v
	return s
}

func (s *FixNodePoolVulsResponse) SetStatusCode(v int32) *FixNodePoolVulsResponse {
	s.StatusCode = &v
	return s
}

func (s *FixNodePoolVulsResponse) SetBody(v *FixNodePoolVulsResponseBody) *FixNodePoolVulsResponse {
	s.Body = v
	return s
}

func (s *FixNodePoolVulsResponse) Validate() error {
	return dara.Validate(s)
}
