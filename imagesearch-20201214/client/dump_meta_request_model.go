// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DumpMetaRequest
	GetInstanceName() *string
}

type DumpMetaRequest struct {
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure you distinguish between them.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DumpMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DumpMetaRequest) SetInstanceName(v string) *DumpMetaRequest {
	s.InstanceName = &v
	return s
}

func (s *DumpMetaRequest) Validate() error {
	return dara.Validate(s)
}
