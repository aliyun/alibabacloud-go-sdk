// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBizMetricResponseBody
	GetCode() *string
	SetData(v *UpdateBizMetricResponseBodyData) *UpdateBizMetricResponseBody
	GetData() *UpdateBizMetricResponseBodyData
	SetHttpStatusCode(v int32) *UpdateBizMetricResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBizMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBizMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBizMetricResponseBody
	GetSuccess() *bool
}

type UpdateBizMetricResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateBizMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBizMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBizMetricResponseBody) GetData() *UpdateBizMetricResponseBodyData {
	return s.Data
}

func (s *UpdateBizMetricResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBizMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBizMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBizMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBizMetricResponseBody) SetCode(v string) *UpdateBizMetricResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizMetricResponseBody) SetData(v *UpdateBizMetricResponseBodyData) *UpdateBizMetricResponseBody {
	s.Data = v
	return s
}

func (s *UpdateBizMetricResponseBody) SetHttpStatusCode(v int32) *UpdateBizMetricResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBizMetricResponseBody) SetMessage(v string) *UpdateBizMetricResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBizMetricResponseBody) SetRequestId(v string) *UpdateBizMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBizMetricResponseBody) SetSuccess(v bool) *UpdateBizMetricResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBizMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBizMetricResponseBodyData struct {
	// example:
	//
	// npe
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBizMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateBizMetricResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UpdateBizMetricResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBizMetricResponseBodyData) SetMessage(v string) *UpdateBizMetricResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateBizMetricResponseBodyData) SetSuccess(v bool) *UpdateBizMetricResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateBizMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
