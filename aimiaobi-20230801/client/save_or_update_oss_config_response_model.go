// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOrUpdateOssConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveOrUpdateOssConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveOrUpdateOssConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveOrUpdateOssConfigResponseBody) *SaveOrUpdateOssConfigResponse
	GetBody() *SaveOrUpdateOssConfigResponseBody
}

type SaveOrUpdateOssConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveOrUpdateOssConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveOrUpdateOssConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveOrUpdateOssConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveOrUpdateOssConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveOrUpdateOssConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveOrUpdateOssConfigResponse) GetBody() *SaveOrUpdateOssConfigResponseBody {
	return s.Body
}

func (s *SaveOrUpdateOssConfigResponse) SetHeaders(v map[string]*string) *SaveOrUpdateOssConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveOrUpdateOssConfigResponse) SetStatusCode(v int32) *SaveOrUpdateOssConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponse) SetBody(v *SaveOrUpdateOssConfigResponseBody) *SaveOrUpdateOssConfigResponse {
	s.Body = v
	return s
}

func (s *SaveOrUpdateOssConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
