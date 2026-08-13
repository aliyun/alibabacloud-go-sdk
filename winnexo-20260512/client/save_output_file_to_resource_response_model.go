// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveOutputFileToResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveOutputFileToResourceResponse
	GetStatusCode() *int32
	SetBody(v *SaveOutputFileToResourceResponseBody) *SaveOutputFileToResourceResponse
	GetBody() *SaveOutputFileToResourceResponseBody
}

type SaveOutputFileToResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveOutputFileToResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveOutputFileToResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceResponse) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveOutputFileToResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveOutputFileToResourceResponse) GetBody() *SaveOutputFileToResourceResponseBody {
	return s.Body
}

func (s *SaveOutputFileToResourceResponse) SetHeaders(v map[string]*string) *SaveOutputFileToResourceResponse {
	s.Headers = v
	return s
}

func (s *SaveOutputFileToResourceResponse) SetStatusCode(v int32) *SaveOutputFileToResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveOutputFileToResourceResponse) SetBody(v *SaveOutputFileToResourceResponseBody) *SaveOutputFileToResourceResponse {
	s.Body = v
	return s
}

func (s *SaveOutputFileToResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
