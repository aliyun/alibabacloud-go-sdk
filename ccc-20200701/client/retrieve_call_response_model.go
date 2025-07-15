// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveCallResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveCallResponseBody) *RetrieveCallResponse
	GetBody() *RetrieveCallResponseBody
}

type RetrieveCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveCallResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallResponse) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveCallResponse) GetBody() *RetrieveCallResponseBody {
	return s.Body
}

func (s *RetrieveCallResponse) SetHeaders(v map[string]*string) *RetrieveCallResponse {
	s.Headers = v
	return s
}

func (s *RetrieveCallResponse) SetStatusCode(v int32) *RetrieveCallResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveCallResponse) SetBody(v *RetrieveCallResponseBody) *RetrieveCallResponse {
	s.Body = v
	return s
}

func (s *RetrieveCallResponse) Validate() error {
	return dara.Validate(s)
}
