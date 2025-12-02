// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFeatureConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyFeatureConfigResponseBody
	GetCode() *int32
	SetData(v bool) *ModifyFeatureConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ModifyFeatureConfigResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *ModifyFeatureConfigResponseBody
	GetMsg() *string
	SetRequestId(v string) *ModifyFeatureConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyFeatureConfigResponseBody
	GetSuccess() *bool
}

type ModifyFeatureConfigResponseBody struct {
	// Success indicator.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// query
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Returned data
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response message of this request.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFeatureConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyFeatureConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyFeatureConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyFeatureConfigResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ModifyFeatureConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFeatureConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyFeatureConfigResponseBody) SetCode(v int32) *ModifyFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetData(v bool) *ModifyFeatureConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetHttpStatusCode(v int32) *ModifyFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetMsg(v string) *ModifyFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetRequestId(v string) *ModifyFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetSuccess(v bool) *ModifyFeatureConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
