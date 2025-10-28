// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLogPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLogPathResponse
	GetStatusCode() *int32
	SetBody(v *AddLogPathResponseBody) *AddLogPathResponse
	GetBody() *AddLogPathResponseBody
}

type AddLogPathResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLogPathResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLogPathResponse) GoString() string {
	return s.String()
}

func (s *AddLogPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLogPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLogPathResponse) GetBody() *AddLogPathResponseBody {
	return s.Body
}

func (s *AddLogPathResponse) SetHeaders(v map[string]*string) *AddLogPathResponse {
	s.Headers = v
	return s
}

func (s *AddLogPathResponse) SetStatusCode(v int32) *AddLogPathResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLogPathResponse) SetBody(v *AddLogPathResponseBody) *AddLogPathResponse {
	s.Body = v
	return s
}

func (s *AddLogPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
