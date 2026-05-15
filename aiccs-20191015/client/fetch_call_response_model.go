// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchCallResponse
	GetStatusCode() *int32
	SetBody(v *FetchCallResponseBody) *FetchCallResponse
	GetBody() *FetchCallResponseBody
}

type FetchCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchCallResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchCallResponse) GoString() string {
	return s.String()
}

func (s *FetchCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchCallResponse) GetBody() *FetchCallResponseBody {
	return s.Body
}

func (s *FetchCallResponse) SetHeaders(v map[string]*string) *FetchCallResponse {
	s.Headers = v
	return s
}

func (s *FetchCallResponse) SetStatusCode(v int32) *FetchCallResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchCallResponse) SetBody(v *FetchCallResponseBody) *FetchCallResponse {
	s.Body = v
	return s
}

func (s *FetchCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
