// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudResourceResponseBody) *ModifyCloudResourceResponse
	GetBody() *ModifyCloudResourceResponseBody
}

type ModifyCloudResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudResourceResponse) GetBody() *ModifyCloudResourceResponseBody {
	return s.Body
}

func (s *ModifyCloudResourceResponse) SetHeaders(v map[string]*string) *ModifyCloudResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudResourceResponse) SetStatusCode(v int32) *ModifyCloudResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudResourceResponse) SetBody(v *ModifyCloudResourceResponseBody) *ModifyCloudResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudResourceResponse) Validate() error {
	return dara.Validate(s)
}
