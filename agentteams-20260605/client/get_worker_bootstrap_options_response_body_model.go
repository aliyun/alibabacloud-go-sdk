// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerBootstrapOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkerBootstrapOptionsResponseBody
	GetCode() *string
	SetData(v *GetWorkerBootstrapOptionsResponseBodyData) *GetWorkerBootstrapOptionsResponseBody
	GetData() *GetWorkerBootstrapOptionsResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkerBootstrapOptionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkerBootstrapOptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkerBootstrapOptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkerBootstrapOptionsResponseBody
	GetSuccess() *bool
}

type GetWorkerBootstrapOptionsResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetWorkerBootstrapOptionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerBootstrapOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerBootstrapOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetData() *GetWorkerBootstrapOptionsResponseBodyData {
	return s.Data
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkerBootstrapOptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetCode(v string) *GetWorkerBootstrapOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetData(v *GetWorkerBootstrapOptionsResponseBodyData) *GetWorkerBootstrapOptionsResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetHttpStatusCode(v int32) *GetWorkerBootstrapOptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetMessage(v string) *GetWorkerBootstrapOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetRequestId(v string) *GetWorkerBootstrapOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) SetSuccess(v bool) *GetWorkerBootstrapOptionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkerBootstrapOptionsResponseBodyData struct {
	InstanceId     *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name           *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkOptions []*GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Repeated"`
}

func (s GetWorkerBootstrapOptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerBootstrapOptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) GetNetworkOptions() []*GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions {
	return s.NetworkOptions
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) SetInstanceId(v string) *GetWorkerBootstrapOptionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) SetName(v string) *GetWorkerBootstrapOptionsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) SetNetworkOptions(v []*GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) *GetWorkerBootstrapOptionsResponseBodyData {
	s.NetworkOptions = v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBodyData) Validate() error {
	if s.NetworkOptions != nil {
		for _, item := range s.NetworkOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions struct {
	Available   *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) GoString() string {
	return s.String()
}

func (s *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) GetAvailable() *bool {
	return s.Available
}

func (s *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) SetAvailable(v bool) *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions {
	s.Available = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) SetNetworkType(v string) *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions {
	s.NetworkType = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponseBodyDataNetworkOptions) Validate() error {
	return dara.Validate(s)
}
