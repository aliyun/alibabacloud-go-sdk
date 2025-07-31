// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFreezeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOssCheckResultsFreezeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOssCheckResultsFreezeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOssCheckResultsFreezeResponseBody) *UpdateOssCheckResultsFreezeResponse
	GetBody() *UpdateOssCheckResultsFreezeResponseBody
}

type UpdateOssCheckResultsFreezeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssCheckResultsFreezeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssCheckResultsFreezeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFreezeResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFreezeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOssCheckResultsFreezeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOssCheckResultsFreezeResponse) GetBody() *UpdateOssCheckResultsFreezeResponseBody {
	return s.Body
}

func (s *UpdateOssCheckResultsFreezeResponse) SetHeaders(v map[string]*string) *UpdateOssCheckResultsFreezeResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponse) SetStatusCode(v int32) *UpdateOssCheckResultsFreezeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponse) SetBody(v *UpdateOssCheckResultsFreezeResponseBody) *UpdateOssCheckResultsFreezeResponse {
	s.Body = v
	return s
}

func (s *UpdateOssCheckResultsFreezeResponse) Validate() error {
	return dara.Validate(s)
}
