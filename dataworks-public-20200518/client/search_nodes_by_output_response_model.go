// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNodesByOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchNodesByOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchNodesByOutputResponse
	GetStatusCode() *int32
	SetBody(v *SearchNodesByOutputResponseBody) *SearchNodesByOutputResponse
	GetBody() *SearchNodesByOutputResponseBody
}

type SearchNodesByOutputResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchNodesByOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchNodesByOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchNodesByOutputResponse) GoString() string {
	return s.String()
}

func (s *SearchNodesByOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchNodesByOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchNodesByOutputResponse) GetBody() *SearchNodesByOutputResponseBody {
	return s.Body
}

func (s *SearchNodesByOutputResponse) SetHeaders(v map[string]*string) *SearchNodesByOutputResponse {
	s.Headers = v
	return s
}

func (s *SearchNodesByOutputResponse) SetStatusCode(v int32) *SearchNodesByOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchNodesByOutputResponse) SetBody(v *SearchNodesByOutputResponseBody) *SearchNodesByOutputResponse {
	s.Body = v
	return s
}

func (s *SearchNodesByOutputResponse) Validate() error {
	return dara.Validate(s)
}
