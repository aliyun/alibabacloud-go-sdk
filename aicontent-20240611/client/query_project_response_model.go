// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProjectResponse
	GetStatusCode() *int32
	SetBody(v *QueryProjectResponseBody) *QueryProjectResponse
	GetBody() *QueryProjectResponseBody
}

type QueryProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProjectResponse) GetBody() *QueryProjectResponseBody {
	return s.Body
}

func (s *QueryProjectResponse) SetHeaders(v map[string]*string) *QueryProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectResponse) SetStatusCode(v int32) *QueryProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectResponse) SetBody(v *QueryProjectResponseBody) *QueryProjectResponse {
	s.Body = v
	return s
}

func (s *QueryProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
