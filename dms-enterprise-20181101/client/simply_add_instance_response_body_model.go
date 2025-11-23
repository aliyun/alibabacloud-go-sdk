// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimplyAddInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SimplyAddInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SimplyAddInstanceResponseBody
	GetErrorMessage() *string
	SetInstance(v *SimplyAddInstanceResponseBodyInstance) *SimplyAddInstanceResponseBody
	GetInstance() *SimplyAddInstanceResponseBodyInstance
	SetRequestId(v string) *SimplyAddInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SimplyAddInstanceResponseBody
	GetSuccess() *bool
}

type SimplyAddInstanceResponseBody struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Instance     *SimplyAddInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SimplyAddInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SimplyAddInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SimplyAddInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SimplyAddInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SimplyAddInstanceResponseBody) GetInstance() *SimplyAddInstanceResponseBodyInstance {
	return s.Instance
}

func (s *SimplyAddInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SimplyAddInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SimplyAddInstanceResponseBody) SetErrorCode(v string) *SimplyAddInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SimplyAddInstanceResponseBody) SetErrorMessage(v string) *SimplyAddInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SimplyAddInstanceResponseBody) SetInstance(v *SimplyAddInstanceResponseBodyInstance) *SimplyAddInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *SimplyAddInstanceResponseBody) SetRequestId(v string) *SimplyAddInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SimplyAddInstanceResponseBody) SetSuccess(v bool) *SimplyAddInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *SimplyAddInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SimplyAddInstanceResponseBodyInstance struct {
	// example:
	//
	// 192.168.XXX.XXX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 188****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SimplyAddInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s SimplyAddInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *SimplyAddInstanceResponseBodyInstance) GetHost() *string {
	return s.Host
}

func (s *SimplyAddInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SimplyAddInstanceResponseBodyInstance) GetPort() *string {
	return s.Port
}

func (s *SimplyAddInstanceResponseBodyInstance) SetHost(v string) *SimplyAddInstanceResponseBodyInstance {
	s.Host = &v
	return s
}

func (s *SimplyAddInstanceResponseBodyInstance) SetInstanceId(v string) *SimplyAddInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *SimplyAddInstanceResponseBodyInstance) SetPort(v string) *SimplyAddInstanceResponseBodyInstance {
	s.Port = &v
	return s
}

func (s *SimplyAddInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
