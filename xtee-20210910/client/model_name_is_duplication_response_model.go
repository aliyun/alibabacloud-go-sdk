// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelNameIsDuplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelNameIsDuplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelNameIsDuplicationResponse
	GetStatusCode() *int32
	SetBody(v *ModelNameIsDuplicationResponseBody) *ModelNameIsDuplicationResponse
	GetBody() *ModelNameIsDuplicationResponseBody
}

type ModelNameIsDuplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelNameIsDuplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelNameIsDuplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelNameIsDuplicationResponse) GoString() string {
	return s.String()
}

func (s *ModelNameIsDuplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelNameIsDuplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelNameIsDuplicationResponse) GetBody() *ModelNameIsDuplicationResponseBody {
	return s.Body
}

func (s *ModelNameIsDuplicationResponse) SetHeaders(v map[string]*string) *ModelNameIsDuplicationResponse {
	s.Headers = v
	return s
}

func (s *ModelNameIsDuplicationResponse) SetStatusCode(v int32) *ModelNameIsDuplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelNameIsDuplicationResponse) SetBody(v *ModelNameIsDuplicationResponseBody) *ModelNameIsDuplicationResponse {
	s.Body = v
	return s
}

func (s *ModelNameIsDuplicationResponse) Validate() error {
	return dara.Validate(s)
}
