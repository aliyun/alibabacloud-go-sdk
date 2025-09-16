// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverIndexResponse
	GetStatusCode() *int32
	SetBody(v *RecoverIndexResponseBody) *RecoverIndexResponse
	GetBody() *RecoverIndexResponseBody
}

type RecoverIndexResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverIndexResponse) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverIndexResponse) GetBody() *RecoverIndexResponseBody {
	return s.Body
}

func (s *RecoverIndexResponse) SetHeaders(v map[string]*string) *RecoverIndexResponse {
	s.Headers = v
	return s
}

func (s *RecoverIndexResponse) SetStatusCode(v int32) *RecoverIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverIndexResponse) SetBody(v *RecoverIndexResponseBody) *RecoverIndexResponse {
	s.Body = v
	return s
}

func (s *RecoverIndexResponse) Validate() error {
	return dara.Validate(s)
}
