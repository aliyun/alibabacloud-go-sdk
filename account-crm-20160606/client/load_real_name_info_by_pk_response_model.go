// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadRealNameInfoByPkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoadRealNameInfoByPkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoadRealNameInfoByPkResponse
	GetStatusCode() *int32
	SetBody(v *LoadRealNameInfoByPkResponseBody) *LoadRealNameInfoByPkResponse
	GetBody() *LoadRealNameInfoByPkResponseBody
}

type LoadRealNameInfoByPkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoadRealNameInfoByPkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoadRealNameInfoByPkResponse) String() string {
	return dara.Prettify(s)
}

func (s LoadRealNameInfoByPkResponse) GoString() string {
	return s.String()
}

func (s *LoadRealNameInfoByPkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoadRealNameInfoByPkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoadRealNameInfoByPkResponse) GetBody() *LoadRealNameInfoByPkResponseBody {
	return s.Body
}

func (s *LoadRealNameInfoByPkResponse) SetHeaders(v map[string]*string) *LoadRealNameInfoByPkResponse {
	s.Headers = v
	return s
}

func (s *LoadRealNameInfoByPkResponse) SetStatusCode(v int32) *LoadRealNameInfoByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadRealNameInfoByPkResponse) SetBody(v *LoadRealNameInfoByPkResponseBody) *LoadRealNameInfoByPkResponse {
	s.Body = v
	return s
}

func (s *LoadRealNameInfoByPkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
