// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlignStoragePrimaryAzoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlignStoragePrimaryAzoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlignStoragePrimaryAzoneResponse
	GetStatusCode() *int32
	SetBody(v *AlignStoragePrimaryAzoneResponseBody) *AlignStoragePrimaryAzoneResponse
	GetBody() *AlignStoragePrimaryAzoneResponseBody
}

type AlignStoragePrimaryAzoneResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlignStoragePrimaryAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlignStoragePrimaryAzoneResponse) String() string {
	return dara.Prettify(s)
}

func (s AlignStoragePrimaryAzoneResponse) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlignStoragePrimaryAzoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlignStoragePrimaryAzoneResponse) GetBody() *AlignStoragePrimaryAzoneResponseBody {
	return s.Body
}

func (s *AlignStoragePrimaryAzoneResponse) SetHeaders(v map[string]*string) *AlignStoragePrimaryAzoneResponse {
	s.Headers = v
	return s
}

func (s *AlignStoragePrimaryAzoneResponse) SetStatusCode(v int32) *AlignStoragePrimaryAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponse) SetBody(v *AlignStoragePrimaryAzoneResponseBody) *AlignStoragePrimaryAzoneResponse {
	s.Body = v
	return s
}

func (s *AlignStoragePrimaryAzoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
