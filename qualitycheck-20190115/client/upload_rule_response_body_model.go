// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadRuleResponseBody
	GetCode() *string
	SetData(v *UploadRuleResponseBodyData) *UploadRuleResponseBody
	GetData() *UploadRuleResponseBodyData
	SetMessage(v string) *UploadRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadRuleResponseBody
	GetSuccess() *bool
}

type UploadRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadRuleResponseBody) GetData() *UploadRuleResponseBodyData {
	return s.Data
}

func (s *UploadRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadRuleResponseBody) SetCode(v string) *UploadRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UploadRuleResponseBody) SetData(v *UploadRuleResponseBodyData) *UploadRuleResponseBody {
	s.Data = v
	return s
}

func (s *UploadRuleResponseBody) SetMessage(v string) *UploadRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRuleResponseBody) SetRequestId(v string) *UploadRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRuleResponseBody) SetSuccess(v bool) *UploadRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UploadRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadRuleResponseBodyData struct {
	RidInfo []*string `json:"RidInfo,omitempty" xml:"RidInfo,omitempty" type:"Repeated"`
}

func (s UploadRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBodyData) GetRidInfo() []*string {
	return s.RidInfo
}

func (s *UploadRuleResponseBodyData) SetRidInfo(v []*string) *UploadRuleResponseBodyData {
	s.RidInfo = v
	return s
}

func (s *UploadRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
