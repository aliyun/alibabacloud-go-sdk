// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetBillDetailFileListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202502
	BillMonth          *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	OssAccessKeyId     *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	OssBucketName      *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssRegion          *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	OssSecurityToken   *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
}

func (s GetBillDetailFileListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBillDetailFileListRequest) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListRequest) SetBillMonth(v string) *GetBillDetailFileListRequest {
	s.BillMonth = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssAccessKeyId(v string) *GetBillDetailFileListRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssAccessKeySecret(v string) *GetBillDetailFileListRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssBucketName(v string) *GetBillDetailFileListRequest {
	s.OssBucketName = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssEndpoint(v string) *GetBillDetailFileListRequest {
	s.OssEndpoint = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssRegion(v string) *GetBillDetailFileListRequest {
	s.OssRegion = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssSecurityToken(v string) *GetBillDetailFileListRequest {
	s.OssSecurityToken = &v
	return s
}

type GetBillDetailFileListResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetBillDetailFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg     *string                                  `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 210bc4b416874189683843905d9f9a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBillDetailFileListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBillDetailFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponseBody) SetCode(v string) *GetBillDetailFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetData(v []*GetBillDetailFileListResponseBodyData) *GetBillDetailFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetMessage(v string) *GetBillDetailFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetMsg(v string) *GetBillDetailFileListResponseBody {
	s.Msg = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetRequestId(v string) *GetBillDetailFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBillDetailFileListResponseBody) SetSuccess(v bool) *GetBillDetailFileListResponseBody {
	s.Success = &v
	return s
}

type GetBillDetailFileListResponseBodyData struct {
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// aps.ailyun.com/file/download?resourceId=1234&type=1
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBillDetailFileListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBillDetailFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponseBodyData) SetBillMonth(v string) *GetBillDetailFileListResponseBodyData {
	s.BillMonth = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetFileName(v string) *GetBillDetailFileListResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetFileUrl(v string) *GetBillDetailFileListResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetStatus(v string) *GetBillDetailFileListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBillDetailFileListResponseBodyData) SetType(v string) *GetBillDetailFileListResponseBodyData {
	s.Type = &v
	return s
}

type GetBillDetailFileListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBillDetailFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBillDetailFileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBillDetailFileListResponse) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListResponse) SetHeaders(v map[string]*string) *GetBillDetailFileListResponse {
	s.Headers = v
	return s
}

func (s *GetBillDetailFileListResponse) SetStatusCode(v int32) *GetBillDetailFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBillDetailFileListResponse) SetBody(v *GetBillDetailFileListResponseBody) *GetBillDetailFileListResponse {
	s.Body = v
	return s
}

type GetCommissionDetailFileListRequest struct {
	// example:
	//
	// 202501
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
}

func (s GetCommissionDetailFileListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommissionDetailFileListRequest) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListRequest) SetBillMonth(v string) *GetCommissionDetailFileListRequest {
	s.BillMonth = &v
	return s
}

type GetCommissionDetailFileListResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code    *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetCommissionDetailFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCommissionDetailFileListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBody) SetCode(v string) *GetCommissionDetailFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetData(v *GetCommissionDetailFileListResponseBodyData) *GetCommissionDetailFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetMessage(v string) *GetCommissionDetailFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetRequestId(v string) *GetCommissionDetailFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBody) SetSuccess(v bool) *GetCommissionDetailFileListResponseBody {
	s.Success = &v
	return s
}

type GetCommissionDetailFileListResponseBodyData struct {
	// example:
	//
	// 202502
	BillMonth *string                                                `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	FileList  []*GetCommissionDetailFileListResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 1112332432
	PartnerUid *string `json:"PartnerUid,omitempty" xml:"PartnerUid,omitempty"`
}

func (s GetCommissionDetailFileListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBodyData) SetBillMonth(v string) *GetCommissionDetailFileListResponseBodyData {
	s.BillMonth = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyData) SetFileList(v []*GetCommissionDetailFileListResponseBodyDataFileList) *GetCommissionDetailFileListResponseBodyData {
	s.FileList = v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyData) SetPartnerUid(v string) *GetCommissionDetailFileListResponseBodyData {
	s.PartnerUid = &v
	return s
}

type GetCommissionDetailFileListResponseBodyDataFileList struct {
	CommissionPolicyName *string `json:"CommissionPolicyName,omitempty" xml:"CommissionPolicyName,omitempty"`
	FileType             *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// aps.ailyun.com/file/download?resourceId=1234&type=1
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetCommissionDetailFileListResponseBodyDataFileList) String() string {
	return tea.Prettify(s)
}

func (s GetCommissionDetailFileListResponseBodyDataFileList) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetCommissionPolicyName(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.CommissionPolicyName = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetFileType(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.FileType = &v
	return s
}

func (s *GetCommissionDetailFileListResponseBodyDataFileList) SetFileUrl(v string) *GetCommissionDetailFileListResponseBodyDataFileList {
	s.FileUrl = &v
	return s
}

type GetCommissionDetailFileListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommissionDetailFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommissionDetailFileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommissionDetailFileListResponse) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListResponse) SetHeaders(v map[string]*string) *GetCommissionDetailFileListResponse {
	s.Headers = v
	return s
}

func (s *GetCommissionDetailFileListResponse) SetStatusCode(v int32) *GetCommissionDetailFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommissionDetailFileListResponse) SetBody(v *GetCommissionDetailFileListResponseBody) *GetCommissionDetailFileListResponse {
	s.Body = v
	return s
}

type GetCustomerOrderListRequest struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// example:
	//
	// 13595216
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// example:
	//
	// 3
	OrderStatus   *int32    `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderTypeList []*string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PayAmountAfter *float64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// example:
	//
	// 1000
	PayAmountBefore *float64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202502002231
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1234532
	RamAccountForCustomerManager *string `json:"RamAccountForCustomerManager,omitempty" xml:"RamAccountForCustomerManager,omitempty"`
}

func (s GetCustomerOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrderListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListRequest) SetCustomerAccount(v string) *GetCustomerOrderListRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetCustomerUid(v int64) *GetCustomerOrderListRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderCreateAfter(v int64) *GetCustomerOrderListRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderCreateBefore(v int64) *GetCustomerOrderListRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderId(v int64) *GetCustomerOrderListRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderPayAfter(v int64) *GetCustomerOrderListRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderPayBefore(v int64) *GetCustomerOrderListRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderStatus(v int32) *GetCustomerOrderListRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderTypeList(v []*string) *GetCustomerOrderListRequest {
	s.OrderTypeList = v
	return s
}

func (s *GetCustomerOrderListRequest) SetPageNo(v int32) *GetCustomerOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPageSize(v int32) *GetCustomerOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayAmountAfter(v float64) *GetCustomerOrderListRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayAmountBefore(v float64) *GetCustomerOrderListRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayType(v int32) *GetCustomerOrderListRequest {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProductCode(v string) *GetCustomerOrderListRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProductName(v string) *GetCustomerOrderListRequest {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProjectId(v int64) *GetCustomerOrderListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetRamAccountForCustomerManager(v string) *GetCustomerOrderListRequest {
	s.RamAccountForCustomerManager = &v
	return s
}

type GetCustomerOrderListShrinkRequest struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// example:
	//
	// 13595216
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// example:
	//
	// 3
	OrderStatus         *int32  `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderTypeListShrink *string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PayAmountAfter *float64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// example:
	//
	// 1000
	PayAmountBefore *float64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202502002231
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1234532
	RamAccountForCustomerManager *string `json:"RamAccountForCustomerManager,omitempty" xml:"RamAccountForCustomerManager,omitempty"`
}

