// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMailTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMailTaskListResponse
	GetStatusCode() *int32
	SetBody(v *GetMailTaskListResponseBody) *GetMailTaskListResponse
	GetBody() *GetMailTaskListResponseBody
}

type GetMailTaskListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMailTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMailTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetMailTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMailTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMailTaskListResponse) GetBody() *GetMailTaskListResponseBody {
	return s.Body
}

func (s *GetMailTaskListResponse) SetHeaders(v map[string]*string) *GetMailTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetMailTaskListResponse) SetStatusCode(v int32) *GetMailTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMailTaskListResponse) SetBody(v *GetMailTaskListResponseBody) *GetMailTaskListResponse {
	s.Body = v
	return s
}

func (s *GetMailTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
