// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToGroupResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveGroupOutputFileToGroupResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveGroupOutputFileToGroupResourceResponse
	GetStatusCode() *int32
	SetBody(v *SaveGroupOutputFileToGroupResourceResponseBody) *SaveGroupOutputFileToGroupResourceResponse
	GetBody() *SaveGroupOutputFileToGroupResourceResponseBody
}

type SaveGroupOutputFileToGroupResourceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveGroupOutputFileToGroupResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveGroupOutputFileToGroupResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceResponse) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveGroupOutputFileToGroupResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveGroupOutputFileToGroupResourceResponse) GetBody() *SaveGroupOutputFileToGroupResourceResponseBody {
	return s.Body
}

func (s *SaveGroupOutputFileToGroupResourceResponse) SetHeaders(v map[string]*string) *SaveGroupOutputFileToGroupResourceResponse {
	s.Headers = v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponse) SetStatusCode(v int32) *SaveGroupOutputFileToGroupResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponse) SetBody(v *SaveGroupOutputFileToGroupResourceResponseBody) *SaveGroupOutputFileToGroupResourceResponse {
	s.Body = v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
