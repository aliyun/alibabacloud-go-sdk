// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataItemsModelDataValue struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Linked        *bool   `json:"Linked,omitempty" xml:"Linked,omitempty"`
	LinkModuleId  *string `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	LinkModelId   *string `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Props         *string `json:"Props,omitempty" xml:"Props,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	ModelDigest   *string `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
}

func (s DataItemsModelDataValue) String() string {
	return tea.Prettify(s)
}

func (s DataItemsModelDataValue) GoString() string {
	return s.String()
}

func (s *DataItemsModelDataValue) SetId(v string) *DataItemsModelDataValue {
	s.Id = &v
	return s
}

func (s *DataItemsModelDataValue) SetModelId(v string) *DataItemsModelDataValue {
	s.ModelId = &v
	return s
}

func (s *DataItemsModelDataValue) SetModelName(v string) *DataItemsModelDataValue {
	s.ModelName = &v
	return s
}

func (s *DataItemsModelDataValue) SetModelStatus(v string) *DataItemsModelDataValue {
	s.ModelStatus = &v
	return s
}

func (s *DataItemsModelDataValue) SetModelType(v string) *DataItemsModelDataValue {
	s.ModelType = &v
	return s
}

func (s *DataItemsModelDataValue) SetSubType(v string) *DataItemsModelDataValue {
	s.SubType = &v
	return s
}

func (s *DataItemsModelDataValue) SetModuleId(v string) *DataItemsModelDataValue {
	s.ModuleId = &v
	return s
}

func (s *DataItemsModelDataValue) SetContent(v string) *DataItemsModelDataValue {
	s.Content = &v
	return s
}

func (s *DataItemsModelDataValue) SetAppId(v string) *DataItemsModelDataValue {
	s.AppId = &v
	return s
}

func (s *DataItemsModelDataValue) SetLinked(v bool) *DataItemsModelDataValue {
	s.Linked = &v
	return s
}

func (s *DataItemsModelDataValue) SetLinkModuleId(v string) *DataItemsModelDataValue {
	s.LinkModuleId = &v
	return s
}

func (s *DataItemsModelDataValue) SetLinkModelId(v string) *DataItemsModelDataValue {
	s.LinkModelId = &v
	return s
}

func (s *DataItemsModelDataValue) SetSchemaVersion(v string) *DataItemsModelDataValue {
	s.SchemaVersion = &v
	return s
}

func (s *DataItemsModelDataValue) SetDescription(v string) *DataItemsModelDataValue {
	s.Description = &v
	return s
}

func (s *DataItemsModelDataValue) SetProps(v string) *DataItemsModelDataValue {
	s.Props = &v
	return s
}

func (s *DataItemsModelDataValue) SetVisibility(v string) *DataItemsModelDataValue {
	s.Visibility = &v
	return s
}

func (s *DataItemsModelDataValue) SetModelDigest(v string) *DataItemsModelDataValue {
	s.ModelDigest = &v
	return s
}

type DataItemsResourceDataValue struct {
	ResourceId    *string                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string                `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Description   *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	SchemaVersion *string                `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	ModuleId      *string                `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Content       map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	AppId         *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Visibility    *string                `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DataItemsResourceDataValue) String() string {
	return tea.Prettify(s)
}

func (s DataItemsResourceDataValue) GoString() string {
	return s.String()
}

func (s *DataItemsResourceDataValue) SetResourceId(v string) *DataItemsResourceDataValue {
	s.ResourceId = &v
	return s
}

func (s *DataItemsResourceDataValue) SetResourceName(v string) *DataItemsResourceDataValue {
	s.ResourceName = &v
	return s
}

func (s *DataItemsResourceDataValue) SetResourceType(v string) *DataItemsResourceDataValue {
	s.ResourceType = &v
	return s
}

func (s *DataItemsResourceDataValue) SetDescription(v string) *DataItemsResourceDataValue {
	s.Description = &v
	return s
}

func (s *DataItemsResourceDataValue) SetSchemaVersion(v string) *DataItemsResourceDataValue {
	s.SchemaVersion = &v
	return s
}

func (s *DataItemsResourceDataValue) SetModuleId(v string) *DataItemsResourceDataValue {
	s.ModuleId = &v
	return s
}

func (s *DataItemsResourceDataValue) SetContent(v map[string]interface{}) *DataItemsResourceDataValue {
	s.Content = v
	return s
}

func (s *DataItemsResourceDataValue) SetAppId(v string) *DataItemsResourceDataValue {
	s.AppId = &v
	return s
}

func (s *DataItemsResourceDataValue) SetVisibility(v string) *DataItemsResourceDataValue {
	s.Visibility = &v
	return s
}

type BatchCreateModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelDataJson *string `json:"ModelDataJson,omitempty" xml:"ModelDataJson,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s BatchCreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateModelRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateModelRequest) SetAppId(v string) *BatchCreateModelRequest {
	s.AppId = &v
	return s
}

func (s *BatchCreateModelRequest) SetModelDataJson(v string) *BatchCreateModelRequest {
	s.ModelDataJson = &v
	return s
}

func (s *BatchCreateModelRequest) SetModuleId(v string) *BatchCreateModelRequest {
	s.ModuleId = &v
	return s
}

func (s *BatchCreateModelRequest) SetSchemaVersion(v string) *BatchCreateModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *BatchCreateModelRequest) SetSource(v string) *BatchCreateModelRequest {
	s.Source = &v
	return s
}

func (s *BatchCreateModelRequest) SetSubType(v string) *BatchCreateModelRequest {
	s.SubType = &v
	return s
}

type BatchCreateModelResponseBody struct {
	Data      *BatchCreateModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateModelResponseBody) SetData(v *BatchCreateModelResponseBodyData) *BatchCreateModelResponseBody {
	s.Data = v
	return s
}

func (s *BatchCreateModelResponseBody) SetRequestId(v string) *BatchCreateModelResponseBody {
	s.RequestId = &v
	return s
}

type BatchCreateModelResponseBodyData struct {
	Items []*BatchCreateModelResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s BatchCreateModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCreateModelResponseBodyData) SetItems(v []*BatchCreateModelResponseBodyDataItems) *BatchCreateModelResponseBodyData {
	s.Items = v
	return s
}

type BatchCreateModelResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s BatchCreateModelResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateModelResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *BatchCreateModelResponseBodyDataItems) SetAppId(v string) *BatchCreateModelResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetAttributes(v []map[string]*string) *BatchCreateModelResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetContent(v map[string]*string) *BatchCreateModelResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetCreateTime(v string) *BatchCreateModelResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetDescription(v string) *BatchCreateModelResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetId(v string) *BatchCreateModelResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetLinkModelId(v string) *BatchCreateModelResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetLinkModuleId(v string) *BatchCreateModelResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetLinked(v bool) *BatchCreateModelResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModelDigest(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModelDigest = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModelId(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModelName(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModelStatus(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModelType(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModifiedTime(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetModuleId(v string) *BatchCreateModelResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetProps(v map[string]*string) *BatchCreateModelResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetRevision(v int32) *BatchCreateModelResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetSchemaVersion(v string) *BatchCreateModelResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetSubType(v string) *BatchCreateModelResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *BatchCreateModelResponseBodyDataItems) SetVisibility(v string) *BatchCreateModelResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type BatchCreateModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateModelResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateModelResponse) SetHeaders(v map[string]*string) *BatchCreateModelResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateModelResponse) SetStatusCode(v int32) *BatchCreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateModelResponse) SetBody(v *BatchCreateModelResponseBody) *BatchCreateModelResponse {
	s.Body = v
	return s
}

type BatchDeleteModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelIdList   *string `json:"ModelIdList,omitempty" xml:"ModelIdList,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s BatchDeleteModelRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteModelRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelRequest) SetAppId(v string) *BatchDeleteModelRequest {
	s.AppId = &v
	return s
}

func (s *BatchDeleteModelRequest) SetModelIdList(v string) *BatchDeleteModelRequest {
	s.ModelIdList = &v
	return s
}

func (s *BatchDeleteModelRequest) SetModuleId(v string) *BatchDeleteModelRequest {
	s.ModuleId = &v
	return s
}

func (s *BatchDeleteModelRequest) SetSchemaVersion(v string) *BatchDeleteModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *BatchDeleteModelRequest) SetSource(v string) *BatchDeleteModelRequest {
	s.Source = &v
	return s
}

type BatchDeleteModelResponseBody struct {
	Data      *BatchDeleteModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelResponseBody) SetData(v *BatchDeleteModelResponseBodyData) *BatchDeleteModelResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteModelResponseBody) SetRequestId(v string) *BatchDeleteModelResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteModelResponseBodyData struct {
	Items []*BatchDeleteModelResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s BatchDeleteModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelResponseBodyData) SetItems(v []*BatchDeleteModelResponseBodyDataItems) *BatchDeleteModelResponseBodyData {
	s.Items = v
	return s
}

type BatchDeleteModelResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s BatchDeleteModelResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteModelResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelResponseBodyDataItems) SetAppId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetAttributes(v []map[string]*string) *BatchDeleteModelResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetContent(v map[string]*string) *BatchDeleteModelResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetCreateTime(v string) *BatchDeleteModelResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetDescription(v string) *BatchDeleteModelResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetLinkModelId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetLinkModuleId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetLinked(v bool) *BatchDeleteModelResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModelId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModelName(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModelStatus(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModelType(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModifiedTime(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetModuleId(v string) *BatchDeleteModelResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetProps(v map[string]*string) *BatchDeleteModelResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetRevision(v int32) *BatchDeleteModelResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetSchemaVersion(v string) *BatchDeleteModelResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetSubType(v string) *BatchDeleteModelResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *BatchDeleteModelResponseBodyDataItems) SetVisibility(v string) *BatchDeleteModelResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type BatchDeleteModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteModelResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelResponse) SetHeaders(v map[string]*string) *BatchDeleteModelResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteModelResponse) SetStatusCode(v int32) *BatchDeleteModelResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteModelResponse) SetBody(v *BatchDeleteModelResponseBody) *BatchDeleteModelResponse {
	s.Body = v
	return s
}

type BatchDeleteResourcesRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModuleId       *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceIdList *string `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s BatchDeleteResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteResourcesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteResourcesRequest) SetAppId(v string) *BatchDeleteResourcesRequest {
	s.AppId = &v
	return s
}

func (s *BatchDeleteResourcesRequest) SetModuleId(v string) *BatchDeleteResourcesRequest {
	s.ModuleId = &v
	return s
}

func (s *BatchDeleteResourcesRequest) SetResourceIdList(v string) *BatchDeleteResourcesRequest {
	s.ResourceIdList = &v
	return s
}

func (s *BatchDeleteResourcesRequest) SetSource(v string) *BatchDeleteResourcesRequest {
	s.Source = &v
	return s
}

type BatchDeleteResourcesResponseBody struct {
	Data      *BatchDeleteResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteResourcesResponseBody) SetData(v *BatchDeleteResourcesResponseBodyData) *BatchDeleteResourcesResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteResourcesResponseBody) SetRequestId(v string) *BatchDeleteResourcesResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteResourcesResponseBodyData struct {
	Items []*BatchDeleteResourcesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s BatchDeleteResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeleteResourcesResponseBodyData) SetItems(v []*BatchDeleteResourcesResponseBodyDataItems) *BatchDeleteResourcesResponseBodyData {
	s.Items = v
	return s
}

type BatchDeleteResourcesResponseBodyDataItems struct {
	AppId         *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime  *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId    *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision      *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s BatchDeleteResourcesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteResourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetAppId(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetContent(v map[string]*string) *BatchDeleteResourcesResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetCreateTime(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetDescription(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetModifiedTime(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetModuleId(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetResourceId(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.ResourceId = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetResourceName(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.ResourceName = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetResourceType(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetRevision(v int32) *BatchDeleteResourcesResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *BatchDeleteResourcesResponseBodyDataItems) SetSchemaVersion(v string) *BatchDeleteResourcesResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

type BatchDeleteResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteResourcesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteResourcesResponse) SetHeaders(v map[string]*string) *BatchDeleteResourcesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteResourcesResponse) SetStatusCode(v int32) *BatchDeleteResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteResourcesResponse) SetBody(v *BatchDeleteResourcesResponseBody) *BatchDeleteResourcesResponse {
	s.Body = v
	return s
}

type BatchRestoreModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelIdList   *string `json:"ModelIdList,omitempty" xml:"ModelIdList,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s BatchRestoreModelRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRestoreModelRequest) GoString() string {
	return s.String()
}

func (s *BatchRestoreModelRequest) SetAppId(v string) *BatchRestoreModelRequest {
	s.AppId = &v
	return s
}

func (s *BatchRestoreModelRequest) SetModelIdList(v string) *BatchRestoreModelRequest {
	s.ModelIdList = &v
	return s
}

func (s *BatchRestoreModelRequest) SetModuleId(v string) *BatchRestoreModelRequest {
	s.ModuleId = &v
	return s
}

func (s *BatchRestoreModelRequest) SetSchemaVersion(v string) *BatchRestoreModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *BatchRestoreModelRequest) SetSource(v string) *BatchRestoreModelRequest {
	s.Source = &v
	return s
}

type BatchRestoreModelResponseBody struct {
	Data      *BatchRestoreModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchRestoreModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRestoreModelResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRestoreModelResponseBody) SetData(v *BatchRestoreModelResponseBodyData) *BatchRestoreModelResponseBody {
	s.Data = v
	return s
}

func (s *BatchRestoreModelResponseBody) SetRequestId(v string) *BatchRestoreModelResponseBody {
	s.RequestId = &v
	return s
}

type BatchRestoreModelResponseBodyData struct {
	Items []*BatchRestoreModelResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s BatchRestoreModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchRestoreModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchRestoreModelResponseBodyData) SetItems(v []*BatchRestoreModelResponseBodyDataItems) *BatchRestoreModelResponseBodyData {
	s.Items = v
	return s
}

type BatchRestoreModelResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s BatchRestoreModelResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s BatchRestoreModelResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *BatchRestoreModelResponseBodyDataItems) SetAppId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetAttributes(v []map[string]*string) *BatchRestoreModelResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetContent(v map[string]*string) *BatchRestoreModelResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetCreateTime(v string) *BatchRestoreModelResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetDescription(v string) *BatchRestoreModelResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetLinkModelId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetLinkModuleId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetLinked(v bool) *BatchRestoreModelResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModelId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModelName(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModelStatus(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModelType(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModifiedTime(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetModuleId(v string) *BatchRestoreModelResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetProps(v map[string]*string) *BatchRestoreModelResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetRevision(v int32) *BatchRestoreModelResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetSchemaVersion(v string) *BatchRestoreModelResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetSubType(v string) *BatchRestoreModelResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *BatchRestoreModelResponseBodyDataItems) SetVisibility(v string) *BatchRestoreModelResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type BatchRestoreModelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRestoreModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRestoreModelResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRestoreModelResponse) GoString() string {
	return s.String()
}

func (s *BatchRestoreModelResponse) SetHeaders(v map[string]*string) *BatchRestoreModelResponse {
	s.Headers = v
	return s
}

func (s *BatchRestoreModelResponse) SetStatusCode(v int32) *BatchRestoreModelResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRestoreModelResponse) SetBody(v *BatchRestoreModelResponseBody) *BatchRestoreModelResponse {
	s.Body = v
	return s
}

type CheckDomainRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) SetAppId(v string) *CheckDomainRequest {
	s.AppId = &v
	return s
}

func (s *CheckDomainRequest) SetDomain(v string) *CheckDomainRequest {
	s.Domain = &v
	return s
}

func (s *CheckDomainRequest) SetDomainType(v string) *CheckDomainRequest {
	s.DomainType = &v
	return s
}

func (s *CheckDomainRequest) SetEnvId(v string) *CheckDomainRequest {
	s.EnvId = &v
	return s
}

func (s *CheckDomainRequest) SetSource(v string) *CheckDomainRequest {
	s.Source = &v
	return s
}

type CheckDomainResponseBody struct {
	Data      *CheckDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) SetData(v *CheckDomainResponseBodyData) *CheckDomainResponseBody {
	s.Data = v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

type CheckDomainResponseBodyData struct {
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s CheckDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBodyData) SetValid(v bool) *CheckDomainResponseBodyData {
	s.Valid = &v
	return s
}

type CheckDomainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainResponse) SetHeaders(v map[string]*string) *CheckDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainResponse) SetStatusCode(v int32) *CheckDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainResponse) SetBody(v *CheckDomainResponseBody) *CheckDomainResponse {
	s.Body = v
	return s
}

type CloneAppRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CloneAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneAppRequest) GoString() string {
	return s.String()
}

func (s *CloneAppRequest) SetAppId(v string) *CloneAppRequest {
	s.AppId = &v
	return s
}

func (s *CloneAppRequest) SetAppName(v string) *CloneAppRequest {
	s.AppName = &v
	return s
}

func (s *CloneAppRequest) SetDescription(v string) *CloneAppRequest {
	s.Description = &v
	return s
}

func (s *CloneAppRequest) SetIcon(v string) *CloneAppRequest {
	s.Icon = &v
	return s
}

func (s *CloneAppRequest) SetSource(v string) *CloneAppRequest {
	s.Source = &v
	return s
}

type CloneAppResponseBody struct {
	Data      *CloneAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneAppResponseBody) GoString() string {
	return s.String()
}

func (s *CloneAppResponseBody) SetData(v *CloneAppResponseBodyData) *CloneAppResponseBody {
	s.Data = v
	return s
}

func (s *CloneAppResponseBody) SetRequestId(v string) *CloneAppResponseBody {
	s.RequestId = &v
	return s
}

type CloneAppResponseBodyData struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus     *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType       *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon          *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate    *bool   `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime  *string `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId  *string `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CloneAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CloneAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloneAppResponseBodyData) SetAppId(v string) *CloneAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CloneAppResponseBodyData) SetAppName(v string) *CloneAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CloneAppResponseBodyData) SetAppStatus(v string) *CloneAppResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *CloneAppResponseBodyData) SetAppType(v string) *CloneAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *CloneAppResponseBodyData) SetCreateTime(v string) *CloneAppResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloneAppResponseBodyData) SetDescription(v string) *CloneAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *CloneAppResponseBodyData) SetIcon(v string) *CloneAppResponseBodyData {
	s.Icon = &v
	return s
}

func (s *CloneAppResponseBodyData) SetIsTemplate(v bool) *CloneAppResponseBodyData {
	s.IsTemplate = &v
	return s
}

func (s *CloneAppResponseBodyData) SetLastEditTime(v string) *CloneAppResponseBodyData {
	s.LastEditTime = &v
	return s
}

func (s *CloneAppResponseBodyData) SetMainModuleId(v string) *CloneAppResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *CloneAppResponseBodyData) SetModifiedTime(v string) *CloneAppResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CloneAppResponseBodyData) SetSchemaVersion(v string) *CloneAppResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *CloneAppResponseBodyData) SetSource(v string) *CloneAppResponseBodyData {
	s.Source = &v
	return s
}

type CloneAppResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneAppResponse) GoString() string {
	return s.String()
}

func (s *CloneAppResponse) SetHeaders(v map[string]*string) *CloneAppResponse {
	s.Headers = v
	return s
}

func (s *CloneAppResponse) SetStatusCode(v int32) *CloneAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneAppResponse) SetBody(v *CloneAppResponseBody) *CloneAppResponse {
	s.Body = v
	return s
}

type CloneModelFromCommitRequest struct {
	ModelId        *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceCommitId *string `json:"SourceCommitId,omitempty" xml:"SourceCommitId,omitempty"`
	SourceModuleId *string `json:"SourceModuleId,omitempty" xml:"SourceModuleId,omitempty"`
	SubType        *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	TargetModuleId *string `json:"TargetModuleId,omitempty" xml:"TargetModuleId,omitempty"`
	TargetName     *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetSubType  *string `json:"TargetSubType,omitempty" xml:"TargetSubType,omitempty"`
}

func (s CloneModelFromCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneModelFromCommitRequest) GoString() string {
	return s.String()
}

func (s *CloneModelFromCommitRequest) SetModelId(v string) *CloneModelFromCommitRequest {
	s.ModelId = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetSource(v string) *CloneModelFromCommitRequest {
	s.Source = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetSourceCommitId(v string) *CloneModelFromCommitRequest {
	s.SourceCommitId = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetSourceModuleId(v string) *CloneModelFromCommitRequest {
	s.SourceModuleId = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetSubType(v string) *CloneModelFromCommitRequest {
	s.SubType = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetTargetModuleId(v string) *CloneModelFromCommitRequest {
	s.TargetModuleId = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetTargetName(v string) *CloneModelFromCommitRequest {
	s.TargetName = &v
	return s
}

func (s *CloneModelFromCommitRequest) SetTargetSubType(v string) *CloneModelFromCommitRequest {
	s.TargetSubType = &v
	return s
}

type CloneModelFromCommitResponseBody struct {
	Data      *CloneModelFromCommitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneModelFromCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneModelFromCommitResponseBody) GoString() string {
	return s.String()
}

func (s *CloneModelFromCommitResponseBody) SetData(v *CloneModelFromCommitResponseBodyData) *CloneModelFromCommitResponseBody {
	s.Data = v
	return s
}

func (s *CloneModelFromCommitResponseBody) SetRequestId(v string) *CloneModelFromCommitResponseBody {
	s.RequestId = &v
	return s
}

type CloneModelFromCommitResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CloneModelFromCommitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CloneModelFromCommitResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloneModelFromCommitResponseBodyData) SetAppId(v string) *CloneModelFromCommitResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetAttributes(v []map[string]*string) *CloneModelFromCommitResponseBodyData {
	s.Attributes = v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetContent(v map[string]*string) *CloneModelFromCommitResponseBodyData {
	s.Content = v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetCreateTime(v string) *CloneModelFromCommitResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetDescription(v string) *CloneModelFromCommitResponseBodyData {
	s.Description = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetId(v string) *CloneModelFromCommitResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetLinkModelId(v string) *CloneModelFromCommitResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetLinkModuleId(v string) *CloneModelFromCommitResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetLinked(v bool) *CloneModelFromCommitResponseBodyData {
	s.Linked = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModelId(v string) *CloneModelFromCommitResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModelName(v string) *CloneModelFromCommitResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModelStatus(v string) *CloneModelFromCommitResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModelType(v string) *CloneModelFromCommitResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModifiedTime(v string) *CloneModelFromCommitResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetModuleId(v string) *CloneModelFromCommitResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetProps(v map[string]*string) *CloneModelFromCommitResponseBodyData {
	s.Props = v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetRevision(v int32) *CloneModelFromCommitResponseBodyData {
	s.Revision = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetSchemaVersion(v string) *CloneModelFromCommitResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetSubType(v string) *CloneModelFromCommitResponseBodyData {
	s.SubType = &v
	return s
}

func (s *CloneModelFromCommitResponseBodyData) SetVisibility(v string) *CloneModelFromCommitResponseBodyData {
	s.Visibility = &v
	return s
}

type CloneModelFromCommitResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneModelFromCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneModelFromCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneModelFromCommitResponse) GoString() string {
	return s.String()
}

func (s *CloneModelFromCommitResponse) SetHeaders(v map[string]*string) *CloneModelFromCommitResponse {
	s.Headers = v
	return s
}

func (s *CloneModelFromCommitResponse) SetStatusCode(v int32) *CloneModelFromCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneModelFromCommitResponse) SetBody(v *CloneModelFromCommitResponseBody) *CloneModelFromCommitResponse {
	s.Body = v
	return s
}

type CloneModelInModuleRequest struct {
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TargetName    *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	TargetSubType *string `json:"TargetSubType,omitempty" xml:"TargetSubType,omitempty"`
}

func (s CloneModelInModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneModelInModuleRequest) GoString() string {
	return s.String()
}

func (s *CloneModelInModuleRequest) SetModelId(v string) *CloneModelInModuleRequest {
	s.ModelId = &v
	return s
}

func (s *CloneModelInModuleRequest) SetModuleId(v string) *CloneModelInModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *CloneModelInModuleRequest) SetSource(v string) *CloneModelInModuleRequest {
	s.Source = &v
	return s
}

func (s *CloneModelInModuleRequest) SetTargetName(v string) *CloneModelInModuleRequest {
	s.TargetName = &v
	return s
}

func (s *CloneModelInModuleRequest) SetTargetSubType(v string) *CloneModelInModuleRequest {
	s.TargetSubType = &v
	return s
}

type CloneModelInModuleResponseBody struct {
	Data      *CloneModelInModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneModelInModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneModelInModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CloneModelInModuleResponseBody) SetData(v *CloneModelInModuleResponseBodyData) *CloneModelInModuleResponseBody {
	s.Data = v
	return s
}

func (s *CloneModelInModuleResponseBody) SetRequestId(v string) *CloneModelInModuleResponseBody {
	s.RequestId = &v
	return s
}

type CloneModelInModuleResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CloneModelInModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CloneModelInModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloneModelInModuleResponseBodyData) SetAppId(v string) *CloneModelInModuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetAttributes(v []map[string]*string) *CloneModelInModuleResponseBodyData {
	s.Attributes = v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetContent(v map[string]*string) *CloneModelInModuleResponseBodyData {
	s.Content = v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetCreateTime(v string) *CloneModelInModuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetDescription(v string) *CloneModelInModuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetId(v string) *CloneModelInModuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetLinkModelId(v string) *CloneModelInModuleResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetLinkModuleId(v string) *CloneModelInModuleResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetLinked(v bool) *CloneModelInModuleResponseBodyData {
	s.Linked = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModelId(v string) *CloneModelInModuleResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModelName(v string) *CloneModelInModuleResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModelStatus(v string) *CloneModelInModuleResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModelType(v string) *CloneModelInModuleResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModifiedTime(v string) *CloneModelInModuleResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetModuleId(v string) *CloneModelInModuleResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetProps(v map[string]*string) *CloneModelInModuleResponseBodyData {
	s.Props = v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetRevision(v int32) *CloneModelInModuleResponseBodyData {
	s.Revision = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetSchemaVersion(v string) *CloneModelInModuleResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetSubType(v string) *CloneModelInModuleResponseBodyData {
	s.SubType = &v
	return s
}

func (s *CloneModelInModuleResponseBodyData) SetVisibility(v string) *CloneModelInModuleResponseBodyData {
	s.Visibility = &v
	return s
}

type CloneModelInModuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneModelInModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneModelInModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneModelInModuleResponse) GoString() string {
	return s.String()
}

func (s *CloneModelInModuleResponse) SetHeaders(v map[string]*string) *CloneModelInModuleResponse {
	s.Headers = v
	return s
}

func (s *CloneModelInModuleResponse) SetStatusCode(v int32) *CloneModelInModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneModelInModuleResponse) SetBody(v *CloneModelInModuleResponseBody) *CloneModelInModuleResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType         *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Asynchronous    *bool   `json:"Asynchronous,omitempty" xml:"Asynchronous,omitempty"`
	CategoryId      *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon            *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	PlatformVersion *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	SchemaVersion   *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceCommitId  *string `json:"SourceCommitId,omitempty" xml:"SourceCommitId,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Templated       *bool   `json:"Templated,omitempty" xml:"Templated,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppType(v string) *CreateAppRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppRequest) SetAsynchronous(v bool) *CreateAppRequest {
	s.Asynchronous = &v
	return s
}

func (s *CreateAppRequest) SetCategoryId(v string) *CreateAppRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetIcon(v string) *CreateAppRequest {
	s.Icon = &v
	return s
}

func (s *CreateAppRequest) SetPlatformVersion(v string) *CreateAppRequest {
	s.PlatformVersion = &v
	return s
}

func (s *CreateAppRequest) SetSchemaVersion(v string) *CreateAppRequest {
	s.SchemaVersion = &v
	return s
}

func (s *CreateAppRequest) SetSource(v string) *CreateAppRequest {
	s.Source = &v
	return s
}

func (s *CreateAppRequest) SetSourceCommitId(v string) *CreateAppRequest {
	s.SourceCommitId = &v
	return s
}

