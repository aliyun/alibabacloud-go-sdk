// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceByTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceByTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceByTagResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceByTagResponseBody) *QueryInstanceByTagResponse
	GetBody() *QueryInstanceByTagResponseBody
}

type QueryInstanceByTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceByTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceByTagResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceByTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceByTagResponse) GetBody() *QueryInstanceByTagResponseBody {
	return s.Body
}

func (s *QueryInstanceByTagResponse) SetHeaders(v map[string]*string) *QueryInstanceByTagResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceByTagResponse) SetStatusCode(v int32) *QueryInstanceByTagResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceByTagResponse) SetBody(v *QueryInstanceByTagResponseBody) *QueryInstanceByTagResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceByTagResponse) Validate() error {
	return dara.Validate(s)
}
