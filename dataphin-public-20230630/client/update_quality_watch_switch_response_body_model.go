// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityWatchSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateQualityWatchSwitchResponseBody
	GetCode() *string
	SetData(v int32) *UpdateQualityWatchSwitchResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *UpdateQualityWatchSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateQualityWatchSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateQualityWatchSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityWatchSwitchResponseBody
	GetSuccess() *bool
}

type UpdateQualityWatchSwitchResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityWatchSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityWatchSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityWatchSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateQualityWatchSwitchResponseBody) GetData() *int32 {
	return s.Data
}

func (s *UpdateQualityWatchSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateQualityWatchSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQualityWatchSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityWatchSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityWatchSwitchResponseBody) SetCode(v string) *UpdateQualityWatchSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) SetData(v int32) *UpdateQualityWatchSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) SetHttpStatusCode(v int32) *UpdateQualityWatchSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) SetMessage(v string) *UpdateQualityWatchSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) SetRequestId(v string) *UpdateQualityWatchSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) SetSuccess(v bool) *UpdateQualityWatchSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
