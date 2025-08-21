// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistApplicationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DistApplicationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DistApplicationDataResponse
	GetStatusCode() *int32
	SetBody(v *DistApplicationDataResponseBody) *DistApplicationDataResponse
	GetBody() *DistApplicationDataResponseBody
}

type DistApplicationDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DistApplicationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DistApplicationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DistApplicationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DistApplicationDataResponse) GetBody() *DistApplicationDataResponseBody {
	return s.Body
}

func (s *DistApplicationDataResponse) SetHeaders(v map[string]*string) *DistApplicationDataResponse {
	s.Headers = v
	return s
}

func (s *DistApplicationDataResponse) SetStatusCode(v int32) *DistApplicationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DistApplicationDataResponse) SetBody(v *DistApplicationDataResponseBody) *DistApplicationDataResponse {
	s.Body = v
	return s
}

func (s *DistApplicationDataResponse) Validate() error {
	return dara.Validate(s)
}
