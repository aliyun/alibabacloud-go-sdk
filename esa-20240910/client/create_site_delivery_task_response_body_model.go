// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCenter(v string) *CreateSiteDeliveryTaskResponseBody
	GetDataCenter() *string
	SetRequestId(v string) *CreateSiteDeliveryTaskResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *CreateSiteDeliveryTaskResponseBody
	GetSiteId() *int64
	SetTaskName(v string) *CreateSiteDeliveryTaskResponseBody
	GetTaskName() *string
}

type CreateSiteDeliveryTaskResponseBody struct {
	// The data center. Valid values:
	//
	// 	- cn: the Chinese mainland.
	//
	// 	- oversea: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9358E852-992D-5BC7-8BD7-975CA02773A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.[](~~2850189~~)
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The name of the delivery task.
	//
	// example:
	//
	// er-oss
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateSiteDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteDeliveryTaskResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteDeliveryTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSiteDeliveryTaskResponseBody) SetDataCenter(v string) *CreateSiteDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetRequestId(v string) *CreateSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetSiteId(v int64) *CreateSiteDeliveryTaskResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) SetTaskName(v string) *CreateSiteDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
