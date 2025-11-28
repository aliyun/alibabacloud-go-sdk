// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelServices(v []*ListModelServicesResponseBodyModelServices) *ListModelServicesResponseBody
	GetModelServices() []*ListModelServicesResponseBodyModelServices
	SetPageNumber(v int32) *ListModelServicesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListModelServicesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListModelServicesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListModelServicesResponseBody
	GetTotalRecordCount() *int32
}

type ListModelServicesResponseBody struct {
	// Model services.
	ModelServices []*ListModelServicesResponseBodyModelServices `json:"ModelServices,omitempty" xml:"ModelServices,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListModelServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelServicesResponseBody) GetModelServices() []*ListModelServicesResponseBodyModelServices {
	return s.ModelServices
}

func (s *ListModelServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelServicesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListModelServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelServicesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListModelServicesResponseBody) SetModelServices(v []*ListModelServicesResponseBodyModelServices) *ListModelServicesResponseBody {
	s.ModelServices = v
	return s
}

func (s *ListModelServicesResponseBody) SetPageNumber(v int32) *ListModelServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelServicesResponseBody) SetPageRecordCount(v int32) *ListModelServicesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListModelServicesResponseBody) SetRequestId(v string) *ListModelServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelServicesResponseBody) SetTotalRecordCount(v int32) *ListModelServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListModelServicesResponseBody) Validate() error {
	if s.ModelServices != nil {
		for _, item := range s.ModelServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelServicesResponseBodyModelServices struct {
	// A list of AI nodes for model deployment.
	AiNodes []*string `json:"AiNodes,omitempty" xml:"AiNodes,omitempty" type:"Repeated"`
	// The API key.
	//
	// example:
	//
	// mI3F7B18vgfqJjUtWmgITw==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-09-28T02:18:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model name.
	//
	// example:
	//
	// sambert-zhiying-v1
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Model service parameters (not available).
	ModelParams map[string]*string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// Model service ID.
	//
	// example:
	//
	// mx-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// Private Endpoint.
	//
	// example:
	//
	// http://ms-xxxxxxx.xxxx.rds.aliyuncs.com
	PrivateConnUrl *string `json:"PrivateConnUrl,omitempty" xml:"PrivateConnUrl,omitempty"`
	// Public endpoint.
	//
	// example:
	//
	// http://ms-xxxxxxx-o.xxxx.rds.aliyuncs.com
	PublicConnUrl *string `json:"PublicConnUrl,omitempty" xml:"PublicConnUrl,omitempty"`
	// The IP addresses listed in the whitelist. Up to 1,000 IP addresses are contained in a whitelist and separated by commas (,). The IP addresses must use one of the following formats:
	//
	// 	- 0.0.0.0/0
	//
	// 	- 10.23.12.24(IP)
	//
	// 	- 10.23.12.24/24 (This is a CIDR block. The value`/24`indicates the network prefix length, which must be an integer and in the range of `[1,32]`.)
	//
	// example:
	//
	// 0.0.0.0/0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListModelServicesResponseBodyModelServices) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesResponseBodyModelServices) GoString() string {
	return s.String()
}

func (s *ListModelServicesResponseBodyModelServices) GetAiNodes() []*string {
	return s.AiNodes
}

func (s *ListModelServicesResponseBodyModelServices) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListModelServicesResponseBodyModelServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModelServicesResponseBodyModelServices) GetDescription() *string {
	return s.Description
}

func (s *ListModelServicesResponseBodyModelServices) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelServicesResponseBodyModelServices) GetModelParams() map[string]*string {
	return s.ModelParams
}

func (s *ListModelServicesResponseBodyModelServices) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ListModelServicesResponseBodyModelServices) GetPrivateConnUrl() *string {
	return s.PrivateConnUrl
}

func (s *ListModelServicesResponseBodyModelServices) GetPublicConnUrl() *string {
	return s.PublicConnUrl
}

func (s *ListModelServicesResponseBodyModelServices) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ListModelServicesResponseBodyModelServices) GetStatus() *string {
	return s.Status
}

func (s *ListModelServicesResponseBodyModelServices) SetAiNodes(v []*string) *ListModelServicesResponseBodyModelServices {
	s.AiNodes = v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetApiKey(v string) *ListModelServicesResponseBodyModelServices {
	s.ApiKey = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetCreateTime(v string) *ListModelServicesResponseBodyModelServices {
	s.CreateTime = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetDescription(v string) *ListModelServicesResponseBodyModelServices {
	s.Description = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetModelName(v string) *ListModelServicesResponseBodyModelServices {
	s.ModelName = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetModelParams(v map[string]*string) *ListModelServicesResponseBodyModelServices {
	s.ModelParams = v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetModelServiceId(v string) *ListModelServicesResponseBodyModelServices {
	s.ModelServiceId = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetPrivateConnUrl(v string) *ListModelServicesResponseBodyModelServices {
	s.PrivateConnUrl = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetPublicConnUrl(v string) *ListModelServicesResponseBodyModelServices {
	s.PublicConnUrl = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetSecurityIPList(v string) *ListModelServicesResponseBodyModelServices {
	s.SecurityIPList = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) SetStatus(v string) *ListModelServicesResponseBodyModelServices {
	s.Status = &v
	return s
}

func (s *ListModelServicesResponseBodyModelServices) Validate() error {
	return dara.Validate(s)
}
