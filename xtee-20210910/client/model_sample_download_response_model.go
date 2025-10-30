// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSampleDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelSampleDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelSampleDownloadResponse
	GetStatusCode() *int32
	SetBody(v *ModelSampleDownloadResponseBody) *ModelSampleDownloadResponse
	GetBody() *ModelSampleDownloadResponseBody
}

type ModelSampleDownloadResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelSampleDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelSampleDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelSampleDownloadResponse) GoString() string {
	return s.String()
}

func (s *ModelSampleDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelSampleDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelSampleDownloadResponse) GetBody() *ModelSampleDownloadResponseBody {
	return s.Body
}

func (s *ModelSampleDownloadResponse) SetHeaders(v map[string]*string) *ModelSampleDownloadResponse {
	s.Headers = v
	return s
}

func (s *ModelSampleDownloadResponse) SetStatusCode(v int32) *ModelSampleDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelSampleDownloadResponse) SetBody(v *ModelSampleDownloadResponseBody) *ModelSampleDownloadResponse {
	s.Body = v
	return s
}

func (s *ModelSampleDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
