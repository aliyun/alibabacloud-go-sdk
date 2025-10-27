// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastOnceTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLastOnceTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLastOnceTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLastOnceTaskInfoResponseBody) *GetLastOnceTaskInfoResponse
	GetBody() *GetLastOnceTaskInfoResponseBody
}

type GetLastOnceTaskInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLastOnceTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLastOnceTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLastOnceTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLastOnceTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLastOnceTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLastOnceTaskInfoResponse) GetBody() *GetLastOnceTaskInfoResponseBody {
	return s.Body
}

func (s *GetLastOnceTaskInfoResponse) SetHeaders(v map[string]*string) *GetLastOnceTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLastOnceTaskInfoResponse) SetStatusCode(v int32) *GetLastOnceTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLastOnceTaskInfoResponse) SetBody(v *GetLastOnceTaskInfoResponseBody) *GetLastOnceTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetLastOnceTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
