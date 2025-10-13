// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishWebApplicationRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *PublishWebApplicationRevisionRequest
	GetNamespaceId() *string
	SetBody(v *PublishWebApplicationRevisionInput) *PublishWebApplicationRevisionRequest
	GetBody() *PublishWebApplicationRevisionInput
}

type PublishWebApplicationRevisionRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configurations of the version.
	//
	// This parameter is required.
	Body *PublishWebApplicationRevisionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishWebApplicationRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishWebApplicationRevisionRequest) GoString() string {
	return s.String()
}

func (s *PublishWebApplicationRevisionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *PublishWebApplicationRevisionRequest) GetBody() *PublishWebApplicationRevisionInput {
	return s.Body
}

func (s *PublishWebApplicationRevisionRequest) SetNamespaceId(v string) *PublishWebApplicationRevisionRequest {
	s.NamespaceId = &v
	return s
}

func (s *PublishWebApplicationRevisionRequest) SetBody(v *PublishWebApplicationRevisionInput) *PublishWebApplicationRevisionRequest {
	s.Body = v
	return s
}

func (s *PublishWebApplicationRevisionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
