// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCloudMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetMessageCloudMonitorConfigResponseBody
	GetRequestId() *string
}

type SetMessageCloudMonitorConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMessageCloudMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCloudMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetMessageCloudMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetMessageCloudMonitorConfigResponseBody) SetRequestId(v string) *SetMessageCloudMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMessageCloudMonitorConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
