// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSyncAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeSyncAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeSyncAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *InvokeSyncAlgorithmResponseBody) *InvokeSyncAlgorithmResponse
	GetBody() *InvokeSyncAlgorithmResponseBody
}

type InvokeSyncAlgorithmResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeSyncAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeSyncAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeSyncAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *InvokeSyncAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeSyncAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeSyncAlgorithmResponse) GetBody() *InvokeSyncAlgorithmResponseBody {
	return s.Body
}

func (s *InvokeSyncAlgorithmResponse) SetHeaders(v map[string]*string) *InvokeSyncAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *InvokeSyncAlgorithmResponse) SetStatusCode(v int32) *InvokeSyncAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeSyncAlgorithmResponse) SetBody(v *InvokeSyncAlgorithmResponseBody) *InvokeSyncAlgorithmResponse {
	s.Body = v
	return s
}

func (s *InvokeSyncAlgorithmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
