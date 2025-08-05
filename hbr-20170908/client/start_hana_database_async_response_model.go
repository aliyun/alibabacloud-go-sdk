// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHanaDatabaseAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartHanaDatabaseAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartHanaDatabaseAsyncResponse
	GetStatusCode() *int32
	SetBody(v *StartHanaDatabaseAsyncResponseBody) *StartHanaDatabaseAsyncResponse
	GetBody() *StartHanaDatabaseAsyncResponseBody
}

type StartHanaDatabaseAsyncResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartHanaDatabaseAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartHanaDatabaseAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s StartHanaDatabaseAsyncResponse) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartHanaDatabaseAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartHanaDatabaseAsyncResponse) GetBody() *StartHanaDatabaseAsyncResponseBody {
	return s.Body
}

func (s *StartHanaDatabaseAsyncResponse) SetHeaders(v map[string]*string) *StartHanaDatabaseAsyncResponse {
	s.Headers = v
	return s
}

func (s *StartHanaDatabaseAsyncResponse) SetStatusCode(v int32) *StartHanaDatabaseAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponse) SetBody(v *StartHanaDatabaseAsyncResponseBody) *StartHanaDatabaseAsyncResponse {
	s.Body = v
	return s
}

func (s *StartHanaDatabaseAsyncResponse) Validate() error {
	return dara.Validate(s)
}