func (s *CreateAppRequest) SetTemplateId(v string) *CreateAppRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateAppRequest) SetTemplated(v bool) *CreateAppRequest {
	s.Templated = &v
	return s
}

type CreateAppResponseBody struct {
	Data      *CreateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppResponseBodyData struct {
	AppId           *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus       *string                                `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType         *string                                `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Categories      []*CreateAppResponseBodyDataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime      *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon            *string                                `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate      *bool                                  `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime    *string                                `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId    *string                                `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime    *string                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PlatformVersion *string                                `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	SchemaVersion   *string                                `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source          *string                                `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyData) SetAppId(v string) *CreateAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBodyData) SetAppName(v string) *CreateAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateAppResponseBodyData) SetAppStatus(v string) *CreateAppResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *CreateAppResponseBodyData) SetAppType(v string) *CreateAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *CreateAppResponseBodyData) SetCategories(v []*CreateAppResponseBodyDataCategories) *CreateAppResponseBodyData {
	s.Categories = v
	return s
}

func (s *CreateAppResponseBodyData) SetCreateTime(v string) *CreateAppResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateAppResponseBodyData) SetDescription(v string) *CreateAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyData) SetIcon(v string) *CreateAppResponseBodyData {
	s.Icon = &v
	return s
}

func (s *CreateAppResponseBodyData) SetIsTemplate(v bool) *CreateAppResponseBodyData {
	s.IsTemplate = &v
	return s
}

func (s *CreateAppResponseBodyData) SetLastEditTime(v string) *CreateAppResponseBodyData {
	s.LastEditTime = &v
	return s
}

func (s *CreateAppResponseBodyData) SetMainModuleId(v string) *CreateAppResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *CreateAppResponseBodyData) SetModifiedTime(v string) *CreateAppResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateAppResponseBodyData) SetPlatformVersion(v string) *CreateAppResponseBodyData {
	s.PlatformVersion = &v
	return s
}

func (s *CreateAppResponseBodyData) SetSchemaVersion(v string) *CreateAppResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *CreateAppResponseBodyData) SetSource(v string) *CreateAppResponseBodyData {
	s.Source = &v
	return s
}

type CreateAppResponseBodyDataCategories struct {
	CategoryId       *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName     *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s CreateAppResponseBodyDataCategories) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyDataCategories) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyDataCategories) SetCategoryId(v string) *CreateAppResponseBodyDataCategories {
	s.CategoryId = &v
	return s
}

func (s *CreateAppResponseBodyDataCategories) SetCategoryName(v string) *CreateAppResponseBodyDataCategories {
	s.CategoryName = &v
	return s
}

func (s *CreateAppResponseBodyDataCategories) SetParentCategoryId(v string) *CreateAppResponseBodyDataCategories {
	s.ParentCategoryId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateCommitRequest struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommitLog          *string `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	MainModuleCommitId *string `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	ModuleId           *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	RollbackToCommitId *string `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommitRequest) GoString() string {
	return s.String()
}

func (s *CreateCommitRequest) SetAppId(v string) *CreateCommitRequest {
	s.AppId = &v
	return s
}

func (s *CreateCommitRequest) SetClientToken(v string) *CreateCommitRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCommitRequest) SetCommitLog(v string) *CreateCommitRequest {
	s.CommitLog = &v
	return s
}

func (s *CreateCommitRequest) SetCommitType(v string) *CreateCommitRequest {
	s.CommitType = &v
	return s
}

func (s *CreateCommitRequest) SetMainModuleCommitId(v string) *CreateCommitRequest {
	s.MainModuleCommitId = &v
	return s
}

func (s *CreateCommitRequest) SetModuleId(v string) *CreateCommitRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateCommitRequest) SetRollbackToCommitId(v string) *CreateCommitRequest {
	s.RollbackToCommitId = &v
	return s
}

func (s *CreateCommitRequest) SetRollbackType(v string) *CreateCommitRequest {
	s.RollbackType = &v
	return s
}

func (s *CreateCommitRequest) SetSchemaVersion(v string) *CreateCommitRequest {
	s.SchemaVersion = &v
	return s
}

func (s *CreateCommitRequest) SetSource(v string) *CreateCommitRequest {
	s.Source = &v
	return s
}

type CreateCommitResponseBody struct {
	Data      *CreateCommitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCommitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommitResponseBody) SetData(v *CreateCommitResponseBodyData) *CreateCommitResponseBody {
	s.Data = v
	return s
}

func (s *CreateCommitResponseBody) SetRequestId(v string) *CreateCommitResponseBody {
	s.RequestId = &v
	return s
}

type CreateCommitResponseBodyData struct {
	AppId              *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId           *string            `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitLog          *string            `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string            `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	CreateTime         *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MainModuleCommitId *string            `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	MainModuleId       *string            `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModelDataPath      *string            `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModifiedTime       *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId           *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDataPath   *string            `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
	ResourceDigest     map[string]*string `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	RollbackToCommitId *string            `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string            `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s CreateCommitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCommitResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCommitResponseBodyData) SetAppId(v string) *CreateCommitResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetCommitId(v string) *CreateCommitResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetCommitLog(v string) *CreateCommitResponseBodyData {
	s.CommitLog = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetCommitType(v string) *CreateCommitResponseBodyData {
	s.CommitType = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetCreateTime(v string) *CreateCommitResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetMainModuleCommitId(v string) *CreateCommitResponseBodyData {
	s.MainModuleCommitId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetMainModuleId(v string) *CreateCommitResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetModelDataPath(v string) *CreateCommitResponseBodyData {
	s.ModelDataPath = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetModifiedTime(v string) *CreateCommitResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetModuleId(v string) *CreateCommitResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetResourceDataPath(v string) *CreateCommitResponseBodyData {
	s.ResourceDataPath = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetResourceDigest(v map[string]*string) *CreateCommitResponseBodyData {
	s.ResourceDigest = v
	return s
}

func (s *CreateCommitResponseBodyData) SetRollbackToCommitId(v string) *CreateCommitResponseBodyData {
	s.RollbackToCommitId = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetRollbackType(v string) *CreateCommitResponseBodyData {
	s.RollbackType = &v
	return s
}

func (s *CreateCommitResponseBodyData) SetSchemaVersion(v string) *CreateCommitResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type CreateCommitResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCommitResponse) GoString() string {
	return s.String()
}

func (s *CreateCommitResponse) SetHeaders(v map[string]*string) *CreateCommitResponse {
	s.Headers = v
	return s
}

func (s *CreateCommitResponse) SetStatusCode(v int32) *CreateCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommitResponse) SetBody(v *CreateCommitResponseBody) *CreateCommitResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType      *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PublicKey       *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WithCertificate *bool   `json:"WithCertificate,omitempty" xml:"WithCertificate,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetAppId(v string) *CreateDomainRequest {
	s.AppId = &v
	return s
}

func (s *CreateDomainRequest) SetClientToken(v string) *CreateDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetDomainType(v string) *CreateDomainRequest {
	s.DomainType = &v
	return s
}

func (s *CreateDomainRequest) SetEnvId(v string) *CreateDomainRequest {
	s.EnvId = &v
	return s
}

func (s *CreateDomainRequest) SetPath(v string) *CreateDomainRequest {
	s.Path = &v
	return s
}

func (s *CreateDomainRequest) SetPrivateKey(v string) *CreateDomainRequest {
	s.PrivateKey = &v
	return s
}

func (s *CreateDomainRequest) SetPublicKey(v string) *CreateDomainRequest {
	s.PublicKey = &v
	return s
}

func (s *CreateDomainRequest) SetSource(v string) *CreateDomainRequest {
	s.Source = &v
	return s
}

func (s *CreateDomainRequest) SetWithCertificate(v bool) *CreateDomainRequest {
	s.WithCertificate = &v
	return s
}

type CreateDomainResponseBody struct {
	Data      *CreateDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetData(v *CreateDomainResponseBodyData) *CreateDomainResponseBody {
	s.Data = v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponseBodyData struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Applied         *bool   `json:"Applied,omitempty" xml:"Applied,omitempty"`
	Checked         *bool   `json:"Checked,omitempty" xml:"Checked,omitempty"`
	Cname           *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Deleted         *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType      *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	WithCertificate *bool   `json:"WithCertificate,omitempty" xml:"WithCertificate,omitempty"`
}

func (s CreateDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyData) SetAppId(v string) *CreateDomainResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetApplied(v bool) *CreateDomainResponseBodyData {
	s.Applied = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetChecked(v bool) *CreateDomainResponseBodyData {
	s.Checked = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetCname(v string) *CreateDomainResponseBodyData {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetDeleted(v bool) *CreateDomainResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetDomain(v string) *CreateDomainResponseBodyData {
	s.Domain = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetDomainType(v string) *CreateDomainResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetEnvId(v string) *CreateDomainResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetPath(v string) *CreateDomainResponseBodyData {
	s.Path = &v
	return s
}

func (s *CreateDomainResponseBodyData) SetWithCertificate(v bool) *CreateDomainResponseBodyData {
	s.WithCertificate = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateLinkEntityAndAssociationRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ModelData   *string `json:"ModelData,omitempty" xml:"ModelData,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateLinkEntityAndAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLinkEntityAndAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateLinkEntityAndAssociationRequest) SetClientToken(v string) *CreateLinkEntityAndAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLinkEntityAndAssociationRequest) SetModelData(v string) *CreateLinkEntityAndAssociationRequest {
	s.ModelData = &v
	return s
}

func (s *CreateLinkEntityAndAssociationRequest) SetSource(v string) *CreateLinkEntityAndAssociationRequest {
	s.Source = &v
	return s
}

type CreateLinkEntityAndAssociationResponseBody struct {
	Data      *CreateLinkEntityAndAssociationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLinkEntityAndAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLinkEntityAndAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLinkEntityAndAssociationResponseBody) SetData(v *CreateLinkEntityAndAssociationResponseBodyData) *CreateLinkEntityAndAssociationResponseBody {
	s.Data = v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBody) SetRequestId(v string) *CreateLinkEntityAndAssociationResponseBody {
	s.RequestId = &v
	return s
}

type CreateLinkEntityAndAssociationResponseBodyData struct {
	Items []*CreateLinkEntityAndAssociationResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s CreateLinkEntityAndAssociationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLinkEntityAndAssociationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLinkEntityAndAssociationResponseBodyData) SetItems(v []*CreateLinkEntityAndAssociationResponseBodyDataItems) *CreateLinkEntityAndAssociationResponseBodyData {
	s.Items = v
	return s
}

type CreateLinkEntityAndAssociationResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateLinkEntityAndAssociationResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s CreateLinkEntityAndAssociationResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetAppId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetAttributes(v []map[string]*string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetContent(v map[string]*string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetCreateTime(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetDescription(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetLinkModelId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetLinkModuleId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetLinked(v bool) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModelId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModelName(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModelStatus(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModelType(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModifiedTime(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetModuleId(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetProps(v map[string]*string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetRevision(v int32) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetSchemaVersion(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetSubType(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponseBodyDataItems) SetVisibility(v string) *CreateLinkEntityAndAssociationResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type CreateLinkEntityAndAssociationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLinkEntityAndAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLinkEntityAndAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLinkEntityAndAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateLinkEntityAndAssociationResponse) SetHeaders(v map[string]*string) *CreateLinkEntityAndAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateLinkEntityAndAssociationResponse) SetStatusCode(v int32) *CreateLinkEntityAndAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLinkEntityAndAssociationResponse) SetBody(v *CreateLinkEntityAndAssociationResponseBody) *CreateLinkEntityAndAssociationResponse {
	s.Body = v
	return s
}

type CreateModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncodeType    *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	LinkModelId   *string `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool   `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) SetAppId(v string) *CreateModelRequest {
	s.AppId = &v
	return s
}

func (s *CreateModelRequest) SetClientToken(v string) *CreateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelRequest) SetContent(v string) *CreateModelRequest {
	s.Content = &v
	return s
}

func (s *CreateModelRequest) SetDescription(v string) *CreateModelRequest {
	s.Description = &v
	return s
}

func (s *CreateModelRequest) SetEncodeType(v string) *CreateModelRequest {
	s.EncodeType = &v
	return s
}

func (s *CreateModelRequest) SetLinkModelId(v string) *CreateModelRequest {
	s.LinkModelId = &v
	return s
}

func (s *CreateModelRequest) SetLinkModuleId(v string) *CreateModelRequest {
	s.LinkModuleId = &v
	return s
}

func (s *CreateModelRequest) SetLinked(v bool) *CreateModelRequest {
	s.Linked = &v
	return s
}

func (s *CreateModelRequest) SetModelId(v string) *CreateModelRequest {
	s.ModelId = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetModelType(v string) *CreateModelRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelRequest) SetModuleId(v string) *CreateModelRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateModelRequest) SetSchemaVersion(v string) *CreateModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *CreateModelRequest) SetSource(v string) *CreateModelRequest {
	s.Source = &v
	return s
}

func (s *CreateModelRequest) SetSubType(v string) *CreateModelRequest {
	s.SubType = &v
	return s
}

func (s *CreateModelRequest) SetVisibility(v string) *CreateModelRequest {
	s.Visibility = &v
	return s
}

type CreateModelResponseBody struct {
	Data      *CreateModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetData(v *CreateModelResponseBodyData) *CreateModelResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

type CreateModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBodyData) SetAppId(v string) *CreateModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetAttributes(v []map[string]*string) *CreateModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *CreateModelResponseBodyData) SetContent(v map[string]*string) *CreateModelResponseBodyData {
	s.Content = v
	return s
}

func (s *CreateModelResponseBodyData) SetCreateTime(v string) *CreateModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateModelResponseBodyData) SetDescription(v string) *CreateModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModelResponseBodyData) SetId(v string) *CreateModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateModelResponseBodyData) SetLinkModelId(v string) *CreateModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetLinkModuleId(v string) *CreateModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetLinked(v bool) *CreateModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelDigest(v string) *CreateModelResponseBodyData {
	s.ModelDigest = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelId(v string) *CreateModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelName(v string) *CreateModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelStatus(v string) *CreateModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelType(v string) *CreateModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModifiedTime(v string) *CreateModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModuleId(v string) *CreateModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetProps(v map[string]*string) *CreateModelResponseBodyData {
	s.Props = v
	return s
}

func (s *CreateModelResponseBodyData) SetRevision(v int32) *CreateModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *CreateModelResponseBodyData) SetSchemaVersion(v string) *CreateModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *CreateModelResponseBodyData) SetSubType(v string) *CreateModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *CreateModelResponseBodyData) SetVisibility(v string) *CreateModelResponseBodyData {
	s.Visibility = &v
	return s
}

type CreateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetStatusCode(v int32) *CreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreateModuleRequest struct {
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType             *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Source                 *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceModuleId         *string `json:"SourceModuleId,omitempty" xml:"SourceModuleId,omitempty"`
	TargetAppSource        *string `json:"TargetAppSource,omitempty" xml:"TargetAppSource,omitempty"`
}

func (s CreateModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleRequest) SetClientToken(v string) *CreateModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModuleRequest) SetDescription(v string) *CreateModuleRequest {
	s.Description = &v
	return s
}

func (s *CreateModuleRequest) SetIcon(v string) *CreateModuleRequest {
	s.Icon = &v
	return s
}

func (s *CreateModuleRequest) SetMinimumPlatformVersion(v string) *CreateModuleRequest {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *CreateModuleRequest) SetModuleName(v string) *CreateModuleRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateModuleRequest) SetModuleType(v string) *CreateModuleRequest {
	s.ModuleType = &v
	return s
}

func (s *CreateModuleRequest) SetPlatform(v string) *CreateModuleRequest {
	s.Platform = &v
	return s
}

func (s *CreateModuleRequest) SetSource(v string) *CreateModuleRequest {
	s.Source = &v
	return s
}

func (s *CreateModuleRequest) SetSourceModuleId(v string) *CreateModuleRequest {
	s.SourceModuleId = &v
	return s
}

func (s *CreateModuleRequest) SetTargetAppSource(v string) *CreateModuleRequest {
	s.TargetAppSource = &v
	return s
}

type CreateModuleResponseBody struct {
	Data      *CreateModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleResponseBody) SetData(v *CreateModuleResponseBodyData) *CreateModuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateModuleResponseBody) SetRequestId(v string) *CreateModuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateModuleResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType             *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PlatformVersion        *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
}

func (s CreateModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModuleResponseBodyData) SetCreateTime(v string) *CreateModuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetDescription(v string) *CreateModuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetIcon(v string) *CreateModuleResponseBodyData {
	s.Icon = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetLatestPublishedCommit(v string) *CreateModuleResponseBodyData {
	s.LatestPublishedCommit = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetLatestPublishedVersion(v string) *CreateModuleResponseBodyData {
	s.LatestPublishedVersion = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetMinimumPlatformVersion(v string) *CreateModuleResponseBodyData {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetModifiedTime(v string) *CreateModuleResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetModuleId(v string) *CreateModuleResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetModuleName(v string) *CreateModuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetModuleType(v string) *CreateModuleResponseBodyData {
	s.ModuleType = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetOwnerAppId(v string) *CreateModuleResponseBodyData {
	s.OwnerAppId = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetOwnerUserId(v string) *CreateModuleResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetPlatform(v string) *CreateModuleResponseBodyData {
	s.Platform = &v
	return s
}

func (s *CreateModuleResponseBodyData) SetPlatformVersion(v string) *CreateModuleResponseBodyData {
	s.PlatformVersion = &v
	return s
}

type CreateModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleResponse) SetHeaders(v map[string]*string) *CreateModuleResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleResponse) SetStatusCode(v int32) *CreateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModuleResponse) SetBody(v *CreateModuleResponseBody) *CreateModuleResponse {
	s.Body = v
	return s
}

type CreateModulePublishRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleId       *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PublishVersion *string `json:"PublishVersion,omitempty" xml:"PublishVersion,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateModulePublishRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModulePublishRequest) GoString() string {
	return s.String()
}

func (s *CreateModulePublishRequest) SetClientToken(v string) *CreateModulePublishRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModulePublishRequest) SetDescription(v string) *CreateModulePublishRequest {
	s.Description = &v
	return s
}

func (s *CreateModulePublishRequest) SetModuleId(v string) *CreateModulePublishRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateModulePublishRequest) SetPublishVersion(v string) *CreateModulePublishRequest {
	s.PublishVersion = &v
	return s
}

func (s *CreateModulePublishRequest) SetSource(v string) *CreateModulePublishRequest {
	s.Source = &v
	return s
}

type CreateModulePublishResponseBody struct {
	Data      *CreateModulePublishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModulePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModulePublishResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModulePublishResponseBody) SetData(v *CreateModulePublishResponseBodyData) *CreateModulePublishResponseBody {
	s.Data = v
	return s
}

func (s *CreateModulePublishResponseBody) SetRequestId(v string) *CreateModulePublishResponseBody {
	s.RequestId = &v
	return s
}

type CreateModulePublishResponseBodyData struct {
	CommitId     *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId     *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PublishId    *string `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateModulePublishResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateModulePublishResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModulePublishResponseBodyData) SetCommitId(v string) *CreateModulePublishResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetCreateTime(v string) *CreateModulePublishResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetDescription(v string) *CreateModulePublishResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetModifiedTime(v string) *CreateModulePublishResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetModuleId(v string) *CreateModulePublishResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetPublishId(v string) *CreateModulePublishResponseBodyData {
	s.PublishId = &v
	return s
}

func (s *CreateModulePublishResponseBodyData) SetVersion(v string) *CreateModulePublishResponseBodyData {
	s.Version = &v
	return s
}

type CreateModulePublishResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModulePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModulePublishResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModulePublishResponse) GoString() string {
	return s.String()
}

func (s *CreateModulePublishResponse) SetHeaders(v map[string]*string) *CreateModulePublishResponse {
	s.Headers = v
	return s
}

func (s *CreateModulePublishResponse) SetStatusCode(v int32) *CreateModulePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModulePublishResponse) SetBody(v *CreateModulePublishResponseBody) *CreateModulePublishResponse {
	s.Body = v
	return s
}

type CreatePublishRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommitId      *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvType       *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	VersionNumber *string `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s CreatePublishRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishRequest) SetAppId(v string) *CreatePublishRequest {
	s.AppId = &v
	return s
}

func (s *CreatePublishRequest) SetClientToken(v string) *CreatePublishRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePublishRequest) SetCommitId(v string) *CreatePublishRequest {
	s.CommitId = &v
	return s
}

func (s *CreatePublishRequest) SetDescription(v string) *CreatePublishRequest {
	s.Description = &v
	return s
}

func (s *CreatePublishRequest) SetEnvType(v string) *CreatePublishRequest {
	s.EnvType = &v
	return s
}

func (s *CreatePublishRequest) SetPublishType(v string) *CreatePublishRequest {
	s.PublishType = &v
	return s
}

func (s *CreatePublishRequest) SetSource(v string) *CreatePublishRequest {
	s.Source = &v
	return s
}

func (s *CreatePublishRequest) SetVersionNumber(v string) *CreatePublishRequest {
	s.VersionNumber = &v
	return s
}

type CreatePublishResponseBody struct {
	Data      *CreatePublishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublishResponseBody) SetData(v *CreatePublishResponseBodyData) *CreatePublishResponseBody {
	s.Data = v
	return s
}

func (s *CreatePublishResponseBody) SetRequestId(v string) *CreatePublishResponseBody {
	s.RequestId = &v
	return s
}

type CreatePublishResponseBodyData struct {
	AppId          *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId       *string              `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CompletionTime *string              `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreateTime     *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvId          *string              `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ModifiedTime   *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishId      *string              `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	PublishStatus  *string              `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishType    *string              `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Reason         *string              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime      *string              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubTasks       []map[string]*string `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	VersionNumber  *string              `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s CreatePublishResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePublishResponseBodyData) SetAppId(v string) *CreatePublishResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetCommitId(v string) *CreatePublishResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetCompletionTime(v string) *CreatePublishResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetCreateTime(v string) *CreatePublishResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetDescription(v string) *CreatePublishResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetEnvId(v string) *CreatePublishResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetModifiedTime(v string) *CreatePublishResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetPublishId(v string) *CreatePublishResponseBodyData {
	s.PublishId = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetPublishStatus(v string) *CreatePublishResponseBodyData {
	s.PublishStatus = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetPublishType(v string) *CreatePublishResponseBodyData {
	s.PublishType = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetReason(v string) *CreatePublishResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetStartTime(v string) *CreatePublishResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreatePublishResponseBodyData) SetSubTasks(v []map[string]*string) *CreatePublishResponseBodyData {
	s.SubTasks = v
	return s
}

func (s *CreatePublishResponseBodyData) SetVersionNumber(v string) *CreatePublishResponseBodyData {
	s.VersionNumber = &v
	return s
}

type CreatePublishResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublishResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishResponse) GoString() string {
	return s.String()
}

func (s *CreatePublishResponse) SetHeaders(v map[string]*string) *CreatePublishResponse {
	s.Headers = v
	return s
}

func (s *CreatePublishResponse) SetStatusCode(v int32) *CreatePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublishResponse) SetBody(v *CreatePublishResponseBody) *CreatePublishResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetAppId(v string) *CreateResourceRequest {
	s.AppId = &v
	return s
}

func (s *CreateResourceRequest) SetClientToken(v string) *CreateResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceRequest) SetContent(v string) *CreateResourceRequest {
	s.Content = &v
	return s
}

func (s *CreateResourceRequest) SetDescription(v string) *CreateResourceRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceRequest) SetModuleId(v string) *CreateResourceRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateResourceRequest) SetResourceId(v string) *CreateResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceRequest) SetResourceName(v string) *CreateResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateResourceRequest) SetResourceType(v string) *CreateResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceRequest) SetSchemaVersion(v string) *CreateResourceRequest {
	s.SchemaVersion = &v
	return s
}

func (s *CreateResourceRequest) SetSource(v string) *CreateResourceRequest {
	s.Source = &v
	return s
}

func (s *CreateResourceRequest) SetVisibility(v string) *CreateResourceRequest {
	s.Visibility = &v
	return s
}

type CreateResourceResponseBody struct {
	Data      *CreateResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetData(v *CreateResourceResponseBodyData) *CreateResourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceResponseBodyData struct {
	AppId          *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime   *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId       *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDigest *string            `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourceId     *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision       *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion  *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s CreateResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBodyData) SetAppId(v string) *CreateResourceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetContent(v map[string]*string) *CreateResourceResponseBodyData {
	s.Content = v
	return s
}

func (s *CreateResourceResponseBodyData) SetCreateTime(v string) *CreateResourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetDescription(v string) *CreateResourceResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetModifiedTime(v string) *CreateResourceResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetModuleId(v string) *CreateResourceResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetResourceDigest(v string) *CreateResourceResponseBodyData {
	s.ResourceDigest = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetResourceId(v string) *CreateResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetResourceName(v string) *CreateResourceResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetResourceType(v string) *CreateResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetRevision(v int32) *CreateResourceResponseBodyData {
	s.Revision = &v
	return s
}

func (s *CreateResourceResponseBodyData) SetSchemaVersion(v string) *CreateResourceResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppId(v string) *DeleteAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRequest) SetSource(v string) *DeleteAppRequest {
	s.Source = &v
	return s
}

