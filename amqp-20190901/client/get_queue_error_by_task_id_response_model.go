// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueErrorByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueErrorByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueErrorByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueErrorByTaskIdResponseBody) *GetQueueErrorByTaskIdResponse
	GetBody() *GetQueueErrorByTaskIdResponseBody
}

type GetQueueErrorByTaskIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueErrorByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueErrorByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueErrorByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueErrorByTaskIdResponse) GetBody() *GetQueueErrorByTaskIdResponseBody {
	return s.Body
}

func (s *GetQueueErrorByTaskIdResponse) SetHeaders(v map[string]*string) *GetQueueErrorByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetQueueErrorByTaskIdResponse) SetStatusCode(v int32) *GetQueueErrorByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueErrorByTaskIdResponse) SetBody(v *GetQueueErrorByTaskIdResponseBody) *GetQueueErrorByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetQueueErrorByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
