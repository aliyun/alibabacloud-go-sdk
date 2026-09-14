// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpdateTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUpdateTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUpdateTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetUpdateTaskResultResponseBody) *GetUpdateTaskResultResponse
	GetBody() *GetUpdateTaskResultResponseBody
}

type GetUpdateTaskResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUpdateTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUpdateTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUpdateTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetUpdateTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUpdateTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUpdateTaskResultResponse) GetBody() *GetUpdateTaskResultResponseBody {
	return s.Body
}

func (s *GetUpdateTaskResultResponse) SetHeaders(v map[string]*string) *GetUpdateTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetUpdateTaskResultResponse) SetStatusCode(v int32) *GetUpdateTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUpdateTaskResultResponse) SetBody(v *GetUpdateTaskResultResponseBody) *GetUpdateTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetUpdateTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
