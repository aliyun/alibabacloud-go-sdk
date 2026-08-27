// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveGroupOutputFileToPersonalResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveGroupOutputFileToPersonalResourceResponse
	GetStatusCode() *int32
	SetBody(v *SaveGroupOutputFileToPersonalResourceResponseBody) *SaveGroupOutputFileToPersonalResourceResponse
	GetBody() *SaveGroupOutputFileToPersonalResourceResponseBody
}

type SaveGroupOutputFileToPersonalResourceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveGroupOutputFileToPersonalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveGroupOutputFileToPersonalResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceResponse) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) GetBody() *SaveGroupOutputFileToPersonalResourceResponseBody {
	return s.Body
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) SetHeaders(v map[string]*string) *SaveGroupOutputFileToPersonalResourceResponse {
	s.Headers = v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) SetStatusCode(v int32) *SaveGroupOutputFileToPersonalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) SetBody(v *SaveGroupOutputFileToPersonalResourceResponseBody) *SaveGroupOutputFileToPersonalResourceResponse {
	s.Body = v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
