// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEngineNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateEngineNamespaceResponseBody
	GetClusterId() *string
	SetData(v *CreateEngineNamespaceResponseBodyData) *CreateEngineNamespaceResponseBody
	GetData() *CreateEngineNamespaceResponseBodyData
	SetErrorCode(v string) *CreateEngineNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateEngineNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEngineNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEngineNamespaceResponseBody
	GetSuccess() *bool
}

type CreateEngineNamespaceResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// mse-892k****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The details of the data.
	Data *CreateEngineNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// F6092602-C357-4750-89D9-E572FBEA****
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

func (s CreateEngineNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateEngineNamespaceResponseBody) GetData() *CreateEngineNamespaceResponseBodyData {
	return s.Data
}

func (s *CreateEngineNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateEngineNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEngineNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEngineNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEngineNamespaceResponseBody) SetClusterId(v string) *CreateEngineNamespaceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetData(v *CreateEngineNamespaceResponseBodyData) *CreateEngineNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetErrorCode(v string) *CreateEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetMessage(v string) *CreateEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetRequestId(v string) *CreateEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetSuccess(v bool) *CreateEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEngineNamespaceResponseBodyData struct {
	// The number of configurations.
	//
	// example:
	//
	// 1
	ConfigCount *int32 `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// f4fa5b81-2f26-4900-833a-7516b315ebb2
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
	// dev
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	// The quota of configurations.
	//
	// example:
	//
	// 1
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The number of active services.
	//
	// example:
	//
	// 3
	ServiceCount *int32 `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
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

func (s CreateEngineNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEngineNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponseBodyData) GetConfigCount() *int32 {
	return s.ConfigCount
}

func (s *CreateEngineNamespaceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateEngineNamespaceResponseBodyData) GetNamespaceDesc() *string {
	return s.NamespaceDesc
}

func (s *CreateEngineNamespaceResponseBodyData) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *CreateEngineNamespaceResponseBodyData) GetQuota() *int32 {
	return s.Quota
}

func (s *CreateEngineNamespaceResponseBodyData) GetServiceCount() *int32 {
	return s.ServiceCount
}

func (s *CreateEngineNamespaceResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *CreateEngineNamespaceResponseBodyData) SetConfigCount(v int32) *CreateEngineNamespaceResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespace(v string) *CreateEngineNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespaceDesc(v string) *CreateEngineNamespaceResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespaceShowName(v string) *CreateEngineNamespaceResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetQuota(v int32) *CreateEngineNamespaceResponseBodyData {
	s.Quota = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetServiceCount(v int32) *CreateEngineNamespaceResponseBodyData {
	s.ServiceCount = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetType(v int32) *CreateEngineNamespaceResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
