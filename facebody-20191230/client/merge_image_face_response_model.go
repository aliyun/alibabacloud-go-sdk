// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeImageFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MergeImageFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MergeImageFaceResponse
	GetStatusCode() *int32
	SetBody(v *MergeImageFaceResponseBody) *MergeImageFaceResponse
	GetBody() *MergeImageFaceResponseBody
}

type MergeImageFaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeImageFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeImageFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MergeImageFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MergeImageFaceResponse) GetBody() *MergeImageFaceResponseBody {
	return s.Body
}

func (s *MergeImageFaceResponse) SetHeaders(v map[string]*string) *MergeImageFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeImageFaceResponse) SetStatusCode(v int32) *MergeImageFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeImageFaceResponse) SetBody(v *MergeImageFaceResponseBody) *MergeImageFaceResponse {
	s.Body = v
	return s
}

func (s *MergeImageFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
