// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUniqueDeviceStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUniqueDeviceStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUniqueDeviceStatResponse
	GetStatusCode() *int32
	SetBody(v *QueryUniqueDeviceStatResponseBody) *QueryUniqueDeviceStatResponse
	GetBody() *QueryUniqueDeviceStatResponseBody
}

type QueryUniqueDeviceStatResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUniqueDeviceStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUniqueDeviceStatResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUniqueDeviceStatResponse) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUniqueDeviceStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUniqueDeviceStatResponse) GetBody() *QueryUniqueDeviceStatResponseBody {
	return s.Body
}

func (s *QueryUniqueDeviceStatResponse) SetHeaders(v map[string]*string) *QueryUniqueDeviceStatResponse {
	s.Headers = v
	return s
}

func (s *QueryUniqueDeviceStatResponse) SetStatusCode(v int32) *QueryUniqueDeviceStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUniqueDeviceStatResponse) SetBody(v *QueryUniqueDeviceStatResponseBody) *QueryUniqueDeviceStatResponse {
	s.Body = v
	return s
}

func (s *QueryUniqueDeviceStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
