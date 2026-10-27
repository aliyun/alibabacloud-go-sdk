// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMemoryRequestBody) *UpdateMemoryRequest
	GetBody() *UpdateMemoryRequestBody
}

type UpdateMemoryRequest struct {
	// The update memory request.
	//
	// This parameter is required.
	Body *UpdateMemoryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetBody() *UpdateMemoryRequestBody {
	return s.Body
}

func (s *UpdateMemoryRequest) SetBody(v *UpdateMemoryRequestBody) *UpdateMemoryRequest {
	s.Body = v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryRequestBody struct {
	// The updated string key-value metadata. You must provide at least one of Metadata or Text. A maximum of 16 keys are allowed. The maximum length of a key is 64 characters, and the maximum length of a value is 1,024 characters.
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The updated memory text. You must provide at least one of Text or Metadata. The maximum size is 32,000 UTF-8 bytes.
	//
	// example:
	//
	// The user enjoys hiking on weekends
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s UpdateMemoryRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequestBody) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *UpdateMemoryRequestBody) GetText() *string {
	return s.Text
}

func (s *UpdateMemoryRequestBody) SetMetadata(v map[string]*string) *UpdateMemoryRequestBody {
	s.Metadata = v
	return s
}

func (s *UpdateMemoryRequestBody) SetText(v string) *UpdateMemoryRequestBody {
	s.Text = &v
	return s
}

func (s *UpdateMemoryRequestBody) Validate() error {
	return dara.Validate(s)
}
