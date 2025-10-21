// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInstanceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelInstanceInfoResponseBody
	GetCode() *string
	SetDescription(v string) *GetModelInstanceInfoResponseBody
	GetDescription() *string
	SetEasServiceId(v string) *GetModelInstanceInfoResponseBody
	GetEasServiceId() *string
	SetEasServiceName(v string) *GetModelInstanceInfoResponseBody
	GetEasServiceName() *string
	SetGmtModified(v int64) *GetModelInstanceInfoResponseBody
	GetGmtModified() *int64
	SetHarmfulCategoryConfigInfoList(v []*GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) *GetModelInstanceInfoResponseBody
	GetHarmfulCategoryConfigInfoList() []*GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList
	SetHttpStatusCode(v int32) *GetModelInstanceInfoResponseBody
	GetHttpStatusCode() *int32
	SetIsSupportImage(v bool) *GetModelInstanceInfoResponseBody
	GetIsSupportImage() *bool
	SetIsSupportText(v bool) *GetModelInstanceInfoResponseBody
	GetIsSupportText() *bool
	SetMessage(v string) *GetModelInstanceInfoResponseBody
	GetMessage() *string
	SetModelCallName(v string) *GetModelInstanceInfoResponseBody
	GetModelCallName() *string
	SetModelCategoryId(v int64) *GetModelInstanceInfoResponseBody
	GetModelCategoryId() *int64
	SetModelInstanceId(v int64) *GetModelInstanceInfoResponseBody
	GetModelInstanceId() *int64
	SetModelSource(v int32) *GetModelInstanceInfoResponseBody
	GetModelSource() *int32
	SetModelToken(v string) *GetModelInstanceInfoResponseBody
	GetModelToken() *string
	SetModelUrl(v string) *GetModelInstanceInfoResponseBody
	GetModelUrl() *string
	SetModelVpcUrl(v string) *GetModelInstanceInfoResponseBody
	GetModelVpcUrl() *string
	SetPromptAttackInfoList(v []*GetModelInstanceInfoResponseBodyPromptAttackInfoList) *GetModelInstanceInfoResponseBody
	GetPromptAttackInfoList() []*GetModelInstanceInfoResponseBodyPromptAttackInfoList
	SetRequestId(v string) *GetModelInstanceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelInstanceInfoResponseBody
	GetSuccess() *bool
}

type GetModelInstanceInfoResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// opencompass-vllm07-acc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eas-m-12345678
	EasServiceId *string `json:"EasServiceId,omitempty" xml:"EasServiceId,omitempty"`
	// example:
	//
	// rai_content_detection_model
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// 1634122349000
	GmtModified                   *int64                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HarmfulCategoryConfigInfoList []*GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList `json:"HarmfulCategoryConfigInfoList,omitempty" xml:"HarmfulCategoryConfigInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// False
	IsSupportImage *bool `json:"IsSupportImage,omitempty" xml:"IsSupportImage,omitempty"`
	// example:
	//
	// True
	IsSupportText *bool `json:"IsSupportText,omitempty" xml:"IsSupportText,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// demo
	ModelCallName *string `json:"ModelCallName,omitempty" xml:"ModelCallName,omitempty"`
	// example:
	//
	// 1
	ModelCategoryId *int64 `json:"ModelCategoryId,omitempty" xml:"ModelCategoryId,omitempty"`
	// example:
	//
	// 123
	ModelInstanceId *int64 `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
	// example:
	//
	// 1
	ModelSource *int32 `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	// example:
	//
	// http://12345*****.cn-shanghai.aliyuncs.com/api/predict/echo
	ModelUrl *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	// example:
	//
	// http://12345*****.vpc.cn-shanghai.aliyuncs.com/api/predict/demo
	ModelVpcUrl          *string                                                 `json:"ModelVpcUrl,omitempty" xml:"ModelVpcUrl,omitempty"`
	PromptAttackInfoList []*GetModelInstanceInfoResponseBodyPromptAttackInfoList `json:"PromptAttackInfoList,omitempty" xml:"PromptAttackInfoList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelInstanceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelInstanceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelInstanceInfoResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetModelInstanceInfoResponseBody) GetEasServiceId() *string {
	return s.EasServiceId
}

func (s *GetModelInstanceInfoResponseBody) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *GetModelInstanceInfoResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetModelInstanceInfoResponseBody) GetHarmfulCategoryConfigInfoList() []*GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	return s.HarmfulCategoryConfigInfoList
}

func (s *GetModelInstanceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelInstanceInfoResponseBody) GetIsSupportImage() *bool {
	return s.IsSupportImage
}

func (s *GetModelInstanceInfoResponseBody) GetIsSupportText() *bool {
	return s.IsSupportText
}

func (s *GetModelInstanceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelInstanceInfoResponseBody) GetModelCallName() *string {
	return s.ModelCallName
}

func (s *GetModelInstanceInfoResponseBody) GetModelCategoryId() *int64 {
	return s.ModelCategoryId
}

func (s *GetModelInstanceInfoResponseBody) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *GetModelInstanceInfoResponseBody) GetModelSource() *int32 {
	return s.ModelSource
}

func (s *GetModelInstanceInfoResponseBody) GetModelToken() *string {
	return s.ModelToken
}

func (s *GetModelInstanceInfoResponseBody) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *GetModelInstanceInfoResponseBody) GetModelVpcUrl() *string {
	return s.ModelVpcUrl
}

func (s *GetModelInstanceInfoResponseBody) GetPromptAttackInfoList() []*GetModelInstanceInfoResponseBodyPromptAttackInfoList {
	return s.PromptAttackInfoList
}

func (s *GetModelInstanceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelInstanceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelInstanceInfoResponseBody) SetCode(v string) *GetModelInstanceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetDescription(v string) *GetModelInstanceInfoResponseBody {
	s.Description = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetEasServiceId(v string) *GetModelInstanceInfoResponseBody {
	s.EasServiceId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetEasServiceName(v string) *GetModelInstanceInfoResponseBody {
	s.EasServiceName = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetGmtModified(v int64) *GetModelInstanceInfoResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetHarmfulCategoryConfigInfoList(v []*GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) *GetModelInstanceInfoResponseBody {
	s.HarmfulCategoryConfigInfoList = v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetHttpStatusCode(v int32) *GetModelInstanceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetIsSupportImage(v bool) *GetModelInstanceInfoResponseBody {
	s.IsSupportImage = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetIsSupportText(v bool) *GetModelInstanceInfoResponseBody {
	s.IsSupportText = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetMessage(v string) *GetModelInstanceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelCallName(v string) *GetModelInstanceInfoResponseBody {
	s.ModelCallName = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelCategoryId(v int64) *GetModelInstanceInfoResponseBody {
	s.ModelCategoryId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelInstanceId(v int64) *GetModelInstanceInfoResponseBody {
	s.ModelInstanceId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelSource(v int32) *GetModelInstanceInfoResponseBody {
	s.ModelSource = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelToken(v string) *GetModelInstanceInfoResponseBody {
	s.ModelToken = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelUrl(v string) *GetModelInstanceInfoResponseBody {
	s.ModelUrl = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetModelVpcUrl(v string) *GetModelInstanceInfoResponseBody {
	s.ModelVpcUrl = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetPromptAttackInfoList(v []*GetModelInstanceInfoResponseBodyPromptAttackInfoList) *GetModelInstanceInfoResponseBody {
	s.PromptAttackInfoList = v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetRequestId(v string) *GetModelInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) SetSuccess(v bool) *GetModelInstanceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelInstanceInfoResponseBody) Validate() error {
	if s.HarmfulCategoryConfigInfoList != nil {
		for _, item := range s.HarmfulCategoryConfigInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PromptAttackInfoList != nil {
		for _, item := range s.PromptAttackInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList struct {
	// example:
	//
	// 1
	CategoryId    *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 0
	InputOutputType *int32 `json:"InputOutputType,omitempty" xml:"InputOutputType,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GetInputOutputType() *int32 {
	return s.InputOutputType
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryId(v int64) *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryLabel(v string) *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) SetCategoryType(v int32) *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) SetInputOutputType(v int32) *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.InputOutputType = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) SetSecurityLevel(v int32) *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyHarmfulCategoryConfigInfoList) Validate() error {
	return dara.Validate(s)
}

type GetModelInstanceInfoResponseBodyPromptAttackInfoList struct {
	// example:
	//
	// 1
	CategoryId    *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInstanceInfoResponseBodyPromptAttackInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetModelInstanceInfoResponseBodyPromptAttackInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) GetCategoryLabel() *string {
	return s.CategoryLabel
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) GetSecurityLevel() *int32 {
	return s.SecurityLevel
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) SetCategoryId(v int64) *GetModelInstanceInfoResponseBodyPromptAttackInfoList {
	s.CategoryId = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) SetCategoryLabel(v string) *GetModelInstanceInfoResponseBodyPromptAttackInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) SetSecurityLevel(v int32) *GetModelInstanceInfoResponseBodyPromptAttackInfoList {
	s.SecurityLevel = &v
	return s
}

func (s *GetModelInstanceInfoResponseBodyPromptAttackInfoList) Validate() error {
	return dara.Validate(s)
}
