// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataServiceApiAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataServiceApiAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataServiceApiAuthorityResponseBody) *DeleteDataServiceApiAuthorityResponse
	GetBody() *DeleteDataServiceApiAuthorityResponseBody
}

type DeleteDataServiceApiAuthorityResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataServiceApiAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataServiceApiAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiAuthorityResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataServiceApiAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataServiceApiAuthorityResponse) GetBody() *DeleteDataServiceApiAuthorityResponseBody {
	return s.Body
}

func (s *DeleteDataServiceApiAuthorityResponse) SetHeaders(v map[string]*string) *DeleteDataServiceApiAuthorityResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataServiceApiAuthorityResponse) SetStatusCode(v int32) *DeleteDataServiceApiAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityResponse) SetBody(v *DeleteDataServiceApiAuthorityResponseBody) *DeleteDataServiceApiAuthorityResponse {
	s.Body = v
	return s
}

func (s *DeleteDataServiceApiAuthorityResponse) Validate() error {
	return dara.Validate(s)
}
