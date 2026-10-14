// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateC2paBucketSignTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateC2paBucketSignTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateC2paBucketSignTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateC2paBucketSignTaskResponseBody) *CreateC2paBucketSignTaskResponse
	GetBody() *CreateC2paBucketSignTaskResponseBody
}

type CreateC2paBucketSignTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateC2paBucketSignTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateC2paBucketSignTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateC2paBucketSignTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateC2paBucketSignTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateC2paBucketSignTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateC2paBucketSignTaskResponse) GetBody() *CreateC2paBucketSignTaskResponseBody {
	return s.Body
}

func (s *CreateC2paBucketSignTaskResponse) SetHeaders(v map[string]*string) *CreateC2paBucketSignTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateC2paBucketSignTaskResponse) SetStatusCode(v int32) *CreateC2paBucketSignTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateC2paBucketSignTaskResponse) SetBody(v *CreateC2paBucketSignTaskResponseBody) *CreateC2paBucketSignTaskResponse {
	s.Body = v
	return s
}

func (s *CreateC2paBucketSignTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
