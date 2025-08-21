// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEnsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEnsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEnsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *AttachEnsInstancesResponseBody) *AttachEnsInstancesResponse
	GetBody() *AttachEnsInstancesResponseBody
}

type AttachEnsInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEnsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEnsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEnsInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEnsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEnsInstancesResponse) GetBody() *AttachEnsInstancesResponseBody {
	return s.Body
}

func (s *AttachEnsInstancesResponse) SetHeaders(v map[string]*string) *AttachEnsInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachEnsInstancesResponse) SetStatusCode(v int32) *AttachEnsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEnsInstancesResponse) SetBody(v *AttachEnsInstancesResponseBody) *AttachEnsInstancesResponse {
	s.Body = v
	return s
}

func (s *AttachEnsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
