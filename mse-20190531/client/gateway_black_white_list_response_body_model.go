// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayBlackWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GatewayBlackWhiteListResponseBody
	GetCode() *int32
	SetData(v *GatewayBlackWhiteListResponseBodyData) *GatewayBlackWhiteListResponseBody
	GetData() *GatewayBlackWhiteListResponseBodyData
	SetDynamicCode(v string) *GatewayBlackWhiteListResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GatewayBlackWhiteListResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GatewayBlackWhiteListResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GatewayBlackWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GatewayBlackWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GatewayBlackWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GatewayBlackWhiteListResponseBody
	GetSuccess() *bool
}

type GatewayBlackWhiteListResponseBody struct {
	// The code that synchronizes with the Success parameter.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GatewayBlackWhiteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The placeholder in the dynamic error message. This parameter is not used and can be ignored.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. This parameter is not used and can be ignored.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58E06A0A-BD2C-47A0-99C2-B100F353****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GatewayBlackWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GatewayBlackWhiteListResponseBody) GetData() *GatewayBlackWhiteListResponseBodyData {
	return s.Data
}

func (s *GatewayBlackWhiteListResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GatewayBlackWhiteListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GatewayBlackWhiteListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GatewayBlackWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GatewayBlackWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GatewayBlackWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GatewayBlackWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GatewayBlackWhiteListResponseBody) SetCode(v int32) *GatewayBlackWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetData(v *GatewayBlackWhiteListResponseBodyData) *GatewayBlackWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetDynamicCode(v string) *GatewayBlackWhiteListResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetDynamicMessage(v string) *GatewayBlackWhiteListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetErrorCode(v string) *GatewayBlackWhiteListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetHttpStatusCode(v int32) *GatewayBlackWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetMessage(v string) *GatewayBlackWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetRequestId(v string) *GatewayBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) SetSuccess(v bool) *GatewayBlackWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GatewayBlackWhiteListResponseBodyData struct {
	// The page number of the instance list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page specified in the request.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned result set.
	Result []*GatewayBlackWhiteListResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of queried instances.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GatewayBlackWhiteListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GatewayBlackWhiteListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GatewayBlackWhiteListResponseBodyData) GetResult() []*GatewayBlackWhiteListResponseBodyDataResult {
	return s.Result
}

func (s *GatewayBlackWhiteListResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GatewayBlackWhiteListResponseBodyData) SetPageNumber(v int32) *GatewayBlackWhiteListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyData) SetPageSize(v int32) *GatewayBlackWhiteListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyData) SetResult(v []*GatewayBlackWhiteListResponseBodyDataResult) *GatewayBlackWhiteListResponseBodyData {
	s.Result = v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyData) SetTotalSize(v int32) *GatewayBlackWhiteListResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GatewayBlackWhiteListResponseBodyDataResult struct {
	// The content of the blacklist.
	//
	// example:
	//
	// 1.1.1.1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 81
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique identifier of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-08-02T02:43:40.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-08-02T02:43:40.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the blacklists and whitelists.
	//
	// example:
	//
	// 549
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether it is a whitelist.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The name of the blacklist.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The note.
	//
	// example:
	//
	// Blocked IP list
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 549
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The JSON string of the resource ID list.
	//
	// - If ResourceType is set to ROUTE, this parameter is a list of route IDs.
	//
	// - If ResourceType is set to DOMAIN, this parameter is a list of domain name IDs.
	//
	// example:
	//
	// [234]
	ResourceIdJsonList *string `json:"ResourceIdJsonList,omitempty" xml:"ResourceIdJsonList,omitempty"`
	// The resource name description.
	//
	// example:
	//
	// {}
	ResourceIdNameJson *string `json:"ResourceIdNameJson,omitempty" xml:"ResourceIdNameJson,omitempty"`
	// The effective granularity. Valid values:
	//
	// - GATEWAY: global gateway level.
	//
	// - DOMAIN: domain name level.
	//
	// - ROUTE: route level.
	//
	// example:
	//
	// GATEWAY
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the blacklists and whitelists. Valid values:
	//
	// - on: enabled.
	//
	// - off: shutdown.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the blacklists and whitelists. The value is fixed to IP blacklists and whitelists.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GatewayBlackWhiteListResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetContent() *string {
	return s.Content
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetNote() *string {
	return s.Note
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetResourceIdJsonList() *string {
	return s.ResourceIdJsonList
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetResourceIdNameJson() *string {
	return s.ResourceIdNameJson
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetStatus() *string {
	return s.Status
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetContent(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetGatewayId(v int64) *GatewayBlackWhiteListResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetGatewayUniqueId(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetGmtCreate(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetGmtModified(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetId(v int64) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetIsWhite(v bool) *GatewayBlackWhiteListResponseBodyDataResult {
	s.IsWhite = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetName(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetNote(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Note = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetResourceId(v int64) *GatewayBlackWhiteListResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetResourceIdJsonList(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.ResourceIdJsonList = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetResourceIdNameJson(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.ResourceIdNameJson = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetResourceType(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.ResourceType = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetStatus(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) SetType(v string) *GatewayBlackWhiteListResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *GatewayBlackWhiteListResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
