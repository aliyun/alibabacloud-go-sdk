// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForDistributeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartTaskForDistributeImageResponseBody
	GetCode() *string
	SetData(v string) *StartTaskForDistributeImageResponseBody
	GetData() *string
	SetMessage(v string) *StartTaskForDistributeImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartTaskForDistributeImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartTaskForDistributeImageResponseBody
	GetSuccess() *bool
}

type StartTaskForDistributeImageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// tid-06xnr5lyp77e7****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 419F31B9-1FDF-5644-ABA3-D00026FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartTaskForDistributeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForDistributeImageResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskForDistributeImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartTaskForDistributeImageResponseBody) GetData() *string {
	return s.Data
}

func (s *StartTaskForDistributeImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartTaskForDistributeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTaskForDistributeImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartTaskForDistributeImageResponseBody) SetCode(v string) *StartTaskForDistributeImageResponseBody {
	s.Code = &v
	return s
}

func (s *StartTaskForDistributeImageResponseBody) SetData(v string) *StartTaskForDistributeImageResponseBody {
	s.Data = &v
	return s
}

func (s *StartTaskForDistributeImageResponseBody) SetMessage(v string) *StartTaskForDistributeImageResponseBody {
	s.Message = &v
	return s
}

func (s *StartTaskForDistributeImageResponseBody) SetRequestId(v string) *StartTaskForDistributeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskForDistributeImageResponseBody) SetSuccess(v bool) *StartTaskForDistributeImageResponseBody {
	s.Success = &v
	return s
}

func (s *StartTaskForDistributeImageResponseBody) Validate() error {
	return dara.Validate(s)
}
