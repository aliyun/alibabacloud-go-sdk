// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkerListResponseBody
	GetCode() *int32
	SetData(v *GetWorkerListResponseBodyData) *GetWorkerListResponseBody
	GetData() *GetWorkerListResponseBodyData
	SetMessage(v string) *GetWorkerListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkerListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkerListResponseBody
	GetSuccess() *bool
}

type GetWorkerListResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The job information.
	Data *GetWorkerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// Cannot find product according to your domain.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkerListResponseBody) GetData() *GetWorkerListResponseBodyData {
	return s.Data
}

func (s *GetWorkerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkerListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkerListResponseBody) SetCode(v int32) *GetWorkerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerListResponseBody) SetData(v *GetWorkerListResponseBodyData) *GetWorkerListResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerListResponseBody) SetMessage(v string) *GetWorkerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerListResponseBody) SetRequestId(v string) *GetWorkerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerListResponseBody) SetSuccess(v bool) *GetWorkerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkerListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkerListResponseBodyData struct {
	// The worker information.
	WorkerInfos []*GetWorkerListResponseBodyDataWorkerInfos `json:"WorkerInfos,omitempty" xml:"WorkerInfos,omitempty" type:"Repeated"`
}

func (s GetWorkerListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBodyData) GetWorkerInfos() []*GetWorkerListResponseBodyDataWorkerInfos {
	return s.WorkerInfos
}

func (s *GetWorkerListResponseBodyData) SetWorkerInfos(v []*GetWorkerListResponseBodyDataWorkerInfos) *GetWorkerListResponseBodyData {
	s.WorkerInfos = v
	return s
}

func (s *GetWorkerListResponseBodyData) Validate() error {
	if s.WorkerInfos != nil {
		for _, item := range s.WorkerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkerListResponseBodyDataWorkerInfos struct {
	// The IP address of the worker.
	//
	// example:
	//
	// 30.225.16.49
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The label of the worker.
	//
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The port number of the worker.
	//
	// example:
	//
	// 60831
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The startup method of the worker.
	//
	// example:
	//
	// springboot
	Starter *string `json:"Starter,omitempty" xml:"Starter,omitempty"`
	// The version of the worker.
	//
	// example:
	//
	// 1.3.4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The address of the worker. The address is in the format of ${worker_id}@${worker_ip}:${worker_port}.
	//
	// example:
	//
	// 030225016049_11734_25917@30.225.16.49:60831
	WorkerAddress *string `json:"WorkerAddress,omitempty" xml:"WorkerAddress,omitempty"`
}

func (s GetWorkerListResponseBodyDataWorkerInfos) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerListResponseBodyDataWorkerInfos) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetIp() *string {
	return s.Ip
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetLabel() *string {
	return s.Label
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetPort() *int32 {
	return s.Port
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetStarter() *string {
	return s.Starter
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetVersion() *string {
	return s.Version
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) GetWorkerAddress() *string {
	return s.WorkerAddress
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetIp(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Ip = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetLabel(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Label = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetPort(v int32) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Port = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetStarter(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Starter = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetVersion(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.Version = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) SetWorkerAddress(v string) *GetWorkerListResponseBodyDataWorkerInfos {
	s.WorkerAddress = &v
	return s
}

func (s *GetWorkerListResponseBodyDataWorkerInfos) Validate() error {
	return dara.Validate(s)
}
