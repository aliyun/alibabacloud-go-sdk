// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadedPartsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUploadedPartsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUploadedPartsResponse
	GetStatusCode() *int32
	SetBody(v *ListUploadedPartsResponseBody) *ListUploadedPartsResponse
	GetBody() *ListUploadedPartsResponseBody
}

type ListUploadedPartsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUploadedPartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUploadedPartsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUploadedPartsResponse) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUploadedPartsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUploadedPartsResponse) GetBody() *ListUploadedPartsResponseBody {
	return s.Body
}

func (s *ListUploadedPartsResponse) SetHeaders(v map[string]*string) *ListUploadedPartsResponse {
	s.Headers = v
	return s
}

func (s *ListUploadedPartsResponse) SetStatusCode(v int32) *ListUploadedPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUploadedPartsResponse) SetBody(v *ListUploadedPartsResponseBody) *ListUploadedPartsResponse {
	s.Body = v
	return s
}

func (s *ListUploadedPartsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
