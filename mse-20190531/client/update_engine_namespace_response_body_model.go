// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateEngineNamespaceResponseBodyData) *UpdateEngineNamespaceResponseBody
	GetData() *UpdateEngineNamespaceResponseBodyData
	SetErrorCode(v string) *UpdateEngineNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateEngineNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEngineNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEngineNamespaceResponseBody
	GetSuccess() *bool
}

type UpdateEngineNamespaceResponseBody struct {
	// The details of the data.
	Data *UpdateEngineNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
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
}

func (s UpdateEngineNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponseBody) GetData() *UpdateEngineNamespaceResponseBodyData {
	return s.Data
}

func (s *UpdateEngineNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateEngineNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEngineNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEngineNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEngineNamespaceResponseBody) SetData(v *UpdateEngineNamespaceResponseBodyData) *UpdateEngineNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetErrorCode(v string) *UpdateEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetMessage(v string) *UpdateEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetRequestId(v string) *UpdateEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetSuccess(v bool) *UpdateEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEngineNamespaceResponseBodyData struct {
	// The quota value.
	//
	// example:
	//
	// 1
	ConfigCount *int32 `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	// The namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// mytest
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	// The display name of the namespace.
	//
	// example:
	//
	// mytestshowname
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	// The quota of configurations.
	//
	// example:
	//
	// 1
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The type of the namespace. Valid values:
	//
	// 	- `0`: global configuration
	//
	// 	- `1`: default namespace
	//
	// 	- `2`: custom namespace
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateEngineNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponseBodyData) GetConfigCount() *int32 {
	return s.ConfigCount
}

func (s *UpdateEngineNamespaceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateEngineNamespaceResponseBodyData) GetNamespaceDesc() *string {
	return s.NamespaceDesc
}

func (s *UpdateEngineNamespaceResponseBodyData) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *UpdateEngineNamespaceResponseBodyData) GetQuota() *int32 {
	return s.Quota
}

func (s *UpdateEngineNamespaceResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *UpdateEngineNamespaceResponseBodyData) SetConfigCount(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespace(v string) *UpdateEngineNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespaceDesc(v string) *UpdateEngineNamespaceResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespaceShowName(v string) *UpdateEngineNamespaceResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetQuota(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.Quota = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetType(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
