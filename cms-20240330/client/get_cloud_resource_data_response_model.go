// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudResourceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudResourceDataResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudResourceDataResponseBody) *GetCloudResourceDataResponse
	GetBody() *GetCloudResourceDataResponseBody
}

type GetCloudResourceDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudResourceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudResourceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataResponse) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudResourceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudResourceDataResponse) GetBody() *GetCloudResourceDataResponseBody {
	return s.Body
}

func (s *GetCloudResourceDataResponse) SetHeaders(v map[string]*string) *GetCloudResourceDataResponse {
	s.Headers = v
	return s
}

func (s *GetCloudResourceDataResponse) SetStatusCode(v int32) *GetCloudResourceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudResourceDataResponse) SetBody(v *GetCloudResourceDataResponseBody) *GetCloudResourceDataResponse {
	s.Body = v
	return s
}

func (s *GetCloudResourceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