type DeleteAppResponseBody struct {
	Data      *DeleteAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetData(v *DeleteAppResponseBodyData) *DeleteAppResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponseBodyData struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus     *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType       *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon          *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate    *bool   `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime  *string `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId  *string `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBodyData) SetAppId(v string) *DeleteAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetAppName(v string) *DeleteAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetAppStatus(v string) *DeleteAppResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetAppType(v string) *DeleteAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetCreateTime(v string) *DeleteAppResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetDescription(v string) *DeleteAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetIcon(v string) *DeleteAppResponseBodyData {
	s.Icon = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetIsTemplate(v bool) *DeleteAppResponseBodyData {
	s.IsTemplate = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetLastEditTime(v string) *DeleteAppResponseBodyData {
	s.LastEditTime = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetMainModuleId(v string) *DeleteAppResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetModifiedTime(v string) *DeleteAppResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetSchemaVersion(v string) *DeleteAppResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *DeleteAppResponseBodyData) SetSource(v string) *DeleteAppResponseBodyData {
	s.Source = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteCommitRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommitRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommitRequest) SetAppId(v string) *DeleteCommitRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCommitRequest) SetCommitId(v string) *DeleteCommitRequest {
	s.CommitId = &v
	return s
}

func (s *DeleteCommitRequest) SetModuleId(v string) *DeleteCommitRequest {
	s.ModuleId = &v
	return s
}

func (s *DeleteCommitRequest) SetSource(v string) *DeleteCommitRequest {
	s.Source = &v
	return s
}

type DeleteCommitResponseBody struct {
	Data      *DeleteCommitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommitResponseBody) SetData(v *DeleteCommitResponseBodyData) *DeleteCommitResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCommitResponseBody) SetRequestId(v string) *DeleteCommitResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCommitResponseBodyData struct {
	AppId              *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId           *string            `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitLog          *string            `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string            `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	CreateTime         *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MainModuleCommitId *string            `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	MainModuleId       *string            `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModelDataPath      *string            `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModifiedTime       *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId           *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDataPath   *string            `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
	ResourceDigest     map[string]*string `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	RollbackToCommitId *string            `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string            `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s DeleteCommitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommitResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCommitResponseBodyData) SetAppId(v string) *DeleteCommitResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetCommitId(v string) *DeleteCommitResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetCommitLog(v string) *DeleteCommitResponseBodyData {
	s.CommitLog = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetCommitType(v string) *DeleteCommitResponseBodyData {
	s.CommitType = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetCreateTime(v string) *DeleteCommitResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetMainModuleCommitId(v string) *DeleteCommitResponseBodyData {
	s.MainModuleCommitId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetMainModuleId(v string) *DeleteCommitResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetModelDataPath(v string) *DeleteCommitResponseBodyData {
	s.ModelDataPath = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetModifiedTime(v string) *DeleteCommitResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetModuleId(v string) *DeleteCommitResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetResourceDataPath(v string) *DeleteCommitResponseBodyData {
	s.ResourceDataPath = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetResourceDigest(v map[string]*string) *DeleteCommitResponseBodyData {
	s.ResourceDigest = v
	return s
}

func (s *DeleteCommitResponseBodyData) SetRollbackToCommitId(v string) *DeleteCommitResponseBodyData {
	s.RollbackToCommitId = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetRollbackType(v string) *DeleteCommitResponseBodyData {
	s.RollbackType = &v
	return s
}

func (s *DeleteCommitResponseBodyData) SetSchemaVersion(v string) *DeleteCommitResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type DeleteCommitResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommitResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommitResponse) SetHeaders(v map[string]*string) *DeleteCommitResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommitResponse) SetStatusCode(v int32) *DeleteCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommitResponse) SetBody(v *DeleteCommitResponseBody) *DeleteCommitResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetAppId(v string) *DeleteDomainRequest {
	s.AppId = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainRequest) SetEnvId(v string) *DeleteDomainRequest {
	s.EnvId = &v
	return s
}

func (s *DeleteDomainRequest) SetSource(v string) *DeleteDomainRequest {
	s.Source = &v
	return s
}

type DeleteDomainResponseBody struct {
	Data      *DeleteDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetData(v *DeleteDomainResponseBodyData) *DeleteDomainResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponseBodyData struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Applied    *bool   `json:"Applied,omitempty" xml:"Applied,omitempty"`
	Deleted    *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBodyData) SetAppId(v string) *DeleteDomainResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetApplied(v bool) *DeleteDomainResponseBodyData {
	s.Applied = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetDeleted(v bool) *DeleteDomainResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetDomain(v string) *DeleteDomainResponseBodyData {
	s.Domain = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetDomainType(v string) *DeleteDomainResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetEnvId(v string) *DeleteDomainResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *DeleteDomainResponseBodyData) SetPath(v string) *DeleteDomainResponseBodyData {
	s.Path = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) SetAppId(v string) *DeleteModelRequest {
	s.AppId = &v
	return s
}

func (s *DeleteModelRequest) SetModelId(v string) *DeleteModelRequest {
	s.ModelId = &v
	return s
}

func (s *DeleteModelRequest) SetModuleId(v string) *DeleteModelRequest {
	s.ModuleId = &v
	return s
}

func (s *DeleteModelRequest) SetSchemaVersion(v string) *DeleteModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *DeleteModelRequest) SetSource(v string) *DeleteModelRequest {
	s.Source = &v
	return s
}

type DeleteModelResponseBody struct {
	Data      *DeleteModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetData(v *DeleteModelResponseBodyData) *DeleteModelResponseBody {
	s.Data = v
	return s
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DeleteModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBodyData) SetAppId(v string) *DeleteModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetAttributes(v []map[string]*string) *DeleteModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *DeleteModelResponseBodyData) SetContent(v map[string]*string) *DeleteModelResponseBodyData {
	s.Content = v
	return s
}

func (s *DeleteModelResponseBodyData) SetCreateTime(v string) *DeleteModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetDescription(v string) *DeleteModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetId(v string) *DeleteModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetLinkModelId(v string) *DeleteModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetLinkModuleId(v string) *DeleteModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetLinked(v bool) *DeleteModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModelId(v string) *DeleteModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModelName(v string) *DeleteModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModelStatus(v string) *DeleteModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModelType(v string) *DeleteModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModifiedTime(v string) *DeleteModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetModuleId(v string) *DeleteModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetProps(v map[string]*string) *DeleteModelResponseBodyData {
	s.Props = v
	return s
}

func (s *DeleteModelResponseBodyData) SetRevision(v int32) *DeleteModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetSchemaVersion(v string) *DeleteModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetSubType(v string) *DeleteModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetVisibility(v string) *DeleteModelResponseBodyData {
	s.Visibility = &v
	return s
}

type DeleteModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetStatusCode(v int32) *DeleteModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeleteModuleRequest struct {
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteModuleRequest) SetModuleId(v string) *DeleteModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *DeleteModuleRequest) SetSource(v string) *DeleteModuleRequest {
	s.Source = &v
	return s
}

type DeleteModuleResponseBody struct {
	Data      *DeleteModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponseBody) SetData(v *DeleteModuleResponseBodyData) *DeleteModuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteModuleResponseBody) SetRequestId(v string) *DeleteModuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModuleResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DeleteModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponseBodyData) SetCreateTime(v string) *DeleteModuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetDescription(v string) *DeleteModuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetIcon(v string) *DeleteModuleResponseBodyData {
	s.Icon = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetLatestPublishedCommit(v string) *DeleteModuleResponseBodyData {
	s.LatestPublishedCommit = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetLatestPublishedVersion(v string) *DeleteModuleResponseBodyData {
	s.LatestPublishedVersion = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetMinimumPlatformVersion(v string) *DeleteModuleResponseBodyData {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetModifiedTime(v string) *DeleteModuleResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetModuleId(v string) *DeleteModuleResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetModuleName(v string) *DeleteModuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetOwnerAppId(v string) *DeleteModuleResponseBodyData {
	s.OwnerAppId = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetOwnerUserId(v string) *DeleteModuleResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *DeleteModuleResponseBodyData) SetPlatform(v string) *DeleteModuleResponseBodyData {
	s.Platform = &v
	return s
}

type DeleteModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponse) SetHeaders(v map[string]*string) *DeleteModuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteModuleResponse) SetStatusCode(v int32) *DeleteModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModuleResponse) SetBody(v *DeleteModuleResponseBody) *DeleteModuleResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModuleId   *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetAppId(v string) *DeleteResourceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteResourceRequest) SetModuleId(v string) *DeleteResourceRequest {
	s.ModuleId = &v
	return s
}

func (s *DeleteResourceRequest) SetResourceId(v string) *DeleteResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteResourceRequest) SetSource(v string) *DeleteResourceRequest {
	s.Source = &v
	return s
}

type DeleteResourceResponseBody struct {
	Data      *DeleteResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetData(v *DeleteResourceResponseBodyData) *DeleteResourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceResponseBodyData struct {
	AppId         *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime  *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId    *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision      *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s DeleteResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBodyData) SetAppId(v string) *DeleteResourceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetContent(v map[string]*string) *DeleteResourceResponseBodyData {
	s.Content = v
	return s
}

func (s *DeleteResourceResponseBodyData) SetCreateTime(v string) *DeleteResourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetDescription(v string) *DeleteResourceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetModifiedTime(v string) *DeleteResourceResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetModuleId(v string) *DeleteResourceResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetResourceId(v string) *DeleteResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetResourceName(v string) *DeleteResourceResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetResourceType(v string) *DeleteResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetRevision(v int32) *DeleteResourceResponseBodyData {
	s.Revision = &v
	return s
}

func (s *DeleteResourceResponseBodyData) SetSchemaVersion(v string) *DeleteResourceResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type GenerateAppUserPasswordRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId    *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GenerateAppUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAppUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *GenerateAppUserPasswordRequest) SetAppId(v string) *GenerateAppUserPasswordRequest {
	s.AppId = &v
	return s
}

func (s *GenerateAppUserPasswordRequest) SetEnvId(v string) *GenerateAppUserPasswordRequest {
	s.EnvId = &v
	return s
}

func (s *GenerateAppUserPasswordRequest) SetSource(v string) *GenerateAppUserPasswordRequest {
	s.Source = &v
	return s
}

func (s *GenerateAppUserPasswordRequest) SetUserName(v string) *GenerateAppUserPasswordRequest {
	s.UserName = &v
	return s
}

type GenerateAppUserPasswordResponseBody struct {
	Data      *GenerateAppUserPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAppUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAppUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAppUserPasswordResponseBody) SetData(v *GenerateAppUserPasswordResponseBodyData) *GenerateAppUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *GenerateAppUserPasswordResponseBody) SetRequestId(v string) *GenerateAppUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAppUserPasswordResponseBodyData struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GenerateAppUserPasswordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateAppUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateAppUserPasswordResponseBodyData) SetPassword(v string) *GenerateAppUserPasswordResponseBodyData {
	s.Password = &v
	return s
}

func (s *GenerateAppUserPasswordResponseBodyData) SetUserName(v string) *GenerateAppUserPasswordResponseBodyData {
	s.UserName = &v
	return s
}

type GenerateAppUserPasswordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAppUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAppUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAppUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *GenerateAppUserPasswordResponse) SetHeaders(v map[string]*string) *GenerateAppUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *GenerateAppUserPasswordResponse) SetStatusCode(v int32) *GenerateAppUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAppUserPasswordResponse) SetBody(v *GenerateAppUserPasswordResponseBody) *GenerateAppUserPasswordResponse {
	s.Body = v
	return s
}

type GenerateAuthTokenRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GenerateAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateAuthTokenRequest) SetAppId(v string) *GenerateAuthTokenRequest {
	s.AppId = &v
	return s
}

func (s *GenerateAuthTokenRequest) SetModuleId(v string) *GenerateAuthTokenRequest {
	s.ModuleId = &v
	return s
}

func (s *GenerateAuthTokenRequest) SetSource(v string) *GenerateAuthTokenRequest {
	s.Source = &v
	return s
}

type GenerateAuthTokenResponseBody struct {
	Data      *GenerateAuthTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAuthTokenResponseBody) SetData(v *GenerateAuthTokenResponseBodyData) *GenerateAuthTokenResponseBody {
	s.Data = v
	return s
}

func (s *GenerateAuthTokenResponseBody) SetRequestId(v string) *GenerateAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAuthTokenResponseBodyData struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GenerateAuthTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateAuthTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateAuthTokenResponseBodyData) SetJwtToken(v string) *GenerateAuthTokenResponseBodyData {
	s.JwtToken = &v
	return s
}

type GenerateAuthTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateAuthTokenResponse) SetHeaders(v map[string]*string) *GenerateAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateAuthTokenResponse) SetStatusCode(v int32) *GenerateAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAuthTokenResponse) SetBody(v *GenerateAuthTokenResponseBody) *GenerateAuthTokenResponse {
	s.Body = v
	return s
}

type GenerateUploadTokenRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MaterialId      *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	ModuleId        *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UploadTokenType *string `json:"UploadTokenType,omitempty" xml:"UploadTokenType,omitempty"`
}

func (s GenerateUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadTokenRequest) SetAppId(v string) *GenerateUploadTokenRequest {
	s.AppId = &v
	return s
}

func (s *GenerateUploadTokenRequest) SetMaterialId(v string) *GenerateUploadTokenRequest {
	s.MaterialId = &v
	return s
}

func (s *GenerateUploadTokenRequest) SetModuleId(v string) *GenerateUploadTokenRequest {
	s.ModuleId = &v
	return s
}

func (s *GenerateUploadTokenRequest) SetSource(v string) *GenerateUploadTokenRequest {
	s.Source = &v
	return s
}

func (s *GenerateUploadTokenRequest) SetUploadTokenType(v string) *GenerateUploadTokenRequest {
	s.UploadTokenType = &v
	return s
}

type GenerateUploadTokenResponseBody struct {
	Data      *GenerateUploadTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadTokenResponseBody) SetData(v *GenerateUploadTokenResponseBodyData) *GenerateUploadTokenResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadTokenResponseBody) SetRequestId(v string) *GenerateUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

type GenerateUploadTokenResponseBodyData struct {
	Key            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ServerURL      *string `json:"ServerURL,omitempty" xml:"ServerURL,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	XAmzAlgorithm  *string `json:"X-Amz-Algorithm,omitempty" xml:"X-Amz-Algorithm,omitempty"`
	XAmzCredential *string `json:"X-Amz-Credential,omitempty" xml:"X-Amz-Credential,omitempty"`
	XAmzDate       *string `json:"X-Amz-Date,omitempty" xml:"X-Amz-Date,omitempty"`
	XAmzSignature  *string `json:"X-Amz-Signature,omitempty" xml:"X-Amz-Signature,omitempty"`
}

func (s GenerateUploadTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadTokenResponseBodyData) SetKey(v string) *GenerateUploadTokenResponseBodyData {
	s.Key = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetOssAccessKeyId(v string) *GenerateUploadTokenResponseBodyData {
	s.OssAccessKeyId = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetPolicy(v string) *GenerateUploadTokenResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetServerURL(v string) *GenerateUploadTokenResponseBodyData {
	s.ServerURL = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetSignature(v string) *GenerateUploadTokenResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetXAmzAlgorithm(v string) *GenerateUploadTokenResponseBodyData {
	s.XAmzAlgorithm = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetXAmzCredential(v string) *GenerateUploadTokenResponseBodyData {
	s.XAmzCredential = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetXAmzDate(v string) *GenerateUploadTokenResponseBodyData {
	s.XAmzDate = &v
	return s
}

func (s *GenerateUploadTokenResponseBodyData) SetXAmzSignature(v string) *GenerateUploadTokenResponseBodyData {
	s.XAmzSignature = &v
	return s
}

type GenerateUploadTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadTokenResponse) SetHeaders(v map[string]*string) *GenerateUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadTokenResponse) SetStatusCode(v int32) *GenerateUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadTokenResponse) SetBody(v *GenerateUploadTokenResponseBody) *GenerateUploadTokenResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetAppId(v string) *GetAppRequest {
	s.AppId = &v
	return s
}

func (s *GetAppRequest) SetSource(v string) *GetAppRequest {
	s.Source = &v
	return s
}

type GetAppResponseBody struct {
	Data      *GetAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetData(v *GetAppResponseBodyData) *GetAppResponseBody {
	s.Data = v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

type GetAppResponseBodyData struct {
	AppId           *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus       *string                             `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType         *string                             `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Categories      []*GetAppResponseBodyDataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime      *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon            *string                             `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate      *bool                               `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime    *string                             `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId    *string                             `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime    *string                             `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PlatformVersion *string                             `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	SchemaVersion   *string                             `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source          *string                             `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyData) SetAppId(v string) *GetAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAppResponseBodyData) SetAppName(v string) *GetAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyData) SetAppStatus(v string) *GetAppResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *GetAppResponseBodyData) SetAppType(v string) *GetAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBodyData) SetCategories(v []*GetAppResponseBodyDataCategories) *GetAppResponseBodyData {
	s.Categories = v
	return s
}

func (s *GetAppResponseBodyData) SetCreateTime(v string) *GetAppResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyData) SetDescription(v string) *GetAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAppResponseBodyData) SetIcon(v string) *GetAppResponseBodyData {
	s.Icon = &v
	return s
}

func (s *GetAppResponseBodyData) SetIsTemplate(v bool) *GetAppResponseBodyData {
	s.IsTemplate = &v
	return s
}

func (s *GetAppResponseBodyData) SetLastEditTime(v string) *GetAppResponseBodyData {
	s.LastEditTime = &v
	return s
}

func (s *GetAppResponseBodyData) SetMainModuleId(v string) *GetAppResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *GetAppResponseBodyData) SetModifiedTime(v string) *GetAppResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetAppResponseBodyData) SetPlatformVersion(v string) *GetAppResponseBodyData {
	s.PlatformVersion = &v
	return s
}

func (s *GetAppResponseBodyData) SetSchemaVersion(v string) *GetAppResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *GetAppResponseBodyData) SetSource(v string) *GetAppResponseBodyData {
	s.Source = &v
	return s
}

type GetAppResponseBodyDataCategories struct {
	CategoryId       *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName     *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s GetAppResponseBodyDataCategories) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyDataCategories) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyDataCategories) SetCategoryId(v string) *GetAppResponseBodyDataCategories {
	s.CategoryId = &v
	return s
}

func (s *GetAppResponseBodyDataCategories) SetCategoryName(v string) *GetAppResponseBodyDataCategories {
	s.CategoryName = &v
	return s
}

func (s *GetAppResponseBodyDataCategories) SetParentCategoryId(v string) *GetAppResponseBodyDataCategories {
	s.ParentCategoryId = &v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetAppModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetAppModelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppModelRequest) GoString() string {
	return s.String()
}

func (s *GetAppModelRequest) SetAppId(v string) *GetAppModelRequest {
	s.AppId = &v
	return s
}

func (s *GetAppModelRequest) SetSchemaVersion(v string) *GetAppModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *GetAppModelRequest) SetSource(v string) *GetAppModelRequest {
	s.Source = &v
	return s
}

func (s *GetAppModelRequest) SetSubType(v string) *GetAppModelRequest {
	s.SubType = &v
	return s
}

type GetAppModelResponseBody struct {
	Data      *GetAppModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppModelResponseBody) SetData(v *GetAppModelResponseBodyData) *GetAppModelResponseBody {
	s.Data = v
	return s
}

func (s *GetAppModelResponseBody) SetRequestId(v string) *GetAppModelResponseBody {
	s.RequestId = &v
	return s
}

type GetAppModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetAppModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppModelResponseBodyData) SetAppId(v string) *GetAppModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetAttributes(v []map[string]*string) *GetAppModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *GetAppModelResponseBodyData) SetContent(v map[string]*string) *GetAppModelResponseBodyData {
	s.Content = v
	return s
}

func (s *GetAppModelResponseBodyData) SetCreateTime(v string) *GetAppModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetDescription(v string) *GetAppModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetId(v string) *GetAppModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetLinkModelId(v string) *GetAppModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetLinkModuleId(v string) *GetAppModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetLinked(v bool) *GetAppModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModelDigest(v string) *GetAppModelResponseBodyData {
	s.ModelDigest = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModelId(v string) *GetAppModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModelName(v string) *GetAppModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModelStatus(v string) *GetAppModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModelType(v string) *GetAppModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModifiedTime(v string) *GetAppModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetModuleId(v string) *GetAppModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetProps(v map[string]*string) *GetAppModelResponseBodyData {
	s.Props = v
	return s
}

func (s *GetAppModelResponseBodyData) SetRevision(v int32) *GetAppModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetSchemaVersion(v string) *GetAppModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetSubType(v string) *GetAppModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *GetAppModelResponseBodyData) SetVisibility(v string) *GetAppModelResponseBodyData {
	s.Visibility = &v
	return s
}

type GetAppModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppModelResponse) GoString() string {
	return s.String()
}

func (s *GetAppModelResponse) SetHeaders(v map[string]*string) *GetAppModelResponse {
	s.Headers = v
	return s
}

func (s *GetAppModelResponse) SetStatusCode(v int32) *GetAppModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppModelResponse) SetBody(v *GetAppModelResponseBody) *GetAppModelResponse {
	s.Body = v
	return s
}

type GetAppSecretRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *GetAppSecretRequest) SetAppId(v string) *GetAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *GetAppSecretRequest) SetSource(v string) *GetAppSecretRequest {
	s.Source = &v
	return s
}

type GetAppSecretResponseBody struct {
	Data      *GetAppSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBody) SetData(v *GetAppSecretResponseBodyData) *GetAppSecretResponseBody {
	s.Data = v
	return s
}

func (s *GetAppSecretResponseBody) SetRequestId(v string) *GetAppSecretResponseBody {
	s.RequestId = &v
	return s
}

type GetAppSecretResponseBodyData struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s GetAppSecretResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBodyData) SetAppId(v string) *GetAppSecretResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAppSecretResponseBodyData) SetAppSecret(v string) *GetAppSecretResponseBodyData {
	s.AppSecret = &v
	return s
}

type GetAppSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponse) SetHeaders(v map[string]*string) *GetAppSecretResponse {
	s.Headers = v
	return s
}

func (s *GetAppSecretResponse) SetStatusCode(v int32) *GetAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSecretResponse) SetBody(v *GetAppSecretResponseBody) *GetAppSecretResponse {
	s.Body = v
	return s
}

type GetArtifactRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) SetAppId(v string) *GetArtifactRequest {
	s.AppId = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactId(v string) *GetArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactRequest) SetSource(v string) *GetArtifactRequest {
	s.Source = &v
	return s
}

type GetArtifactResponseBody struct {
	Data      *GetArtifactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBody) SetData(v *GetArtifactResponseBodyData) *GetArtifactResponseBody {
	s.Data = v
	return s
}

func (s *GetArtifactResponseBody) SetRequestId(v string) *GetArtifactResponseBody {
	s.RequestId = &v
	return s
}

type GetArtifactResponseBodyData struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ArtifactId   *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Available    *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetArtifactResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBodyData) SetAppId(v string) *GetArtifactResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetArtifactId(v string) *GetArtifactResponseBodyData {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetArtifactType(v string) *GetArtifactResponseBodyData {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetAvailable(v bool) *GetArtifactResponseBodyData {
	s.Available = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetCreateTime(v string) *GetArtifactResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetModifiedTime(v string) *GetArtifactResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetArtifactResponseBodyData) SetUrl(v string) *GetArtifactResponseBodyData {
	s.Url = &v
	return s
}

type GetArtifactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactResponse) SetHeaders(v map[string]*string) *GetArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactResponse) SetStatusCode(v int32) *GetArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactResponse) SetBody(v *GetArtifactResponseBody) *GetArtifactResponse {
	s.Body = v
	return s
}

type GetCommitRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId      *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommitRequest) GoString() string {
	return s.String()
}

func (s *GetCommitRequest) SetAppId(v string) *GetCommitRequest {
	s.AppId = &v
	return s
}

func (s *GetCommitRequest) SetCommitId(v string) *GetCommitRequest {
	s.CommitId = &v
	return s
}

func (s *GetCommitRequest) SetModuleId(v string) *GetCommitRequest {
	s.ModuleId = &v
	return s
}

func (s *GetCommitRequest) SetSchemaVersion(v string) *GetCommitRequest {
	s.SchemaVersion = &v
	return s
}

func (s *GetCommitRequest) SetSource(v string) *GetCommitRequest {
	s.Source = &v
	return s
}

type GetCommitResponseBody struct {
	Data      *GetCommitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommitResponseBody) SetData(v *GetCommitResponseBodyData) *GetCommitResponseBody {
	s.Data = v
	return s
}

func (s *GetCommitResponseBody) SetRequestId(v string) *GetCommitResponseBody {
	s.RequestId = &v
	return s
}

type GetCommitResponseBodyData struct {
	AppId              *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitDigest       *string              `json:"CommitDigest,omitempty" xml:"CommitDigest,omitempty"`
	CommitId           *string              `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitLog          *string              `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string              `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	CreateTime         *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MainModuleCommitId *string              `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	MainModuleId       *string              `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModelDataPath      *string              `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModelDigest        map[string]*string   `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelPack          []interface{}        `json:"ModelPack,omitempty" xml:"ModelPack,omitempty" type:"Repeated"`
	ModifiedTime       *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId           *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDataPath   *string              `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
	ResourceDigest     map[string]*string   `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourcePack       []map[string]*string `json:"ResourcePack,omitempty" xml:"ResourcePack,omitempty" type:"Repeated"`
	RollbackToCommitId *string              `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string              `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s GetCommitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCommitResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommitResponseBodyData) SetAppId(v string) *GetCommitResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetCommitDigest(v string) *GetCommitResponseBodyData {
	s.CommitDigest = &v
	return s
}

func (s *GetCommitResponseBodyData) SetCommitId(v string) *GetCommitResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetCommitLog(v string) *GetCommitResponseBodyData {
	s.CommitLog = &v
	return s
}

func (s *GetCommitResponseBodyData) SetCommitType(v string) *GetCommitResponseBodyData {
	s.CommitType = &v
	return s
}

func (s *GetCommitResponseBodyData) SetCreateTime(v string) *GetCommitResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCommitResponseBodyData) SetMainModuleCommitId(v string) *GetCommitResponseBodyData {
	s.MainModuleCommitId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetMainModuleId(v string) *GetCommitResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetModelDataPath(v string) *GetCommitResponseBodyData {
	s.ModelDataPath = &v
	return s
}

func (s *GetCommitResponseBodyData) SetModelDigest(v map[string]*string) *GetCommitResponseBodyData {
	s.ModelDigest = v
	return s
}

func (s *GetCommitResponseBodyData) SetModelPack(v []interface{}) *GetCommitResponseBodyData {
	s.ModelPack = v
	return s
}

func (s *GetCommitResponseBodyData) SetModifiedTime(v string) *GetCommitResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetCommitResponseBodyData) SetModuleId(v string) *GetCommitResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetResourceDataPath(v string) *GetCommitResponseBodyData {
	s.ResourceDataPath = &v
	return s
}

func (s *GetCommitResponseBodyData) SetResourceDigest(v map[string]*string) *GetCommitResponseBodyData {
	s.ResourceDigest = v
	return s
}

func (s *GetCommitResponseBodyData) SetResourcePack(v []map[string]*string) *GetCommitResponseBodyData {
	s.ResourcePack = v
	return s
}

func (s *GetCommitResponseBodyData) SetRollbackToCommitId(v string) *GetCommitResponseBodyData {
	s.RollbackToCommitId = &v
	return s
}

func (s *GetCommitResponseBodyData) SetRollbackType(v string) *GetCommitResponseBodyData {
	s.RollbackType = &v
	return s
}

func (s *GetCommitResponseBodyData) SetSchemaVersion(v string) *GetCommitResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type GetCommitResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommitResponse) GoString() string {
	return s.String()
}

func (s *GetCommitResponse) SetHeaders(v map[string]*string) *GetCommitResponse {
	s.Headers = v
	return s
}

func (s *GetCommitResponse) SetStatusCode(v int32) *GetCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommitResponse) SetBody(v *GetCommitResponseBody) *GetCommitResponse {
	s.Body = v
	return s
}

type GetDefaultAppUserRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetDefaultAppUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAppUserRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultAppUserRequest) SetAppId(v string) *GetDefaultAppUserRequest {
	s.AppId = &v
	return s
}

func (s *GetDefaultAppUserRequest) SetEnvId(v string) *GetDefaultAppUserRequest {
	s.EnvId = &v
	return s
}

func (s *GetDefaultAppUserRequest) SetSource(v string) *GetDefaultAppUserRequest {
	s.Source = &v
	return s
}

type GetDefaultAppUserResponseBody struct {
	Data      *GetDefaultAppUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDefaultAppUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAppUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultAppUserResponseBody) SetData(v *GetDefaultAppUserResponseBodyData) *GetDefaultAppUserResponseBody {
	s.Data = v
	return s
}

func (s *GetDefaultAppUserResponseBody) SetRequestId(v string) *GetDefaultAppUserResponseBody {
	s.RequestId = &v
	return s
}

type GetDefaultAppUserResponseBodyData struct {
	HasPassword *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDefaultAppUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAppUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDefaultAppUserResponseBodyData) SetHasPassword(v bool) *GetDefaultAppUserResponseBodyData {
	s.HasPassword = &v
	return s
}

func (s *GetDefaultAppUserResponseBodyData) SetUserName(v string) *GetDefaultAppUserResponseBodyData {
	s.UserName = &v
	return s
}

type GetDefaultAppUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultAppUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultAppUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAppUserResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultAppUserResponse) SetHeaders(v map[string]*string) *GetDefaultAppUserResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultAppUserResponse) SetStatusCode(v int32) *GetDefaultAppUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultAppUserResponse) SetBody(v *GetDefaultAppUserResponseBody) *GetDefaultAppUserResponse {
	s.Body = v
	return s
}

type GetDomainCnameRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetDomainCnameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *GetDomainCnameRequest) SetAppId(v string) *GetDomainCnameRequest {
	s.AppId = &v
	return s
}

func (s *GetDomainCnameRequest) SetDomain(v string) *GetDomainCnameRequest {
	s.Domain = &v
	return s
}

func (s *GetDomainCnameRequest) SetDomainType(v string) *GetDomainCnameRequest {
	s.DomainType = &v
	return s
}

func (s *GetDomainCnameRequest) SetEnvId(v string) *GetDomainCnameRequest {
	s.EnvId = &v
	return s
}

func (s *GetDomainCnameRequest) SetSource(v string) *GetDomainCnameRequest {
	s.Source = &v
	return s
}

type GetDomainCnameResponseBody struct {
	Data      *GetDomainCnameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainCnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainCnameResponseBody) SetData(v *GetDomainCnameResponseBodyData) *GetDomainCnameResponseBody {
	s.Data = v
	return s
}

func (s *GetDomainCnameResponseBody) SetRequestId(v string) *GetDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainCnameResponseBodyData struct {
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
}

func (s GetDomainCnameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCnameResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDomainCnameResponseBodyData) SetCname(v string) *GetDomainCnameResponseBodyData {
	s.Cname = &v
	return s
}

type GetDomainCnameResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainCnameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *GetDomainCnameResponse) SetHeaders(v map[string]*string) *GetDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *GetDomainCnameResponse) SetStatusCode(v int32) *GetDomainCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainCnameResponse) SetBody(v *GetDomainCnameResponseBody) *GetDomainCnameResponse {
	s.Body = v
	return s
}

type GetDomainOverviewRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetDomainOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetDomainOverviewRequest) SetAppId(v string) *GetDomainOverviewRequest {
	s.AppId = &v
	return s
}

func (s *GetDomainOverviewRequest) SetDomain(v string) *GetDomainOverviewRequest {
	s.Domain = &v
	return s
}

func (s *GetDomainOverviewRequest) SetEnvId(v string) *GetDomainOverviewRequest {
	s.EnvId = &v
	return s
}

func (s *GetDomainOverviewRequest) SetSource(v string) *GetDomainOverviewRequest {
	s.Source = &v
	return s
}

type GetDomainOverviewResponseBody struct {
	Data      *GetDomainOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainOverviewResponseBody) SetData(v *GetDomainOverviewResponseBodyData) *GetDomainOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetDomainOverviewResponseBody) SetRequestId(v string) *GetDomainOverviewResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainOverviewResponseBodyData struct {
	AppId           *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Applied         *bool              `json:"Applied,omitempty" xml:"Applied,omitempty"`
	Certificate     map[string]*string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Cname           *string            `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Deleted         *bool              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Domain          *string            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType      *string            `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId           *string            `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Path            *string            `json:"Path,omitempty" xml:"Path,omitempty"`
	WithCertificate *bool              `json:"WithCertificate,omitempty" xml:"WithCertificate,omitempty"`
}

