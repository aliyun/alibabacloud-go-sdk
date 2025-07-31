// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSourceBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSourceBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSourceBasicInfoResponseBody) *UpdateDataSourceBasicInfoResponse
	GetBody() *UpdateDataSourceBasicInfoResponseBody
}

type UpdateDataSourceBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSourceBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSourceBasicInfoResponse) GetBody() *UpdateDataSourceBasicInfoResponseBody {
	return s.Body
}

func (s *UpdateDataSourceBasicInfoResponse) SetHeaders(v map[string]*string) *UpdateDataSourceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceBasicInfoResponse) SetStatusCode(v int32) *UpdateDataSourceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponse) SetBody(v *UpdateDataSourceBasicInfoResponseBody) *UpdateDataSourceBasicInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSourceBasicInfoResponse) Validate() error {
	return dara.Validate(s)
}
