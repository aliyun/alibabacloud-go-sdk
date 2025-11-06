// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListNacosConfigsResponseBody
	GetCode() *int32
	SetConfigurations(v []*ListNacosConfigsResponseBodyConfigurations) *ListNacosConfigsResponseBody
	GetConfigurations() []*ListNacosConfigsResponseBodyConfigurations
	SetErrorCode(v string) *ListNacosConfigsResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListNacosConfigsResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListNacosConfigsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListNacosConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNacosConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNacosConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNacosConfigsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListNacosConfigsResponseBody
	GetTotalCount() *int32
}

type ListNacosConfigsResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations.
	Configurations []*ListNacosConfigsResponseBodyConfigurations `json:"Configurations,omitempty" xml:"Configurations,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 4081087F-3429-5873-A1E7-D4B5479D0B84
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
	// The total number of returned instances.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNacosConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNacosConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListNacosConfigsResponseBody) GetConfigurations() []*ListNacosConfigsResponseBodyConfigurations {
	return s.Configurations
}

func (s *ListNacosConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNacosConfigsResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListNacosConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNacosConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNacosConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNacosConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNacosConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNacosConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNacosConfigsResponseBody) SetCode(v int32) *ListNacosConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetConfigurations(v []*ListNacosConfigsResponseBodyConfigurations) *ListNacosConfigsResponseBody {
	s.Configurations = v
	return s
}

func (s *ListNacosConfigsResponseBody) SetErrorCode(v string) *ListNacosConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetHttpCode(v string) *ListNacosConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetMessage(v string) *ListNacosConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetPageNumber(v int32) *ListNacosConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetPageSize(v int32) *ListNacosConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetRequestId(v string) *ListNacosConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetSuccess(v bool) *ListNacosConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetTotalCount(v int32) *ListNacosConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNacosConfigsResponseBody) Validate() error {
	if s.Configurations != nil {
		for _, item := range s.Configurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNacosConfigsResponseBodyConfigurations struct {
	// The name of the application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// log.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// public
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 132****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNacosConfigsResponseBodyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListNacosConfigsResponseBodyConfigurations) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponseBodyConfigurations) GetAppName() *string {
	return s.AppName
}

func (s *ListNacosConfigsResponseBodyConfigurations) GetDataId() *string {
	return s.DataId
}

func (s *ListNacosConfigsResponseBodyConfigurations) GetGroup() *string {
	return s.Group
}

func (s *ListNacosConfigsResponseBodyConfigurations) GetId() *string {
	return s.Id
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetAppName(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.AppName = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetDataId(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.DataId = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetGroup(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.Group = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetId(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.Id = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) Validate() error {
	return dara.Validate(s)
}
