// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDistributionChannelResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopDistributionChannelResult
	GetCode() *string
	SetMessage(v string) *StopDistributionChannelResult
	GetMessage() *string
	SetRequestId(v string) *StopDistributionChannelResult
	GetRequestId() *string
}

type StopDistributionChannelResult struct {
	// The return code.
	//
	// example:
	//
	// CODE_SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the return message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopDistributionChannelResult) String() string {
	return dara.Prettify(s)
}

func (s StopDistributionChannelResult) GoString() string {
	return s.String()
}

func (s *StopDistributionChannelResult) GetCode() *string {
	return s.Code
}

func (s *StopDistributionChannelResult) GetMessage() *string {
	return s.Message
}

func (s *StopDistributionChannelResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDistributionChannelResult) SetCode(v string) *StopDistributionChannelResult {
	s.Code = &v
	return s
}

func (s *StopDistributionChannelResult) SetMessage(v string) *StopDistributionChannelResult {
	s.Message = &v
	return s
}

func (s *StopDistributionChannelResult) SetRequestId(v string) *StopDistributionChannelResult {
	s.RequestId = &v
	return s
}

func (s *StopDistributionChannelResult) Validate() error {
	return dara.Validate(s)
}
