// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *GetServiceConfResponseBody
	GetClassify() *string
	SetCode(v int32) *GetServiceConfResponseBody
	GetCode() *int32
	SetCustomServiceConf(v map[string]interface{}) *GetServiceConfResponseBody
	GetCustomServiceConf() map[string]interface{}
	SetGmtModified(v string) *GetServiceConfResponseBody
	GetGmtModified() *string
	SetMsg(v string) *GetServiceConfResponseBody
	GetMsg() *string
	SetOption(v map[string]interface{}) *GetServiceConfResponseBody
	GetOption() map[string]interface{}
	SetRequestId(v string) *GetServiceConfResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetServiceConfResponseBody
	GetResourceType() *string
	SetServiceCode(v string) *GetServiceConfResponseBody
	GetServiceCode() *string
	SetServiceType(v string) *GetServiceConfResponseBody
	GetServiceType() *string
	SetSuccess(v bool) *GetServiceConfResponseBody
	GetSuccess() *bool
	SetUid(v string) *GetServiceConfResponseBody
	GetUid() *string
}

type GetServiceConfResponseBody struct {
	// Classification.
	//
	// example:
	//
	// guard-scene
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// Error code, consistent with HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Service details
	//
	// example:
	//
	// {}
	CustomServiceConf map[string]interface{} `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2023-01-17 12:29:56
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Options.
	//
	// example:
	//
	// {}
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Service type.
	//
	// example:
	//
	// plus
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Success indicator
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// UID.
	//
	// example:
	//
	// 17726*****370735
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetServiceConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConfResponseBody) GetClassify() *string {
	return s.Classify
}

func (s *GetServiceConfResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceConfResponseBody) GetCustomServiceConf() map[string]interface{} {
	return s.CustomServiceConf
}

func (s *GetServiceConfResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetServiceConfResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetServiceConfResponseBody) GetOption() map[string]interface{} {
	return s.Option
}

func (s *GetServiceConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceConfResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceConfResponseBody) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetServiceConfResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceConfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceConfResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetServiceConfResponseBody) SetClassify(v string) *GetServiceConfResponseBody {
	s.Classify = &v
	return s
}

func (s *GetServiceConfResponseBody) SetCode(v int32) *GetServiceConfResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceConfResponseBody) SetCustomServiceConf(v map[string]interface{}) *GetServiceConfResponseBody {
	s.CustomServiceConf = v
	return s
}

func (s *GetServiceConfResponseBody) SetGmtModified(v string) *GetServiceConfResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetServiceConfResponseBody) SetMsg(v string) *GetServiceConfResponseBody {
	s.Msg = &v
	return s
}

func (s *GetServiceConfResponseBody) SetOption(v map[string]interface{}) *GetServiceConfResponseBody {
	s.Option = v
	return s
}

func (s *GetServiceConfResponseBody) SetRequestId(v string) *GetServiceConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceConfResponseBody) SetResourceType(v string) *GetServiceConfResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfResponseBody) SetServiceCode(v string) *GetServiceConfResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfResponseBody) SetServiceType(v string) *GetServiceConfResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceConfResponseBody) SetSuccess(v bool) *GetServiceConfResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceConfResponseBody) SetUid(v string) *GetServiceConfResponseBody {
	s.Uid = &v
	return s
}

func (s *GetServiceConfResponseBody) Validate() error {
	return dara.Validate(s)
}
