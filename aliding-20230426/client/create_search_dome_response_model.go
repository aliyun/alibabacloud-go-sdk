// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSearchDomeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSearchDomeResponse
	GetStatusCode() *int32
	SetBody(v *CreateSearchDomeResponseBody) *CreateSearchDomeResponse
	GetBody() *CreateSearchDomeResponseBody
}

type CreateSearchDomeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchDomeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchDomeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSearchDomeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSearchDomeResponse) GetBody() *CreateSearchDomeResponseBody {
	return s.Body
}

func (s *CreateSearchDomeResponse) SetHeaders(v map[string]*string) *CreateSearchDomeResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchDomeResponse) SetStatusCode(v int32) *CreateSearchDomeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchDomeResponse) SetBody(v *CreateSearchDomeResponseBody) *CreateSearchDomeResponse {
	s.Body = v
	return s
}

func (s *CreateSearchDomeResponse) Validate() error {
	return dara.Validate(s)
}
