// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGalleryModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelGalleryModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelGalleryModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelGalleryModelsResponseBody) *ListModelGalleryModelsResponse
	GetBody() *ListModelGalleryModelsResponseBody
}

type ListModelGalleryModelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelGalleryModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelGalleryModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelGalleryModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelGalleryModelsResponse) GetBody() *ListModelGalleryModelsResponseBody {
	return s.Body
}

func (s *ListModelGalleryModelsResponse) SetHeaders(v map[string]*string) *ListModelGalleryModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelGalleryModelsResponse) SetStatusCode(v int32) *ListModelGalleryModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelGalleryModelsResponse) SetBody(v *ListModelGalleryModelsResponseBody) *ListModelGalleryModelsResponse {
	s.Body = v
	return s
}

func (s *ListModelGalleryModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
