// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceOverviewResponseBody
	GetCode() *string
	SetData(v *GetResourceOverviewResponseBodyData) *GetResourceOverviewResponseBody
	GetData() *GetResourceOverviewResponseBodyData
	SetMessage(v string) *GetResourceOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceOverviewResponseBody
	GetRequestId() *string
}

type GetResourceOverviewResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Resource information.
	Data *GetResourceOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DD19A442-93C5-5C97-AFA0-B9C57EBD781B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetResourceOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceOverviewResponseBody) GetData() *GetResourceOverviewResponseBodyData {
	return s.Data
}

func (s *GetResourceOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceOverviewResponseBody) SetCode(v string) *GetResourceOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetData(v *GetResourceOverviewResponseBodyData) *GetResourceOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceOverviewResponseBody) SetMessage(v string) *GetResourceOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetRequestId(v string) *GetResourceOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceOverviewResponseBodyData struct {
	// API information.
	Api *GetResourceOverviewResponseBodyDataApi `json:"api,omitempty" xml:"api,omitempty" type:"Struct"`
	// Gateway information.
	Gateway *GetResourceOverviewResponseBodyDataGateway `json:"gateway,omitempty" xml:"gateway,omitempty" type:"Struct"`
}

func (s GetResourceOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyData) GetApi() *GetResourceOverviewResponseBodyDataApi {
	return s.Api
}

func (s *GetResourceOverviewResponseBodyData) GetGateway() *GetResourceOverviewResponseBodyDataGateway {
	return s.Gateway
}

func (s *GetResourceOverviewResponseBodyData) SetApi(v *GetResourceOverviewResponseBodyDataApi) *GetResourceOverviewResponseBodyData {
	s.Api = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) SetGateway(v *GetResourceOverviewResponseBodyDataGateway) *GetResourceOverviewResponseBodyData {
	s.Gateway = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetResourceOverviewResponseBodyDataApi struct {
	// Number of published APIs.
	//
	// example:
	//
	// 1
	PublishedCount *int64 `json:"publishedCount,omitempty" xml:"publishedCount,omitempty"`
	// Number of APIs.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataApi) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataApi) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataApi) GetPublishedCount() *int64 {
	return s.PublishedCount
}

func (s *GetResourceOverviewResponseBodyDataApi) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetResourceOverviewResponseBodyDataApi) SetPublishedCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.PublishedCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataApi) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.TotalCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataApi) Validate() error {
	return dara.Validate(s)
}

type GetResourceOverviewResponseBodyDataGateway struct {
	// Number of running gateways.
	//
	// example:
	//
	// 1
	RunningCount *int64 `json:"runningCount,omitempty" xml:"runningCount,omitempty"`
	// Number of gateway instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataGateway) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataGateway) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataGateway) GetRunningCount() *int64 {
	return s.RunningCount
}

func (s *GetResourceOverviewResponseBodyDataGateway) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetRunningCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.RunningCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.TotalCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataGateway) Validate() error {
	return dara.Validate(s)
}
