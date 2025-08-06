// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVodServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVodServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVodServiceResponseBody) *ModifyVodServiceResponse
	GetBody() *ModifyVodServiceResponseBody
}

type ModifyVodServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVodServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVodServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyVodServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVodServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVodServiceResponse) GetBody() *ModifyVodServiceResponseBody {
	return s.Body
}

func (s *ModifyVodServiceResponse) SetHeaders(v map[string]*string) *ModifyVodServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyVodServiceResponse) SetStatusCode(v int32) *ModifyVodServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVodServiceResponse) SetBody(v *ModifyVodServiceResponseBody) *ModifyVodServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyVodServiceResponse) Validate() error {
	return dara.Validate(s)
}
