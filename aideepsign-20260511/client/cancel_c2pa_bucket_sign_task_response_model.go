// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelC2paBucketSignTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelC2paBucketSignTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelC2paBucketSignTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelC2paBucketSignTaskResponseBody) *CancelC2paBucketSignTaskResponse
	GetBody() *CancelC2paBucketSignTaskResponseBody
}

type CancelC2paBucketSignTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelC2paBucketSignTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelC2paBucketSignTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelC2paBucketSignTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelC2paBucketSignTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelC2paBucketSignTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelC2paBucketSignTaskResponse) GetBody() *CancelC2paBucketSignTaskResponseBody {
	return s.Body
}

func (s *CancelC2paBucketSignTaskResponse) SetHeaders(v map[string]*string) *CancelC2paBucketSignTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelC2paBucketSignTaskResponse) SetStatusCode(v int32) *CancelC2paBucketSignTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelC2paBucketSignTaskResponse) SetBody(v *CancelC2paBucketSignTaskResponseBody) *CancelC2paBucketSignTaskResponse {
	s.Body = v
	return s
}

func (s *CancelC2paBucketSignTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