func (s GetDomainOverviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDomainOverviewResponseBodyData) SetAppId(v string) *GetDomainOverviewResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetApplied(v bool) *GetDomainOverviewResponseBodyData {
	s.Applied = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetCertificate(v map[string]*string) *GetDomainOverviewResponseBodyData {
	s.Certificate = v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetCname(v string) *GetDomainOverviewResponseBodyData {
	s.Cname = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetDeleted(v bool) *GetDomainOverviewResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetDomain(v string) *GetDomainOverviewResponseBodyData {
	s.Domain = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetDomainType(v string) *GetDomainOverviewResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetEnvId(v string) *GetDomainOverviewResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetPath(v string) *GetDomainOverviewResponseBodyData {
	s.Path = &v
	return s
}

func (s *GetDomainOverviewResponseBodyData) SetWithCertificate(v bool) *GetDomainOverviewResponseBodyData {
	s.WithCertificate = &v
	return s
}

type GetDomainOverviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetDomainOverviewResponse) SetHeaders(v map[string]*string) *GetDomainOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetDomainOverviewResponse) SetStatusCode(v int32) *GetDomainOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainOverviewResponse) SetBody(v *GetDomainOverviewResponseBody) *GetDomainOverviewResponse {
	s.Body = v
	return s
}

type GetEnvironmentRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *GetEnvironmentRequest) SetAppId(v string) *GetEnvironmentRequest {
	s.AppId = &v
	return s
}

func (s *GetEnvironmentRequest) SetEnvId(v string) *GetEnvironmentRequest {
	s.EnvId = &v
	return s
}

func (s *GetEnvironmentRequest) SetSource(v string) *GetEnvironmentRequest {
	s.Source = &v
	return s
}

type GetEnvironmentResponseBody struct {
	Data      *GetEnvironmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBody) SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentResponseBody) SetRequestId(v string) *GetEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

type GetEnvironmentResponseBodyData struct {
	AccountOpsEndpoint *string `json:"AccountOpsEndpoint,omitempty" xml:"AccountOpsEndpoint,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentPublishId   *string `json:"CurrentPublishId,omitempty" xml:"CurrentPublishId,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EnvId              *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType            *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	ModifiedTime       *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishingId       *string `json:"PublishingId,omitempty" xml:"PublishingId,omitempty"`
}

func (s GetEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyData) SetAccountOpsEndpoint(v string) *GetEnvironmentResponseBodyData {
	s.AccountOpsEndpoint = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetAppId(v string) *GetEnvironmentResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCreateTime(v string) *GetEnvironmentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCurrentPublishId(v string) *GetEnvironmentResponseBodyData {
	s.CurrentPublishId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetEndpoint(v string) *GetEnvironmentResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetEnvId(v string) *GetEnvironmentResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetEnvType(v string) *GetEnvironmentResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetModifiedTime(v string) *GetEnvironmentResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetPublishingId(v string) *GetEnvironmentResponseBodyData {
	s.PublishingId = &v
	return s
}

type GetEnvironmentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetStatusCode(v int32) *GetEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type GetHistoryStatsRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHistoryStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryStatsRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryStatsRequest) SetAppId(v string) *GetHistoryStatsRequest {
	s.AppId = &v
	return s
}

func (s *GetHistoryStatsRequest) SetEndDate(v string) *GetHistoryStatsRequest {
	s.EndDate = &v
	return s
}

func (s *GetHistoryStatsRequest) SetSource(v string) *GetHistoryStatsRequest {
	s.Source = &v
	return s
}

func (s *GetHistoryStatsRequest) SetStartDate(v string) *GetHistoryStatsRequest {
	s.StartDate = &v
	return s
}

type GetHistoryStatsResponseBody struct {
	Data      *GetHistoryStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoryStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryStatsResponseBody) SetData(v *GetHistoryStatsResponseBodyData) *GetHistoryStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryStatsResponseBody) SetRequestId(v string) *GetHistoryStatsResponseBody {
	s.RequestId = &v
	return s
}

type GetHistoryStatsResponseBodyData struct {
	HistoryPv map[string]*string `json:"HistoryPv,omitempty" xml:"HistoryPv,omitempty"`
	HistoryUv map[string]*string `json:"HistoryUv,omitempty" xml:"HistoryUv,omitempty"`
}

func (s GetHistoryStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryStatsResponseBodyData) SetHistoryPv(v map[string]*string) *GetHistoryStatsResponseBodyData {
	s.HistoryPv = v
	return s
}

func (s *GetHistoryStatsResponseBodyData) SetHistoryUv(v map[string]*string) *GetHistoryStatsResponseBodyData {
	s.HistoryUv = v
	return s
}

type GetHistoryStatsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryStatsResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryStatsResponse) SetHeaders(v map[string]*string) *GetHistoryStatsResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryStatsResponse) SetStatusCode(v int32) *GetHistoryStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryStatsResponse) SetBody(v *GetHistoryStatsResponseBody) *GetHistoryStatsResponse {
	s.Body = v
	return s
}

type GetLatestCommitRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetLatestCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLatestCommitRequest) GoString() string {
	return s.String()
}

func (s *GetLatestCommitRequest) SetAppId(v string) *GetLatestCommitRequest {
	s.AppId = &v
	return s
}

func (s *GetLatestCommitRequest) SetModuleId(v string) *GetLatestCommitRequest {
	s.ModuleId = &v
	return s
}

func (s *GetLatestCommitRequest) SetSource(v string) *GetLatestCommitRequest {
	s.Source = &v
	return s
}

type GetLatestCommitResponseBody struct {
	Data      *GetLatestCommitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLatestCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestCommitResponseBody) SetData(v *GetLatestCommitResponseBodyData) *GetLatestCommitResponseBody {
	s.Data = v
	return s
}

func (s *GetLatestCommitResponseBody) SetRequestId(v string) *GetLatestCommitResponseBody {
	s.RequestId = &v
	return s
}

type GetLatestCommitResponseBodyData struct {
	AppId              *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId           *string            `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitLog          *string            `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string            `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	CreateTime         *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MainModuleCommitId *string            `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	MainModuleId       *string            `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModelDataPath      *string            `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModifiedTime       *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId           *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDataPath   *string            `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
	ResourceDigest     map[string]*string `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	RollbackToCommitId *string            `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string            `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s GetLatestCommitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLatestCommitResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLatestCommitResponseBodyData) SetAppId(v string) *GetLatestCommitResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetCommitId(v string) *GetLatestCommitResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetCommitLog(v string) *GetLatestCommitResponseBodyData {
	s.CommitLog = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetCommitType(v string) *GetLatestCommitResponseBodyData {
	s.CommitType = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetCreateTime(v string) *GetLatestCommitResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetMainModuleCommitId(v string) *GetLatestCommitResponseBodyData {
	s.MainModuleCommitId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetMainModuleId(v string) *GetLatestCommitResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetModelDataPath(v string) *GetLatestCommitResponseBodyData {
	s.ModelDataPath = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetModifiedTime(v string) *GetLatestCommitResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetModuleId(v string) *GetLatestCommitResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetResourceDataPath(v string) *GetLatestCommitResponseBodyData {
	s.ResourceDataPath = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetResourceDigest(v map[string]*string) *GetLatestCommitResponseBodyData {
	s.ResourceDigest = v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetRollbackToCommitId(v string) *GetLatestCommitResponseBodyData {
	s.RollbackToCommitId = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetRollbackType(v string) *GetLatestCommitResponseBodyData {
	s.RollbackType = &v
	return s
}

func (s *GetLatestCommitResponseBodyData) SetSchemaVersion(v string) *GetLatestCommitResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type GetLatestCommitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLatestCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLatestCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestCommitResponse) GoString() string {
	return s.String()
}

func (s *GetLatestCommitResponse) SetHeaders(v map[string]*string) *GetLatestCommitResponse {
	s.Headers = v
	return s
}

func (s *GetLatestCommitResponse) SetStatusCode(v int32) *GetLatestCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestCommitResponse) SetBody(v *GetLatestCommitResponseBody) *GetLatestCommitResponse {
	s.Body = v
	return s
}

type GetModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetModelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelRequest) GoString() string {
	return s.String()
}

func (s *GetModelRequest) SetAppId(v string) *GetModelRequest {
	s.AppId = &v
	return s
}

func (s *GetModelRequest) SetModelId(v string) *GetModelRequest {
	s.ModelId = &v
	return s
}

func (s *GetModelRequest) SetModuleId(v string) *GetModelRequest {
	s.ModuleId = &v
	return s
}

func (s *GetModelRequest) SetSchemaVersion(v string) *GetModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *GetModelRequest) SetSource(v string) *GetModelRequest {
	s.Source = &v
	return s
}

type GetModelResponseBody struct {
	Data      *GetModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) SetData(v *GetModelResponseBodyData) *GetModelResponseBody {
	s.Data = v
	return s
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

type GetModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyData) SetAppId(v string) *GetModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetModelResponseBodyData) SetAttributes(v []map[string]*string) *GetModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *GetModelResponseBodyData) SetContent(v map[string]*string) *GetModelResponseBodyData {
	s.Content = v
	return s
}

func (s *GetModelResponseBodyData) SetCreateTime(v string) *GetModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetModelResponseBodyData) SetDescription(v string) *GetModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelResponseBodyData) SetId(v string) *GetModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetModelResponseBodyData) SetLinkModelId(v string) *GetModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *GetModelResponseBodyData) SetLinkModuleId(v string) *GetModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *GetModelResponseBodyData) SetLinked(v bool) *GetModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelId(v string) *GetModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelName(v string) *GetModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelStatus(v string) *GetModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelType(v string) *GetModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *GetModelResponseBodyData) SetModifiedTime(v string) *GetModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetModelResponseBodyData) SetModuleId(v string) *GetModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetModelResponseBodyData) SetProps(v map[string]*string) *GetModelResponseBodyData {
	s.Props = v
	return s
}

func (s *GetModelResponseBodyData) SetRevision(v int32) *GetModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *GetModelResponseBodyData) SetSchemaVersion(v string) *GetModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *GetModelResponseBodyData) SetSubType(v string) *GetModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *GetModelResponseBodyData) SetVisibility(v string) *GetModelResponseBodyData {
	s.Visibility = &v
	return s
}

type GetModelResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponse) GoString() string {
	return s.String()
}

func (s *GetModelResponse) SetHeaders(v map[string]*string) *GetModelResponse {
	s.Headers = v
	return s
}

func (s *GetModelResponse) SetStatusCode(v int32) *GetModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelResponse) SetBody(v *GetModelResponseBody) *GetModelResponse {
	s.Body = v
	return s
}

type GetModuleRequest struct {
	ModuleId   *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleType *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModuleRequest) GoString() string {
	return s.String()
}

func (s *GetModuleRequest) SetModuleId(v string) *GetModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *GetModuleRequest) SetModuleType(v string) *GetModuleRequest {
	s.ModuleType = &v
	return s
}

func (s *GetModuleRequest) SetSource(v string) *GetModuleRequest {
	s.Source = &v
	return s
}

type GetModuleResponseBody struct {
	Data      *GetModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBody) SetData(v *GetModuleResponseBodyData) *GetModuleResponseBody {
	s.Data = v
	return s
}

func (s *GetModuleResponseBody) SetRequestId(v string) *GetModuleResponseBody {
	s.RequestId = &v
	return s
}

type GetModuleResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PlatformVersion        *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
}

func (s GetModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBodyData) SetCreateTime(v string) *GetModuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetModuleResponseBodyData) SetDescription(v string) *GetModuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModuleResponseBodyData) SetIcon(v string) *GetModuleResponseBodyData {
	s.Icon = &v
	return s
}

func (s *GetModuleResponseBodyData) SetLatestPublishedCommit(v string) *GetModuleResponseBodyData {
	s.LatestPublishedCommit = &v
	return s
}

func (s *GetModuleResponseBodyData) SetLatestPublishedVersion(v string) *GetModuleResponseBodyData {
	s.LatestPublishedVersion = &v
	return s
}

func (s *GetModuleResponseBodyData) SetMinimumPlatformVersion(v string) *GetModuleResponseBodyData {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *GetModuleResponseBodyData) SetModifiedTime(v string) *GetModuleResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetModuleResponseBodyData) SetModuleId(v string) *GetModuleResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetModuleResponseBodyData) SetModuleName(v string) *GetModuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *GetModuleResponseBodyData) SetOwnerAppId(v string) *GetModuleResponseBodyData {
	s.OwnerAppId = &v
	return s
}

func (s *GetModuleResponseBodyData) SetOwnerUserId(v string) *GetModuleResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *GetModuleResponseBodyData) SetPlatform(v string) *GetModuleResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetModuleResponseBodyData) SetPlatformVersion(v string) *GetModuleResponseBodyData {
	s.PlatformVersion = &v
	return s
}

type GetModuleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponse) GoString() string {
	return s.String()
}

func (s *GetModuleResponse) SetHeaders(v map[string]*string) *GetModuleResponse {
	s.Headers = v
	return s
}

func (s *GetModuleResponse) SetStatusCode(v int32) *GetModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleResponse) SetBody(v *GetModuleResponseBody) *GetModuleResponse {
	s.Body = v
	return s
}

type GetPublishRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PublishId *string `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetPublishRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublishRequest) GoString() string {
	return s.String()
}

func (s *GetPublishRequest) SetAppId(v string) *GetPublishRequest {
	s.AppId = &v
	return s
}

func (s *GetPublishRequest) SetPublishId(v string) *GetPublishRequest {
	s.PublishId = &v
	return s
}

func (s *GetPublishRequest) SetSource(v string) *GetPublishRequest {
	s.Source = &v
	return s
}

type GetPublishResponseBody struct {
	Data      *GetPublishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublishResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishResponseBody) SetData(v *GetPublishResponseBodyData) *GetPublishResponseBody {
	s.Data = v
	return s
}

func (s *GetPublishResponseBody) SetRequestId(v string) *GetPublishResponseBody {
	s.RequestId = &v
	return s
}

type GetPublishResponseBodyData struct {
	AppId          *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId       *string              `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CompletionTime *string              `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreateTime     *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvId          *string              `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ModifiedTime   *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishId      *string              `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	PublishStatus  *string              `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishType    *string              `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Reason         *string              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime      *string              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubTasks       []map[string]*string `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	VersionNumber  *string              `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s GetPublishResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPublishResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublishResponseBodyData) SetAppId(v string) *GetPublishResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetPublishResponseBodyData) SetCommitId(v string) *GetPublishResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *GetPublishResponseBodyData) SetCompletionTime(v string) *GetPublishResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *GetPublishResponseBodyData) SetCreateTime(v string) *GetPublishResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetPublishResponseBodyData) SetDescription(v string) *GetPublishResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetPublishResponseBodyData) SetEnvId(v string) *GetPublishResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *GetPublishResponseBodyData) SetModifiedTime(v string) *GetPublishResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetPublishResponseBodyData) SetPublishId(v string) *GetPublishResponseBodyData {
	s.PublishId = &v
	return s
}

func (s *GetPublishResponseBodyData) SetPublishStatus(v string) *GetPublishResponseBodyData {
	s.PublishStatus = &v
	return s
}

func (s *GetPublishResponseBodyData) SetPublishType(v string) *GetPublishResponseBodyData {
	s.PublishType = &v
	return s
}

func (s *GetPublishResponseBodyData) SetReason(v string) *GetPublishResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetPublishResponseBodyData) SetStartTime(v string) *GetPublishResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetPublishResponseBodyData) SetSubTasks(v []map[string]*string) *GetPublishResponseBodyData {
	s.SubTasks = v
	return s
}

func (s *GetPublishResponseBodyData) SetVersionNumber(v string) *GetPublishResponseBodyData {
	s.VersionNumber = &v
	return s
}

type GetPublishResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublishResponse) GoString() string {
	return s.String()
}

func (s *GetPublishResponse) SetHeaders(v map[string]*string) *GetPublishResponse {
	s.Headers = v
	return s
}

func (s *GetPublishResponse) SetStatusCode(v int32) *GetPublishResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishResponse) SetBody(v *GetPublishResponseBody) *GetPublishResponse {
	s.Body = v
	return s
}

type GetRealtimeStatsRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetRealtimeStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeStatsRequest) SetAppId(v string) *GetRealtimeStatsRequest {
	s.AppId = &v
	return s
}

func (s *GetRealtimeStatsRequest) SetSource(v string) *GetRealtimeStatsRequest {
	s.Source = &v
	return s
}

type GetRealtimeStatsResponseBody struct {
	Data      *GetRealtimeStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeStatsResponseBody) SetData(v *GetRealtimeStatsResponseBodyData) *GetRealtimeStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeStatsResponseBody) SetRequestId(v string) *GetRealtimeStatsResponseBody {
	s.RequestId = &v
	return s
}

type GetRealtimeStatsResponseBodyData struct {
	TodayPvCount map[string]*string `json:"TodayPvCount,omitempty" xml:"TodayPvCount,omitempty"`
	TodayUvCount map[string]*string `json:"TodayUvCount,omitempty" xml:"TodayUvCount,omitempty"`
	TotalPvCount map[string]*string `json:"TotalPvCount,omitempty" xml:"TotalPvCount,omitempty"`
	TotalUvCount map[string]*string `json:"TotalUvCount,omitempty" xml:"TotalUvCount,omitempty"`
}

func (s GetRealtimeStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeStatsResponseBodyData) SetTodayPvCount(v map[string]*string) *GetRealtimeStatsResponseBodyData {
	s.TodayPvCount = v
	return s
}

func (s *GetRealtimeStatsResponseBodyData) SetTodayUvCount(v map[string]*string) *GetRealtimeStatsResponseBodyData {
	s.TodayUvCount = v
	return s
}

func (s *GetRealtimeStatsResponseBodyData) SetTotalPvCount(v map[string]*string) *GetRealtimeStatsResponseBodyData {
	s.TotalPvCount = v
	return s
}

func (s *GetRealtimeStatsResponseBodyData) SetTotalUvCount(v map[string]*string) *GetRealtimeStatsResponseBodyData {
	s.TotalUvCount = v
	return s
}

type GetRealtimeStatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeStatsResponse) SetHeaders(v map[string]*string) *GetRealtimeStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeStatsResponse) SetStatusCode(v int32) *GetRealtimeStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeStatsResponse) SetBody(v *GetRealtimeStatsResponseBody) *GetRealtimeStatsResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ImageProcessParameter *string `json:"ImageProcessParameter,omitempty" xml:"ImageProcessParameter,omitempty"`
	ModuleId              *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetAppId(v string) *GetResourceRequest {
	s.AppId = &v
	return s
}

func (s *GetResourceRequest) SetImageProcessParameter(v string) *GetResourceRequest {
	s.ImageProcessParameter = &v
	return s
}

func (s *GetResourceRequest) SetModuleId(v string) *GetResourceRequest {
	s.ModuleId = &v
	return s
}

func (s *GetResourceRequest) SetResourceId(v string) *GetResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceRequest) SetSource(v string) *GetResourceRequest {
	s.Source = &v
	return s
}

type GetResourceResponseBody struct {
	Data      *GetResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetData(v *GetResourceResponseBodyData) *GetResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceResponseBodyData struct {
	AppId          *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime   *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId       *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDigest *string            `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourceId     *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision       *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion  *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s GetResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyData) SetAppId(v string) *GetResourceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetResourceResponseBodyData) SetContent(v map[string]*string) *GetResourceResponseBodyData {
	s.Content = v
	return s
}

func (s *GetResourceResponseBodyData) SetCreateTime(v string) *GetResourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetResourceResponseBodyData) SetDescription(v string) *GetResourceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetResourceResponseBodyData) SetModifiedTime(v string) *GetResourceResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetResourceResponseBodyData) SetModuleId(v string) *GetResourceResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceDigest(v string) *GetResourceResponseBodyData {
	s.ResourceDigest = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceId(v string) *GetResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceName(v string) *GetResourceResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceType(v string) *GetResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetResourceResponseBodyData) SetRevision(v int32) *GetResourceResponseBodyData {
	s.Revision = &v
	return s
}

func (s *GetResourceResponseBodyData) SetSchemaVersion(v string) *GetResourceResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetSource(v string) *GetUserRequest {
	s.Source = &v
	return s
}

type GetUserResponseBody struct {
	Data      *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyData struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PlatformVersion *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	UserSecret      *string `json:"UserSecret,omitempty" xml:"UserSecret,omitempty"`
	UserStatus      *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	UserType        *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetCreateTime(v string) *GetUserResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyData) SetDescription(v string) *GetUserResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyData) SetModifiedTime(v string) *GetUserResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetUserResponseBodyData) SetPlatformVersion(v string) *GetUserResponseBodyData {
	s.PlatformVersion = &v
	return s
}

func (s *GetUserResponseBodyData) SetUserSecret(v string) *GetUserResponseBodyData {
	s.UserSecret = &v
	return s
}

func (s *GetUserResponseBodyData) SetUserStatus(v string) *GetUserResponseBodyData {
	s.UserStatus = &v
	return s
}

func (s *GetUserResponseBodyData) SetUserType(v string) *GetUserResponseBodyData {
	s.UserType = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListAppModulesRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Recursive *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListAppModulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppModulesRequest) GoString() string {
	return s.String()
}

func (s *ListAppModulesRequest) SetAppId(v string) *ListAppModulesRequest {
	s.AppId = &v
	return s
}

func (s *ListAppModulesRequest) SetRecursive(v bool) *ListAppModulesRequest {
	s.Recursive = &v
	return s
}

func (s *ListAppModulesRequest) SetSource(v string) *ListAppModulesRequest {
	s.Source = &v
	return s
}

type ListAppModulesResponseBody struct {
	Data      *ListAppModulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppModulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppModulesResponseBody) SetData(v *ListAppModulesResponseBodyData) *ListAppModulesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppModulesResponseBody) SetRequestId(v string) *ListAppModulesResponseBody {
	s.RequestId = &v
	return s
}

type ListAppModulesResponseBodyData struct {
	Items []*ListAppModulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListAppModulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppModulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppModulesResponseBodyData) SetItems(v []*ListAppModulesResponseBodyDataItems) *ListAppModulesResponseBodyData {
	s.Items = v
	return s
}

type ListAppModulesResponseBodyDataItems struct {
	CommitId               *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectDependency       *bool   `json:"DirectDependency,omitempty" xml:"DirectDependency,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version                *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppModulesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListAppModulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAppModulesResponseBodyDataItems) SetCommitId(v string) *ListAppModulesResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetDescription(v string) *ListAppModulesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetDirectDependency(v bool) *ListAppModulesResponseBodyDataItems {
	s.DirectDependency = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetIcon(v string) *ListAppModulesResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetMinimumPlatformVersion(v string) *ListAppModulesResponseBodyDataItems {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetModuleId(v string) *ListAppModulesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetModuleName(v string) *ListAppModulesResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetOwnerUserId(v string) *ListAppModulesResponseBodyDataItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetPlatform(v string) *ListAppModulesResponseBodyDataItems {
	s.Platform = &v
	return s
}

func (s *ListAppModulesResponseBodyDataItems) SetVersion(v string) *ListAppModulesResponseBodyDataItems {
	s.Version = &v
	return s
}

type ListAppModulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppModulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppModulesResponse) GoString() string {
	return s.String()
}

func (s *ListAppModulesResponse) SetHeaders(v map[string]*string) *ListAppModulesResponse {
	s.Headers = v
	return s
}

func (s *ListAppModulesResponse) SetStatusCode(v int32) *ListAppModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppModulesResponse) SetBody(v *ListAppModulesResponseBody) *ListAppModulesResponse {
	s.Body = v
	return s
}

type ListAppTemplatesRequest struct {
	AppType      *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListAppTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesRequest) SetAppType(v string) *ListAppTemplatesRequest {
	s.AppType = &v
	return s
}

func (s *ListAppTemplatesRequest) SetSource(v string) *ListAppTemplatesRequest {
	s.Source = &v
	return s
}

func (s *ListAppTemplatesRequest) SetTemplateType(v string) *ListAppTemplatesRequest {
	s.TemplateType = &v
	return s
}

type ListAppTemplatesResponseBody struct {
	Data      *ListAppTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBody) SetData(v *ListAppTemplatesResponseBodyData) *ListAppTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppTemplatesResponseBody) SetRequestId(v string) *ListAppTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListAppTemplatesResponseBodyData struct {
	Items []*ListAppTemplatesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListAppTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyData) SetItems(v []*ListAppTemplatesResponseBodyDataItems) *ListAppTemplatesResponseBodyData {
	s.Items = v
	return s
}

type ListAppTemplatesResponseBodyDataItems struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType       *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon          *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LastEditTime  *string `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId  *string `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType  *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListAppTemplatesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyDataItems) SetAppName(v string) *ListAppTemplatesResponseBodyDataItems {
	s.AppName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetAppType(v string) *ListAppTemplatesResponseBodyDataItems {
	s.AppType = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetCategoryName(v string) *ListAppTemplatesResponseBodyDataItems {
	s.CategoryName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetCreateTime(v string) *ListAppTemplatesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetDescription(v string) *ListAppTemplatesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetIcon(v string) *ListAppTemplatesResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetLastEditTime(v string) *ListAppTemplatesResponseBodyDataItems {
	s.LastEditTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetMainModuleId(v string) *ListAppTemplatesResponseBodyDataItems {
	s.MainModuleId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetModifiedTime(v string) *ListAppTemplatesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetSchemaVersion(v string) *ListAppTemplatesResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetSource(v string) *ListAppTemplatesResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetTemplateId(v string) *ListAppTemplatesResponseBodyDataItems {
	s.TemplateId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyDataItems) SetTemplateType(v string) *ListAppTemplatesResponseBodyDataItems {
	s.TemplateType = &v
	return s
}

type ListAppTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponse) SetHeaders(v map[string]*string) *ListAppTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAppTemplatesResponse) SetStatusCode(v int32) *ListAppTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppTemplatesResponse) SetBody(v *ListAppTemplatesResponseBody) *ListAppTemplatesResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus      *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType        *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CustomParentId *string `json:"CustomParentId,omitempty" xml:"CustomParentId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MainModuleId   *string `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Template       *bool   `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppId(v string) *ListAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetAppStatus(v string) *ListAppsRequest {
	s.AppStatus = &v
	return s
}

