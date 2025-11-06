// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListListenersByIpResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListListenersByIpResponseBody
	GetHttpCode() *string
	SetListeners(v []*ListListenersByIpResponseBodyListeners) *ListListenersByIpResponseBody
	GetListeners() []*ListListenersByIpResponseBodyListeners
	SetMessage(v string) *ListListenersByIpResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListListenersByIpResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListenersByIpResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListListenersByIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListListenersByIpResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListListenersByIpResponseBody
	GetTotalCount() *int32
}

type ListListenersByIpResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// MSE-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The information about listeners.
	Listeners []*ListListenersByIpResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
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
	// 54973C90-F379-4372-9AA5-053A3F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersByIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListListenersByIpResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListListenersByIpResponseBody) GetListeners() []*ListListenersByIpResponseBodyListeners {
	return s.Listeners
}

func (s *ListListenersByIpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListListenersByIpResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListenersByIpResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListenersByIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenersByIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListListenersByIpResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenersByIpResponseBody) SetErrorCode(v string) *ListListenersByIpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetHttpCode(v string) *ListListenersByIpResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetListeners(v []*ListListenersByIpResponseBodyListeners) *ListListenersByIpResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersByIpResponseBody) SetMessage(v string) *ListListenersByIpResponseBody {
	s.Message = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetPageNumber(v int32) *ListListenersByIpResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetPageSize(v int32) *ListListenersByIpResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetRequestId(v string) *ListListenersByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetSuccess(v bool) *ListListenersByIpResponseBody {
	s.Success = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetTotalCount(v int32) *ListListenersByIpResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersByIpResponseBody) Validate() error {
	if s.Listeners != nil {
		for _, item := range s.Listeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenersByIpResponseBodyListeners struct {
	// The ID of the data.
	//
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The group.
	//
	// example:
	//
	// default
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The verification string.
	//
	// example:
	//
	// 23sdfdf
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListListenersByIpResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByIpResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponseBodyListeners) GetDataId() *string {
	return s.DataId
}

func (s *ListListenersByIpResponseBodyListeners) GetGroup() *string {
	return s.Group
}

func (s *ListListenersByIpResponseBodyListeners) GetMd5() *string {
	return s.Md5
}

func (s *ListListenersByIpResponseBodyListeners) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListListenersByIpResponseBodyListeners) SetDataId(v string) *ListListenersByIpResponseBodyListeners {
	s.DataId = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) SetGroup(v string) *ListListenersByIpResponseBodyListeners {
	s.Group = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) SetMd5(v string) *ListListenersByIpResponseBodyListeners {
	s.Md5 = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) SetNamespaceId(v string) *ListListenersByIpResponseBodyListeners {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}
