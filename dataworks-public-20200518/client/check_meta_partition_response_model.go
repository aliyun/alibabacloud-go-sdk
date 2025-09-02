// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaPartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMetaPartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMetaPartitionResponse
	GetStatusCode() *int32
	SetBody(v *CheckMetaPartitionResponseBody) *CheckMetaPartitionResponse
	GetBody() *CheckMetaPartitionResponseBody
}

type CheckMetaPartitionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMetaPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMetaPartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaPartitionResponse) GoString() string {
	return s.String()
}

func (s *CheckMetaPartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMetaPartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMetaPartitionResponse) GetBody() *CheckMetaPartitionResponseBody {
	return s.Body
}

func (s *CheckMetaPartitionResponse) SetHeaders(v map[string]*string) *CheckMetaPartitionResponse {
	s.Headers = v
	return s
}

func (s *CheckMetaPartitionResponse) SetStatusCode(v int32) *CheckMetaPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMetaPartitionResponse) SetBody(v *CheckMetaPartitionResponseBody) *CheckMetaPartitionResponse {
	s.Body = v
	return s
}

func (s *CheckMetaPartitionResponse) Validate() error {
	return dara.Validate(s)
}
