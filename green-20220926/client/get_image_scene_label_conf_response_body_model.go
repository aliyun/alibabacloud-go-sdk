// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetImageSceneLabelConfResponseBody
	GetCode() *int32
	SetData(v []map[string]interface{}) *GetImageSceneLabelConfResponseBody
	GetData() []map[string]interface{}
	SetHttpStatusCode(v int32) *GetImageSceneLabelConfResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *GetImageSceneLabelConfResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetImageSceneLabelConfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageSceneLabelConfResponseBody
	GetSuccess() *bool
}

type GetImageSceneLabelConfResponseBody struct {
	// example:
	//
	// 200
	Code *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageSceneLabelConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetImageSceneLabelConfResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *GetImageSceneLabelConfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageSceneLabelConfResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetImageSceneLabelConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageSceneLabelConfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageSceneLabelConfResponseBody) SetCode(v int32) *GetImageSceneLabelConfResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetData(v []map[string]interface{}) *GetImageSceneLabelConfResponseBody {
	s.Data = v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetHttpStatusCode(v int32) *GetImageSceneLabelConfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetMsg(v string) *GetImageSceneLabelConfResponseBody {
	s.Msg = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetRequestId(v string) *GetImageSceneLabelConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetSuccess(v bool) *GetImageSceneLabelConfResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) Validate() error {
	return dara.Validate(s)
}