func (s *ListAppsRequest) SetAppType(v string) *ListAppsRequest {
	s.AppType = &v
	return s
}

func (s *ListAppsRequest) SetCustomParentId(v string) *ListAppsRequest {
	s.CustomParentId = &v
	return s
}

func (s *ListAppsRequest) SetDescription(v string) *ListAppsRequest {
	s.Description = &v
	return s
}

func (s *ListAppsRequest) SetMainModuleId(v string) *ListAppsRequest {
	s.MainModuleId = &v
	return s
}

func (s *ListAppsRequest) SetPageNumber(v int32) *ListAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetSource(v string) *ListAppsRequest {
	s.Source = &v
	return s
}

func (s *ListAppsRequest) SetTemplate(v bool) *ListAppsRequest {
	s.Template = &v
	return s
}

type ListAppsResponseBody struct {
	Data      *ListAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppsResponseBodyData struct {
	Items      []*ListAppsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyData) SetItems(v []*ListAppsResponseBodyDataItems) *ListAppsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListAppsResponseBodyData) SetPageNumber(v int32) *ListAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAppsResponseBodyData) SetPageSize(v int32) *ListAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppsResponseBodyData) SetTotalCount(v int32) *ListAppsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyDataItems struct {
	AppId           *string                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string                                    `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus       *string                                    `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType         *string                                    `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Categories      []*ListAppsResponseBodyDataItemsCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime      *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon            *string                                    `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate      *bool                                      `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime    *string                                    `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId    *string                                    `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime    *string                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PlatformVersion *string                                    `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	SchemaVersion   *string                                    `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source          *string                                    `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListAppsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataItems) SetAppId(v string) *ListAppsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetAppName(v string) *ListAppsResponseBodyDataItems {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetAppStatus(v string) *ListAppsResponseBodyDataItems {
	s.AppStatus = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetAppType(v string) *ListAppsResponseBodyDataItems {
	s.AppType = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetCategories(v []*ListAppsResponseBodyDataItemsCategories) *ListAppsResponseBodyDataItems {
	s.Categories = v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetCreateTime(v string) *ListAppsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDescription(v string) *ListAppsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetIcon(v string) *ListAppsResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetIsTemplate(v bool) *ListAppsResponseBodyDataItems {
	s.IsTemplate = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetLastEditTime(v string) *ListAppsResponseBodyDataItems {
	s.LastEditTime = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetMainModuleId(v string) *ListAppsResponseBodyDataItems {
	s.MainModuleId = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetModifiedTime(v string) *ListAppsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetPlatformVersion(v string) *ListAppsResponseBodyDataItems {
	s.PlatformVersion = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetSchemaVersion(v string) *ListAppsResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetSource(v string) *ListAppsResponseBodyDataItems {
	s.Source = &v
	return s
}

type ListAppsResponseBodyDataItemsCategories struct {
	CategoryId       *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName     *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListAppsResponseBodyDataItemsCategories) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyDataItemsCategories) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataItemsCategories) SetCategoryId(v string) *ListAppsResponseBodyDataItemsCategories {
	s.CategoryId = &v
	return s
}

func (s *ListAppsResponseBodyDataItemsCategories) SetCategoryName(v string) *ListAppsResponseBodyDataItemsCategories {
	s.CategoryName = &v
	return s
}

func (s *ListAppsResponseBodyDataItemsCategories) SetParentCategoryId(v string) *ListAppsResponseBodyDataItemsCategories {
	s.ParentCategoryId = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListArtifactsRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PublishId *string `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListArtifactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) SetAppId(v string) *ListArtifactsRequest {
	s.AppId = &v
	return s
}

func (s *ListArtifactsRequest) SetPublishId(v string) *ListArtifactsRequest {
	s.PublishId = &v
	return s
}

func (s *ListArtifactsRequest) SetSource(v string) *ListArtifactsRequest {
	s.Source = &v
	return s
}

type ListArtifactsResponseBody struct {
	Data      *ListArtifactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) SetData(v *ListArtifactsResponseBodyData) *ListArtifactsResponseBody {
	s.Data = v
	return s
}

func (s *ListArtifactsResponseBody) SetRequestId(v string) *ListArtifactsResponseBody {
	s.RequestId = &v
	return s
}

type ListArtifactsResponseBodyData struct {
	Items []*ListArtifactsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListArtifactsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyData) SetItems(v []*ListArtifactsResponseBodyDataItems) *ListArtifactsResponseBodyData {
	s.Items = v
	return s
}

type ListArtifactsResponseBodyDataItems struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ArtifactId   *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Available    *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListArtifactsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyDataItems) SetAppId(v string) *ListArtifactsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetArtifactId(v string) *ListArtifactsResponseBodyDataItems {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetArtifactType(v string) *ListArtifactsResponseBodyDataItems {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetAvailable(v bool) *ListArtifactsResponseBodyDataItems {
	s.Available = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetCreateTime(v string) *ListArtifactsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetModifiedTime(v string) *ListArtifactsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListArtifactsResponseBodyDataItems) SetUrl(v string) *ListArtifactsResponseBodyDataItems {
	s.Url = &v
	return s
}

type ListArtifactsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponse) SetHeaders(v map[string]*string) *ListArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactsResponse) SetStatusCode(v int32) *ListArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactsResponse) SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse {
	s.Body = v
	return s
}

type ListCommitsRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitLog      *string `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CustomParentId *string `json:"CustomParentId,omitempty" xml:"CustomParentId,omitempty"`
	ModuleId       *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListCommitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommitsRequest) GoString() string {
	return s.String()
}

func (s *ListCommitsRequest) SetAppId(v string) *ListCommitsRequest {
	s.AppId = &v
	return s
}

func (s *ListCommitsRequest) SetCommitLog(v string) *ListCommitsRequest {
	s.CommitLog = &v
	return s
}

func (s *ListCommitsRequest) SetCustomParentId(v string) *ListCommitsRequest {
	s.CustomParentId = &v
	return s
}

func (s *ListCommitsRequest) SetModuleId(v string) *ListCommitsRequest {
	s.ModuleId = &v
	return s
}

func (s *ListCommitsRequest) SetPageNumber(v int32) *ListCommitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommitsRequest) SetPageSize(v int32) *ListCommitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommitsRequest) SetSource(v string) *ListCommitsRequest {
	s.Source = &v
	return s
}

type ListCommitsResponseBody struct {
	Data      *ListCommitsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCommitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommitsResponseBody) SetData(v *ListCommitsResponseBodyData) *ListCommitsResponseBody {
	s.Data = v
	return s
}

func (s *ListCommitsResponseBody) SetRequestId(v string) *ListCommitsResponseBody {
	s.RequestId = &v
	return s
}

type ListCommitsResponseBodyData struct {
	Items      []*ListCommitsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommitsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCommitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCommitsResponseBodyData) SetItems(v []*ListCommitsResponseBodyDataItems) *ListCommitsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListCommitsResponseBodyData) SetPageNumber(v int32) *ListCommitsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCommitsResponseBodyData) SetPageSize(v int32) *ListCommitsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCommitsResponseBodyData) SetTotalCount(v int32) *ListCommitsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCommitsResponseBodyDataItems struct {
	AppId              *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitDigest       *string            `json:"CommitDigest,omitempty" xml:"CommitDigest,omitempty"`
	CommitId           *string            `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitLog          *string            `json:"CommitLog,omitempty" xml:"CommitLog,omitempty"`
	CommitType         *string            `json:"CommitType,omitempty" xml:"CommitType,omitempty"`
	CreateTime         *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MainModuleCommitId *string            `json:"MainModuleCommitId,omitempty" xml:"MainModuleCommitId,omitempty"`
	MainModuleId       *string            `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModelDataPath      *string            `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModelDigest        map[string]*string `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModifiedTime       *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId           *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDataPath   *string            `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
	ResourceDigest     map[string]*string `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	RollbackToCommitId *string            `json:"RollbackToCommitId,omitempty" xml:"RollbackToCommitId,omitempty"`
	RollbackType       *string            `json:"RollbackType,omitempty" xml:"RollbackType,omitempty"`
	SchemaVersion      *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s ListCommitsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListCommitsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListCommitsResponseBodyDataItems) SetAppId(v string) *ListCommitsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetCommitDigest(v string) *ListCommitsResponseBodyDataItems {
	s.CommitDigest = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetCommitId(v string) *ListCommitsResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetCommitLog(v string) *ListCommitsResponseBodyDataItems {
	s.CommitLog = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetCommitType(v string) *ListCommitsResponseBodyDataItems {
	s.CommitType = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetCreateTime(v string) *ListCommitsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetMainModuleCommitId(v string) *ListCommitsResponseBodyDataItems {
	s.MainModuleCommitId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetMainModuleId(v string) *ListCommitsResponseBodyDataItems {
	s.MainModuleId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetModelDataPath(v string) *ListCommitsResponseBodyDataItems {
	s.ModelDataPath = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetModelDigest(v map[string]*string) *ListCommitsResponseBodyDataItems {
	s.ModelDigest = v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetModifiedTime(v string) *ListCommitsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetModuleId(v string) *ListCommitsResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetResourceDataPath(v string) *ListCommitsResponseBodyDataItems {
	s.ResourceDataPath = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetResourceDigest(v map[string]*string) *ListCommitsResponseBodyDataItems {
	s.ResourceDigest = v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetRollbackToCommitId(v string) *ListCommitsResponseBodyDataItems {
	s.RollbackToCommitId = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetRollbackType(v string) *ListCommitsResponseBodyDataItems {
	s.RollbackType = &v
	return s
}

func (s *ListCommitsResponseBodyDataItems) SetSchemaVersion(v string) *ListCommitsResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

type ListCommitsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommitsResponse) GoString() string {
	return s.String()
}

func (s *ListCommitsResponse) SetHeaders(v map[string]*string) *ListCommitsResponse {
	s.Headers = v
	return s
}

func (s *ListCommitsResponse) SetStatusCode(v int32) *ListCommitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommitsResponse) SetBody(v *ListCommitsResponseBody) *ListCommitsResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetAppId(v string) *ListDomainsRequest {
	s.AppId = &v
	return s
}

func (s *ListDomainsRequest) SetEnvId(v string) *ListDomainsRequest {
	s.EnvId = &v
	return s
}

func (s *ListDomainsRequest) SetSource(v string) *ListDomainsRequest {
	s.Source = &v
	return s
}

type ListDomainsResponseBody struct {
	Data      *ListDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetData(v *ListDomainsResponseBodyData) *ListDomainsResponseBody {
	s.Data = v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListDomainsResponseBodyData struct {
	Items []*ListDomainsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListDomainsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyData) SetItems(v []*ListDomainsResponseBodyDataItems) *ListDomainsResponseBodyData {
	s.Items = v
	return s
}

type ListDomainsResponseBodyDataItems struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Applied         *bool   `json:"Applied,omitempty" xml:"Applied,omitempty"`
	Checked         *bool   `json:"Checked,omitempty" xml:"Checked,omitempty"`
	Cname           *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Deleted         *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType      *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	WithCertificate *bool   `json:"WithCertificate,omitempty" xml:"WithCertificate,omitempty"`
}

func (s ListDomainsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDataItems) SetAppId(v string) *ListDomainsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetApplied(v bool) *ListDomainsResponseBodyDataItems {
	s.Applied = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetChecked(v bool) *ListDomainsResponseBodyDataItems {
	s.Checked = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetCname(v string) *ListDomainsResponseBodyDataItems {
	s.Cname = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetDeleted(v bool) *ListDomainsResponseBodyDataItems {
	s.Deleted = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetDomain(v string) *ListDomainsResponseBodyDataItems {
	s.Domain = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetDomainType(v string) *ListDomainsResponseBodyDataItems {
	s.DomainType = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetEnvId(v string) *ListDomainsResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetPath(v string) *ListDomainsResponseBodyDataItems {
	s.Path = &v
	return s
}

func (s *ListDomainsResponseBodyDataItems) SetWithCertificate(v bool) *ListDomainsResponseBodyDataItems {
	s.WithCertificate = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListEnvironmentOverviewsRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListEnvironmentOverviewsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentOverviewsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentOverviewsRequest) SetAppId(v string) *ListEnvironmentOverviewsRequest {
	s.AppId = &v
	return s
}

func (s *ListEnvironmentOverviewsRequest) SetSource(v string) *ListEnvironmentOverviewsRequest {
	s.Source = &v
	return s
}

type ListEnvironmentOverviewsResponseBody struct {
	Data      *ListEnvironmentOverviewsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvironmentOverviewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentOverviewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentOverviewsResponseBody) SetData(v *ListEnvironmentOverviewsResponseBodyData) *ListEnvironmentOverviewsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentOverviewsResponseBody) SetRequestId(v string) *ListEnvironmentOverviewsResponseBody {
	s.RequestId = &v
	return s
}

type ListEnvironmentOverviewsResponseBodyData struct {
	Items []*ListEnvironmentOverviewsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListEnvironmentOverviewsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentOverviewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentOverviewsResponseBodyData) SetItems(v []*ListEnvironmentOverviewsResponseBodyDataItems) *ListEnvironmentOverviewsResponseBodyData {
	s.Items = v
	return s
}

type ListEnvironmentOverviewsResponseBodyDataItems struct {
	AppId               *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Config              map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
	CreateTime          *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentPublish      map[string]*string `json:"CurrentPublish,omitempty" xml:"CurrentPublish,omitempty"`
	Endpoint            *string            `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EnvId               *string            `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvStatus           *string            `json:"EnvStatus,omitempty" xml:"EnvStatus,omitempty"`
	EnvType             *string            `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	LatestAppAccessTime *string            `json:"LatestAppAccessTime,omitempty" xml:"LatestAppAccessTime,omitempty"`
	ModifiedTime        *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OpsRecord           map[string]*string `json:"OpsRecord,omitempty" xml:"OpsRecord,omitempty"`
	Publishing          map[string]*string `json:"Publishing,omitempty" xml:"Publishing,omitempty"`
}

func (s ListEnvironmentOverviewsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentOverviewsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetAppId(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetConfig(v map[string]*string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.Config = v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetCreateTime(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetCurrentPublish(v map[string]*string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.CurrentPublish = v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetEndpoint(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.Endpoint = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetEnvId(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetEnvStatus(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.EnvStatus = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetEnvType(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.EnvType = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetLatestAppAccessTime(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.LatestAppAccessTime = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetModifiedTime(v string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetOpsRecord(v map[string]*string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.OpsRecord = v
	return s
}

func (s *ListEnvironmentOverviewsResponseBodyDataItems) SetPublishing(v map[string]*string) *ListEnvironmentOverviewsResponseBodyDataItems {
	s.Publishing = v
	return s
}

type ListEnvironmentOverviewsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentOverviewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentOverviewsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentOverviewsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentOverviewsResponse) SetHeaders(v map[string]*string) *ListEnvironmentOverviewsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentOverviewsResponse) SetStatusCode(v int32) *ListEnvironmentOverviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentOverviewsResponse) SetBody(v *ListEnvironmentOverviewsResponseBody) *ListEnvironmentOverviewsResponse {
	s.Body = v
	return s
}

type ListEnvironmentsRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) SetAppId(v string) *ListEnvironmentsRequest {
	s.AppId = &v
	return s
}

func (s *ListEnvironmentsRequest) SetEnvType(v string) *ListEnvironmentsRequest {
	s.EnvType = &v
	return s
}

func (s *ListEnvironmentsRequest) SetSource(v string) *ListEnvironmentsRequest {
	s.Source = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	Data      *ListEnvironmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetRequestId(v string) *ListEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

type ListEnvironmentsResponseBodyData struct {
	Items []*ListEnvironmentsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) SetItems(v []*ListEnvironmentsResponseBodyDataItems) *ListEnvironmentsResponseBodyData {
	s.Items = v
	return s
}

type ListEnvironmentsResponseBodyDataItems struct {
	AccountOpsEndpoint *string `json:"AccountOpsEndpoint,omitempty" xml:"AccountOpsEndpoint,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentPublishId   *string `json:"CurrentPublishId,omitempty" xml:"CurrentPublishId,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EnvId              *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType            *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	ModifiedTime       *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishingId       *string `json:"PublishingId,omitempty" xml:"PublishingId,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataItems) SetAccountOpsEndpoint(v string) *ListEnvironmentsResponseBodyDataItems {
	s.AccountOpsEndpoint = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetAppId(v string) *ListEnvironmentsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetCreateTime(v string) *ListEnvironmentsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetCurrentPublishId(v string) *ListEnvironmentsResponseBodyDataItems {
	s.CurrentPublishId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetEndpoint(v string) *ListEnvironmentsResponseBodyDataItems {
	s.Endpoint = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetEnvId(v string) *ListEnvironmentsResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetEnvType(v string) *ListEnvironmentsResponseBodyDataItems {
	s.EnvType = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetModifiedTime(v string) *ListEnvironmentsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataItems) SetPublishingId(v string) *ListEnvironmentsResponseBodyDataItems {
	s.PublishingId = &v
	return s
}

type ListEnvironmentsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetStatusCode(v int32) *ListEnvironmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type ListModelsRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	WithContent   *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) SetAppId(v string) *ListModelsRequest {
	s.AppId = &v
	return s
}

func (s *ListModelsRequest) SetModelId(v string) *ListModelsRequest {
	s.ModelId = &v
	return s
}

func (s *ListModelsRequest) SetModelName(v string) *ListModelsRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsRequest) SetModelType(v string) *ListModelsRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsRequest) SetModuleId(v string) *ListModelsRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModelsRequest) SetSchemaVersion(v string) *ListModelsRequest {
	s.SchemaVersion = &v
	return s
}

func (s *ListModelsRequest) SetSource(v string) *ListModelsRequest {
	s.Source = &v
	return s
}

func (s *ListModelsRequest) SetSubType(v string) *ListModelsRequest {
	s.SubType = &v
	return s
}

func (s *ListModelsRequest) SetWithContent(v bool) *ListModelsRequest {
	s.WithContent = &v
	return s
}

type ListModelsResponseBody struct {
	Data      *ListModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) SetData(v *ListModelsResponseBodyData) *ListModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

type ListModelsResponseBodyData struct {
	Items []*ListModelsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListModelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyData) SetItems(v []*ListModelsResponseBodyDataItems) *ListModelsResponseBodyData {
	s.Items = v
	return s
}

type ListModelsResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListModelsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyDataItems) SetAppId(v string) *ListModelsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetAttributes(v []map[string]*string) *ListModelsResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetContent(v map[string]*string) *ListModelsResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetCreateTime(v string) *ListModelsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetDescription(v string) *ListModelsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetId(v string) *ListModelsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetLinkModelId(v string) *ListModelsResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetLinkModuleId(v string) *ListModelsResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetLinked(v bool) *ListModelsResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModelDigest(v string) *ListModelsResponseBodyDataItems {
	s.ModelDigest = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModelId(v string) *ListModelsResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModelName(v string) *ListModelsResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModelStatus(v string) *ListModelsResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModelType(v string) *ListModelsResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModifiedTime(v string) *ListModelsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetModuleId(v string) *ListModelsResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetProps(v map[string]*string) *ListModelsResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetRevision(v int32) *ListModelsResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetSchemaVersion(v string) *ListModelsResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetSubType(v string) *ListModelsResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *ListModelsResponseBodyDataItems) SetVisibility(v string) *ListModelsResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type ListModelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetStatusCode(v int32) *ListModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

type ListModelsByPageRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	WithContent   *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListModelsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageRequest) GoString() string {
	return s.String()
}

func (s *ListModelsByPageRequest) SetAppId(v string) *ListModelsByPageRequest {
	s.AppId = &v
	return s
}

func (s *ListModelsByPageRequest) SetModelName(v string) *ListModelsByPageRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsByPageRequest) SetModelType(v string) *ListModelsByPageRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsByPageRequest) SetModuleId(v string) *ListModelsByPageRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModelsByPageRequest) SetPageNumber(v int32) *ListModelsByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsByPageRequest) SetPageSize(v int32) *ListModelsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsByPageRequest) SetSchemaVersion(v string) *ListModelsByPageRequest {
	s.SchemaVersion = &v
	return s
}

func (s *ListModelsByPageRequest) SetSource(v string) *ListModelsByPageRequest {
	s.Source = &v
	return s
}

func (s *ListModelsByPageRequest) SetSubType(v string) *ListModelsByPageRequest {
	s.SubType = &v
	return s
}

func (s *ListModelsByPageRequest) SetWithContent(v bool) *ListModelsByPageRequest {
	s.WithContent = &v
	return s
}

type ListModelsByPageResponseBody struct {
	Data      *ListModelsByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModelsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponseBody) SetData(v *ListModelsByPageResponseBodyData) *ListModelsByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListModelsByPageResponseBody) SetRequestId(v string) *ListModelsByPageResponseBody {
	s.RequestId = &v
	return s
}

type ListModelsByPageResponseBodyData struct {
	Items      []*ListModelsByPageResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelsByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponseBodyData) SetItems(v []*ListModelsByPageResponseBodyDataItems) *ListModelsByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *ListModelsByPageResponseBodyData) SetPageNumber(v int32) *ListModelsByPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListModelsByPageResponseBodyData) SetPageSize(v int32) *ListModelsByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListModelsByPageResponseBodyData) SetTotalCount(v int32) *ListModelsByPageResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListModelsByPageResponseBodyDataItems struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListModelsByPageResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponseBodyDataItems) SetAppId(v string) *ListModelsByPageResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetAttributes(v []map[string]*string) *ListModelsByPageResponseBodyDataItems {
	s.Attributes = v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetContent(v map[string]*string) *ListModelsByPageResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetCreateTime(v string) *ListModelsByPageResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetDescription(v string) *ListModelsByPageResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetId(v string) *ListModelsByPageResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetLinkModelId(v string) *ListModelsByPageResponseBodyDataItems {
	s.LinkModelId = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetLinkModuleId(v string) *ListModelsByPageResponseBodyDataItems {
	s.LinkModuleId = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetLinked(v bool) *ListModelsByPageResponseBodyDataItems {
	s.Linked = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModelId(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModelId = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModelName(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModelName = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModelStatus(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModelStatus = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModelType(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModifiedTime(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetModuleId(v string) *ListModelsByPageResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetProps(v map[string]*string) *ListModelsByPageResponseBodyDataItems {
	s.Props = v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetRevision(v int32) *ListModelsByPageResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetSchemaVersion(v string) *ListModelsByPageResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetSubType(v string) *ListModelsByPageResponseBodyDataItems {
	s.SubType = &v
	return s
}

func (s *ListModelsByPageResponseBodyDataItems) SetVisibility(v string) *ListModelsByPageResponseBodyDataItems {
	s.Visibility = &v
	return s
}

type ListModelsByPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsByPageResponse) GoString() string {
	return s.String()
}

func (s *ListModelsByPageResponse) SetHeaders(v map[string]*string) *ListModelsByPageResponse {
	s.Headers = v
	return s
}

func (s *ListModelsByPageResponse) SetStatusCode(v int32) *ListModelsByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelsByPageResponse) SetBody(v *ListModelsByPageResponseBody) *ListModelsByPageResponse {
	s.Body = v
	return s
}

type ListModuleDependenciesRequest struct {
	ModuleId  *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Recursive *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListModuleDependenciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModuleDependenciesRequest) GoString() string {
	return s.String()
}

func (s *ListModuleDependenciesRequest) SetModuleId(v string) *ListModuleDependenciesRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModuleDependenciesRequest) SetRecursive(v bool) *ListModuleDependenciesRequest {
	s.Recursive = &v
	return s
}

func (s *ListModuleDependenciesRequest) SetSource(v string) *ListModuleDependenciesRequest {
	s.Source = &v
	return s
}

type ListModuleDependenciesResponseBody struct {
	Data      *ListModuleDependenciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModuleDependenciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModuleDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModuleDependenciesResponseBody) SetData(v *ListModuleDependenciesResponseBodyData) *ListModuleDependenciesResponseBody {
	s.Data = v
	return s
}

func (s *ListModuleDependenciesResponseBody) SetRequestId(v string) *ListModuleDependenciesResponseBody {
	s.RequestId = &v
	return s
}

type ListModuleDependenciesResponseBodyData struct {
	Items []*ListModuleDependenciesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListModuleDependenciesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModuleDependenciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModuleDependenciesResponseBodyData) SetItems(v []*ListModuleDependenciesResponseBodyDataItems) *ListModuleDependenciesResponseBodyData {
	s.Items = v
	return s
}

type ListModuleDependenciesResponseBodyDataItems struct {
	CommitId               *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DirectDependency       *bool   `json:"DirectDependency,omitempty" xml:"DirectDependency,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	Origin                 *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version                *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListModuleDependenciesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModuleDependenciesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetCommitId(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetDescription(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetDirectDependency(v bool) *ListModuleDependenciesResponseBodyDataItems {
	s.DirectDependency = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetIcon(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetMinimumPlatformVersion(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetModuleId(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetModuleName(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetOrigin(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.Origin = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetOwnerUserId(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetPlatform(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.Platform = &v
	return s
}

func (s *ListModuleDependenciesResponseBodyDataItems) SetVersion(v string) *ListModuleDependenciesResponseBodyDataItems {
	s.Version = &v
	return s
}

type ListModuleDependenciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModuleDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModuleDependenciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModuleDependenciesResponse) GoString() string {
	return s.String()
}

func (s *ListModuleDependenciesResponse) SetHeaders(v map[string]*string) *ListModuleDependenciesResponse {
	s.Headers = v
	return s
}

func (s *ListModuleDependenciesResponse) SetStatusCode(v int32) *ListModuleDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModuleDependenciesResponse) SetBody(v *ListModuleDependenciesResponseBody) *ListModuleDependenciesResponse {
	s.Body = v
	return s
}

type ListModuleModelsRequest struct {
	ModuleList  *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubTypes    *string `json:"SubTypes,omitempty" xml:"SubTypes,omitempty"`
	WithContent *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListModuleModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModuleModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModuleModelsRequest) SetModuleList(v string) *ListModuleModelsRequest {
	s.ModuleList = &v
	return s
}

func (s *ListModuleModelsRequest) SetSource(v string) *ListModuleModelsRequest {
	s.Source = &v
	return s
}

func (s *ListModuleModelsRequest) SetSubTypes(v string) *ListModuleModelsRequest {
	s.SubTypes = &v
	return s
}

func (s *ListModuleModelsRequest) SetWithContent(v bool) *ListModuleModelsRequest {
	s.WithContent = &v
	return s
}

type ListModuleModelsResponseBody struct {
	Data      *ListModuleModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModuleModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModuleModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModuleModelsResponseBody) SetData(v *ListModuleModelsResponseBodyData) *ListModuleModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListModuleModelsResponseBody) SetRequestId(v string) *ListModuleModelsResponseBody {
	s.RequestId = &v
	return s
}

type ListModuleModelsResponseBodyData struct {
	Items []*ListModuleModelsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListModuleModelsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModuleModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModuleModelsResponseBodyData) SetItems(v []*ListModuleModelsResponseBodyDataItems) *ListModuleModelsResponseBodyData {
	s.Items = v
	return s
}

type ListModuleModelsResponseBodyDataItems struct {
	CommitId         *string                               `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	ModelData        map[string][]*DataItemsModelDataValue `json:"ModelData,omitempty" xml:"ModelData,omitempty"`
	ModelDataPath    map[string]*string                    `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModelDigest      map[string]*string                    `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModuleId         *string                               `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceData     map[string]*string                    `json:"ResourceData,omitempty" xml:"ResourceData,omitempty"`
	ResourceDataPath map[string]*string                    `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
}

func (s ListModuleModelsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModuleModelsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModuleModelsResponseBodyDataItems) SetCommitId(v string) *ListModuleModelsResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetModelData(v map[string][]*DataItemsModelDataValue) *ListModuleModelsResponseBodyDataItems {
	s.ModelData = v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetModelDataPath(v map[string]*string) *ListModuleModelsResponseBodyDataItems {
	s.ModelDataPath = v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetModelDigest(v map[string]*string) *ListModuleModelsResponseBodyDataItems {
	s.ModelDigest = v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetModuleId(v string) *ListModuleModelsResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetResourceData(v map[string]*string) *ListModuleModelsResponseBodyDataItems {
	s.ResourceData = v
	return s
}

func (s *ListModuleModelsResponseBodyDataItems) SetResourceDataPath(v map[string]*string) *ListModuleModelsResponseBodyDataItems {
	s.ResourceDataPath = v
	return s
}

type ListModuleModelsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModuleModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModuleModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModuleModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModuleModelsResponse) SetHeaders(v map[string]*string) *ListModuleModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModuleModelsResponse) SetStatusCode(v int32) *ListModuleModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModuleModelsResponse) SetBody(v *ListModuleModelsResponseBody) *ListModuleModelsResponse {
	s.Body = v
	return s
}

type ListModulePublishVersionsRequest struct {
	CustomParentId *string `json:"CustomParentId,omitempty" xml:"CustomParentId,omitempty"`
	ModuleId       *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleVersion  *string `json:"ModuleVersion,omitempty" xml:"ModuleVersion,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListModulePublishVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModulePublishVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListModulePublishVersionsRequest) SetCustomParentId(v string) *ListModulePublishVersionsRequest {
	s.CustomParentId = &v
	return s
}

func (s *ListModulePublishVersionsRequest) SetModuleId(v string) *ListModulePublishVersionsRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModulePublishVersionsRequest) SetModuleVersion(v string) *ListModulePublishVersionsRequest {
	s.ModuleVersion = &v
	return s
}

func (s *ListModulePublishVersionsRequest) SetPageNumber(v int32) *ListModulePublishVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulePublishVersionsRequest) SetPageSize(v int32) *ListModulePublishVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulePublishVersionsRequest) SetSource(v string) *ListModulePublishVersionsRequest {
	s.Source = &v
	return s
}

type ListModulePublishVersionsResponseBody struct {
	Data      *ListModulePublishVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModulePublishVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModulePublishVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModulePublishVersionsResponseBody) SetData(v *ListModulePublishVersionsResponseBodyData) *ListModulePublishVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListModulePublishVersionsResponseBody) SetRequestId(v string) *ListModulePublishVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListModulePublishVersionsResponseBodyData struct {
	Items      []*ListModulePublishVersionsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModulePublishVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModulePublishVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModulePublishVersionsResponseBodyData) SetItems(v []*ListModulePublishVersionsResponseBodyDataItems) *ListModulePublishVersionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListModulePublishVersionsResponseBodyData) SetPageNumber(v int32) *ListModulePublishVersionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyData) SetPageSize(v int32) *ListModulePublishVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyData) SetTotalCount(v int32) *ListModulePublishVersionsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListModulePublishVersionsResponseBodyDataItems struct {
	CommitId        *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId        *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PlatformVersion *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
	PublishId       *string `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListModulePublishVersionsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModulePublishVersionsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetCommitId(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetCreateTime(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetDescription(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetModifiedTime(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetModuleId(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetPlatformVersion(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.PlatformVersion = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetPublishId(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.PublishId = &v
	return s
}

func (s *ListModulePublishVersionsResponseBodyDataItems) SetVersion(v string) *ListModulePublishVersionsResponseBodyDataItems {
	s.Version = &v
	return s
}

type ListModulePublishVersionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModulePublishVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModulePublishVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModulePublishVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListModulePublishVersionsResponse) SetHeaders(v map[string]*string) *ListModulePublishVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListModulePublishVersionsResponse) SetStatusCode(v int32) *ListModulePublishVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModulePublishVersionsResponse) SetBody(v *ListModulePublishVersionsResponseBody) *ListModulePublishVersionsResponse {
	s.Body = v
	return s
}

type ListModuleResourcesRequest struct {
	ModuleList  *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Types       *string `json:"Types,omitempty" xml:"Types,omitempty"`
	WithContent *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListModuleResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListModuleResourcesRequest) SetModuleList(v string) *ListModuleResourcesRequest {
	s.ModuleList = &v
	return s
}

func (s *ListModuleResourcesRequest) SetSource(v string) *ListModuleResourcesRequest {
	s.Source = &v
	return s
}

func (s *ListModuleResourcesRequest) SetTypes(v string) *ListModuleResourcesRequest {
	s.Types = &v
	return s
}

func (s *ListModuleResourcesRequest) SetWithContent(v bool) *ListModuleResourcesRequest {
	s.WithContent = &v
	return s
}

type ListModuleResourcesResponseBody struct {
	Data      *ListModuleResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModuleResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModuleResourcesResponseBody) SetData(v *ListModuleResourcesResponseBodyData) *ListModuleResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListModuleResourcesResponseBody) SetRequestId(v string) *ListModuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListModuleResourcesResponseBodyData struct {
	Items []*ListModuleResourcesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListModuleResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModuleResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModuleResourcesResponseBodyData) SetItems(v []*ListModuleResourcesResponseBodyDataItems) *ListModuleResourcesResponseBodyData {
	s.Items = v
	return s
}

type ListModuleResourcesResponseBodyDataItems struct {
	CommitId         *string                                  `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	ModelData        map[string]*string                       `json:"ModelData,omitempty" xml:"ModelData,omitempty"`
	ModelDataPath    map[string]*string                       `json:"ModelDataPath,omitempty" xml:"ModelDataPath,omitempty"`
	ModuleId         *string                                  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceData     map[string][]*DataItemsResourceDataValue `json:"ResourceData,omitempty" xml:"ResourceData,omitempty"`
	ResourceDataPath map[string]*string                       `json:"ResourceDataPath,omitempty" xml:"ResourceDataPath,omitempty"`
}

func (s ListModuleResourcesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModuleResourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModuleResourcesResponseBodyDataItems) SetCommitId(v string) *ListModuleResourcesResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListModuleResourcesResponseBodyDataItems) SetModelData(v map[string]*string) *ListModuleResourcesResponseBodyDataItems {
	s.ModelData = v
	return s
}

func (s *ListModuleResourcesResponseBodyDataItems) SetModelDataPath(v map[string]*string) *ListModuleResourcesResponseBodyDataItems {
	s.ModelDataPath = v
	return s
}

func (s *ListModuleResourcesResponseBodyDataItems) SetModuleId(v string) *ListModuleResourcesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModuleResourcesResponseBodyDataItems) SetResourceData(v map[string][]*DataItemsResourceDataValue) *ListModuleResourcesResponseBodyDataItems {
	s.ResourceData = v
	return s
}

func (s *ListModuleResourcesResponseBodyDataItems) SetResourceDataPath(v map[string]*string) *ListModuleResourcesResponseBodyDataItems {
	s.ResourceDataPath = v
	return s
}

type ListModuleResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModuleResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListModuleResourcesResponse) SetHeaders(v map[string]*string) *ListModuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListModuleResourcesResponse) SetStatusCode(v int32) *ListModuleResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModuleResourcesResponse) SetBody(v *ListModuleResourcesResponseBody) *ListModuleResourcesResponse {
	s.Body = v
	return s
}

type ListModulesRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasOwnerApp *bool   `json:"HasOwnerApp,omitempty" xml:"HasOwnerApp,omitempty"`
	ModuleId    *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName  *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListModulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModulesRequest) GoString() string {
	return s.String()
}

func (s *ListModulesRequest) SetDescription(v string) *ListModulesRequest {
	s.Description = &v
	return s
}

func (s *ListModulesRequest) SetHasOwnerApp(v bool) *ListModulesRequest {
	s.HasOwnerApp = &v
	return s
}

func (s *ListModulesRequest) SetModuleId(v string) *ListModulesRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModulesRequest) SetModuleName(v string) *ListModulesRequest {
	s.ModuleName = &v
	return s
}

func (s *ListModulesRequest) SetPlatform(v string) *ListModulesRequest {
	s.Platform = &v
	return s
}

func (s *ListModulesRequest) SetSource(v string) *ListModulesRequest {
	s.Source = &v
	return s
}

type ListModulesResponseBody struct {
	Data      *ListModulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBody) SetData(v *ListModulesResponseBodyData) *ListModulesResponseBody {
	s.Data = v
	return s
}

func (s *ListModulesResponseBody) SetRequestId(v string) *ListModulesResponseBody {
	s.RequestId = &v
	return s
}

type ListModulesResponseBodyData struct {
	Items []*ListModulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListModulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyData) SetItems(v []*ListModulesResponseBodyDataItems) *ListModulesResponseBodyData {
	s.Items = v
	return s
}

type ListModulesResponseBodyDataItems struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PlatformVersion        *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
}

func (s ListModulesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyDataItems) SetCreateTime(v string) *ListModulesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetDescription(v string) *ListModulesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetIcon(v string) *ListModulesResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetLatestPublishedCommit(v string) *ListModulesResponseBodyDataItems {
	s.LatestPublishedCommit = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetLatestPublishedVersion(v string) *ListModulesResponseBodyDataItems {
	s.LatestPublishedVersion = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetMinimumPlatformVersion(v string) *ListModulesResponseBodyDataItems {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetModifiedTime(v string) *ListModulesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetModuleId(v string) *ListModulesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetModuleName(v string) *ListModulesResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetOwnerAppId(v string) *ListModulesResponseBodyDataItems {
	s.OwnerAppId = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetOwnerUserId(v string) *ListModulesResponseBodyDataItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetPlatform(v string) *ListModulesResponseBodyDataItems {
	s.Platform = &v
	return s
}

func (s *ListModulesResponseBodyDataItems) SetPlatformVersion(v string) *ListModulesResponseBodyDataItems {
	s.PlatformVersion = &v
	return s
}

type ListModulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponse) GoString() string {
	return s.String()
}

func (s *ListModulesResponse) SetHeaders(v map[string]*string) *ListModulesResponse {
	s.Headers = v
	return s
}

func (s *ListModulesResponse) SetStatusCode(v int32) *ListModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModulesResponse) SetBody(v *ListModulesResponseBody) *ListModulesResponse {
	s.Body = v
	return s
}

type ListModulesByPageRequest struct {
	CustomParentId *string `json:"CustomParentId,omitempty" xml:"CustomParentId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasOwnerApp    *bool   `json:"HasOwnerApp,omitempty" xml:"HasOwnerApp,omitempty"`
	ModuleId       *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName     *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType     *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platform       *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListModulesByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModulesByPageRequest) GoString() string {
	return s.String()
}

func (s *ListModulesByPageRequest) SetCustomParentId(v string) *ListModulesByPageRequest {
	s.CustomParentId = &v
	return s
}

func (s *ListModulesByPageRequest) SetDescription(v string) *ListModulesByPageRequest {
	s.Description = &v
	return s
}

func (s *ListModulesByPageRequest) SetHasOwnerApp(v bool) *ListModulesByPageRequest {
	s.HasOwnerApp = &v
	return s
}

func (s *ListModulesByPageRequest) SetModuleId(v string) *ListModulesByPageRequest {
	s.ModuleId = &v
	return s
}

func (s *ListModulesByPageRequest) SetModuleName(v string) *ListModulesByPageRequest {
	s.ModuleName = &v
	return s
}

func (s *ListModulesByPageRequest) SetModuleType(v string) *ListModulesByPageRequest {
	s.ModuleType = &v
	return s
}

func (s *ListModulesByPageRequest) SetPageNumber(v int32) *ListModulesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesByPageRequest) SetPageSize(v int32) *ListModulesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesByPageRequest) SetPlatform(v string) *ListModulesByPageRequest {
	s.Platform = &v
	return s
}

func (s *ListModulesByPageRequest) SetSource(v string) *ListModulesByPageRequest {
	s.Source = &v
	return s
}

type ListModulesByPageResponseBody struct {
	Data      *ListModulesByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModulesByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModulesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListModulesByPageResponseBody) SetData(v *ListModulesByPageResponseBodyData) *ListModulesByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListModulesByPageResponseBody) SetRequestId(v string) *ListModulesByPageResponseBody {
	s.RequestId = &v
	return s
}

type ListModulesByPageResponseBodyData struct {
	Items      []*ListModulesByPageResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModulesByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListModulesByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModulesByPageResponseBodyData) SetItems(v []*ListModulesByPageResponseBodyDataItems) *ListModulesByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *ListModulesByPageResponseBodyData) SetPageNumber(v int32) *ListModulesByPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListModulesByPageResponseBodyData) SetPageSize(v int32) *ListModulesByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListModulesByPageResponseBodyData) SetTotalCount(v int32) *ListModulesByPageResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListModulesByPageResponseBodyDataItems struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType             *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PlatformVersion        *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
}

func (s ListModulesByPageResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListModulesByPageResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModulesByPageResponseBodyDataItems) SetCreateTime(v string) *ListModulesByPageResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetDescription(v string) *ListModulesByPageResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetIcon(v string) *ListModulesByPageResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetLatestPublishedCommit(v string) *ListModulesByPageResponseBodyDataItems {
	s.LatestPublishedCommit = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetLatestPublishedVersion(v string) *ListModulesByPageResponseBodyDataItems {
	s.LatestPublishedVersion = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetMinimumPlatformVersion(v string) *ListModulesByPageResponseBodyDataItems {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetModifiedTime(v string) *ListModulesByPageResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetModuleId(v string) *ListModulesByPageResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetModuleName(v string) *ListModulesByPageResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetModuleType(v string) *ListModulesByPageResponseBodyDataItems {
	s.ModuleType = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetOwnerAppId(v string) *ListModulesByPageResponseBodyDataItems {
	s.OwnerAppId = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetOwnerUserId(v string) *ListModulesByPageResponseBodyDataItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetPlatform(v string) *ListModulesByPageResponseBodyDataItems {
	s.Platform = &v
	return s
}

func (s *ListModulesByPageResponseBodyDataItems) SetPlatformVersion(v string) *ListModulesByPageResponseBodyDataItems {
	s.PlatformVersion = &v
	return s
}

type ListModulesByPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModulesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModulesByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModulesByPageResponse) GoString() string {
	return s.String()
}

func (s *ListModulesByPageResponse) SetHeaders(v map[string]*string) *ListModulesByPageResponse {
	s.Headers = v
	return s
}

func (s *ListModulesByPageResponse) SetStatusCode(v int32) *ListModulesByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModulesByPageResponse) SetBody(v *ListModulesByPageResponseBody) *ListModulesByPageResponse {
	s.Body = v
	return s
}

type ListPublishVersionsRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListPublishVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPublishVersionsRequest) SetAppId(v string) *ListPublishVersionsRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishVersionsRequest) SetEnvId(v string) *ListPublishVersionsRequest {
	s.EnvId = &v
	return s
}

func (s *ListPublishVersionsRequest) SetPageNumber(v int32) *ListPublishVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishVersionsRequest) SetPageSize(v int32) *ListPublishVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishVersionsRequest) SetSource(v string) *ListPublishVersionsRequest {
	s.Source = &v
	return s
}

type ListPublishVersionsResponseBody struct {
	Data      *ListPublishVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishVersionsResponseBody) SetData(v *ListPublishVersionsResponseBodyData) *ListPublishVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishVersionsResponseBody) SetRequestId(v string) *ListPublishVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListPublishVersionsResponseBodyData struct {
	Items      []*ListPublishVersionsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPublishVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishVersionsResponseBodyData) SetItems(v []*ListPublishVersionsResponseBodyDataItems) *ListPublishVersionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPublishVersionsResponseBodyData) SetPageNumber(v int32) *ListPublishVersionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPublishVersionsResponseBodyData) SetPageSize(v int32) *ListPublishVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishVersionsResponseBodyData) SetTotalCount(v int32) *ListPublishVersionsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListPublishVersionsResponseBodyDataItems struct {
	AppId          *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId       *string              `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CompletionTime *string              `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreateTime     *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvId          *string              `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ModifiedTime   *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishId      *string              `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	PublishStatus  *string              `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishType    *string              `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Reason         *string              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime      *string              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubTasks       []map[string]*string `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	VersionNumber  *string              `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s ListPublishVersionsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListPublishVersionsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPublishVersionsResponseBodyDataItems) SetAppId(v string) *ListPublishVersionsResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetCommitId(v string) *ListPublishVersionsResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetCompletionTime(v string) *ListPublishVersionsResponseBodyDataItems {
	s.CompletionTime = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetCreateTime(v string) *ListPublishVersionsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetDescription(v string) *ListPublishVersionsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetEnvId(v string) *ListPublishVersionsResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetModifiedTime(v string) *ListPublishVersionsResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetPublishId(v string) *ListPublishVersionsResponseBodyDataItems {
	s.PublishId = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetPublishStatus(v string) *ListPublishVersionsResponseBodyDataItems {
	s.PublishStatus = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetPublishType(v string) *ListPublishVersionsResponseBodyDataItems {
	s.PublishType = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetReason(v string) *ListPublishVersionsResponseBodyDataItems {
	s.Reason = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetStartTime(v string) *ListPublishVersionsResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetSubTasks(v []map[string]*string) *ListPublishVersionsResponseBodyDataItems {
	s.SubTasks = v
	return s
}

func (s *ListPublishVersionsResponseBodyDataItems) SetVersionNumber(v string) *ListPublishVersionsResponseBodyDataItems {
	s.VersionNumber = &v
	return s
}

type ListPublishVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPublishVersionsResponse) SetHeaders(v map[string]*string) *ListPublishVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPublishVersionsResponse) SetStatusCode(v int32) *ListPublishVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishVersionsResponse) SetBody(v *ListPublishVersionsResponseBody) *ListPublishVersionsResponse {
	s.Body = v
	return s
}

type ListPublishedModulesRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAppId    *string `json:"ExcludeAppId,omitempty" xml:"ExcludeAppId,omitempty"`
	ExcludeModuleId *string `json:"ExcludeModuleId,omitempty" xml:"ExcludeModuleId,omitempty"`
	HasOwnerApp     *bool   `json:"HasOwnerApp,omitempty" xml:"HasOwnerApp,omitempty"`
	ModuleId        *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName      *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType      *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListPublishedModulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedModulesRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedModulesRequest) SetDescription(v string) *ListPublishedModulesRequest {
	s.Description = &v
	return s
}

func (s *ListPublishedModulesRequest) SetExcludeAppId(v string) *ListPublishedModulesRequest {
	s.ExcludeAppId = &v
	return s
}

func (s *ListPublishedModulesRequest) SetExcludeModuleId(v string) *ListPublishedModulesRequest {
	s.ExcludeModuleId = &v
	return s
}

func (s *ListPublishedModulesRequest) SetHasOwnerApp(v bool) *ListPublishedModulesRequest {
	s.HasOwnerApp = &v
	return s
}

func (s *ListPublishedModulesRequest) SetModuleId(v string) *ListPublishedModulesRequest {
	s.ModuleId = &v
	return s
}

func (s *ListPublishedModulesRequest) SetModuleName(v string) *ListPublishedModulesRequest {
	s.ModuleName = &v
	return s
}

func (s *ListPublishedModulesRequest) SetModuleType(v string) *ListPublishedModulesRequest {
	s.ModuleType = &v
	return s
}

func (s *ListPublishedModulesRequest) SetPageNumber(v int32) *ListPublishedModulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedModulesRequest) SetPageSize(v int32) *ListPublishedModulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedModulesRequest) SetPlatform(v string) *ListPublishedModulesRequest {
	s.Platform = &v
	return s
}

func (s *ListPublishedModulesRequest) SetSource(v string) *ListPublishedModulesRequest {
	s.Source = &v
	return s
}

type ListPublishedModulesResponseBody struct {
	Data      *ListPublishedModulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishedModulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedModulesResponseBody) SetData(v *ListPublishedModulesResponseBodyData) *ListPublishedModulesResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishedModulesResponseBody) SetRequestId(v string) *ListPublishedModulesResponseBody {
	s.RequestId = &v
	return s
}

type ListPublishedModulesResponseBodyData struct {
	Items      []*ListPublishedModulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishedModulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedModulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishedModulesResponseBodyData) SetItems(v []*ListPublishedModulesResponseBodyDataItems) *ListPublishedModulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPublishedModulesResponseBodyData) SetPageNumber(v int32) *ListPublishedModulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedModulesResponseBodyData) SetPageSize(v int32) *ListPublishedModulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishedModulesResponseBodyData) SetTotalCount(v int32) *ListPublishedModulesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListPublishedModulesResponseBodyDataItems struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleType             *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PlatformVersion        *string `json:"PlatformVersion,omitempty" xml:"PlatformVersion,omitempty"`
}

func (s ListPublishedModulesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedModulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPublishedModulesResponseBodyDataItems) SetCreateTime(v string) *ListPublishedModulesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetDescription(v string) *ListPublishedModulesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetIcon(v string) *ListPublishedModulesResponseBodyDataItems {
	s.Icon = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetLatestPublishedCommit(v string) *ListPublishedModulesResponseBodyDataItems {
	s.LatestPublishedCommit = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetLatestPublishedVersion(v string) *ListPublishedModulesResponseBodyDataItems {
	s.LatestPublishedVersion = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetMinimumPlatformVersion(v string) *ListPublishedModulesResponseBodyDataItems {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetModifiedTime(v string) *ListPublishedModulesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetModuleId(v string) *ListPublishedModulesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetModuleName(v string) *ListPublishedModulesResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetModuleType(v string) *ListPublishedModulesResponseBodyDataItems {
	s.ModuleType = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetOwnerAppId(v string) *ListPublishedModulesResponseBodyDataItems {
	s.OwnerAppId = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetOwnerUserId(v string) *ListPublishedModulesResponseBodyDataItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetPlatform(v string) *ListPublishedModulesResponseBodyDataItems {
	s.Platform = &v
	return s
}

func (s *ListPublishedModulesResponseBodyDataItems) SetPlatformVersion(v string) *ListPublishedModulesResponseBodyDataItems {
	s.PlatformVersion = &v
	return s
}

type ListPublishedModulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedModulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedModulesResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedModulesResponse) SetHeaders(v map[string]*string) *ListPublishedModulesResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedModulesResponse) SetStatusCode(v int32) *ListPublishedModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedModulesResponse) SetBody(v *ListPublishedModulesResponseBody) *ListPublishedModulesResponse {
	s.Body = v
	return s
}

type ListPublishesRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId         *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListPublishesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishesRequest) GoString() string {
	return s.String()
}

func (s *ListPublishesRequest) SetAppId(v string) *ListPublishesRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishesRequest) SetEnvId(v string) *ListPublishesRequest {
	s.EnvId = &v
	return s
}

func (s *ListPublishesRequest) SetPageNumber(v int32) *ListPublishesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishesRequest) SetPageSize(v int32) *ListPublishesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishesRequest) SetPublishStatus(v string) *ListPublishesRequest {
	s.PublishStatus = &v
	return s
}

func (s *ListPublishesRequest) SetPublishType(v string) *ListPublishesRequest {
	s.PublishType = &v
	return s
}

func (s *ListPublishesRequest) SetSource(v string) *ListPublishesRequest {
	s.Source = &v
	return s
}

type ListPublishesResponseBody struct {
	Data      *ListPublishesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishesResponseBody) SetData(v *ListPublishesResponseBodyData) *ListPublishesResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishesResponseBody) SetRequestId(v string) *ListPublishesResponseBody {
	s.RequestId = &v
	return s
}

type ListPublishesResponseBodyData struct {
	Items      []*ListPublishesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPublishesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishesResponseBodyData) SetItems(v []*ListPublishesResponseBodyDataItems) *ListPublishesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPublishesResponseBodyData) SetPageNumber(v int32) *ListPublishesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPublishesResponseBodyData) SetPageSize(v int32) *ListPublishesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishesResponseBodyData) SetTotalCount(v int32) *ListPublishesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListPublishesResponseBodyDataItems struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId       *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CompletionTime *string `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvId          *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	ModifiedTime   *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PublishId      *string `json:"PublishId,omitempty" xml:"PublishId,omitempty"`
	PublishStatus  *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishType    *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VersionNumber  *string `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s ListPublishesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListPublishesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPublishesResponseBodyDataItems) SetAppId(v string) *ListPublishesResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetCommitId(v string) *ListPublishesResponseBodyDataItems {
	s.CommitId = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetCompletionTime(v string) *ListPublishesResponseBodyDataItems {
	s.CompletionTime = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetCreateTime(v string) *ListPublishesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetDescription(v string) *ListPublishesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetEnvId(v string) *ListPublishesResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetModifiedTime(v string) *ListPublishesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetPublishId(v string) *ListPublishesResponseBodyDataItems {
	s.PublishId = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetPublishStatus(v string) *ListPublishesResponseBodyDataItems {
	s.PublishStatus = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetPublishType(v string) *ListPublishesResponseBodyDataItems {
	s.PublishType = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetReason(v string) *ListPublishesResponseBodyDataItems {
	s.Reason = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetStartTime(v string) *ListPublishesResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *ListPublishesResponseBodyDataItems) SetVersionNumber(v string) *ListPublishesResponseBodyDataItems {
	s.VersionNumber = &v
	return s
}

type ListPublishesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishesResponse) GoString() string {
	return s.String()
}

func (s *ListPublishesResponse) SetHeaders(v map[string]*string) *ListPublishesResponse {
	s.Headers = v
	return s
}

func (s *ListPublishesResponse) SetStatusCode(v int32) *ListPublishesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishesResponse) SetBody(v *ListPublishesResponseBody) *ListPublishesResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageProcessParameter *string `json:"ImageProcessParameter,omitempty" xml:"ImageProcessParameter,omitempty"`
	ModuleId              *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName          *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType          *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WithContent           *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetAppId(v string) *ListResourcesRequest {
	s.AppId = &v
	return s
}

func (s *ListResourcesRequest) SetDescription(v string) *ListResourcesRequest {
	s.Description = &v
	return s
}

func (s *ListResourcesRequest) SetImageProcessParameter(v string) *ListResourcesRequest {
	s.ImageProcessParameter = &v
	return s
}

func (s *ListResourcesRequest) SetModuleId(v string) *ListResourcesRequest {
	s.ModuleId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequest) SetSource(v string) *ListResourcesRequest {
	s.Source = &v
	return s
}

func (s *ListResourcesRequest) SetWithContent(v bool) *ListResourcesRequest {
	s.WithContent = &v
	return s
}

type ListResourcesResponseBody struct {
	Data      *ListResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetData(v *ListResourcesResponseBodyData) *ListResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyData struct {
	Items []*ListResourcesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyData) SetItems(v []*ListResourcesResponseBodyDataItems) *ListResourcesResponseBodyData {
	s.Items = v
	return s
}

type ListResourcesResponseBodyDataItems struct {
	AppId          *string     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string     `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime   *string     `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId       *string     `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDigest *string     `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourceId     *string     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string     `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision       *int32      `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion  *string     `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s ListResourcesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyDataItems) SetAppId(v string) *ListResourcesResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetContent(v interface{}) *ListResourcesResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetCreateTime(v string) *ListResourcesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetDescription(v string) *ListResourcesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetModifiedTime(v string) *ListResourcesResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetModuleId(v string) *ListResourcesResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetResourceDigest(v string) *ListResourcesResponseBodyDataItems {
	s.ResourceDigest = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetResourceId(v string) *ListResourcesResponseBodyDataItems {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetResourceName(v string) *ListResourcesResponseBodyDataItems {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetResourceType(v string) *ListResourcesResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetRevision(v int32) *ListResourcesResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *ListResourcesResponseBodyDataItems) SetSchemaVersion(v string) *ListResourcesResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListResourcesByPageRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageProcessParameter *string `json:"ImageProcessParameter,omitempty" xml:"ImageProcessParameter,omitempty"`
	ModuleId              *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName          *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType          *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WithContent           *bool   `json:"WithContent,omitempty" xml:"WithContent,omitempty"`
}

func (s ListResourcesByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByPageRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesByPageRequest) SetAppId(v string) *ListResourcesByPageRequest {
	s.AppId = &v
	return s
}

func (s *ListResourcesByPageRequest) SetDescription(v string) *ListResourcesByPageRequest {
	s.Description = &v
	return s
}

func (s *ListResourcesByPageRequest) SetImageProcessParameter(v string) *ListResourcesByPageRequest {
	s.ImageProcessParameter = &v
	return s
}

func (s *ListResourcesByPageRequest) SetModuleId(v string) *ListResourcesByPageRequest {
	s.ModuleId = &v
	return s
}

func (s *ListResourcesByPageRequest) SetPageNumber(v int32) *ListResourcesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesByPageRequest) SetPageSize(v int32) *ListResourcesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesByPageRequest) SetResourceId(v string) *ListResourcesByPageRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesByPageRequest) SetResourceName(v string) *ListResourcesByPageRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesByPageRequest) SetResourceType(v string) *ListResourcesByPageRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesByPageRequest) SetSource(v string) *ListResourcesByPageRequest {
	s.Source = &v
	return s
}

func (s *ListResourcesByPageRequest) SetWithContent(v bool) *ListResourcesByPageRequest {
	s.WithContent = &v
	return s
}

type ListResourcesByPageResponseBody struct {
	Data      *ListResourcesByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesByPageResponseBody) SetData(v *ListResourcesByPageResponseBodyData) *ListResourcesByPageResponseBody {
	s.Data = v
	return s
}

func (s *ListResourcesByPageResponseBody) SetRequestId(v string) *ListResourcesByPageResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesByPageResponseBodyData struct {
	Items      []*ListResourcesByPageResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourcesByPageResponseBodyData) SetItems(v []*ListResourcesByPageResponseBodyDataItems) *ListResourcesByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *ListResourcesByPageResponseBodyData) SetPageNumber(v int32) *ListResourcesByPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesByPageResponseBodyData) SetPageSize(v int32) *ListResourcesByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListResourcesByPageResponseBodyData) SetTotalCount(v int32) *ListResourcesByPageResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListResourcesByPageResponseBodyDataItems struct {
	AppId          *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime   *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId       *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDigest *string            `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourceId     *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision       *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion  *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s ListResourcesByPageResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByPageResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListResourcesByPageResponseBodyDataItems) SetAppId(v string) *ListResourcesByPageResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetContent(v map[string]*string) *ListResourcesByPageResponseBodyDataItems {
	s.Content = v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetCreateTime(v string) *ListResourcesByPageResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetDescription(v string) *ListResourcesByPageResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetModifiedTime(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ModifiedTime = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetModuleId(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ModuleId = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetResourceDigest(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ResourceDigest = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetResourceId(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetResourceName(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetResourceType(v string) *ListResourcesByPageResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetRevision(v int32) *ListResourcesByPageResponseBodyDataItems {
	s.Revision = &v
	return s
}

func (s *ListResourcesByPageResponseBodyDataItems) SetSchemaVersion(v string) *ListResourcesByPageResponseBodyDataItems {
	s.SchemaVersion = &v
	return s
}

type ListResourcesByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByPageResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesByPageResponse) SetHeaders(v map[string]*string) *ListResourcesByPageResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesByPageResponse) SetStatusCode(v int32) *ListResourcesByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesByPageResponse) SetBody(v *ListResourcesByPageResponseBody) *ListResourcesByPageResponse {
	s.Body = v
	return s
}

type ResetAppUserPasswordRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId    *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResetAppUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAppUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAppUserPasswordRequest) SetAppId(v string) *ResetAppUserPasswordRequest {
	s.AppId = &v
	return s
}

func (s *ResetAppUserPasswordRequest) SetEnvId(v string) *ResetAppUserPasswordRequest {
	s.EnvId = &v
	return s
}

func (s *ResetAppUserPasswordRequest) SetSource(v string) *ResetAppUserPasswordRequest {
	s.Source = &v
	return s
}

func (s *ResetAppUserPasswordRequest) SetUserName(v string) *ResetAppUserPasswordRequest {
	s.UserName = &v
	return s
}

type ResetAppUserPasswordResponseBody struct {
	Data      *ResetAppUserPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAppUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppUserPasswordResponseBody) SetData(v *ResetAppUserPasswordResponseBodyData) *ResetAppUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetAppUserPasswordResponseBody) SetRequestId(v string) *ResetAppUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAppUserPasswordResponseBodyData struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResetAppUserPasswordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ResetAppUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAppUserPasswordResponseBodyData) SetPassword(v string) *ResetAppUserPasswordResponseBodyData {
	s.Password = &v
	return s
}

func (s *ResetAppUserPasswordResponseBodyData) SetUserName(v string) *ResetAppUserPasswordResponseBodyData {
	s.UserName = &v
	return s
}

type ResetAppUserPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAppUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAppUserPasswordResponse) SetHeaders(v map[string]*string) *ResetAppUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAppUserPasswordResponse) SetStatusCode(v int32) *ResetAppUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppUserPasswordResponse) SetBody(v *ResetAppUserPasswordResponseBody) *ResetAppUserPasswordResponse {
	s.Body = v
	return s
}

type RestoreModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s RestoreModelRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreModelRequest) GoString() string {
	return s.String()
}

func (s *RestoreModelRequest) SetAppId(v string) *RestoreModelRequest {
	s.AppId = &v
	return s
}

func (s *RestoreModelRequest) SetModelId(v string) *RestoreModelRequest {
	s.ModelId = &v
	return s
}

func (s *RestoreModelRequest) SetModuleId(v string) *RestoreModelRequest {
	s.ModuleId = &v
	return s
}

func (s *RestoreModelRequest) SetSchemaVersion(v string) *RestoreModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *RestoreModelRequest) SetSource(v string) *RestoreModelRequest {
	s.Source = &v
	return s
}

type RestoreModelResponseBody struct {
	Data      *RestoreModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreModelResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreModelResponseBody) SetData(v *RestoreModelResponseBodyData) *RestoreModelResponseBody {
	s.Data = v
	return s
}

func (s *RestoreModelResponseBody) SetRequestId(v string) *RestoreModelResponseBody {
	s.RequestId = &v
	return s
}

type RestoreModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s RestoreModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestoreModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestoreModelResponseBodyData) SetAppId(v string) *RestoreModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetAttributes(v []map[string]*string) *RestoreModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *RestoreModelResponseBodyData) SetContent(v map[string]*string) *RestoreModelResponseBodyData {
	s.Content = v
	return s
}

func (s *RestoreModelResponseBodyData) SetCreateTime(v string) *RestoreModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetDescription(v string) *RestoreModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetId(v string) *RestoreModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetLinkModelId(v string) *RestoreModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetLinkModuleId(v string) *RestoreModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetLinked(v bool) *RestoreModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModelId(v string) *RestoreModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModelName(v string) *RestoreModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModelStatus(v string) *RestoreModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModelType(v string) *RestoreModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModifiedTime(v string) *RestoreModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetModuleId(v string) *RestoreModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetProps(v map[string]*string) *RestoreModelResponseBodyData {
	s.Props = v
	return s
}

func (s *RestoreModelResponseBodyData) SetRevision(v int32) *RestoreModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetSchemaVersion(v string) *RestoreModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetSubType(v string) *RestoreModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *RestoreModelResponseBodyData) SetVisibility(v string) *RestoreModelResponseBodyData {
	s.Visibility = &v
	return s
}

type RestoreModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreModelResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreModelResponse) GoString() string {
	return s.String()
}

func (s *RestoreModelResponse) SetHeaders(v map[string]*string) *RestoreModelResponse {
	s.Headers = v
	return s
}

func (s *RestoreModelResponse) SetStatusCode(v int32) *RestoreModelResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreModelResponse) SetBody(v *RestoreModelResponseBody) *RestoreModelResponse {
	s.Body = v
	return s
}

type RunLogicModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CommitId      *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EncodeType    *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Parameters    *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s RunLogicModelRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLogicModelRequest) GoString() string {
	return s.String()
}

func (s *RunLogicModelRequest) SetAppId(v string) *RunLogicModelRequest {
	s.AppId = &v
	return s
}

func (s *RunLogicModelRequest) SetCommitId(v string) *RunLogicModelRequest {
	s.CommitId = &v
	return s
}

func (s *RunLogicModelRequest) SetContent(v string) *RunLogicModelRequest {
	s.Content = &v
	return s
}

func (s *RunLogicModelRequest) SetEncodeType(v string) *RunLogicModelRequest {
	s.EncodeType = &v
	return s
}

func (s *RunLogicModelRequest) SetModuleId(v string) *RunLogicModelRequest {
	s.ModuleId = &v
	return s
}

func (s *RunLogicModelRequest) SetParameters(v string) *RunLogicModelRequest {
	s.Parameters = &v
	return s
}

func (s *RunLogicModelRequest) SetSchemaVersion(v string) *RunLogicModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *RunLogicModelRequest) SetSource(v string) *RunLogicModelRequest {
	s.Source = &v
	return s
}

func (s *RunLogicModelRequest) SetSubType(v string) *RunLogicModelRequest {
	s.SubType = &v
	return s
}

type RunLogicModelResponseBody struct {
	Data      *RunLogicModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunLogicModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunLogicModelResponseBody) GoString() string {
	return s.String()
}

func (s *RunLogicModelResponseBody) SetData(v *RunLogicModelResponseBodyData) *RunLogicModelResponseBody {
	s.Data = v
	return s
}

func (s *RunLogicModelResponseBody) SetRequestId(v string) *RunLogicModelResponseBody {
	s.RequestId = &v
	return s
}

type RunLogicModelResponseBodyData struct {
	Body    *string            `json:"Body,omitempty" xml:"Body,omitempty"`
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Status  *int32             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RunLogicModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunLogicModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunLogicModelResponseBodyData) SetBody(v string) *RunLogicModelResponseBodyData {
	s.Body = &v
	return s
}

func (s *RunLogicModelResponseBodyData) SetHeaders(v map[string]*string) *RunLogicModelResponseBodyData {
	s.Headers = v
	return s
}

func (s *RunLogicModelResponseBodyData) SetStatus(v int32) *RunLogicModelResponseBodyData {
	s.Status = &v
	return s
}

type RunLogicModelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLogicModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLogicModelResponse) String() string {
	return tea.Prettify(s)
}

func (s RunLogicModelResponse) GoString() string {
	return s.String()
}

func (s *RunLogicModelResponse) SetHeaders(v map[string]*string) *RunLogicModelResponse {
	s.Headers = v
	return s
}

func (s *RunLogicModelResponse) SetStatusCode(v int32) *RunLogicModelResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLogicModelResponse) SetBody(v *RunLogicModelResponseBody) *RunLogicModelResponse {
	s.Body = v
	return s
}

type SetEnvironmentDefaultDomainRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EnvId      *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SetEnvironmentDefaultDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetEnvironmentDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetEnvironmentDefaultDomainRequest) SetAppId(v string) *SetEnvironmentDefaultDomainRequest {
	s.AppId = &v
	return s
}

func (s *SetEnvironmentDefaultDomainRequest) SetDomain(v string) *SetEnvironmentDefaultDomainRequest {
	s.Domain = &v
	return s
}

func (s *SetEnvironmentDefaultDomainRequest) SetDomainType(v string) *SetEnvironmentDefaultDomainRequest {
	s.DomainType = &v
	return s
}

func (s *SetEnvironmentDefaultDomainRequest) SetEnvId(v string) *SetEnvironmentDefaultDomainRequest {
	s.EnvId = &v
	return s
}

func (s *SetEnvironmentDefaultDomainRequest) SetSource(v string) *SetEnvironmentDefaultDomainRequest {
	s.Source = &v
	return s
}

type SetEnvironmentDefaultDomainResponseBody struct {
	Data      *SetEnvironmentDefaultDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetEnvironmentDefaultDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetEnvironmentDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetEnvironmentDefaultDomainResponseBody) SetData(v *SetEnvironmentDefaultDomainResponseBodyData) *SetEnvironmentDefaultDomainResponseBody {
	s.Data = v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBody) SetRequestId(v string) *SetEnvironmentDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

type SetEnvironmentDefaultDomainResponseBodyData struct {
	ConfigChanged       *bool   `json:"ConfigChanged,omitempty" xml:"ConfigChanged,omitempty"`
	DefaultMasterDomain *string `json:"DefaultMasterDomain,omitempty" xml:"DefaultMasterDomain,omitempty"`
	DefaultStaticDomain *string `json:"DefaultStaticDomain,omitempty" xml:"DefaultStaticDomain,omitempty"`
	MasterDomain        *string `json:"MasterDomain,omitempty" xml:"MasterDomain,omitempty"`
	MasterDomainApplied *bool   `json:"MasterDomainApplied,omitempty" xml:"MasterDomainApplied,omitempty"`
	StaticDomain        *string `json:"StaticDomain,omitempty" xml:"StaticDomain,omitempty"`
	StaticDomainApplied *bool   `json:"StaticDomainApplied,omitempty" xml:"StaticDomainApplied,omitempty"`
}

func (s SetEnvironmentDefaultDomainResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetEnvironmentDefaultDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetConfigChanged(v bool) *SetEnvironmentDefaultDomainResponseBodyData {
	s.ConfigChanged = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetDefaultMasterDomain(v string) *SetEnvironmentDefaultDomainResponseBodyData {
	s.DefaultMasterDomain = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetDefaultStaticDomain(v string) *SetEnvironmentDefaultDomainResponseBodyData {
	s.DefaultStaticDomain = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetMasterDomain(v string) *SetEnvironmentDefaultDomainResponseBodyData {
	s.MasterDomain = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetMasterDomainApplied(v bool) *SetEnvironmentDefaultDomainResponseBodyData {
	s.MasterDomainApplied = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetStaticDomain(v string) *SetEnvironmentDefaultDomainResponseBodyData {
	s.StaticDomain = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponseBodyData) SetStaticDomainApplied(v bool) *SetEnvironmentDefaultDomainResponseBodyData {
	s.StaticDomainApplied = &v
	return s
}

type SetEnvironmentDefaultDomainResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEnvironmentDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEnvironmentDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetEnvironmentDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *SetEnvironmentDefaultDomainResponse) SetHeaders(v map[string]*string) *SetEnvironmentDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *SetEnvironmentDefaultDomainResponse) SetStatusCode(v int32) *SetEnvironmentDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEnvironmentDefaultDomainResponse) SetBody(v *SetEnvironmentDefaultDomainResponseBody) *SetEnvironmentDefaultDomainResponse {
	s.Body = v
	return s
}

type StartAppServerRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s StartAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAppServerRequest) GoString() string {
	return s.String()
}

func (s *StartAppServerRequest) SetAppId(v string) *StartAppServerRequest {
	s.AppId = &v
	return s
}

func (s *StartAppServerRequest) SetEnvId(v string) *StartAppServerRequest {
	s.EnvId = &v
	return s
}

func (s *StartAppServerRequest) SetSource(v string) *StartAppServerRequest {
	s.Source = &v
	return s
}

type StartAppServerResponseBody struct {
	Data      *StartAppServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *StartAppServerResponseBody) SetData(v *StartAppServerResponseBodyData) *StartAppServerResponseBody {
	s.Data = v
	return s
}

func (s *StartAppServerResponseBody) SetRequestId(v string) *StartAppServerResponseBody {
	s.RequestId = &v
	return s
}

type StartAppServerResponseBodyData struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppServerStatus *string `json:"AppServerStatus,omitempty" xml:"AppServerStatus,omitempty"`
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StartAppServerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartAppServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartAppServerResponseBodyData) SetAppId(v string) *StartAppServerResponseBodyData {
	s.AppId = &v
	return s
}

func (s *StartAppServerResponseBodyData) SetAppServerStatus(v string) *StartAppServerResponseBodyData {
	s.AppServerStatus = &v
	return s
}

func (s *StartAppServerResponseBodyData) SetEnvId(v string) *StartAppServerResponseBodyData {
	s.EnvId = &v
	return s
}

type StartAppServerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAppServerResponse) GoString() string {
	return s.String()
}

func (s *StartAppServerResponse) SetHeaders(v map[string]*string) *StartAppServerResponse {
	s.Headers = v
	return s
}

func (s *StartAppServerResponse) SetStatusCode(v int32) *StartAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAppServerResponse) SetBody(v *StartAppServerResponseBody) *StartAppServerResponse {
	s.Body = v
	return s
}

type StopAppServerRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvId  *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s StopAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppServerRequest) GoString() string {
	return s.String()
}

func (s *StopAppServerRequest) SetAppId(v string) *StopAppServerRequest {
	s.AppId = &v
	return s
}

func (s *StopAppServerRequest) SetEnvId(v string) *StopAppServerRequest {
	s.EnvId = &v
	return s
}

func (s *StopAppServerRequest) SetSource(v string) *StopAppServerRequest {
	s.Source = &v
	return s
}

type StopAppServerResponseBody struct {
	Data      *StopAppServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppServerResponseBody) SetData(v *StopAppServerResponseBodyData) *StopAppServerResponseBody {
	s.Data = v
	return s
}

func (s *StopAppServerResponseBody) SetRequestId(v string) *StopAppServerResponseBody {
	s.RequestId = &v
	return s
}

type StopAppServerResponseBodyData struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppServerStatus *string `json:"AppServerStatus,omitempty" xml:"AppServerStatus,omitempty"`
	EnvId           *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s StopAppServerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopAppServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopAppServerResponseBodyData) SetAppId(v string) *StopAppServerResponseBodyData {
	s.AppId = &v
	return s
}

func (s *StopAppServerResponseBodyData) SetAppServerStatus(v string) *StopAppServerResponseBodyData {
	s.AppServerStatus = &v
	return s
}

func (s *StopAppServerResponseBodyData) SetEnvId(v string) *StopAppServerResponseBodyData {
	s.EnvId = &v
	return s
}

type StopAppServerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppServerResponse) GoString() string {
	return s.String()
}

func (s *StopAppServerResponse) SetHeaders(v map[string]*string) *StopAppServerResponse {
	s.Headers = v
	return s
}

func (s *StopAppServerResponse) SetStatusCode(v int32) *StopAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppServerResponse) SetBody(v *StopAppServerResponseBody) *StopAppServerResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetAppId(v string) *UpdateAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppRequest) SetAppName(v string) *UpdateAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppRequest) SetDescription(v string) *UpdateAppRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppRequest) SetIcon(v string) *UpdateAppRequest {
	s.Icon = &v
	return s
}

func (s *UpdateAppRequest) SetSource(v string) *UpdateAppRequest {
	s.Source = &v
	return s
}

type UpdateAppResponseBody struct {
	Data      *UpdateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetData(v *UpdateAppResponseBodyData) *UpdateAppResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppResponseBodyData struct {
	AppId         *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus     *string                                `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppType       *string                                `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Categories    []*UpdateAppResponseBodyDataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreateTime    *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon          *string                                `json:"Icon,omitempty" xml:"Icon,omitempty"`
	IsTemplate    *bool                                  `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	LastEditTime  *string                                `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	MainModuleId  *string                                `json:"MainModuleId,omitempty" xml:"MainModuleId,omitempty"`
	ModifiedTime  *string                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SchemaVersion *string                                `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string                                `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBodyData) SetAppId(v string) *UpdateAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetAppName(v string) *UpdateAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetAppStatus(v string) *UpdateAppResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetAppType(v string) *UpdateAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetCategories(v []*UpdateAppResponseBodyDataCategories) *UpdateAppResponseBodyData {
	s.Categories = v
	return s
}

func (s *UpdateAppResponseBodyData) SetCreateTime(v string) *UpdateAppResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetDescription(v string) *UpdateAppResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetIcon(v string) *UpdateAppResponseBodyData {
	s.Icon = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetIsTemplate(v bool) *UpdateAppResponseBodyData {
	s.IsTemplate = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetLastEditTime(v string) *UpdateAppResponseBodyData {
	s.LastEditTime = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetMainModuleId(v string) *UpdateAppResponseBodyData {
	s.MainModuleId = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetModifiedTime(v string) *UpdateAppResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetSchemaVersion(v string) *UpdateAppResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateAppResponseBodyData) SetSource(v string) *UpdateAppResponseBodyData {
	s.Source = &v
	return s
}

type UpdateAppResponseBodyDataCategories struct {
	CategoryId       *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName     *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s UpdateAppResponseBodyDataCategories) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBodyDataCategories) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBodyDataCategories) SetCategoryId(v string) *UpdateAppResponseBodyDataCategories {
	s.CategoryId = &v
	return s
}

func (s *UpdateAppResponseBodyDataCategories) SetCategoryName(v string) *UpdateAppResponseBodyDataCategories {
	s.CategoryName = &v
	return s
}

func (s *UpdateAppResponseBodyDataCategories) SetParentCategoryId(v string) *UpdateAppResponseBodyDataCategories {
	s.ParentCategoryId = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateAppModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EncodeType    *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s UpdateAppModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppModelRequest) SetAppId(v string) *UpdateAppModelRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppModelRequest) SetContent(v string) *UpdateAppModelRequest {
	s.Content = &v
	return s
}

func (s *UpdateAppModelRequest) SetEncodeType(v string) *UpdateAppModelRequest {
	s.EncodeType = &v
	return s
}

func (s *UpdateAppModelRequest) SetSchemaVersion(v string) *UpdateAppModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateAppModelRequest) SetSource(v string) *UpdateAppModelRequest {
	s.Source = &v
	return s
}

func (s *UpdateAppModelRequest) SetSubType(v string) *UpdateAppModelRequest {
	s.SubType = &v
	return s
}

type UpdateAppModelResponseBody struct {
	Data      *UpdateAppModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppModelResponseBody) SetData(v *UpdateAppModelResponseBodyData) *UpdateAppModelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAppModelResponseBody) SetRequestId(v string) *UpdateAppModelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateAppModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAppModelResponseBodyData) SetAppId(v string) *UpdateAppModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetAttributes(v []map[string]*string) *UpdateAppModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetContent(v map[string]*string) *UpdateAppModelResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetCreateTime(v string) *UpdateAppModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetDescription(v string) *UpdateAppModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetId(v string) *UpdateAppModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetLinkModelId(v string) *UpdateAppModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetLinkModuleId(v string) *UpdateAppModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetLinked(v bool) *UpdateAppModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModelDigest(v string) *UpdateAppModelResponseBodyData {
	s.ModelDigest = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModelId(v string) *UpdateAppModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModelName(v string) *UpdateAppModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModelStatus(v string) *UpdateAppModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModelType(v string) *UpdateAppModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModifiedTime(v string) *UpdateAppModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetModuleId(v string) *UpdateAppModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetProps(v map[string]*string) *UpdateAppModelResponseBodyData {
	s.Props = v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetRevision(v int32) *UpdateAppModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetSchemaVersion(v string) *UpdateAppModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetSubType(v string) *UpdateAppModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *UpdateAppModelResponseBodyData) SetVisibility(v string) *UpdateAppModelResponseBodyData {
	s.Visibility = &v
	return s
}

type UpdateAppModelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppModelResponse) SetHeaders(v map[string]*string) *UpdateAppModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppModelResponse) SetStatusCode(v int32) *UpdateAppModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppModelResponse) SetBody(v *UpdateAppModelResponseBody) *UpdateAppModelResponse {
	s.Body = v
	return s
}

type UpdateModelRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncodeType    *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModuleId      *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) SetAppId(v string) *UpdateModelRequest {
	s.AppId = &v
	return s
}

func (s *UpdateModelRequest) SetContent(v string) *UpdateModelRequest {
	s.Content = &v
	return s
}

func (s *UpdateModelRequest) SetDescription(v string) *UpdateModelRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelRequest) SetEncodeType(v string) *UpdateModelRequest {
	s.EncodeType = &v
	return s
}

func (s *UpdateModelRequest) SetModelId(v string) *UpdateModelRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateModelRequest) SetModelName(v string) *UpdateModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateModelRequest) SetModuleId(v string) *UpdateModelRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateModelRequest) SetSchemaVersion(v string) *UpdateModelRequest {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateModelRequest) SetSource(v string) *UpdateModelRequest {
	s.Source = &v
	return s
}

type UpdateModelResponseBody struct {
	Data      *UpdateModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBody) SetData(v *UpdateModelResponseBodyData) *UpdateModelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelResponseBody) SetRequestId(v string) *UpdateModelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModelResponseBodyData struct {
	AppId         *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Attributes    []map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Content       map[string]*string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string              `json:"Id,omitempty" xml:"Id,omitempty"`
	LinkModelId   *string              `json:"LinkModelId,omitempty" xml:"LinkModelId,omitempty"`
	LinkModuleId  *string              `json:"LinkModuleId,omitempty" xml:"LinkModuleId,omitempty"`
	Linked        *bool                `json:"Linked,omitempty" xml:"Linked,omitempty"`
	ModelDigest   *string              `json:"ModelDigest,omitempty" xml:"ModelDigest,omitempty"`
	ModelId       *string              `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string              `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus   *string              `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	ModelType     *string              `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ModifiedTime  *string              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string              `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	Props         map[string]*string   `json:"Props,omitempty" xml:"Props,omitempty"`
	Revision      *int32               `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string              `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	SubType       *string              `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Visibility    *string              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBodyData) SetAppId(v string) *UpdateModelResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetAttributes(v []map[string]*string) *UpdateModelResponseBodyData {
	s.Attributes = v
	return s
}

func (s *UpdateModelResponseBodyData) SetContent(v map[string]*string) *UpdateModelResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateModelResponseBodyData) SetCreateTime(v string) *UpdateModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetDescription(v string) *UpdateModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetId(v string) *UpdateModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetLinkModelId(v string) *UpdateModelResponseBodyData {
	s.LinkModelId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetLinkModuleId(v string) *UpdateModelResponseBodyData {
	s.LinkModuleId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetLinked(v bool) *UpdateModelResponseBodyData {
	s.Linked = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelDigest(v string) *UpdateModelResponseBodyData {
	s.ModelDigest = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelId(v string) *UpdateModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelName(v string) *UpdateModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelStatus(v string) *UpdateModelResponseBodyData {
	s.ModelStatus = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelType(v string) *UpdateModelResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModifiedTime(v string) *UpdateModelResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModuleId(v string) *UpdateModelResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetProps(v map[string]*string) *UpdateModelResponseBodyData {
	s.Props = v
	return s
}

func (s *UpdateModelResponseBodyData) SetRevision(v int32) *UpdateModelResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetSchemaVersion(v string) *UpdateModelResponseBodyData {
	s.SchemaVersion = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetSubType(v string) *UpdateModelResponseBodyData {
	s.SubType = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetVisibility(v string) *UpdateModelResponseBodyData {
	s.Visibility = &v
	return s
}

type UpdateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelResponse) SetHeaders(v map[string]*string) *UpdateModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelResponse) SetStatusCode(v int32) *UpdateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelResponse) SetBody(v *UpdateModelResponseBody) *UpdateModelResponse {
	s.Body = v
	return s
}

type UpdateModuleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleId    *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName  *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateModuleRequest) SetDescription(v string) *UpdateModuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateModuleRequest) SetModuleId(v string) *UpdateModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateModuleRequest) SetModuleName(v string) *UpdateModuleRequest {
	s.ModuleName = &v
	return s
}

func (s *UpdateModuleRequest) SetSource(v string) *UpdateModuleRequest {
	s.Source = &v
	return s
}

type UpdateModuleResponseBody struct {
	Data      *UpdateModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModuleResponseBody) SetData(v *UpdateModuleResponseBodyData) *UpdateModuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModuleResponseBody) SetRequestId(v string) *UpdateModuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModuleResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon                   *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	LatestPublishedCommit  *string `json:"LatestPublishedCommit,omitempty" xml:"LatestPublishedCommit,omitempty"`
	LatestPublishedVersion *string `json:"LatestPublishedVersion,omitempty" xml:"LatestPublishedVersion,omitempty"`
	MinimumPlatformVersion *string `json:"MinimumPlatformVersion,omitempty" xml:"MinimumPlatformVersion,omitempty"`
	ModifiedTime           *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId               *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName             *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	OwnerAppId             *string `json:"OwnerAppId,omitempty" xml:"OwnerAppId,omitempty"`
	OwnerUserId            *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s UpdateModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModuleResponseBodyData) SetCreateTime(v string) *UpdateModuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetDescription(v string) *UpdateModuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetIcon(v string) *UpdateModuleResponseBodyData {
	s.Icon = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetLatestPublishedCommit(v string) *UpdateModuleResponseBodyData {
	s.LatestPublishedCommit = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetLatestPublishedVersion(v string) *UpdateModuleResponseBodyData {
	s.LatestPublishedVersion = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetMinimumPlatformVersion(v string) *UpdateModuleResponseBodyData {
	s.MinimumPlatformVersion = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetModifiedTime(v string) *UpdateModuleResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetModuleId(v string) *UpdateModuleResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetModuleName(v string) *UpdateModuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetOwnerAppId(v string) *UpdateModuleResponseBodyData {
	s.OwnerAppId = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetOwnerUserId(v string) *UpdateModuleResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateModuleResponseBodyData) SetPlatform(v string) *UpdateModuleResponseBodyData {
	s.Platform = &v
	return s
}

type UpdateModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateModuleResponse) SetHeaders(v map[string]*string) *UpdateModuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateModuleResponse) SetStatusCode(v int32) *UpdateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModuleResponse) SetBody(v *UpdateModuleResponseBody) *UpdateModuleResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleId     *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetAppId(v string) *UpdateResourceRequest {
	s.AppId = &v
	return s
}

func (s *UpdateResourceRequest) SetContent(v string) *UpdateResourceRequest {
	s.Content = &v
	return s
}

func (s *UpdateResourceRequest) SetDescription(v string) *UpdateResourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceRequest) SetModuleId(v string) *UpdateResourceRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceRequest) SetResourceId(v string) *UpdateResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceRequest) SetResourceName(v string) *UpdateResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceRequest) SetSource(v string) *UpdateResourceRequest {
	s.Source = &v
	return s
}

type UpdateResourceResponseBody struct {
	Data      *UpdateResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetData(v *UpdateResourceResponseBodyData) *UpdateResourceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceResponseBodyData struct {
	AppId          *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime     *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime   *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId       *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceDigest *string            `json:"ResourceDigest,omitempty" xml:"ResourceDigest,omitempty"`
	ResourceId     *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision       *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion  *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s UpdateResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBodyData) SetAppId(v string) *UpdateResourceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetContent(v map[string]*string) *UpdateResourceResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateResourceResponseBodyData) SetCreateTime(v string) *UpdateResourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetDescription(v string) *UpdateResourceResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetModifiedTime(v string) *UpdateResourceResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetModuleId(v string) *UpdateResourceResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetResourceDigest(v string) *UpdateResourceResponseBodyData {
	s.ResourceDigest = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetResourceId(v string) *UpdateResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetResourceName(v string) *UpdateResourceResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetResourceType(v string) *UpdateResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetRevision(v int32) *UpdateResourceResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateResourceResponseBodyData) SetSchemaVersion(v string) *UpdateResourceResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateResourceContentRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ModuleId   *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateResourceContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceContentRequest) SetAppId(v string) *UpdateResourceContentRequest {
	s.AppId = &v
	return s
}

func (s *UpdateResourceContentRequest) SetContent(v string) *UpdateResourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateResourceContentRequest) SetModuleId(v string) *UpdateResourceContentRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceContentRequest) SetResourceId(v string) *UpdateResourceContentRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceContentRequest) SetSource(v string) *UpdateResourceContentRequest {
	s.Source = &v
	return s
}

type UpdateResourceContentResponseBody struct {
	Data      *UpdateResourceContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceContentResponseBody) SetData(v *UpdateResourceContentResponseBodyData) *UpdateResourceContentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateResourceContentResponseBody) SetRequestId(v string) *UpdateResourceContentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceContentResponseBodyData struct {
	AppId         *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime  *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId    *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision      *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s UpdateResourceContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateResourceContentResponseBodyData) SetAppId(v string) *UpdateResourceContentResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetContent(v map[string]*string) *UpdateResourceContentResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetCreateTime(v string) *UpdateResourceContentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetDescription(v string) *UpdateResourceContentResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetModifiedTime(v string) *UpdateResourceContentResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetModuleId(v string) *UpdateResourceContentResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetResourceId(v string) *UpdateResourceContentResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetResourceName(v string) *UpdateResourceContentResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetResourceType(v string) *UpdateResourceContentResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetRevision(v int32) *UpdateResourceContentResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateResourceContentResponseBodyData) SetSchemaVersion(v string) *UpdateResourceContentResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type UpdateResourceContentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceContentResponse) SetHeaders(v map[string]*string) *UpdateResourceContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceContentResponse) SetStatusCode(v int32) *UpdateResourceContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceContentResponse) SetBody(v *UpdateResourceContentResponseBody) *UpdateResourceContentResponse {
	s.Body = v
	return s
}

type UpdateResourceInfoRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleId     *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInfoRequest) SetAppId(v string) *UpdateResourceInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateResourceInfoRequest) SetDescription(v string) *UpdateResourceInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceInfoRequest) SetModuleId(v string) *UpdateResourceInfoRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceInfoRequest) SetResourceId(v string) *UpdateResourceInfoRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceInfoRequest) SetResourceName(v string) *UpdateResourceInfoRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceInfoRequest) SetSource(v string) *UpdateResourceInfoRequest {
	s.Source = &v
	return s
}

type UpdateResourceInfoResponseBody struct {
	Data      *UpdateResourceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInfoResponseBody) SetData(v *UpdateResourceInfoResponseBodyData) *UpdateResourceInfoResponseBody {
	s.Data = v
	return s
}

func (s *UpdateResourceInfoResponseBody) SetRequestId(v string) *UpdateResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceInfoResponseBodyData struct {
	AppId         *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content       map[string]*string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime  *string            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModuleId      *string            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ResourceId    *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName  *string            `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Revision      *int32             `json:"Revision,omitempty" xml:"Revision,omitempty"`
	SchemaVersion *string            `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s UpdateResourceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateResourceInfoResponseBodyData) SetAppId(v string) *UpdateResourceInfoResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetContent(v map[string]*string) *UpdateResourceInfoResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetCreateTime(v string) *UpdateResourceInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetDescription(v string) *UpdateResourceInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetModifiedTime(v string) *UpdateResourceInfoResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetModuleId(v string) *UpdateResourceInfoResponseBodyData {
	s.ModuleId = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetResourceId(v string) *UpdateResourceInfoResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetResourceName(v string) *UpdateResourceInfoResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetResourceType(v string) *UpdateResourceInfoResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetRevision(v int32) *UpdateResourceInfoResponseBodyData {
	s.Revision = &v
	return s
}

func (s *UpdateResourceInfoResponseBodyData) SetSchemaVersion(v string) *UpdateResourceInfoResponseBodyData {
	s.SchemaVersion = &v
	return s
}

type UpdateResourceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInfoResponse) SetHeaders(v map[string]*string) *UpdateResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInfoResponse) SetStatusCode(v int32) *UpdateResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceInfoResponse) SetBody(v *UpdateResourceInfoResponseBody) *UpdateResourceInfoResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("miniapplcdp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchCreateModelWithOptions(request *BatchCreateModelRequest, runtime *util.RuntimeOptions) (_result *BatchCreateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDataJson)) {
		query["ModelDataJson"] = request.ModelDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreateModel(request *BatchCreateModelRequest) (_result *BatchCreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCreateModelResponse{}
	_body, _err := client.BatchCreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteModelWithOptions(request *BatchDeleteModelRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelIdList)) {
		query["ModelIdList"] = request.ModelIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteModel(request *BatchDeleteModelRequest) (_result *BatchDeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteModelResponse{}
	_body, _err := client.BatchDeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteResourcesWithOptions(request *BatchDeleteResourcesRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdList)) {
		query["ResourceIdList"] = request.ResourceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteResources"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteResources(request *BatchDeleteResourcesRequest) (_result *BatchDeleteResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteResourcesResponse{}
	_body, _err := client.BatchDeleteResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRestoreModelWithOptions(request *BatchRestoreModelRequest, runtime *util.RuntimeOptions) (_result *BatchRestoreModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelIdList)) {
		query["ModelIdList"] = request.ModelIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRestoreModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRestoreModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRestoreModel(request *BatchRestoreModelRequest) (_result *BatchRestoreModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchRestoreModelResponse{}
	_body, _err := client.BatchRestoreModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *util.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomain"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDomain(request *CheckDomainRequest) (_result *CheckDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainResponse{}
	_body, _err := client.CheckDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneAppWithOptions(request *CloneAppRequest, runtime *util.RuntimeOptions) (_result *CloneAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		query["Icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneApp"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneApp(request *CloneAppRequest) (_result *CloneAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneAppResponse{}
	_body, _err := client.CloneAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneModelFromCommitWithOptions(request *CloneModelFromCommitRequest, runtime *util.RuntimeOptions) (_result *CloneModelFromCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCommitId)) {
		query["SourceCommitId"] = request.SourceCommitId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceModuleId)) {
		query["SourceModuleId"] = request.SourceModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetModuleId)) {
		query["TargetModuleId"] = request.TargetModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetName)) {
		query["TargetName"] = request.TargetName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSubType)) {
		query["TargetSubType"] = request.TargetSubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneModelFromCommit"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneModelFromCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneModelFromCommit(request *CloneModelFromCommitRequest) (_result *CloneModelFromCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneModelFromCommitResponse{}
	_body, _err := client.CloneModelFromCommitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneModelInModuleWithOptions(request *CloneModelInModuleRequest, runtime *util.RuntimeOptions) (_result *CloneModelInModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TargetName)) {
		query["TargetName"] = request.TargetName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSubType)) {
		query["TargetSubType"] = request.TargetSubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneModelInModule"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneModelInModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneModelInModule(request *CloneModelInModuleRequest) (_result *CloneModelInModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneModelInModuleResponse{}
	_body, _err := client.CloneModelInModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Asynchronous)) {
		query["Asynchronous"] = request.Asynchronous
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		query["Icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformVersion)) {
		query["PlatformVersion"] = request.PlatformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCommitId)) {
		query["SourceCommitId"] = request.SourceCommitId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Templated)) {
		query["Templated"] = request.Templated
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCommitWithOptions(request *CreateCommitRequest, runtime *util.RuntimeOptions) (_result *CreateCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommitLog)) {
		query["CommitLog"] = request.CommitLog
	}

	if !tea.BoolValue(util.IsUnset(request.CommitType)) {
		query["CommitType"] = request.CommitType
	}

	if !tea.BoolValue(util.IsUnset(request.MainModuleCommitId)) {
		query["MainModuleCommitId"] = request.MainModuleCommitId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RollbackToCommitId)) {
		query["RollbackToCommitId"] = request.RollbackToCommitId
	}

	if !tea.BoolValue(util.IsUnset(request.RollbackType)) {
		query["RollbackType"] = request.RollbackType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCommit"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCommit(request *CreateCommitRequest) (_result *CreateCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCommitResponse{}
	_body, _err := client.CreateCommitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKey)) {
		query["PublicKey"] = request.PublicKey
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WithCertificate)) {
		query["WithCertificate"] = request.WithCertificate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLinkEntityAndAssociationWithOptions(request *CreateLinkEntityAndAssociationRequest, runtime *util.RuntimeOptions) (_result *CreateLinkEntityAndAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModelData)) {
		query["ModelData"] = request.ModelData
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLinkEntityAndAssociation"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLinkEntityAndAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLinkEntityAndAssociation(request *CreateLinkEntityAndAssociationRequest) (_result *CreateLinkEntityAndAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLinkEntityAndAssociationResponse{}
	_body, _err := client.CreateLinkEntityAndAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelWithOptions(request *CreateModelRequest, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.LinkModelId)) {
		query["LinkModelId"] = request.LinkModelId
	}

	if !tea.BoolValue(util.IsUnset(request.LinkModuleId)) {
		query["LinkModuleId"] = request.LinkModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Linked)) {
		query["Linked"] = request.Linked
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModuleWithOptions(request *CreateModuleRequest, runtime *util.RuntimeOptions) (_result *CreateModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		query["Icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.MinimumPlatformVersion)) {
		query["MinimumPlatformVersion"] = request.MinimumPlatformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		query["ModuleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceModuleId)) {
		query["SourceModuleId"] = request.SourceModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAppSource)) {
		query["TargetAppSource"] = request.TargetAppSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModule"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModule(request *CreateModuleRequest) (_result *CreateModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModuleResponse{}
	_body, _err := client.CreateModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModulePublishWithOptions(request *CreateModulePublishRequest, runtime *util.RuntimeOptions) (_result *CreateModulePublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PublishVersion)) {
		query["PublishVersion"] = request.PublishVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModulePublish"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModulePublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModulePublish(request *CreateModulePublishRequest) (_result *CreateModulePublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModulePublishResponse{}
	_body, _err := client.CreateModulePublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePublishWithOptions(request *CreatePublishRequest, runtime *util.RuntimeOptions) (_result *CreatePublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommitId)) {
		query["CommitId"] = request.CommitId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.PublishType)) {
		query["PublishType"] = request.PublishType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.VersionNumber)) {
		query["VersionNumber"] = request.VersionNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublish"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePublish(request *CreatePublishRequest) (_result *CreatePublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublishResponse{}
	_body, _err := client.CreatePublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommitWithOptions(request *DeleteCommitRequest, runtime *util.RuntimeOptions) (_result *DeleteCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CommitId)) {
		query["CommitId"] = request.CommitId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCommit"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCommit(request *DeleteCommitRequest) (_result *DeleteCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommitResponse{}
	_body, _err := client.DeleteCommitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModuleWithOptions(request *DeleteModuleRequest, runtime *util.RuntimeOptions) (_result *DeleteModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModule"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModule(request *DeleteModuleRequest) (_result *DeleteModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteModuleResponse{}
	_body, _err := client.DeleteModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAppUserPasswordWithOptions(request *GenerateAppUserPasswordRequest, runtime *util.RuntimeOptions) (_result *GenerateAppUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAppUserPassword"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAppUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAppUserPassword(request *GenerateAppUserPasswordRequest) (_result *GenerateAppUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAppUserPasswordResponse{}
	_body, _err := client.GenerateAppUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAuthTokenWithOptions(request *GenerateAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAuthToken"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAuthToken(request *GenerateAuthTokenRequest) (_result *GenerateAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAuthTokenResponse{}
	_body, _err := client.GenerateAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadTokenWithOptions(request *GenerateUploadTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialId)) {
		query["MaterialId"] = request.MaterialId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UploadTokenType)) {
		query["UploadTokenType"] = request.UploadTokenType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadToken"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadToken(request *GenerateUploadTokenRequest) (_result *GenerateUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadTokenResponse{}
	_body, _err := client.GenerateUploadTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppModelWithOptions(request *GetAppModelRequest, runtime *util.RuntimeOptions) (_result *GetAppModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppModel(request *GetAppModelRequest) (_result *GetAppModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppModelResponse{}
	_body, _err := client.GetAppModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppSecretWithOptions(request *GetAppSecretRequest, runtime *util.RuntimeOptions) (_result *GetAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppSecret"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppSecret(request *GetAppSecretRequest) (_result *GetAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSecretResponse{}
	_body, _err := client.GetAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetArtifactWithOptions(request *GetArtifactRequest, runtime *util.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifact"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetArtifact(request *GetArtifactRequest) (_result *GetArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactResponse{}
	_body, _err := client.GetArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCommitWithOptions(request *GetCommitRequest, runtime *util.RuntimeOptions) (_result *GetCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CommitId)) {
		query["CommitId"] = request.CommitId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCommit"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCommit(request *GetCommitRequest) (_result *GetCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommitResponse{}
	_body, _err := client.GetCommitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultAppUserWithOptions(request *GetDefaultAppUserRequest, runtime *util.RuntimeOptions) (_result *GetDefaultAppUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDefaultAppUser"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultAppUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultAppUser(request *GetDefaultAppUserRequest) (_result *GetDefaultAppUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultAppUserResponse{}
	_body, _err := client.GetDefaultAppUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainCnameWithOptions(request *GetDomainCnameRequest, runtime *util.RuntimeOptions) (_result *GetDomainCnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainCname"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainCname(request *GetDomainCnameRequest) (_result *GetDomainCnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainCnameResponse{}
	_body, _err := client.GetDomainCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainOverviewWithOptions(request *GetDomainOverviewRequest, runtime *util.RuntimeOptions) (_result *GetDomainOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainOverview"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainOverview(request *GetDomainOverviewRequest) (_result *GetDomainOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainOverviewResponse{}
	_body, _err := client.GetDomainOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentWithOptions(request *GetEnvironmentRequest, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnvironment"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnvironment(request *GetEnvironmentRequest) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoryStatsWithOptions(request *GetHistoryStatsRequest, runtime *util.RuntimeOptions) (_result *GetHistoryStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryStats"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoryStats(request *GetHistoryStatsRequest) (_result *GetHistoryStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoryStatsResponse{}
	_body, _err := client.GetHistoryStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLatestCommitWithOptions(request *GetLatestCommitRequest, runtime *util.RuntimeOptions) (_result *GetLatestCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLatestCommit"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLatestCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLatestCommit(request *GetLatestCommitRequest) (_result *GetLatestCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLatestCommitResponse{}
	_body, _err := client.GetLatestCommitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelWithOptions(request *GetModelRequest, runtime *util.RuntimeOptions) (_result *GetModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModel(request *GetModelRequest) (_result *GetModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelResponse{}
	_body, _err := client.GetModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModuleWithOptions(request *GetModuleRequest, runtime *util.RuntimeOptions) (_result *GetModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		query["ModuleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetModule"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModule(request *GetModuleRequest) (_result *GetModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModuleResponse{}
	_body, _err := client.GetModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublishWithOptions(request *GetPublishRequest, runtime *util.RuntimeOptions) (_result *GetPublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PublishId)) {
		query["PublishId"] = request.PublishId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublish"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublish(request *GetPublishRequest) (_result *GetPublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublishResponse{}
	_body, _err := client.GetPublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealtimeStatsWithOptions(request *GetRealtimeStatsRequest, runtime *util.RuntimeOptions) (_result *GetRealtimeStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealtimeStats"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealtimeStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealtimeStats(request *GetRealtimeStatsRequest) (_result *GetRealtimeStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealtimeStatsResponse{}
	_body, _err := client.GetRealtimeStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageProcessParameter)) {
		query["ImageProcessParameter"] = request.ImageProcessParameter
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppModulesWithOptions(request *ListAppModulesRequest, runtime *util.RuntimeOptions) (_result *ListAppModulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppModules"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppModulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppModules(request *ListAppModulesRequest) (_result *ListAppModulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppModulesResponse{}
	_body, _err := client.ListAppModulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppTemplatesWithOptions(request *ListAppTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListAppTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppTemplates"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppTemplates(request *ListAppTemplatesRequest) (_result *ListAppTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.ListAppTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppStatus)) {
		query["AppStatus"] = request.AppStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CustomParentId)) {
		query["CustomParentId"] = request.CustomParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MainModuleId)) {
		query["MainModuleId"] = request.MainModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		query["Template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListArtifactsWithOptions(request *ListArtifactsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PublishId)) {
		query["PublishId"] = request.PublishId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifacts"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListArtifacts(request *ListArtifactsRequest) (_result *ListArtifactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactsResponse{}
	_body, _err := client.ListArtifactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommitsWithOptions(request *ListCommitsRequest, runtime *util.RuntimeOptions) (_result *ListCommitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CommitLog)) {
		query["CommitLog"] = request.CommitLog
	}

	if !tea.BoolValue(util.IsUnset(request.CustomParentId)) {
		query["CustomParentId"] = request.CustomParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommits"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommits(request *ListCommitsRequest) (_result *ListCommitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommitsResponse{}
	_body, _err := client.ListCommitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentOverviewsWithOptions(request *ListEnvironmentOverviewsRequest, runtime *util.RuntimeOptions) (_result *ListEnvironmentOverviewsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironmentOverviews"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentOverviewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentOverviews(request *ListEnvironmentOverviewsRequest) (_result *ListEnvironmentOverviewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnvironmentOverviewsResponse{}
	_body, _err := client.ListEnvironmentOverviewsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithOptions(request *ListEnvironmentsRequest, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnvironments"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironments(request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsWithOptions(request *ListModelsRequest, runtime *util.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModels"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModels(request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsByPageWithOptions(request *ListModelsByPageRequest, runtime *util.RuntimeOptions) (_result *ListModelsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModelsByPage"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelsByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModelsByPage(request *ListModelsByPageRequest) (_result *ListModelsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModelsByPageResponse{}
	_body, _err := client.ListModelsByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModuleDependenciesWithOptions(request *ListModuleDependenciesRequest, runtime *util.RuntimeOptions) (_result *ListModuleDependenciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModuleDependencies"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModuleDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModuleDependencies(request *ListModuleDependenciesRequest) (_result *ListModuleDependenciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModuleDependenciesResponse{}
	_body, _err := client.ListModuleDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModuleModelsWithOptions(request *ListModuleModelsRequest, runtime *util.RuntimeOptions) (_result *ListModuleModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleList)) {
		query["ModuleList"] = request.ModuleList
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubTypes)) {
		query["SubTypes"] = request.SubTypes
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModuleModels"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModuleModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModuleModels(request *ListModuleModelsRequest) (_result *ListModuleModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModuleModelsResponse{}
	_body, _err := client.ListModuleModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModulePublishVersionsWithOptions(request *ListModulePublishVersionsRequest, runtime *util.RuntimeOptions) (_result *ListModulePublishVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomParentId)) {
		query["CustomParentId"] = request.CustomParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		query["ModuleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModulePublishVersions"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModulePublishVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModulePublishVersions(request *ListModulePublishVersionsRequest) (_result *ListModulePublishVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModulePublishVersionsResponse{}
	_body, _err := client.ListModulePublishVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModuleResourcesWithOptions(request *ListModuleResourcesRequest, runtime *util.RuntimeOptions) (_result *ListModuleResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleList)) {
		query["ModuleList"] = request.ModuleList
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["Types"] = request.Types
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModuleResources"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModuleResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModuleResources(request *ListModuleResourcesRequest) (_result *ListModuleResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModuleResourcesResponse{}
	_body, _err := client.ListModuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModulesWithOptions(request *ListModulesRequest, runtime *util.RuntimeOptions) (_result *ListModulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HasOwnerApp)) {
		query["HasOwnerApp"] = request.HasOwnerApp
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModules"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModules(request *ListModulesRequest) (_result *ListModulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModulesResponse{}
	_body, _err := client.ListModulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModulesByPageWithOptions(request *ListModulesByPageRequest, runtime *util.RuntimeOptions) (_result *ListModulesByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomParentId)) {
		query["CustomParentId"] = request.CustomParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HasOwnerApp)) {
		query["HasOwnerApp"] = request.HasOwnerApp
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		query["ModuleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModulesByPage"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModulesByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModulesByPage(request *ListModulesByPageRequest) (_result *ListModulesByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModulesByPageResponse{}
	_body, _err := client.ListModulesByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishVersionsWithOptions(request *ListPublishVersionsRequest, runtime *util.RuntimeOptions) (_result *ListPublishVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishVersions"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishVersions(request *ListPublishVersionsRequest) (_result *ListPublishVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishVersionsResponse{}
	_body, _err := client.ListPublishVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishedModulesWithOptions(request *ListPublishedModulesRequest, runtime *util.RuntimeOptions) (_result *ListPublishedModulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeAppId)) {
		query["ExcludeAppId"] = request.ExcludeAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeModuleId)) {
		query["ExcludeModuleId"] = request.ExcludeModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.HasOwnerApp)) {
		query["HasOwnerApp"] = request.HasOwnerApp
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleType)) {
		query["ModuleType"] = request.ModuleType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishedModules"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishedModulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishedModules(request *ListPublishedModulesRequest) (_result *ListPublishedModulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishedModulesResponse{}
	_body, _err := client.ListPublishedModulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishesWithOptions(request *ListPublishesRequest, runtime *util.RuntimeOptions) (_result *ListPublishesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublishStatus)) {
		query["PublishStatus"] = request.PublishStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PublishType)) {
		query["PublishType"] = request.PublishType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishes"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishes(request *ListPublishesRequest) (_result *ListPublishesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishesResponse{}
	_body, _err := client.ListPublishesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageProcessParameter)) {
		query["ImageProcessParameter"] = request.ImageProcessParameter
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesByPageWithOptions(request *ListResourcesByPageRequest, runtime *util.RuntimeOptions) (_result *ListResourcesByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageProcessParameter)) {
		query["ImageProcessParameter"] = request.ImageProcessParameter
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WithContent)) {
		query["WithContent"] = request.WithContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcesByPage"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourcesByPage(request *ListResourcesByPageRequest) (_result *ListResourcesByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesByPageResponse{}
	_body, _err := client.ListResourcesByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAppUserPasswordWithOptions(request *ResetAppUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAppUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppUserPassword"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAppUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAppUserPassword(request *ResetAppUserPasswordRequest) (_result *ResetAppUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAppUserPasswordResponse{}
	_body, _err := client.ResetAppUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreModelWithOptions(request *RestoreModelRequest, runtime *util.RuntimeOptions) (_result *RestoreModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreModel(request *RestoreModelRequest) (_result *RestoreModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreModelResponse{}
	_body, _err := client.RestoreModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunLogicModelWithOptions(request *RunLogicModelRequest, runtime *util.RuntimeOptions) (_result *RunLogicModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CommitId)) {
		query["CommitId"] = request.CommitId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunLogicModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunLogicModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunLogicModel(request *RunLogicModelRequest) (_result *RunLogicModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunLogicModelResponse{}
	_body, _err := client.RunLogicModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetEnvironmentDefaultDomainWithOptions(request *SetEnvironmentDefaultDomainRequest, runtime *util.RuntimeOptions) (_result *SetEnvironmentDefaultDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainType)) {
		query["DomainType"] = request.DomainType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetEnvironmentDefaultDomain"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetEnvironmentDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetEnvironmentDefaultDomain(request *SetEnvironmentDefaultDomainRequest) (_result *SetEnvironmentDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetEnvironmentDefaultDomainResponse{}
	_body, _err := client.SetEnvironmentDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAppServerWithOptions(request *StartAppServerRequest, runtime *util.RuntimeOptions) (_result *StartAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAppServer"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAppServer(request *StartAppServerRequest) (_result *StartAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAppServerResponse{}
	_body, _err := client.StartAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppServerWithOptions(request *StopAppServerRequest, runtime *util.RuntimeOptions) (_result *StopAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["EnvId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAppServer"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppServer(request *StopAppServerRequest) (_result *StopAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppServerResponse{}
	_body, _err := client.StopAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		query["Icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppModelWithOptions(request *UpdateAppModelRequest, runtime *util.RuntimeOptions) (_result *UpdateAppModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		query["SubType"] = request.SubType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppModel(request *UpdateAppModelRequest) (_result *UpdateAppModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppModelResponse{}
	_body, _err := client.UpdateAppModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateModelWithOptions(request *UpdateModelRequest, runtime *util.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncodeType)) {
		query["EncodeType"] = request.EncodeType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModel"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateModel(request *UpdateModelRequest) (_result *UpdateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateModelResponse{}
	_body, _err := client.UpdateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateModuleWithOptions(request *UpdateModuleRequest, runtime *util.RuntimeOptions) (_result *UpdateModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModule"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateModule(request *UpdateModuleRequest) (_result *UpdateModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateModuleResponse{}
	_body, _err := client.UpdateModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceWithOptions(request *UpdateResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceContentWithOptions(request *UpdateResourceContentRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceContent"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceContent(request *UpdateResourceContentRequest) (_result *UpdateResourceContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceContentResponse{}
	_body, _err := client.UpdateResourceContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceInfoWithOptions(request *UpdateResourceInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["ModuleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceInfo"),
		Version:     tea.String("2020-01-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceInfo(request *UpdateResourceInfoRequest) (_result *UpdateResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceInfoResponse{}
	_body, _err := client.UpdateResourceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
