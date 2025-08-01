// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineNamepaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigCount(v string) *GetEngineNamepaceResponseBody
	GetConfigCount() *string
	SetErrorCode(v string) *GetEngineNamepaceResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *GetEngineNamepaceResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GetEngineNamepaceResponseBody
	GetMessage() *string
	SetNamespace(v string) *GetEngineNamepaceResponseBody
	GetNamespace() *string
	SetNamespaceDesc(v string) *GetEngineNamepaceResponseBody
	GetNamespaceDesc() *string
	SetNamespaceShowName(v string) *GetEngineNamepaceResponseBody
	GetNamespaceShowName() *string
	SetQuota(v string) *GetEngineNamepaceResponseBody
	GetQuota() *string
	SetRequestId(v string) *GetEngineNamepaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEngineNamepaceResponseBody
	GetSuccess() *bool
	SetType(v string) *GetEngineNamepaceResponseBody
	GetType() *string
}

type GetEngineNamepaceResponseBody struct {
	// The number of configurations.
	//
	// example:
	//
	// 1
	ConfigCount *string `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
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
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the namespace.
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
	// public
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	// The quota of configurations.
	//
	// example:
	//
	// 200
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FA8F966F-420C-52F5-B49E-6ED7CCE02697
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
	// The type of the namespace. Valid values:
	//
	// 	- 0: global configuration
	//
	// 	- 1: default namespace
	//
	// 	- 2: custom namespace
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEngineNamepaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEngineNamepaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceResponseBody) GetConfigCount() *string {
	return s.ConfigCount
}

func (s *GetEngineNamepaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetEngineNamepaceResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetEngineNamepaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEngineNamepaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *GetEngineNamepaceResponseBody) GetNamespaceDesc() *string {
	return s.NamespaceDesc
}

func (s *GetEngineNamepaceResponseBody) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *GetEngineNamepaceResponseBody) GetQuota() *string {
	return s.Quota
}

func (s *GetEngineNamepaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEngineNamepaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEngineNamepaceResponseBody) GetType() *string {
	return s.Type
}

func (s *GetEngineNamepaceResponseBody) SetConfigCount(v string) *GetEngineNamepaceResponseBody {
	s.ConfigCount = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetErrorCode(v string) *GetEngineNamepaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetHttpCode(v string) *GetEngineNamepaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetMessage(v string) *GetEngineNamepaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespace(v string) *GetEngineNamepaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespaceDesc(v string) *GetEngineNamepaceResponseBody {
	s.NamespaceDesc = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespaceShowName(v string) *GetEngineNamepaceResponseBody {
	s.NamespaceShowName = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetQuota(v string) *GetEngineNamepaceResponseBody {
	s.Quota = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetRequestId(v string) *GetEngineNamepaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetSuccess(v bool) *GetEngineNamepaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetType(v string) *GetEngineNamepaceResponseBody {
	s.Type = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) Validate() error {
	return dara.Validate(s)
}
