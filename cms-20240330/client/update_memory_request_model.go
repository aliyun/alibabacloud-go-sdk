// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v map[string]interface{}) *UpdateMemoryRequest
	GetMetadata() map[string]interface{}
	SetText(v string) *UpdateMemoryRequest
	GetText() *string
}

type UpdateMemoryRequest struct {
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// Likes to play tennis on weekends
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateMemoryRequest) GetText() *string {
	return s.Text
}

func (s *UpdateMemoryRequest) SetMetadata(v map[string]interface{}) *UpdateMemoryRequest {
	s.Metadata = v
	return s
}

func (s *UpdateMemoryRequest) SetText(v string) *UpdateMemoryRequest {
	s.Text = &v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
