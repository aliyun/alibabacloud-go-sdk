// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryTransferOwnershipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryTransferOwnershipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryTransferOwnershipResponse
	GetStatusCode() *int32
	SetBody(v *RetryTransferOwnershipResponseBody) *RetryTransferOwnershipResponse
	GetBody() *RetryTransferOwnershipResponseBody
}

type RetryTransferOwnershipResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryTransferOwnershipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryTransferOwnershipResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryTransferOwnershipResponse) GoString() string {
	return s.String()
}

func (s *RetryTransferOwnershipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryTransferOwnershipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryTransferOwnershipResponse) GetBody() *RetryTransferOwnershipResponseBody {
	return s.Body
}

func (s *RetryTransferOwnershipResponse) SetHeaders(v map[string]*string) *RetryTransferOwnershipResponse {
	s.Headers = v
	return s
}

func (s *RetryTransferOwnershipResponse) SetStatusCode(v int32) *RetryTransferOwnershipResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryTransferOwnershipResponse) SetBody(v *RetryTransferOwnershipResponseBody) *RetryTransferOwnershipResponse {
	s.Body = v
	return s
}

func (s *RetryTransferOwnershipResponse) Validate() error {
	return dara.Validate(s)
}
