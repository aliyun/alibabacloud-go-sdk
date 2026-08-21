// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOnlineHeatmapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceOnlineHeatmap(v [][]*int64) *GetDeviceOnlineHeatmapResponseBody
	GetDeviceOnlineHeatmap() [][]*int64
	SetRequestId(v string) *GetDeviceOnlineHeatmapResponseBody
	GetRequestId() *string
}

type GetDeviceOnlineHeatmapResponseBody struct {
	// The online time distribution.
	DeviceOnlineHeatmap [][]*int64 `json:"DeviceOnlineHeatmap,omitempty" xml:"DeviceOnlineHeatmap,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOnlineHeatmapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOnlineHeatmapResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOnlineHeatmapResponseBody) GetDeviceOnlineHeatmap() [][]*int64 {
	return s.DeviceOnlineHeatmap
}

func (s *GetDeviceOnlineHeatmapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceOnlineHeatmapResponseBody) SetDeviceOnlineHeatmap(v [][]*int64) *GetDeviceOnlineHeatmapResponseBody {
	s.DeviceOnlineHeatmap = v
	return s
}

func (s *GetDeviceOnlineHeatmapResponseBody) SetRequestId(v string) *GetDeviceOnlineHeatmapResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceOnlineHeatmapResponseBody) Validate() error {
	return dara.Validate(s)
}
