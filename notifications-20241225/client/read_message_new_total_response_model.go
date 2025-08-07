// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageNewTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMessageNewTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMessageNewTotalResponse
	GetStatusCode() *int32
	SetBody(v *ReadMessageNewTotalResponseBody) *ReadMessageNewTotalResponse
	GetBody() *ReadMessageNewTotalResponseBody
}

type ReadMessageNewTotalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageNewTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageNewTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageNewTotalResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMessageNewTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMessageNewTotalResponse) GetBody() *ReadMessageNewTotalResponseBody {
	return s.Body
}

func (s *ReadMessageNewTotalResponse) SetHeaders(v map[string]*string) *ReadMessageNewTotalResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageNewTotalResponse) SetStatusCode(v int32) *ReadMessageNewTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageNewTotalResponse) SetBody(v *ReadMessageNewTotalResponseBody) *ReadMessageNewTotalResponse {
	s.Body = v
	return s
}

func (s *ReadMessageNewTotalResponse) Validate() error {
	return dara.Validate(s)
}
