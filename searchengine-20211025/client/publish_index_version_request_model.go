// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishIndexVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *PublishIndexVersionRequest
	GetBody() map[string]interface{}
}

type PublishIndexVersionRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishIndexVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *PublishIndexVersionRequest) SetBody(v map[string]interface{}) *PublishIndexVersionRequest {
	s.Body = v
	return s
}

func (s *PublishIndexVersionRequest) Validate() error {
	return dara.Validate(s)
}
