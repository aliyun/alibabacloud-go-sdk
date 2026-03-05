// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUnionTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUnionTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryUnionTaskInfoResponseBody) *QueryUnionTaskInfoResponse
	GetBody() *QueryUnionTaskInfoResponseBody
}

type QueryUnionTaskInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnionTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnionTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUnionTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUnionTaskInfoResponse) GetBody() *QueryUnionTaskInfoResponseBody {
	return s.Body
}

func (s *QueryUnionTaskInfoResponse) SetHeaders(v map[string]*string) *QueryUnionTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUnionTaskInfoResponse) SetStatusCode(v int32) *QueryUnionTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnionTaskInfoResponse) SetBody(v *QueryUnionTaskInfoResponseBody) *QueryUnionTaskInfoResponse {
	s.Body = v
	return s
}

func (s *QueryUnionTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
