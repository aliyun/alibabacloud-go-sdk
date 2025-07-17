// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteTraceAppShrinkRequest
	GetAppId() *string
	SetDeleteReasonShrink(v string) *DeleteTraceAppShrinkRequest
	GetDeleteReasonShrink() *string
	SetPid(v string) *DeleteTraceAppShrinkRequest
	GetPid() *string
	SetRegionId(v string) *DeleteTraceAppShrinkRequest
	GetRegionId() *string
	SetType(v string) *DeleteTraceAppShrinkRequest
	GetType() *string
}

type DeleteTraceAppShrinkRequest struct {
	// The ID of the application that you want to delete. You can call the SearchTraceAppByName operation to query the application ID. For more information, see [SearchTraceAppByName](https://help.aliyun.com/document_detail/130676.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5406**
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The reason(s) to delete application.
	DeleteReasonShrink *string `json:"DeleteReason,omitempty" xml:"DeleteReason,omitempty"`
	// The PID of the application. For more information about how to query the PID, see [QueryMetricByPage](https://www.alibabacloud.com/help/zh/doc-detail/186100.htm?spm=a2cdw.13409063.0.0.7a72281f0bkTfx#title-imy-7gj-qhr).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9w0sc5gxxz@edcsd447c2f****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region in which the application is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the application that you want to delete. You can call the SearchTraceAppByName operation to query the application type. For more information, see [SearchTraceAppByName](https://help.aliyun.com/document_detail/130676.html). Valid values:
	//
	// 	- `TRACE`: Application Monitoring
	//
	// 	- `RETCODE`: frontend monitoring
	//
	// This parameter is required.
	//
	// example:
	//
	// TRACE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteTraceAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteTraceAppShrinkRequest) GetDeleteReasonShrink() *string {
	return s.DeleteReasonShrink
}

func (s *DeleteTraceAppShrinkRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteTraceAppShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTraceAppShrinkRequest) GetType() *string {
	return s.Type
}

func (s *DeleteTraceAppShrinkRequest) SetAppId(v string) *DeleteTraceAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTraceAppShrinkRequest) SetDeleteReasonShrink(v string) *DeleteTraceAppShrinkRequest {
	s.DeleteReasonShrink = &v
	return s
}

func (s *DeleteTraceAppShrinkRequest) SetPid(v string) *DeleteTraceAppShrinkRequest {
	s.Pid = &v
	return s
}

func (s *DeleteTraceAppShrinkRequest) SetRegionId(v string) *DeleteTraceAppShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppShrinkRequest) SetType(v string) *DeleteTraceAppShrinkRequest {
	s.Type = &v
	return s
}

func (s *DeleteTraceAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
