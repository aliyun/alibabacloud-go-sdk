// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupExistResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupExistResponseBody) *GetGroupExistResponse
	GetBody() *GetGroupExistResponseBody
}

type GetGroupExistResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupExistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupExistResponse) GoString() string {
	return s.String()
}

func (s *GetGroupExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupExistResponse) GetBody() *GetGroupExistResponseBody {
	return s.Body
}

func (s *GetGroupExistResponse) SetHeaders(v map[string]*string) *GetGroupExistResponse {
	s.Headers = v
	return s
}

func (s *GetGroupExistResponse) SetStatusCode(v int32) *GetGroupExistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupExistResponse) SetBody(v *GetGroupExistResponseBody) *GetGroupExistResponse {
	s.Body = v
	return s
}

func (s *GetGroupExistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
