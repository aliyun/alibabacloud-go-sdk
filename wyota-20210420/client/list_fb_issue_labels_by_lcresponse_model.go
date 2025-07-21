// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFbIssueLabelsByLCResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFbIssueLabelsByLCResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFbIssueLabelsByLCResponse
	GetStatusCode() *int32
	SetBody(v *ListFbIssueLabelsByLCResponseBody) *ListFbIssueLabelsByLCResponse
	GetBody() *ListFbIssueLabelsByLCResponseBody
}

type ListFbIssueLabelsByLCResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFbIssueLabelsByLCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFbIssueLabelsByLCResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponse) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFbIssueLabelsByLCResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFbIssueLabelsByLCResponse) GetBody() *ListFbIssueLabelsByLCResponseBody {
	return s.Body
}

func (s *ListFbIssueLabelsByLCResponse) SetHeaders(v map[string]*string) *ListFbIssueLabelsByLCResponse {
	s.Headers = v
	return s
}

func (s *ListFbIssueLabelsByLCResponse) SetStatusCode(v int32) *ListFbIssueLabelsByLCResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponse) SetBody(v *ListFbIssueLabelsByLCResponseBody) *ListFbIssueLabelsByLCResponse {
	s.Body = v
	return s
}

func (s *ListFbIssueLabelsByLCResponse) Validate() error {
	return dara.Validate(s)
}
