// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExecutorConfigResponseBody
	GetCode() *int32
	SetData(v *GetExecutorConfigResponseBodyData) *GetExecutorConfigResponseBody
	GetData() *GetExecutorConfigResponseBodyData
	SetMessage(v string) *GetExecutorConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExecutorConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExecutorConfigResponseBody
	GetSuccess() *bool
}

type GetExecutorConfigResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetExecutorConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetExecutorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExecutorConfigResponseBody) GetData() *GetExecutorConfigResponseBodyData {
	return s.Data
}

func (s *GetExecutorConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExecutorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecutorConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExecutorConfigResponseBody) SetCode(v int32) *GetExecutorConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetData(v *GetExecutorConfigResponseBodyData) *GetExecutorConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetExecutorConfigResponseBody) SetMessage(v string) *GetExecutorConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetRequestId(v string) *GetExecutorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetSuccess(v bool) *GetExecutorConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetExecutorConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExecutorConfigResponseBodyData struct {
	// example:
	//
	// [{"cluster":"c2c619b5129e0400fa3df263b249622aa","namespace":"default","service":"xxljob-http-demo1-svc"}]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// k8s_service
	ExecutorType *string `json:"ExecutorType,omitempty" xml:"ExecutorType,omitempty"`
}

func (s GetExecutorConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetExecutorConfigResponseBodyData) GetExecutorType() *string {
	return s.ExecutorType
}

func (s *GetExecutorConfigResponseBodyData) SetConfig(v string) *GetExecutorConfigResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetExecutorConfigResponseBodyData) SetExecutorType(v string) *GetExecutorConfigResponseBodyData {
	s.ExecutorType = &v
	return s
}

func (s *GetExecutorConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
