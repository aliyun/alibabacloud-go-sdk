// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PublishImageRequest
	GetId() *string
	SetProcessId(v string) *PublishImageRequest
	GetProcessId() *string
}

type PublishImageRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image publish execution ID, which is used as an idempotence identifier.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s PublishImageRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishImageRequest) GoString() string {
	return s.String()
}

func (s *PublishImageRequest) GetId() *string {
	return s.Id
}

func (s *PublishImageRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *PublishImageRequest) SetId(v string) *PublishImageRequest {
	s.Id = &v
	return s
}

func (s *PublishImageRequest) SetProcessId(v string) *PublishImageRequest {
	s.ProcessId = &v
	return s
}

func (s *PublishImageRequest) Validate() error {
	return dara.Validate(s)
}
