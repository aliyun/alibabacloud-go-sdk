// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumersResponseBody
	GetCode() *string
	SetData(v *ListConsumersResponseBodyData) *ListConsumersResponseBody
	GetData() *ListConsumersResponseBodyData
	SetMessage(v string) *ListConsumersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumersResponseBody
	GetRequestId() *string
}

type ListConsumersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListConsumersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9BDD6A7C-CBA7-504F-B8C5-51B9F16590F7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumersResponseBody) GetData() *ListConsumersResponseBodyData {
	return s.Data
}

func (s *ListConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumersResponseBody) SetCode(v string) *ListConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumersResponseBody) SetData(v *ListConsumersResponseBodyData) *ListConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumersResponseBody) SetMessage(v string) *ListConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumersResponseBody) SetRequestId(v string) *ListConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConsumersResponseBodyData struct {
	Items []*ListConsumersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 18
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumersResponseBodyData) GetItems() []*ListConsumersResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumersResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListConsumersResponseBodyData) SetItems(v []*ListConsumersResponseBodyDataItems) *ListConsumersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumersResponseBodyData) SetPageNumber(v int32) *ListConsumersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumersResponseBodyData) SetPageSize(v int32) *ListConsumersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumersResponseBodyData) SetTotalSize(v int32) *ListConsumersResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListConsumersResponseBodyDataItems struct {
	// example:
	//
	// cs-csheiftlhtgmp0j0hp4g
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// 1721097861050
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// {}
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// user-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1721123855214
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s ListConsumersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumersResponseBodyDataItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListConsumersResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListConsumersResponseBodyDataItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListConsumersResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListConsumersResponseBodyDataItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListConsumersResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListConsumersResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListConsumersResponseBodyDataItems) SetConsumerId(v string) *ListConsumersResponseBodyDataItems {
	s.ConsumerId = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetCreateTimestamp(v int64) *ListConsumersResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetDeployStatus(v string) *ListConsumersResponseBodyDataItems {
	s.DeployStatus = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetDescription(v string) *ListConsumersResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetEnable(v bool) *ListConsumersResponseBodyDataItems {
	s.Enable = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetName(v string) *ListConsumersResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListConsumersResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListConsumersResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
