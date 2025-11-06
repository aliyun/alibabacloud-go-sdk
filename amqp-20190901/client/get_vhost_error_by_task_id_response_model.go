// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostErrorByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVhostErrorByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVhostErrorByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetVhostErrorByTaskIdResponseBody) *GetVhostErrorByTaskIdResponse
	GetBody() *GetVhostErrorByTaskIdResponseBody
}

type GetVhostErrorByTaskIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVhostErrorByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVhostErrorByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVhostErrorByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVhostErrorByTaskIdResponse) GetBody() *GetVhostErrorByTaskIdResponseBody {
	return s.Body
}

func (s *GetVhostErrorByTaskIdResponse) SetHeaders(v map[string]*string) *GetVhostErrorByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetVhostErrorByTaskIdResponse) SetStatusCode(v int32) *GetVhostErrorByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponse) SetBody(v *GetVhostErrorByTaskIdResponseBody) *GetVhostErrorByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetVhostErrorByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
