// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceNcdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceNcdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceNcdResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceNcdResponseBody) *QueryInstanceNcdResponse
	GetBody() *QueryInstanceNcdResponseBody
}

type QueryInstanceNcdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceNcdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceNcdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceNcdResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceNcdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceNcdResponse) GetBody() *QueryInstanceNcdResponseBody {
	return s.Body
}

func (s *QueryInstanceNcdResponse) SetHeaders(v map[string]*string) *QueryInstanceNcdResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceNcdResponse) SetStatusCode(v int32) *QueryInstanceNcdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceNcdResponse) SetBody(v *QueryInstanceNcdResponseBody) *QueryInstanceNcdResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceNcdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
