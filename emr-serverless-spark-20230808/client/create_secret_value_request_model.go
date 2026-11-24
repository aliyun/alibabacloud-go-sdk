// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateSecretValueRequest
	GetComment() *string
	SetContent(v string) *CreateSecretValueRequest
	GetContent() *string
	SetName(v string) *CreateSecretValueRequest
	GetName() *string
}

type CreateSecretValueRequest struct {
	// The comment.
	//
	// example:
	//
	// My secret
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The content of the ciphertext.
	//
	// example:
	//
	// password
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the ciphertext.
	//
	// example:
	//
	// mysecret
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretValueRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretValueRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateSecretValueRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSecretValueRequest) GetName() *string {
	return s.Name
}

func (s *CreateSecretValueRequest) SetComment(v string) *CreateSecretValueRequest {
	s.Comment = &v
	return s
}

func (s *CreateSecretValueRequest) SetContent(v string) *CreateSecretValueRequest {
	s.Content = &v
	return s
}

func (s *CreateSecretValueRequest) SetName(v string) *CreateSecretValueRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretValueRequest) Validate() error {
	return dara.Validate(s)
}
