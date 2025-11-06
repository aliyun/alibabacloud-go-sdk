// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosHistoryConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListNacosHistoryConfigsResponseBody
	GetErrorCode() *string
	SetHistoryItems(v []*ListNacosHistoryConfigsResponseBodyHistoryItems) *ListNacosHistoryConfigsResponseBody
	GetHistoryItems() []*ListNacosHistoryConfigsResponseBodyHistoryItems
	SetHttpCode(v string) *ListNacosHistoryConfigsResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListNacosHistoryConfigsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListNacosHistoryConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNacosHistoryConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNacosHistoryConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNacosHistoryConfigsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListNacosHistoryConfigsResponseBody
	GetTotalCount() *int32
}

type ListNacosHistoryConfigsResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The configuration items.
	HistoryItems []*ListNacosHistoryConfigsResponseBodyHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 53338ECA-F880-54D8-A9B3-5606355A1B89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNacosHistoryConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNacosHistoryConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNacosHistoryConfigsResponseBody) GetHistoryItems() []*ListNacosHistoryConfigsResponseBodyHistoryItems {
	return s.HistoryItems
}

func (s *ListNacosHistoryConfigsResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListNacosHistoryConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNacosHistoryConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNacosHistoryConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNacosHistoryConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNacosHistoryConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNacosHistoryConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNacosHistoryConfigsResponseBody) SetErrorCode(v string) *ListNacosHistoryConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetHistoryItems(v []*ListNacosHistoryConfigsResponseBodyHistoryItems) *ListNacosHistoryConfigsResponseBody {
	s.HistoryItems = v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetHttpCode(v string) *ListNacosHistoryConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetMessage(v string) *ListNacosHistoryConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetPageNumber(v int32) *ListNacosHistoryConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetPageSize(v int32) *ListNacosHistoryConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetRequestId(v string) *ListNacosHistoryConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetSuccess(v bool) *ListNacosHistoryConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetTotalCount(v int32) *ListNacosHistoryConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) Validate() error {
	if s.HistoryItems != nil {
		for _, item := range s.HistoryItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNacosHistoryConfigsResponseBodyHistoryItems struct {
	// The application tag.
	//
	// example:
	//
	// gateway
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the data.
	//
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// default
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// 23fdsf
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timestamp when the configuration was last modified.
	//
	// example:
	//
	// 16434400
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The format of the configuration file.
	//
	// example:
	//
	// yaml
	OpType  *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	SrcUser *string `json:"SrcUser,omitempty" xml:"SrcUser,omitempty"`
}

func (s ListNacosHistoryConfigsResponseBodyHistoryItems) String() string {
	return dara.Prettify(s)
}

func (s ListNacosHistoryConfigsResponseBodyHistoryItems) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetAppName() *string {
	return s.AppName
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetDataId() *string {
	return s.DataId
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetGroup() *string {
	return s.Group
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetId() *int64 {
	return s.Id
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetOpType() *string {
	return s.OpType
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) GetSrcUser() *string {
	return s.SrcUser
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetAppName(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.AppName = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetDataId(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.DataId = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetGroup(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.Group = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetId(v int64) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.Id = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetLastModifiedTime(v int64) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.LastModifiedTime = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetOpType(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.OpType = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetSrcUser(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.SrcUser = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) Validate() error {
	return dara.Validate(s)
}
