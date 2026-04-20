// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMmAppRagWeightResponseBody
	GetCode() *string
	SetData(v *UpdateMmAppRagWeightResponseBodyData) *UpdateMmAppRagWeightResponseBody
	GetData() *UpdateMmAppRagWeightResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMmAppRagWeightResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMmAppRagWeightResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMmAppRagWeightResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppRagWeightResponseBody
	GetSuccess() *bool
}

type UpdateMmAppRagWeightResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateMmAppRagWeightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5979B783-0AF5-547E-A549-8659F8A2A12A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagWeightResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppRagWeightResponseBody) GetData() *UpdateMmAppRagWeightResponseBodyData {
	return s.Data
}

func (s *UpdateMmAppRagWeightResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMmAppRagWeightResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMmAppRagWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppRagWeightResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagWeightResponseBody) SetCode(v string) *UpdateMmAppRagWeightResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) SetData(v *UpdateMmAppRagWeightResponseBodyData) *UpdateMmAppRagWeightResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) SetHttpStatusCode(v int32) *UpdateMmAppRagWeightResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) SetMessage(v string) *UpdateMmAppRagWeightResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) SetRequestId(v string) *UpdateMmAppRagWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) SetSuccess(v bool) *UpdateMmAppRagWeightResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppRagWeightResponseBodyData struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagWeightResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagWeightResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagWeightResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagWeightResponseBodyData) SetSuccess(v bool) *UpdateMmAppRagWeightResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagWeightResponseBodyData) Validate() error {
	return dara.Validate(s)
}
