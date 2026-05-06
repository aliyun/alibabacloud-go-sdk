// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSubTaskResultResponseBody) *GetSubTaskResultResponse
	GetBody() *GetSubTaskResultResponseBody
}

type GetSubTaskResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubTaskResultResponse) GetBody() *GetSubTaskResultResponseBody {
	return s.Body
}

func (s *GetSubTaskResultResponse) SetHeaders(v map[string]*string) *GetSubTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetSubTaskResultResponse) SetStatusCode(v int32) *GetSubTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubTaskResultResponse) SetBody(v *GetSubTaskResultResponseBody) *GetSubTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetSubTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
