// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaDataComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMetaDataComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMetaDataComponentResponse
	GetStatusCode() *int32
	SetBody(v *AddMetaDataComponentResponseBody) *AddMetaDataComponentResponse
	GetBody() *AddMetaDataComponentResponseBody
}

type AddMetaDataComponentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMetaDataComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMetaDataComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMetaDataComponentResponse) GoString() string {
	return s.String()
}

func (s *AddMetaDataComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMetaDataComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMetaDataComponentResponse) GetBody() *AddMetaDataComponentResponseBody {
	return s.Body
}

func (s *AddMetaDataComponentResponse) SetHeaders(v map[string]*string) *AddMetaDataComponentResponse {
	s.Headers = v
	return s
}

func (s *AddMetaDataComponentResponse) SetStatusCode(v int32) *AddMetaDataComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMetaDataComponentResponse) SetBody(v *AddMetaDataComponentResponseBody) *AddMetaDataComponentResponse {
	s.Body = v
	return s
}

func (s *AddMetaDataComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
