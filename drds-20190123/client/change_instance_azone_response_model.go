// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeInstanceAzoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeInstanceAzoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeInstanceAzoneResponse
	GetStatusCode() *int32
	SetBody(v *ChangeInstanceAzoneResponseBody) *ChangeInstanceAzoneResponse
	GetBody() *ChangeInstanceAzoneResponseBody
}

type ChangeInstanceAzoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeInstanceAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeInstanceAzoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeInstanceAzoneResponse) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeInstanceAzoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeInstanceAzoneResponse) GetBody() *ChangeInstanceAzoneResponseBody {
	return s.Body
}

func (s *ChangeInstanceAzoneResponse) SetHeaders(v map[string]*string) *ChangeInstanceAzoneResponse {
	s.Headers = v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetStatusCode(v int32) *ChangeInstanceAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeInstanceAzoneResponse) SetBody(v *ChangeInstanceAzoneResponseBody) *ChangeInstanceAzoneResponse {
	s.Body = v
	return s
}

func (s *ChangeInstanceAzoneResponse) Validate() error {
	return dara.Validate(s)
}
