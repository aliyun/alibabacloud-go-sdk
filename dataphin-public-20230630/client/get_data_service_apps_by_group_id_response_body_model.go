// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppsByGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfoList(v []*GetDataServiceAppsByGroupIdResponseBodyAppInfoList) *GetDataServiceAppsByGroupIdResponseBody
	GetAppInfoList() []*GetDataServiceAppsByGroupIdResponseBodyAppInfoList
	SetCode(v string) *GetDataServiceAppsByGroupIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceAppsByGroupIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAppsByGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAppsByGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAppsByGroupIdResponseBody
	GetSuccess() *bool
}

type GetDataServiceAppsByGroupIdResponseBody struct {
	AppInfoList []*GetDataServiceAppsByGroupIdResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceAppsByGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppsByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetAppInfoList() []*GetDataServiceAppsByGroupIdResponseBodyAppInfoList {
	return s.AppInfoList
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAppsByGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetAppInfoList(v []*GetDataServiceAppsByGroupIdResponseBodyAppInfoList) *GetDataServiceAppsByGroupIdResponseBody {
	s.AppInfoList = v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetCode(v string) *GetDataServiceAppsByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAppsByGroupIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetMessage(v string) *GetDataServiceAppsByGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetRequestId(v string) *GetDataServiceAppsByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) SetSuccess(v bool) *GetDataServiceAppsByGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBody) Validate() error {
	if s.AppInfoList != nil {
		for _, item := range s.AppInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataServiceAppsByGroupIdResponseBodyAppInfoList struct {
	// AppId
	//
	// example:
	//
	// 1021
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDataServiceAppsByGroupIdResponseBodyAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppsByGroupIdResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppsByGroupIdResponseBodyAppInfoList) GetId() *int32 {
	return s.Id
}

func (s *GetDataServiceAppsByGroupIdResponseBodyAppInfoList) GetName() *string {
	return s.Name
}

func (s *GetDataServiceAppsByGroupIdResponseBodyAppInfoList) SetId(v int32) *GetDataServiceAppsByGroupIdResponseBodyAppInfoList {
	s.Id = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBodyAppInfoList) SetName(v string) *GetDataServiceAppsByGroupIdResponseBodyAppInfoList {
	s.Name = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdResponseBodyAppInfoList) Validate() error {
	return dara.Validate(s)
}
