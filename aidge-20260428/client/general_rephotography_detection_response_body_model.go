// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralRephotographyDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GeneralRephotographyDetectionResponseBody
	GetCode() *string
	SetData(v *GeneralRephotographyDetectionResponseBodyData) *GeneralRephotographyDetectionResponseBody
	GetData() *GeneralRephotographyDetectionResponseBodyData
	SetMessage(v string) *GeneralRephotographyDetectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GeneralRephotographyDetectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GeneralRephotographyDetectionResponseBody
	GetSuccess() *bool
}

type GeneralRephotographyDetectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recapture detection result.
	Data *GeneralRephotographyDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message or failure description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 70CBEFDF-BB17-1EB3-8A21-569F3124738F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GeneralRephotographyDetectionResponseBody) GetData() *GeneralRephotographyDetectionResponseBodyData {
	return s.Data
}

func (s *GeneralRephotographyDetectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GeneralRephotographyDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GeneralRephotographyDetectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GeneralRephotographyDetectionResponseBody) SetCode(v string) *GeneralRephotographyDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetData(v *GeneralRephotographyDetectionResponseBodyData) *GeneralRephotographyDetectionResponseBody {
	s.Data = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetMessage(v string) *GeneralRephotographyDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetRequestId(v string) *GeneralRephotographyDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) SetSuccess(v bool) *GeneralRephotographyDetectionResponseBody {
	s.Success = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneralRephotographyDetectionResponseBodyData struct {
	// The business result. This is an empty object if the call fails.
	Result *GeneralRephotographyDetectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The usage information. The value is `{"ProcessingCount":1}` if the call is successful, or an empty object if the call fails.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyData) GetResult() *GeneralRephotographyDetectionResponseBodyDataResult {
	return s.Result
}

func (s *GeneralRephotographyDetectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *GeneralRephotographyDetectionResponseBodyData) SetResult(v *GeneralRephotographyDetectionResponseBodyDataResult) *GeneralRephotographyDetectionResponseBodyData {
	s.Result = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyData) SetUsageMap(v map[string]*int64) *GeneralRephotographyDetectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneralRephotographyDetectionResponseBodyDataResult struct {
	// Indicates whether the image is recaptured.
	//
	// example:
	//
	// true
	IsFake *bool `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// Fixed to general.
	//
	// example:
	//
	// general
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GeneralRephotographyDetectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) GetIsFake() *bool {
	return s.IsFake
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) SetIsFake(v bool) *GeneralRephotographyDetectionResponseBodyDataResult {
	s.IsFake = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) SetType(v string) *GeneralRephotographyDetectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *GeneralRephotographyDetectionResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
