// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *FetchImageTaskResponseBody) *FetchImageTaskResponse
	GetBody() *FetchImageTaskResponseBody
}

type FetchImageTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchImageTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchImageTaskResponse) GetBody() *FetchImageTaskResponseBody {
	return s.Body
}

func (s *FetchImageTaskResponse) SetHeaders(v map[string]*string) *FetchImageTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchImageTaskResponse) SetStatusCode(v int32) *FetchImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchImageTaskResponse) SetBody(v *FetchImageTaskResponseBody) *FetchImageTaskResponse {
	s.Body = v
	return s
}

func (s *FetchImageTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
