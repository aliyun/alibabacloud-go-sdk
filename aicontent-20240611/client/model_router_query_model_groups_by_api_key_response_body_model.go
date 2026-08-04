// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupsByApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetData() *ModelRouterQueryModelGroupsByApiKeyResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupsByApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupsByApiKeyResponseBody struct {
	// The data object.
	Data *ModelRouterQueryModelGroupsByApiKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message code.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryModelGroupsByApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupsByApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetData() *ModelRouterQueryModelGroupsByApiKeyResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetData(v *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelGroupsByApiKeyResponseBodyData struct {
	// example:
	//
	// department
	BindType *string `json:"bindType,omitempty" xml:"bindType,omitempty"`
	// The department ID to which the key belongs.
	//
	// example:
	//
	// 1001
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The list of bound model groups.
	//
	// example:
	//
	// []
	Groups []*ModelGroupDTO `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// The list of individually authorized model IDs.
	//
	// example:
	//
	// [200]
	StandaloneModelList []*int64 `json:"standaloneModelList,omitempty" xml:"standaloneModelList,omitempty" type:"Repeated"`
}

func (s ModelRouterQueryModelGroupsByApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupsByApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) GetBindType() *string {
	return s.BindType
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) GetGroups() []*ModelGroupDTO {
	return s.Groups
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) GetStandaloneModelList() []*int64 {
	return s.StandaloneModelList
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) SetBindType(v string) *ModelRouterQueryModelGroupsByApiKeyResponseBodyData {
	s.BindType = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) SetClientId(v int64) *ModelRouterQueryModelGroupsByApiKeyResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) SetGroups(v []*ModelGroupDTO) *ModelRouterQueryModelGroupsByApiKeyResponseBodyData {
	s.Groups = v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) SetStandaloneModelList(v []*int64) *ModelRouterQueryModelGroupsByApiKeyResponseBodyData {
	s.StandaloneModelList = v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponseBodyData) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
