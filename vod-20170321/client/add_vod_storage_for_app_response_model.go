// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodStorageForAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVodStorageForAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVodStorageForAppResponse
	GetStatusCode() *int32
	SetBody(v *AddVodStorageForAppResponseBody) *AddVodStorageForAppResponse
	GetBody() *AddVodStorageForAppResponseBody
}

type AddVodStorageForAppResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVodStorageForAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVodStorageForAppResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVodStorageForAppResponse) GoString() string {
	return s.String()
}

func (s *AddVodStorageForAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVodStorageForAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVodStorageForAppResponse) GetBody() *AddVodStorageForAppResponseBody {
	return s.Body
}

func (s *AddVodStorageForAppResponse) SetHeaders(v map[string]*string) *AddVodStorageForAppResponse {
	s.Headers = v
	return s
}

func (s *AddVodStorageForAppResponse) SetStatusCode(v int32) *AddVodStorageForAppResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVodStorageForAppResponse) SetBody(v *AddVodStorageForAppResponseBody) *AddVodStorageForAppResponse {
	s.Body = v
	return s
}

func (s *AddVodStorageForAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
