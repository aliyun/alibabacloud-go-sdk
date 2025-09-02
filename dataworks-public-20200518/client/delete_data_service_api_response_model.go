// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataServiceApiResponseBody) *DeleteDataServiceApiResponse
	GetBody() *DeleteDataServiceApiResponseBody
}

type DeleteDataServiceApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataServiceApiResponse) GetBody() *DeleteDataServiceApiResponseBody {
	return s.Body
}

func (s *DeleteDataServiceApiResponse) SetHeaders(v map[string]*string) *DeleteDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataServiceApiResponse) SetStatusCode(v int32) *DeleteDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataServiceApiResponse) SetBody(v *DeleteDataServiceApiResponseBody) *DeleteDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *DeleteDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
