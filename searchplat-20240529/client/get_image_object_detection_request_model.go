// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageObjectDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *GetImageObjectDetectionRequestImage) *GetImageObjectDetectionRequest
	GetImage() *GetImageObjectDetectionRequestImage
	SetOptions(v map[string]interface{}) *GetImageObjectDetectionRequest
	GetOptions() map[string]interface{}
}

type GetImageObjectDetectionRequest struct {
	Image   *GetImageObjectDetectionRequestImage `json:"image,omitempty" xml:"image,omitempty" type:"Struct"`
	Options map[string]interface{}               `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GetImageObjectDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionRequest) GetImage() *GetImageObjectDetectionRequestImage {
	return s.Image
}

func (s *GetImageObjectDetectionRequest) GetOptions() map[string]interface{} {
	return s.Options
}

func (s *GetImageObjectDetectionRequest) SetImage(v *GetImageObjectDetectionRequestImage) *GetImageObjectDetectionRequest {
	s.Image = v
	return s
}

func (s *GetImageObjectDetectionRequest) SetOptions(v map[string]interface{}) *GetImageObjectDetectionRequest {
	s.Options = v
	return s
}

func (s *GetImageObjectDetectionRequest) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageObjectDetectionRequestImage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Url     *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetImageObjectDetectionRequestImage) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionRequestImage) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionRequestImage) GetContent() *string {
	return s.Content
}

func (s *GetImageObjectDetectionRequestImage) GetUrl() *string {
	return s.Url
}

func (s *GetImageObjectDetectionRequestImage) SetContent(v string) *GetImageObjectDetectionRequestImage {
	s.Content = &v
	return s
}

func (s *GetImageObjectDetectionRequestImage) SetUrl(v string) *GetImageObjectDetectionRequestImage {
	s.Url = &v
	return s
}

func (s *GetImageObjectDetectionRequestImage) Validate() error {
	return dara.Validate(s)
}
