// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterListMemberSubscriptionsResponseBodyData) *ModelRouterListMemberSubscriptionsResponseBody
	GetData() *ModelRouterListMemberSubscriptionsResponseBodyData
	SetErrCode(v string) *ModelRouterListMemberSubscriptionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterListMemberSubscriptionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterListMemberSubscriptionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterListMemberSubscriptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterListMemberSubscriptionsResponseBody
	GetSuccess() *bool
}

type ModelRouterListMemberSubscriptionsResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterListMemberSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault information code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterListMemberSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetData() *ModelRouterListMemberSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetData(v *ModelRouterListMemberSubscriptionsResponseBodyData) *ModelRouterListMemberSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetErrCode(v string) *ModelRouterListMemberSubscriptionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetErrMessage(v string) *ModelRouterListMemberSubscriptionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ModelRouterListMemberSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetRequestId(v string) *ModelRouterListMemberSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) SetSuccess(v bool) *ModelRouterListMemberSubscriptionsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterListMemberSubscriptionsResponseBodyData struct {
	// The response data list.
	//
	// example:
	//
	// []
	List []*SubscriptionDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ModelRouterListMemberSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberSubscriptionsResponseBodyData) GetList() []*SubscriptionDTO {
	return s.List
}

func (s *ModelRouterListMemberSubscriptionsResponseBodyData) SetList(v []*SubscriptionDTO) *ModelRouterListMemberSubscriptionsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