func (s GetCustomerOrderListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrderListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListShrinkRequest) SetCustomerAccount(v string) *GetCustomerOrderListShrinkRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetCustomerUid(v int64) *GetCustomerOrderListShrinkRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderCreateAfter(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderCreateBefore(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderId(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderPayAfter(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderPayBefore(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderStatus(v int32) *GetCustomerOrderListShrinkRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderTypeListShrink(v string) *GetCustomerOrderListShrinkRequest {
	s.OrderTypeListShrink = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPageNo(v int32) *GetCustomerOrderListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPageSize(v int32) *GetCustomerOrderListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayAmountAfter(v float64) *GetCustomerOrderListShrinkRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayAmountBefore(v float64) *GetCustomerOrderListShrinkRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayType(v int32) *GetCustomerOrderListShrinkRequest {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProductCode(v string) *GetCustomerOrderListShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProductName(v string) *GetCustomerOrderListShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProjectId(v int64) *GetCustomerOrderListShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetRamAccountForCustomerManager(v string) *GetCustomerOrderListShrinkRequest {
	s.RamAccountForCustomerManager = &v
	return s
}

type GetCustomerOrderListResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetCustomerOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCustomerOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponseBody) SetAccessDeniedDetail(v string) *GetCustomerOrderListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetCode(v string) *GetCustomerOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetData(v []*GetCustomerOrderListResponseBodyData) *GetCustomerOrderListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetHttpStatusCode(v int32) *GetCustomerOrderListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetMessage(v string) *GetCustomerOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetPageNo(v int32) *GetCustomerOrderListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetPageSize(v int32) *GetCustomerOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetRequestId(v string) *GetCustomerOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetSuccess(v bool) *GetCustomerOrderListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetTotal(v int32) *GetCustomerOrderListResponseBody {
	s.Total = &v
	return s
}

type GetCustomerOrderListResponseBodyData struct {
	// example:
	//
	// 1
	AmountDiscount *float64 `json:"AmountDiscount,omitempty" xml:"AmountDiscount,omitempty"`
	// example:
	//
	// 29137
	AmountDue *float64 `json:"AmountDue,omitempty" xml:"AmountDue,omitempty"`
	// example:
	//
	// 2019-01-24 14:20:40
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// test_123
	CustomerAccount        *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	CustomerClassification *string `json:"CustomerClassification,omitempty" xml:"CustomerClassification,omitempty"`
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 0
	DeductedAmountByCoupons *float64 `json:"DeductedAmountByCoupons,omitempty" xml:"DeductedAmountByCoupons,omitempty"`
	// example:
	//
	// 29137
	DiscountedPrice *float64 `json:"DiscountedPrice,omitempty" xml:"DiscountedPrice,omitempty"`
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 2019-01-24 14:20:40
	PaidAt *string `json:"PaidAt,omitempty" xml:"PaidAt,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 29137
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// slb
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// slb
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202502230421
	ProjectId                     *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RamAccountForCustomerManagers []*string `json:"RamAccountForCustomerManagers,omitempty" xml:"RamAccountForCustomerManagers,omitempty" type:"Repeated"`
}

func (s GetCustomerOrderListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponseBodyData) SetAmountDiscount(v float64) *GetCustomerOrderListResponseBodyData {
	s.AmountDiscount = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetAmountDue(v float64) *GetCustomerOrderListResponseBodyData {
	s.AmountDue = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCreatedAt(v string) *GetCustomerOrderListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerAccount(v string) *GetCustomerOrderListResponseBodyData {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerClassification(v string) *GetCustomerOrderListResponseBodyData {
	s.CustomerClassification = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerUid(v int64) *GetCustomerOrderListResponseBodyData {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetDeductedAmountByCoupons(v float64) *GetCustomerOrderListResponseBodyData {
	s.DeductedAmountByCoupons = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetDiscountedPrice(v float64) *GetCustomerOrderListResponseBodyData {
	s.DiscountedPrice = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderId(v int64) *GetCustomerOrderListResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderStatus(v int32) *GetCustomerOrderListResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderType(v string) *GetCustomerOrderListResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPaidAt(v string) *GetCustomerOrderListResponseBodyData {
	s.PaidAt = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPayType(v int32) *GetCustomerOrderListResponseBodyData {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPrice(v float64) *GetCustomerOrderListResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProductCode(v string) *GetCustomerOrderListResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProductName(v string) *GetCustomerOrderListResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProjectId(v int64) *GetCustomerOrderListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetRamAccountForCustomerManagers(v []*string) *GetCustomerOrderListResponseBodyData {
	s.RamAccountForCustomerManagers = v
	return s
}

type GetCustomerOrderListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrderListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponse) SetHeaders(v map[string]*string) *GetCustomerOrderListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerOrderListResponse) SetStatusCode(v int32) *GetCustomerOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerOrderListResponse) SetBody(v *GetCustomerOrderListResponseBody) *GetCustomerOrderListResponse {
	s.Body = v
	return s
}

type GetRenewalRateListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025Q4
	FiscalYearAndQuarter *string `json:"FiscalYearAndQuarter,omitempty" xml:"FiscalYearAndQuarter,omitempty"`
}

func (s GetRenewalRateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRenewalRateListRequest) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListRequest) SetFiscalYearAndQuarter(v string) *GetRenewalRateListRequest {
	s.FiscalYearAndQuarter = &v
	return s
}

type GetRenewalRateListResponseBody struct {
	// example:
	//
	// 200
	Code    *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetRenewalRateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRenewalRateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRenewalRateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponseBody) SetCode(v string) *GetRenewalRateListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetData(v []*GetRenewalRateListResponseBodyData) *GetRenewalRateListResponseBody {
	s.Data = v
	return s
}

func (s *GetRenewalRateListResponseBody) SetMessage(v string) *GetRenewalRateListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetRequestId(v string) *GetRenewalRateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetSuccess(v bool) *GetRenewalRateListResponseBody {
	s.Success = &v
	return s
}

type GetRenewalRateListResponseBodyData struct {
	// example:
	//
	// 100
	CustomerAdjustedRenewalAmountDue *float64 `json:"CustomerAdjustedRenewalAmountDue,omitempty" xml:"CustomerAdjustedRenewalAmountDue,omitempty"`
	// example:
	//
	// 100
	CustomerOtherBillAmount *float64 `json:"CustomerOtherBillAmount,omitempty" xml:"CustomerOtherBillAmount,omitempty"`
	// example:
	//
	// 100
	FinalCustomerRenewalAmountDue *float64 `json:"FinalCustomerRenewalAmountDue,omitempty" xml:"FinalCustomerRenewalAmountDue,omitempty"`
	// example:
	//
	// 0.9
	FinalCustomerRenewalRate *float64 `json:"FinalCustomerRenewalRate,omitempty" xml:"FinalCustomerRenewalRate,omitempty"`
	// example:
	//
	// 100
	FinalCustomerRenewedAmount *float64 `json:"FinalCustomerRenewedAmount,omitempty" xml:"FinalCustomerRenewedAmount,omitempty"`
	// example:
	//
	// 100
	FinalOtherBillAmount *float64 `json:"FinalOtherBillAmount,omitempty" xml:"FinalOtherBillAmount,omitempty"`
	// example:
	//
	// 100
	FinalRenewalAmountDue *float64 `json:"FinalRenewalAmountDue,omitempty" xml:"FinalRenewalAmountDue,omitempty"`
	// example:
	//
	// 0.9
	FinalRenewalRate *float64 `json:"FinalRenewalRate,omitempty" xml:"FinalRenewalRate,omitempty"`
	// example:
	//
	// 100
	FinalRenewedAmount *float64 `json:"FinalRenewedAmount,omitempty" xml:"FinalRenewedAmount,omitempty"`
	// example:
	//
	// 100
	FinalSubPartnerRenewalAmountDue *float64 `json:"FinalSubPartnerRenewalAmountDue,omitempty" xml:"FinalSubPartnerRenewalAmountDue,omitempty"`
	// example:
	//
	// 0.85
	FinalSubPartnerRenewalRate *float64 `json:"FinalSubPartnerRenewalRate,omitempty" xml:"FinalSubPartnerRenewalRate,omitempty"`
	// example:
	//
	// 100
	FinalSubPartnerRenewedAmount *float64 `json:"FinalSubPartnerRenewedAmount,omitempty" xml:"FinalSubPartnerRenewedAmount,omitempty"`
	// example:
	//
	// 2025Q4
	FiscalYearAndQuarter *string `json:"FiscalYearAndQuarter,omitempty" xml:"FiscalYearAndQuarter,omitempty"`
	// example:
	//
	// P123423453
	MasterPid     *string `json:"MasterPid,omitempty" xml:"MasterPid,omitempty"`
	MasterPidName *string `json:"MasterPidName,omitempty" xml:"MasterPidName,omitempty"`
	// example:
	//
	// 0.7
	SpecialCustomerRenewRatio *float64 `json:"SpecialCustomerRenewRatio,omitempty" xml:"SpecialCustomerRenewRatio,omitempty"`
	// example:
	//
	// 100
	SpecialCustomerRenewalAmountDue *float64 `json:"SpecialCustomerRenewalAmountDue,omitempty" xml:"SpecialCustomerRenewalAmountDue,omitempty"`
	// example:
	//
	// 100
	SpecialCustomerRenewedAmount *float64 `json:"SpecialCustomerRenewedAmount,omitempty" xml:"SpecialCustomerRenewedAmount,omitempty"`
	// example:
	//
	// 0.7
	SpecialFinalRenewRatio *float64 `json:"SpecialFinalRenewRatio,omitempty" xml:"SpecialFinalRenewRatio,omitempty"`
	// example:
	//
	// 100
	SpecialFinalRenewalAmountDue *float64 `json:"SpecialFinalRenewalAmountDue,omitempty" xml:"SpecialFinalRenewalAmountDue,omitempty"`
	// example:
	//
	// 100
	SpecialFinalRenewedAmount *float64 `json:"SpecialFinalRenewedAmount,omitempty" xml:"SpecialFinalRenewedAmount,omitempty"`
	// example:
	//
	// 0.8
	SpecialSubPartnerRenewRatio *float64 `json:"SpecialSubPartnerRenewRatio,omitempty" xml:"SpecialSubPartnerRenewRatio,omitempty"`
	// example:
	//
	// 100
	SpecialSubPartnerRenewalAmountDue *float64 `json:"SpecialSubPartnerRenewalAmountDue,omitempty" xml:"SpecialSubPartnerRenewalAmountDue,omitempty"`
	// example:
	//
	// 100
	SpecialSubPartnerRenewedAmount *float64 `json:"SpecialSubPartnerRenewedAmount,omitempty" xml:"SpecialSubPartnerRenewedAmount,omitempty"`
	// example:
	//
	// 100
	SubPartnerAdjustedRenewalAmountDue *float64 `json:"SubPartnerAdjustedRenewalAmountDue,omitempty" xml:"SubPartnerAdjustedRenewalAmountDue,omitempty"`
	// example:
	//
	// 100
	SubPartnerOtherBillAmount *float64 `json:"SubPartnerOtherBillAmount,omitempty" xml:"SubPartnerOtherBillAmount,omitempty"`
}

func (s GetRenewalRateListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRenewalRateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponseBodyData) SetCustomerAdjustedRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.CustomerAdjustedRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetCustomerOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.CustomerOtherBillAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalOtherBillAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFiscalYearAndQuarter(v string) *GetRenewalRateListResponseBodyData {
	s.FiscalYearAndQuarter = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetMasterPid(v string) *GetRenewalRateListResponseBodyData {
	s.MasterPid = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetMasterPidName(v string) *GetRenewalRateListResponseBodyData {
	s.MasterPidName = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSubPartnerAdjustedRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SubPartnerAdjustedRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSubPartnerOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SubPartnerOtherBillAmount = &v
	return s
}

type GetRenewalRateListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRenewalRateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRenewalRateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRenewalRateListResponse) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponse) SetHeaders(v map[string]*string) *GetRenewalRateListResponse {
	s.Headers = v
	return s
}

func (s *GetRenewalRateListResponse) SetStatusCode(v int32) *GetRenewalRateListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRenewalRateListResponse) SetBody(v *GetRenewalRateListResponseBody) *GetRenewalRateListResponse {
	s.Body = v
	return s
}

type GetSubPartnerListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubPartnerCompanyName *string `json:"SubPartnerCompanyName,omitempty" xml:"SubPartnerCompanyName,omitempty"`
	// example:
	//
	// 2323431211
	SubPartnerPid *string `json:"SubPartnerPid,omitempty" xml:"SubPartnerPid,omitempty"`
}

func (s GetSubPartnerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerListRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListRequest) SetPageNo(v int32) *GetSubPartnerListRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerListRequest) SetPageSize(v int32) *GetSubPartnerListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerListRequest) SetSubPartnerCompanyName(v string) *GetSubPartnerListRequest {
	s.SubPartnerCompanyName = &v
	return s
}

func (s *GetSubPartnerListRequest) SetSubPartnerPid(v string) *GetSubPartnerListRequest {
	s.SubPartnerPid = &v
	return s
}

type GetSubPartnerListResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubPartnerList []*GetSubPartnerListResponseBodySubPartnerList `json:"SubPartnerList,omitempty" xml:"SubPartnerList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSubPartnerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponseBody) SetMessage(v string) *GetSubPartnerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetPageNo(v string) *GetSubPartnerListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetPageSize(v string) *GetSubPartnerListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetRequestId(v string) *GetSubPartnerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetSubPartnerList(v []*GetSubPartnerListResponseBodySubPartnerList) *GetSubPartnerListResponseBody {
	s.SubPartnerList = v
	return s
}

func (s *GetSubPartnerListResponseBody) SetSuccess(v bool) *GetSubPartnerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetTotal(v int32) *GetSubPartnerListResponseBody {
	s.Total = &v
	return s
}

type GetSubPartnerListResponseBodySubPartnerList struct {
	Address             *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AgreementStatus     *string `json:"AgreementStatus,omitempty" xml:"AgreementStatus,omitempty"`
	AgreementStatusDesc *string `json:"AgreementStatusDesc,omitempty" xml:"AgreementStatusDesc,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	CompanyName         *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	Contact             *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	District            *string `json:"District,omitempty" xml:"District,omitempty"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	MasterAccount       *string `json:"MasterAccount,omitempty" xml:"MasterAccount,omitempty"`
	MasterUid           *string `json:"MasterUid,omitempty" xml:"MasterUid,omitempty"`
	Pid                 *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetSubPartnerListResponseBodySubPartnerList) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerListResponseBodySubPartnerList) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAddress(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Address = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAgreementStatus(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.AgreementStatus = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAgreementStatusDesc(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.AgreementStatusDesc = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetCity(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.City = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetCompanyName(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.CompanyName = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetContact(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Contact = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetDistrict(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.District = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetJoinTime(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.JoinTime = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetMasterAccount(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.MasterAccount = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetMasterUid(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.MasterUid = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetPid(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Pid = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetProvince(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Province = &v
	return s
}

type GetSubPartnerListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubPartnerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubPartnerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerListResponse) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponse) SetHeaders(v map[string]*string) *GetSubPartnerListResponse {
	s.Headers = v
	return s
}

func (s *GetSubPartnerListResponse) SetStatusCode(v int32) *GetSubPartnerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubPartnerListResponse) SetBody(v *GetSubPartnerListResponseBody) *GetSubPartnerListResponse {
	s.Body = v
	return s
}

type GetSubPartnerOrderListRequest struct {
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// example:
	//
	// 3
	OrderStatus   *int64    `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderTypeList []*string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PayAmountAfter *int64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// example:
	//
	// 100
	PayAmountBefore *int64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// example:
	//
	// 1
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202501101023
	ProjectId      *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// example:
	//
	// 123432311
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerOrderListRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListRequest) SetOrderCreateAfter(v int64) *GetSubPartnerOrderListRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderCreateBefore(v int64) *GetSubPartnerOrderListRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderId(v int64) *GetSubPartnerOrderListRequest {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderPayAfter(v int64) *GetSubPartnerOrderListRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderPayBefore(v int64) *GetSubPartnerOrderListRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderStatus(v int64) *GetSubPartnerOrderListRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderTypeList(v []*string) *GetSubPartnerOrderListRequest {
	s.OrderTypeList = v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPageNo(v int32) *GetSubPartnerOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPageSize(v int32) *GetSubPartnerOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayAmountAfter(v int64) *GetSubPartnerOrderListRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayAmountBefore(v int64) *GetSubPartnerOrderListRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayType(v int64) *GetSubPartnerOrderListRequest {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProductCode(v string) *GetSubPartnerOrderListRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProductName(v string) *GetSubPartnerOrderListRequest {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProjectId(v int64) *GetSubPartnerOrderListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetSubPartnerName(v string) *GetSubPartnerOrderListRequest {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetSubPartnerUid(v int64) *GetSubPartnerOrderListRequest {
	s.SubPartnerUid = &v
	return s
}

type GetSubPartnerOrderListShrinkRequest struct {
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// example:
	//
	// 3
	OrderStatus         *int64  `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderTypeListShrink *string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PayAmountAfter *int64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// example:
	//
	// 100
	PayAmountBefore *int64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// example:
	//
	// 1
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202501101023
	ProjectId      *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// example:
	//
	// 123432311
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerOrderListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderCreateAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderCreateBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderId(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderPayAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderPayBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderStatus(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderTypeListShrink(v string) *GetSubPartnerOrderListShrinkRequest {
	s.OrderTypeListShrink = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPageNo(v int32) *GetSubPartnerOrderListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPageSize(v int32) *GetSubPartnerOrderListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayAmountAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayAmountBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayType(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProductCode(v string) *GetSubPartnerOrderListShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProductName(v string) *GetSubPartnerOrderListShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProjectId(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetSubPartnerName(v string) *GetSubPartnerOrderListShrinkRequest {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetSubPartnerUid(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.SubPartnerUid = &v
	return s
}

type GetSubPartnerOrderListResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetSubPartnerOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSubPartnerOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponseBody) SetCode(v string) *GetSubPartnerOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetData(v []*GetSubPartnerOrderListResponseBodyData) *GetSubPartnerOrderListResponseBody {
	s.Data = v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetMessage(v string) *GetSubPartnerOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetPageNo(v int32) *GetSubPartnerOrderListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetPageSize(v int32) *GetSubPartnerOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetRequestId(v string) *GetSubPartnerOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetSuccess(v bool) *GetSubPartnerOrderListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetTotal(v int32) *GetSubPartnerOrderListResponseBody {
	s.Total = &v
	return s
}

type GetSubPartnerOrderListResponseBodyData struct {
	// example:
	//
	// 0.9
	AmountDiscount *float64 `json:"AmountDiscount,omitempty" xml:"AmountDiscount,omitempty"`
	// example:
	//
	// 3750
	AmountDue *float64 `json:"AmountDue,omitempty" xml:"AmountDue,omitempty"`
	// example:
	//
	// 2024-07-07 07:52:22
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 0
	DeductedAmountByCoupons *float64 `json:"DeductedAmountByCoupons,omitempty" xml:"DeductedAmountByCoupons,omitempty"`
	// example:
	//
	// 3750
	DiscountedPrice *float64 `json:"DiscountedPrice,omitempty" xml:"DiscountedPrice,omitempty"`
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 2024-07-07 07:52:22
	PaidAt *string `json:"PaidAt,omitempty" xml:"PaidAt,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 3750
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// 202502233443
	ProjectId      *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// example:
	//
	// 1123132
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponseBodyData) SetAmountDiscount(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.AmountDiscount = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetAmountDue(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.AmountDue = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetCreatedAt(v string) *GetSubPartnerOrderListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetDeductedAmountByCoupons(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.DeductedAmountByCoupons = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetDiscountedPrice(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.DiscountedPrice = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderId(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderStatus(v int32) *GetSubPartnerOrderListResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderType(v string) *GetSubPartnerOrderListResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPaidAt(v string) *GetSubPartnerOrderListResponseBodyData {
	s.PaidAt = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPayType(v int32) *GetSubPartnerOrderListResponseBodyData {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPrice(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProductCode(v string) *GetSubPartnerOrderListResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProductName(v string) *GetSubPartnerOrderListResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProjectId(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetSubPartnerName(v string) *GetSubPartnerOrderListResponseBodyData {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetSubPartnerUid(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.SubPartnerUid = &v
	return s
}

type GetSubPartnerOrderListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubPartnerOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubPartnerOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubPartnerOrderListResponse) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponse) SetHeaders(v map[string]*string) *GetSubPartnerOrderListResponse {
	s.Headers = v
	return s
}

func (s *GetSubPartnerOrderListResponse) SetStatusCode(v int32) *GetSubPartnerOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubPartnerOrderListResponse) SetBody(v *GetSubPartnerOrderListResponseBody) *GetSubPartnerOrderListResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("agency.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("agency.aliyuncs.com"),
		"ap-south-1":                  tea.String("agency.aliyuncs.com"),
		"ap-southeast-2":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-3":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-5":              tea.String("agency.aliyuncs.com"),
		"cn-beijing":                  tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("agency.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("agency.aliyuncs.com"),
		"cn-chengdu":                  tea.String("agency.aliyuncs.com"),
		"cn-edge-1":                   tea.String("agency.aliyuncs.com"),
		"cn-fujian":                   tea.String("agency.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("agency.aliyuncs.com"),
		"cn-hongkong":                 tea.String("agency.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("agency.aliyuncs.com"),
		"cn-huhehaote":                tea.String("agency.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("agency.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-qingdao":                  tea.String("agency.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai":                 tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("agency.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("agency.aliyuncs.com"),
		"cn-wuhan":                    tea.String("agency.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("agency.aliyuncs.com"),
		"cn-yushanfang":               tea.String("agency.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("agency.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("agency.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("agency.aliyuncs.com"),
		"eu-central-1":                tea.String("agency.aliyuncs.com"),
		"eu-west-1":                   tea.String("agency.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("agency.aliyuncs.com"),
		"me-east-1":                   tea.String("agency.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("agency.aliyuncs.com"),
		"us-east-1":                   tea.String("agency.aliyuncs.com"),
		"us-west-1":                   tea.String("agency.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("agency"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账单导出文件
//
// @param request - GetBillDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBillDetailFileListResponse
func (client *Client) GetBillDetailFileListWithOptions(request *GetBillDetailFileListRequest, runtime *util.RuntimeOptions) (_result *GetBillDetailFileListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillMonth)) {
		query["BillMonth"] = request.BillMonth
	}

	if !tea.BoolValue(util.IsUnset(request.OssAccessKeyId)) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.OssAccessKeySecret)) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OssRegion)) {
		query["OssRegion"] = request.OssRegion
	}

	if !tea.BoolValue(util.IsUnset(request.OssSecurityToken)) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBillDetailFileList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBillDetailFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账单导出文件
//
// @param request - GetBillDetailFileListRequest
//
// @return GetBillDetailFileListResponse
func (client *Client) GetBillDetailFileList(request *GetBillDetailFileListRequest) (_result *GetBillDetailFileListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBillDetailFileListResponse{}
	_body, _err := client.GetBillDetailFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询伙伴佣金明细
//
// @param request - GetCommissionDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommissionDetailFileListResponse
func (client *Client) GetCommissionDetailFileListWithOptions(request *GetCommissionDetailFileListRequest, runtime *util.RuntimeOptions) (_result *GetCommissionDetailFileListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillMonth)) {
		query["BillMonth"] = request.BillMonth
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCommissionDetailFileList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommissionDetailFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询伙伴佣金明细
//
// @param request - GetCommissionDetailFileListRequest
//
// @return GetCommissionDetailFileListResponse
func (client *Client) GetCommissionDetailFileList(request *GetCommissionDetailFileListRequest) (_result *GetCommissionDetailFileListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommissionDetailFileListResponse{}
	_body, _err := client.GetCommissionDetailFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询伙伴拓客订单
//
// @param tmpReq - GetCustomerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrderListResponse
func (client *Client) GetCustomerOrderListWithOptions(tmpReq *GetCustomerOrderListRequest, runtime *util.RuntimeOptions) (_result *GetCustomerOrderListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetCustomerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderTypeList)) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, tea.String("OrderTypeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerAccount)) {
		query["CustomerAccount"] = request.CustomerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerUid)) {
		query["CustomerUid"] = request.CustomerUid
	}

	if !tea.BoolValue(util.IsUnset(request.OrderCreateAfter)) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.OrderCreateBefore)) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderPayAfter)) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !tea.BoolValue(util.IsUnset(request.OrderPayBefore)) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStatus)) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OrderTypeListShrink)) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PayAmountAfter)) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !tea.BoolValue(util.IsUnset(request.PayAmountBefore)) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RamAccountForCustomerManager)) {
		query["RamAccountForCustomerManager"] = request.RamAccountForCustomerManager
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomerOrderList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomerOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询伙伴拓客订单
//
// @param request - GetCustomerOrderListRequest
//
// @return GetCustomerOrderListResponse
func (client *Client) GetCustomerOrderList(request *GetCustomerOrderListRequest) (_result *GetCustomerOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerOrderListResponse{}
	_body, _err := client.GetCustomerOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询伙伴续费率
//
// @param request - GetRenewalRateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenewalRateListResponse
func (client *Client) GetRenewalRateListWithOptions(request *GetRenewalRateListRequest, runtime *util.RuntimeOptions) (_result *GetRenewalRateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FiscalYearAndQuarter)) {
		query["FiscalYearAndQuarter"] = request.FiscalYearAndQuarter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRenewalRateList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRenewalRateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询伙伴续费率
//
// @param request - GetRenewalRateListRequest
//
// @return GetRenewalRateListResponse
func (client *Client) GetRenewalRateList(request *GetRenewalRateListRequest) (_result *GetRenewalRateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRenewalRateListResponse{}
	_body, _err := client.GetRenewalRateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询二级分销商列表
//
// @param request - GetSubPartnerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerListResponse
func (client *Client) GetSubPartnerListWithOptions(request *GetSubPartnerListRequest, runtime *util.RuntimeOptions) (_result *GetSubPartnerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubPartnerCompanyName)) {
		query["SubPartnerCompanyName"] = request.SubPartnerCompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.SubPartnerPid)) {
		query["SubPartnerPid"] = request.SubPartnerPid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubPartnerList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubPartnerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询二级分销商列表
//
// @param request - GetSubPartnerListRequest
//
// @return GetSubPartnerListResponse
func (client *Client) GetSubPartnerList(request *GetSubPartnerListRequest) (_result *GetSubPartnerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubPartnerListResponse{}
	_body, _err := client.GetSubPartnerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询拓渠订单
//
// @param tmpReq - GetSubPartnerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerOrderListResponse
func (client *Client) GetSubPartnerOrderListWithOptions(tmpReq *GetSubPartnerOrderListRequest, runtime *util.RuntimeOptions) (_result *GetSubPartnerOrderListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSubPartnerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderTypeList)) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, tea.String("OrderTypeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderCreateAfter)) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.OrderCreateBefore)) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderPayAfter)) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !tea.BoolValue(util.IsUnset(request.OrderPayBefore)) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStatus)) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OrderTypeListShrink)) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PayAmountAfter)) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !tea.BoolValue(util.IsUnset(request.PayAmountBefore)) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SubPartnerName)) {
		query["SubPartnerName"] = request.SubPartnerName
	}

	if !tea.BoolValue(util.IsUnset(request.SubPartnerUid)) {
		query["SubPartnerUid"] = request.SubPartnerUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubPartnerOrderList"),
		Version:     tea.String("2025-02-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubPartnerOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询拓渠订单
//
// @param request - GetSubPartnerOrderListRequest
//
// @return GetSubPartnerOrderListResponse
func (client *Client) GetSubPartnerOrderList(request *GetSubPartnerOrderListRequest) (_result *GetSubPartnerOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubPartnerOrderListResponse{}
	_body, _err := client.GetSubPartnerOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
