// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelieveAccountRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RelieveAccountRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RelieveAccountRelationResponse
	GetStatusCode() *int32
	SetBody(v *RelieveAccountRelationResponseBody) *RelieveAccountRelationResponse
	GetBody() *RelieveAccountRelationResponseBody
}

type RelieveAccountRelationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelieveAccountRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RelieveAccountRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s RelieveAccountRelationResponse) GoString() string {
	return s.String()
}

func (s *RelieveAccountRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RelieveAccountRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RelieveAccountRelationResponse) GetBody() *RelieveAccountRelationResponseBody {
	return s.Body
}

func (s *RelieveAccountRelationResponse) SetHeaders(v map[string]*string) *RelieveAccountRelationResponse {
	s.Headers = v
	return s
}

func (s *RelieveAccountRelationResponse) SetStatusCode(v int32) *RelieveAccountRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *RelieveAccountRelationResponse) SetBody(v *RelieveAccountRelationResponseBody) *RelieveAccountRelationResponse {
	s.Body = v
	return s
}

func (s *RelieveAccountRelationResponse) Validate() error {
	return dara.Validate(s)
}
