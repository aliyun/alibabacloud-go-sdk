// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueRateResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueRateResponseBody) *GetQueueRateResponse
	GetBody() *GetQueueRateResponseBody
}

type GetQueueRateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueRateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRateResponse) GoString() string {
	return s.String()
}

func (s *GetQueueRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueRateResponse) GetBody() *GetQueueRateResponseBody {
	return s.Body
}

func (s *GetQueueRateResponse) SetHeaders(v map[string]*string) *GetQueueRateResponse {
	s.Headers = v
	return s
}

func (s *GetQueueRateResponse) SetStatusCode(v int32) *GetQueueRateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueRateResponse) SetBody(v *GetQueueRateResponseBody) *GetQueueRateResponse {
	s.Body = v
	return s
}

func (s *GetQueueRateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
