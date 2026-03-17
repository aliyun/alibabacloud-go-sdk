// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAGDpiAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDpiMonitorStatus(v string) *GetSmartAGDpiAttributeResponseBody
	GetDpiMonitorStatus() *string
	SetDpiStatus(v string) *GetSmartAGDpiAttributeResponseBody
	GetDpiStatus() *string
	SetLogstoreName(v string) *GetSmartAGDpiAttributeResponseBody
	GetLogstoreName() *string
	SetProjectName(v string) *GetSmartAGDpiAttributeResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetSmartAGDpiAttributeResponseBody
	GetRequestId() *string
	SetSlsRegion(v string) *GetSmartAGDpiAttributeResponseBody
	GetSlsRegion() *string
}

type GetSmartAGDpiAttributeResponseBody struct {
	// The status of the DPI-based monitoring feature. Valid values:
	//
	// 	- **Active**: enabled
	//
	// 	- **Inactive**: disabled
	//
	// example:
	//
	// Inactive
	DpiMonitorStatus *string `json:"DpiMonitorStatus,omitempty" xml:"DpiMonitorStatus,omitempty"`
	// The status of the DPI feature. Valid values:
	//
	// 	- **On**: enabled
	//
	// 	- **Off**: disabled
	//
	// example:
	//
	// On
	DpiStatus *string `json:"DpiStatus,omitempty" xml:"DpiStatus,omitempty"`
	// The name of the Log Service Logstore that is associated with the DPI feature.
	//
	// example:
	//
	// test1
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the Log Service project that is associated with the DPI feature.
	//
	// example:
	//
	// test2
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2997DC4-F1A2-418B-81FC-C8892CD31CFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region where Log Service is deployed.
	//
	// example:
	//
	// cn-shanghai
	SlsRegion *string `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
}

func (s GetSmartAGDpiAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAGDpiAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartAGDpiAttributeResponseBody) GetDpiMonitorStatus() *string {
	return s.DpiMonitorStatus
}

func (s *GetSmartAGDpiAttributeResponseBody) GetDpiStatus() *string {
	return s.DpiStatus
}

func (s *GetSmartAGDpiAttributeResponseBody) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *GetSmartAGDpiAttributeResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSmartAGDpiAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartAGDpiAttributeResponseBody) GetSlsRegion() *string {
	return s.SlsRegion
}

func (s *GetSmartAGDpiAttributeResponseBody) SetDpiMonitorStatus(v string) *GetSmartAGDpiAttributeResponseBody {
	s.DpiMonitorStatus = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) SetDpiStatus(v string) *GetSmartAGDpiAttributeResponseBody {
	s.DpiStatus = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) SetLogstoreName(v string) *GetSmartAGDpiAttributeResponseBody {
	s.LogstoreName = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) SetProjectName(v string) *GetSmartAGDpiAttributeResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) SetRequestId(v string) *GetSmartAGDpiAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) SetSlsRegion(v string) *GetSmartAGDpiAttributeResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *GetSmartAGDpiAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
