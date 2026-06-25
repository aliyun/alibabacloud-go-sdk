// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DetailRequest
	GetInstanceName() *string
}

type DetailRequest struct {
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Instance names must be unique within the same region. Make sure you distinguish between the two.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s *DetailRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DetailRequest) SetInstanceName(v string) *DetailRequest {
	s.InstanceName = &v
	return s
}

func (s *DetailRequest) Validate() error {
	return dara.Validate(s)
}
