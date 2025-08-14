// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSampleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSampleDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSampleDataResponseBody) *DeleteSampleDataResponse
	GetBody() *DeleteSampleDataResponseBody
}

type DeleteSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSampleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteSampleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSampleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSampleDataResponse) GetBody() *DeleteSampleDataResponseBody {
	return s.Body
}

func (s *DeleteSampleDataResponse) SetHeaders(v map[string]*string) *DeleteSampleDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteSampleDataResponse) SetStatusCode(v int32) *DeleteSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSampleDataResponse) SetBody(v *DeleteSampleDataResponseBody) *DeleteSampleDataResponse {
	s.Body = v
	return s
}

func (s *DeleteSampleDataResponse) Validate() error {
	return dara.Validate(s)
}
