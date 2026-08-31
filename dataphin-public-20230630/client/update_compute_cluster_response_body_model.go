// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateComputeClusterResponseBody
	GetCode() *string
	SetData(v *UpdateComputeClusterResponseBodyData) *UpdateComputeClusterResponseBody
	GetData() *UpdateComputeClusterResponseBodyData
	SetHttpStatusCode(v int32) *UpdateComputeClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateComputeClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateComputeClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeClusterResponseBody
	GetSuccess() *bool
}

type UpdateComputeClusterResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateComputeClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateComputeClusterResponseBody) GetData() *UpdateComputeClusterResponseBodyData {
	return s.Data
}

func (s *UpdateComputeClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateComputeClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateComputeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeClusterResponseBody) SetCode(v string) *UpdateComputeClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateComputeClusterResponseBody) SetData(v *UpdateComputeClusterResponseBodyData) *UpdateComputeClusterResponseBody {
	s.Data = v
	return s
}

func (s *UpdateComputeClusterResponseBody) SetHttpStatusCode(v int32) *UpdateComputeClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateComputeClusterResponseBody) SetMessage(v string) *UpdateComputeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateComputeClusterResponseBody) SetRequestId(v string) *UpdateComputeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeClusterResponseBody) SetSuccess(v bool) *UpdateComputeClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeClusterResponseBodyData struct {
	// example:
	//
	// 102111
	DsId *int64 `json:"DsId,omitempty" xml:"DsId,omitempty"`
}

func (s UpdateComputeClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterResponseBodyData) GetDsId() *int64 {
	return s.DsId
}

func (s *UpdateComputeClusterResponseBodyData) SetDsId(v int64) *UpdateComputeClusterResponseBodyData {
	s.DsId = &v
	return s
}

func (s *UpdateComputeClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
