// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKeyDeletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelKeyDeletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelKeyDeletionResponse
	GetStatusCode() *int32
	SetBody(v *CancelKeyDeletionResponseBody) *CancelKeyDeletionResponse
	GetBody() *CancelKeyDeletionResponseBody
}

type CancelKeyDeletionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelKeyDeletionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelKeyDeletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelKeyDeletionResponse) GetBody() *CancelKeyDeletionResponseBody {
	return s.Body
}

func (s *CancelKeyDeletionResponse) SetHeaders(v map[string]*string) *CancelKeyDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelKeyDeletionResponse) SetStatusCode(v int32) *CancelKeyDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelKeyDeletionResponse) SetBody(v *CancelKeyDeletionResponseBody) *CancelKeyDeletionResponse {
	s.Body = v
	return s
}

func (s *CancelKeyDeletionResponse) Validate() error {
	return dara.Validate(s)
}
