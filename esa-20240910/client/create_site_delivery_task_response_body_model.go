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
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
