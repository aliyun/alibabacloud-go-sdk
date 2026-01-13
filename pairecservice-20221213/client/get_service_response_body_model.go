// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetServiceResponseBody
	GetDescription() *string
	SetEngineConfigId(v string) *GetServiceResponseBody
	GetEngineConfigId() *string
	SetGmtReleasedTime(v string) *GetServiceResponseBody
	GetGmtReleasedTime() *string
	SetImageAuth(v string) *GetServiceResponseBody
	GetImageAuth() *string
	SetImageName(v string) *GetServiceResponseBody
	GetImageName() *string
	SetLatestProdReleaseOrder(v *GetServiceResponseBodyLatestProdReleaseOrder) *GetServiceResponseBody
	GetLatestProdReleaseOrder() *GetServiceResponseBodyLatestProdReleaseOrder
	SetName(v string) *GetServiceResponseBody
	GetName() *string
	SetRegion(v string) *GetServiceResponseBody
	GetRegion() *string
	SetRequestId(v string) *GetServiceResponseBody
	GetRequestId() *string
	SetServiceConfig(v string) *GetServiceResponseBody
	GetServiceConfig() *string
	SetServiceResourceUri(v string) *GetServiceResponseBody
	GetServiceResourceUri() *string
}

type GetServiceResponseBody struct {
	// example:
	//
	// this is a test rec engine
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtReleasedTime *string `json:"GmtReleasedTime,omitempty" xml:"GmtReleasedTime,omitempty"`
	// example:
	//
	// ********
	ImageAuth              *string                                       `json:"ImageAuth,omitempty" xml:"ImageAuth,omitempty"`
	ImageName              *string                                       `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	LatestProdReleaseOrder *GetServiceResponseBodyLatestProdReleaseOrder `json:"LatestProdReleaseOrder,omitempty" xml:"LatestProdReleaseOrder,omitempty" type:"Struct"`
	// example:
	//
	// test_rec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// D75C43DC-3D3A-5CC8-9AAC-8C77306C433B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"Port":8080}
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// example:
	//
	// eas-resource-xxx
	ServiceResourceUri *string `json:"ServiceResourceUri,omitempty" xml:"ServiceResourceUri,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetServiceResponseBody) GetEngineConfigId() *string {
	return s.EngineConfigId
}

func (s *GetServiceResponseBody) GetGmtReleasedTime() *string {
	return s.GmtReleasedTime
}

func (s *GetServiceResponseBody) GetImageAuth() *string {
	return s.ImageAuth
}

func (s *GetServiceResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *GetServiceResponseBody) GetLatestProdReleaseOrder() *GetServiceResponseBodyLatestProdReleaseOrder {
	return s.LatestProdReleaseOrder
}

func (s *GetServiceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBody) GetRegion() *string {
	return s.Region
}

func (s *GetServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceResponseBody) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *GetServiceResponseBody) GetServiceResourceUri() *string {
	return s.ServiceResourceUri
}

func (s *GetServiceResponseBody) SetDescription(v string) *GetServiceResponseBody {
	s.Description = &v
	return s
}

func (s *GetServiceResponseBody) SetEngineConfigId(v string) *GetServiceResponseBody {
	s.EngineConfigId = &v
	return s
}

func (s *GetServiceResponseBody) SetGmtReleasedTime(v string) *GetServiceResponseBody {
	s.GmtReleasedTime = &v
	return s
}

func (s *GetServiceResponseBody) SetImageAuth(v string) *GetServiceResponseBody {
	s.ImageAuth = &v
	return s
}

func (s *GetServiceResponseBody) SetImageName(v string) *GetServiceResponseBody {
	s.ImageName = &v
	return s
}

func (s *GetServiceResponseBody) SetLatestProdReleaseOrder(v *GetServiceResponseBodyLatestProdReleaseOrder) *GetServiceResponseBody {
	s.LatestProdReleaseOrder = v
	return s
}

func (s *GetServiceResponseBody) SetName(v string) *GetServiceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBody) SetRegion(v string) *GetServiceResponseBody {
	s.Region = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceConfig(v string) *GetServiceResponseBody {
	s.ServiceConfig = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceResourceUri(v string) *GetServiceResponseBody {
	s.ServiceResourceUri = &v
	return s
}

func (s *GetServiceResponseBody) Validate() error {
	if s.LatestProdReleaseOrder != nil {
		if err := s.LatestProdReleaseOrder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceResponseBodyLatestProdReleaseOrder struct {
	// example:
	//
	// update golang version to 1.22
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2.0.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// example:
	//
	// {
	//
	// 	"Pre": {
	//
	//     "Status": "Released",
	//
	//     "GmtReleasedTime" : "2021-12-15T23:24:33.132+08:00",
	//
	//   },
	//
	//   "Prod": {
	//
	//     "Status": "Released",
	//
	//     "GmtReleasedTime" : "2021-12-15T23:24:33.132+08:00",
	//
	//   }
	//
	// }
	ReleaseInfo *string `json:"ReleaseInfo,omitempty" xml:"ReleaseInfo,omitempty"`
	// example:
	//
	// 3
	ReleaseOrderId *string `json:"ReleaseOrderId,omitempty" xml:"ReleaseOrderId,omitempty"`
	// example:
	//
	// E-xxx.xxx-@xxx.onaliyun.com
	Releaser *string `json:"Releaser,omitempty" xml:"Releaser,omitempty"`
	// example:
	//
	// update version
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetServiceResponseBodyLatestProdReleaseOrder) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyLatestProdReleaseOrder) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetContent() *string {
	return s.Content
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetReleaseInfo() *string {
	return s.ReleaseInfo
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetReleaseOrderId() *string {
	return s.ReleaseOrderId
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetReleaser() *string {
	return s.Releaser
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) GetTopic() *string {
	return s.Topic
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetContent(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.Content = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetImageVersion(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.ImageVersion = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetReleaseInfo(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.ReleaseInfo = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetReleaseOrderId(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.ReleaseOrderId = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetReleaser(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.Releaser = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) SetTopic(v string) *GetServiceResponseBodyLatestProdReleaseOrder {
	s.Topic = &v
	return s
}

func (s *GetServiceResponseBodyLatestProdReleaseOrder) Validate() error {
	return dara.Validate(s)
}
