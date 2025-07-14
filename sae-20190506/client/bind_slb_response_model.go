// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSlbResponse
	GetStatusCode() *int32
	SetBody(v *BindSlbResponseBody) *BindSlbResponse
	GetBody() *BindSlbResponseBody
}

type BindSlbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSlbResponse) GoString() string {
	return s.String()
}

func (s *BindSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSlbResponse) GetBody() *BindSlbResponseBody {
	return s.Body
}

func (s *BindSlbResponse) SetHeaders(v map[string]*string) *BindSlbResponse {
	s.Headers = v
	return s
}

func (s *BindSlbResponse) SetStatusCode(v int32) *BindSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSlbResponse) SetBody(v *BindSlbResponseBody) *BindSlbResponse {
	s.Body = v
	return s
}

func (s *BindSlbResponse) Validate() error {
	return dara.Validate(s)
}
