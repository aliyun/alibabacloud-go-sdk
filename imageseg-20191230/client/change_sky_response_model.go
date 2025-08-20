// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSkyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeSkyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeSkyResponse
	GetStatusCode() *int32
	SetBody(v *ChangeSkyResponseBody) *ChangeSkyResponse
	GetBody() *ChangeSkyResponseBody
}

type ChangeSkyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeSkyResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeSkyResponse) GoString() string {
	return s.String()
}

func (s *ChangeSkyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeSkyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeSkyResponse) GetBody() *ChangeSkyResponseBody {
	return s.Body
}

func (s *ChangeSkyResponse) SetHeaders(v map[string]*string) *ChangeSkyResponse {
	s.Headers = v
	return s
}

func (s *ChangeSkyResponse) SetStatusCode(v int32) *ChangeSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeSkyResponse) SetBody(v *ChangeSkyResponseBody) *ChangeSkyResponse {
	s.Body = v
	return s
}

func (s *ChangeSkyResponse) Validate() error {
	return dara.Validate(s)
}
