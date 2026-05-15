// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotlineSessionQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotlineSessionQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotlineSessionQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotlineSessionQueryResponseBody) *HotlineSessionQueryResponse
	GetBody() *HotlineSessionQueryResponseBody
}

type HotlineSessionQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotlineSessionQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotlineSessionQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotlineSessionQueryResponse) GoString() string {
	return s.String()
}

func (s *HotlineSessionQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotlineSessionQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotlineSessionQueryResponse) GetBody() *HotlineSessionQueryResponseBody {
	return s.Body
}

func (s *HotlineSessionQueryResponse) SetHeaders(v map[string]*string) *HotlineSessionQueryResponse {
	s.Headers = v
	return s
}

func (s *HotlineSessionQueryResponse) SetStatusCode(v int32) *HotlineSessionQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotlineSessionQueryResponse) SetBody(v *HotlineSessionQueryResponseBody) *HotlineSessionQueryResponse {
	s.Body = v
	return s
}

func (s *HotlineSessionQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
