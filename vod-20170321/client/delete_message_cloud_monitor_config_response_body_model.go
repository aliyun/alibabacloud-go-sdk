// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCloudMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMessageCloudMonitorConfigResponseBody
	GetRequestId() *string
}

type DeleteMessageCloudMonitorConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMessageCloudMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCloudMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageCloudMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMessageCloudMonitorConfigResponseBody) SetRequestId(v string) *DeleteMessageCloudMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageCloudMonitorConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
