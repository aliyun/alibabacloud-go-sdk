// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerMaxVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkerMaxVersionResponseBody
	GetCode() *string
	SetData(v *GetWorkerMaxVersionResponseBodyData) *GetWorkerMaxVersionResponseBody
	GetData() *GetWorkerMaxVersionResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkerMaxVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkerMaxVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkerMaxVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkerMaxVersionResponseBody
	GetSuccess() *bool
}

type GetWorkerMaxVersionResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetWorkerMaxVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerMaxVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerMaxVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerMaxVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkerMaxVersionResponseBody) GetData() *GetWorkerMaxVersionResponseBodyData {
	return s.Data
}

func (s *GetWorkerMaxVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkerMaxVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkerMaxVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkerMaxVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkerMaxVersionResponseBody) SetCode(v string) *GetWorkerMaxVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) SetData(v *GetWorkerMaxVersionResponseBodyData) *GetWorkerMaxVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) SetHttpStatusCode(v int32) *GetWorkerMaxVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) SetMessage(v string) *GetWorkerMaxVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) SetRequestId(v string) *GetWorkerMaxVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) SetSuccess(v bool) *GetWorkerMaxVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkerMaxVersionResponseBodyData struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s GetWorkerMaxVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerMaxVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerMaxVersionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerMaxVersionResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetWorkerMaxVersionResponseBodyData) SetInstanceId(v string) *GetWorkerMaxVersionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBodyData) SetVersionCode(v string) *GetWorkerMaxVersionResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *GetWorkerMaxVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
