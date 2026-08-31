// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateComputeClusterResponseBody
	GetCode() *string
	SetData(v *CreateComputeClusterResponseBodyData) *CreateComputeClusterResponseBody
	GetData() *CreateComputeClusterResponseBodyData
	SetHttpStatusCode(v int32) *CreateComputeClusterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateComputeClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateComputeClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeClusterResponseBody
	GetSuccess() *bool
}

type CreateComputeClusterResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result.
	Data *CreateComputeClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
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
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateComputeClusterResponseBody) GetData() *CreateComputeClusterResponseBodyData {
	return s.Data
}

func (s *CreateComputeClusterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateComputeClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateComputeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeClusterResponseBody) SetCode(v string) *CreateComputeClusterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeClusterResponseBody) SetData(v *CreateComputeClusterResponseBodyData) *CreateComputeClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateComputeClusterResponseBody) SetHttpStatusCode(v int32) *CreateComputeClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateComputeClusterResponseBody) SetMessage(v string) *CreateComputeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComputeClusterResponseBody) SetRequestId(v string) *CreateComputeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeClusterResponseBody) SetSuccess(v bool) *CreateComputeClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeClusterResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// 102111
	DsId *int64 `json:"DsId,omitempty" xml:"DsId,omitempty"`
}

func (s CreateComputeClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterResponseBodyData) GetDsId() *int64 {
	return s.DsId
}

func (s *CreateComputeClusterResponseBodyData) SetDsId(v int64) *CreateComputeClusterResponseBodyData {
	s.DsId = &v
	return s
}

func (s *CreateComputeClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
