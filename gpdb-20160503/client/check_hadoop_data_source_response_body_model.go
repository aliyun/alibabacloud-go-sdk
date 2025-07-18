// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CheckHadoopDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckHadoopDataSourceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckHadoopDataSourceResponseBody
	GetStatus() *string
}

type CheckHadoopDataSourceResponseBody struct {
	// The returned message. If the service failed, an error message is returned. Otherwise, a pair of double quotation marks ("") is returned.
	//
	// example:
	//
	// serivce unavaliable
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Running
	//
	// 	- Failed
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckHadoopDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckHadoopDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckHadoopDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckHadoopDataSourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckHadoopDataSourceResponseBody) SetMessage(v string) *CheckHadoopDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CheckHadoopDataSourceResponseBody) SetRequestId(v string) *CheckHadoopDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckHadoopDataSourceResponseBody) SetStatus(v string) *CheckHadoopDataSourceResponseBody {
	s.Status = &v
	return s
}

func (s *CheckHadoopDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
