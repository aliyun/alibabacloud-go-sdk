// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v []*GetMultiModalEmbeddingRequestInput) *GetMultiModalEmbeddingRequest
	GetInput() []*GetMultiModalEmbeddingRequestInput
	SetOptions(v map[string]interface{}) *GetMultiModalEmbeddingRequest
	GetOptions() map[string]interface{}
}

type GetMultiModalEmbeddingRequest struct {
	Input   []*GetMultiModalEmbeddingRequestInput `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	Options map[string]interface{}                `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GetMultiModalEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingRequest) GetInput() []*GetMultiModalEmbeddingRequestInput {
	return s.Input
}

func (s *GetMultiModalEmbeddingRequest) GetOptions() map[string]interface{} {
	return s.Options
}

func (s *GetMultiModalEmbeddingRequest) SetInput(v []*GetMultiModalEmbeddingRequestInput) *GetMultiModalEmbeddingRequest {
	s.Input = v
	return s
}

func (s *GetMultiModalEmbeddingRequest) SetOptions(v map[string]interface{}) *GetMultiModalEmbeddingRequest {
	s.Options = v
	return s
}

func (s *GetMultiModalEmbeddingRequest) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMultiModalEmbeddingRequestInput struct {
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetMultiModalEmbeddingRequestInput) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingRequestInput) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingRequestInput) GetImage() *string {
	return s.Image
}

func (s *GetMultiModalEmbeddingRequestInput) GetText() *string {
	return s.Text
}

func (s *GetMultiModalEmbeddingRequestInput) SetImage(v string) *GetMultiModalEmbeddingRequestInput {
	s.Image = &v
	return s
}

func (s *GetMultiModalEmbeddingRequestInput) SetText(v string) *GetMultiModalEmbeddingRequestInput {
	s.Text = &v
	return s
}

func (s *GetMultiModalEmbeddingRequestInput) Validate() error {
	return dara.Validate(s)
}
