// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlogTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlogTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTlogTaskInfoResponseBody) *GetTlogTaskInfoResponse
	GetBody() *GetTlogTaskInfoResponseBody
}

type GetTlogTaskInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlogTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlogTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTlogTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlogTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlogTaskInfoResponse) GetBody() *GetTlogTaskInfoResponseBody {
	return s.Body
}

func (s *GetTlogTaskInfoResponse) SetHeaders(v map[string]*string) *GetTlogTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTlogTaskInfoResponse) SetStatusCode(v int32) *GetTlogTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlogTaskInfoResponse) SetBody(v *GetTlogTaskInfoResponseBody) *GetTlogTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetTlogTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
