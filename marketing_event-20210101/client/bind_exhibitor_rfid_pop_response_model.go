// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindExhibitorRfidPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindExhibitorRfidPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindExhibitorRfidPopResponse
	GetStatusCode() *int32
	SetBody(v *BindExhibitorRfidPopResponseBody) *BindExhibitorRfidPopResponse
	GetBody() *BindExhibitorRfidPopResponseBody
}

type BindExhibitorRfidPopResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindExhibitorRfidPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindExhibitorRfidPopResponse) String() string {
	return dara.Prettify(s)
}

func (s BindExhibitorRfidPopResponse) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindExhibitorRfidPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindExhibitorRfidPopResponse) GetBody() *BindExhibitorRfidPopResponseBody {
	return s.Body
}

func (s *BindExhibitorRfidPopResponse) SetHeaders(v map[string]*string) *BindExhibitorRfidPopResponse {
	s.Headers = v
	return s
}

func (s *BindExhibitorRfidPopResponse) SetStatusCode(v int32) *BindExhibitorRfidPopResponse {
	s.StatusCode = &v
	return s
}

func (s *BindExhibitorRfidPopResponse) SetBody(v *BindExhibitorRfidPopResponseBody) *BindExhibitorRfidPopResponse {
	s.Body = v
	return s
}

func (s *BindExhibitorRfidPopResponse) Validate() error {
	return dara.Validate(s)
}
