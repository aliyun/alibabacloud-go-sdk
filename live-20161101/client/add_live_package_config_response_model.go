// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLivePackageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLivePackageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLivePackageConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLivePackageConfigResponseBody) *AddLivePackageConfigResponse
	GetBody() *AddLivePackageConfigResponseBody
}

type AddLivePackageConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLivePackageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLivePackageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLivePackageConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLivePackageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLivePackageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLivePackageConfigResponse) GetBody() *AddLivePackageConfigResponseBody {
	return s.Body
}

func (s *AddLivePackageConfigResponse) SetHeaders(v map[string]*string) *AddLivePackageConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLivePackageConfigResponse) SetStatusCode(v int32) *AddLivePackageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLivePackageConfigResponse) SetBody(v *AddLivePackageConfigResponseBody) *AddLivePackageConfigResponse {
	s.Body = v
	return s
}

func (s *AddLivePackageConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
