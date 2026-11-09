// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQuerySessionByClientIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQuerySessionByClientIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQuerySessionByClientIdsResponse
	GetStatusCode() *int32
	SetBody(v *BatchQuerySessionByClientIdsResponseBody) *BatchQuerySessionByClientIdsResponse
	GetBody() *BatchQuerySessionByClientIdsResponseBody
}

type BatchQuerySessionByClientIdsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQuerySessionByClientIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponse) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQuerySessionByClientIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQuerySessionByClientIdsResponse) GetBody() *BatchQuerySessionByClientIdsResponseBody {
	return s.Body
}

func (s *BatchQuerySessionByClientIdsResponse) SetHeaders(v map[string]*string) *BatchQuerySessionByClientIdsResponse {
	s.Headers = v
	return s
}

func (s *BatchQuerySessionByClientIdsResponse) SetStatusCode(v int32) *BatchQuerySessionByClientIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponse) SetBody(v *BatchQuerySessionByClientIdsResponseBody) *BatchQuerySessionByClientIdsResponse {
	s.Body = v
	return s
}

func (s *BatchQuerySessionByClientIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
