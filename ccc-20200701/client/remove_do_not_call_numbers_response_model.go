// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDoNotCallNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDoNotCallNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDoNotCallNumbersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDoNotCallNumbersResponseBody) *RemoveDoNotCallNumbersResponse
	GetBody() *RemoveDoNotCallNumbersResponseBody
}

type RemoveDoNotCallNumbersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDoNotCallNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDoNotCallNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDoNotCallNumbersResponse) GoString() string {
	return s.String()
}

func (s *RemoveDoNotCallNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDoNotCallNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDoNotCallNumbersResponse) GetBody() *RemoveDoNotCallNumbersResponseBody {
	return s.Body
}

func (s *RemoveDoNotCallNumbersResponse) SetHeaders(v map[string]*string) *RemoveDoNotCallNumbersResponse {
	s.Headers = v
	return s
}

func (s *RemoveDoNotCallNumbersResponse) SetStatusCode(v int32) *RemoveDoNotCallNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDoNotCallNumbersResponse) SetBody(v *RemoveDoNotCallNumbersResponseBody) *RemoveDoNotCallNumbersResponse {
	s.Body = v
	return s
}

func (s *RemoveDoNotCallNumbersResponse) Validate() error {
	return dara.Validate(s)
}
