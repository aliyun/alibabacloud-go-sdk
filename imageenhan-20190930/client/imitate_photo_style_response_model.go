// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImitatePhotoStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImitatePhotoStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImitatePhotoStyleResponse
	GetStatusCode() *int32
	SetBody(v *ImitatePhotoStyleResponseBody) *ImitatePhotoStyleResponse
	GetBody() *ImitatePhotoStyleResponseBody
}

type ImitatePhotoStyleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImitatePhotoStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImitatePhotoStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s ImitatePhotoStyleResponse) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImitatePhotoStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImitatePhotoStyleResponse) GetBody() *ImitatePhotoStyleResponseBody {
	return s.Body
}

func (s *ImitatePhotoStyleResponse) SetHeaders(v map[string]*string) *ImitatePhotoStyleResponse {
	s.Headers = v
	return s
}

func (s *ImitatePhotoStyleResponse) SetStatusCode(v int32) *ImitatePhotoStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *ImitatePhotoStyleResponse) SetBody(v *ImitatePhotoStyleResponseBody) *ImitatePhotoStyleResponse {
	s.Body = v
	return s
}

func (s *ImitatePhotoStyleResponse) Validate() error {
	return dara.Validate(s)
}
