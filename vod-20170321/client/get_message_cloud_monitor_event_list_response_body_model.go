// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v string) *GetMessageCloudMonitorEventListResponseBody
	GetEventList() *string
	SetRequestId(v string) *GetMessageCloudMonitorEventListResponseBody
	GetRequestId() *string
}

type GetMessageCloudMonitorEventListResponseBody struct {
	EventList *string `json:"EventList,omitempty" xml:"EventList,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageCloudMonitorEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorEventListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorEventListResponseBody) GetEventList() *string {
	return s.EventList
}

func (s *GetMessageCloudMonitorEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageCloudMonitorEventListResponseBody) SetEventList(v string) *GetMessageCloudMonitorEventListResponseBody {
	s.EventList = &v
	return s
}

func (s *GetMessageCloudMonitorEventListResponseBody) SetRequestId(v string) *GetMessageCloudMonitorEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageCloudMonitorEventListResponseBody) Validate() error {
	return dara.Validate(s)
}
