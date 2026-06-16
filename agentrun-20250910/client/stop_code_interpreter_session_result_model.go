// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCodeInterpreterSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopCodeInterpreterSessionResult
	GetCode() *string
	SetRequestId(v string) *StopCodeInterpreterSessionResult
	GetRequestId() *string
}

type StopCodeInterpreterSessionResult struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Request ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopCodeInterpreterSessionResult) String() string {
	return dara.Prettify(s)
}

func (s StopCodeInterpreterSessionResult) GoString() string {
	return s.String()
}

func (s *StopCodeInterpreterSessionResult) GetCode() *string {
	return s.Code
}

func (s *StopCodeInterpreterSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCodeInterpreterSessionResult) SetCode(v string) *StopCodeInterpreterSessionResult {
	s.Code = &v
	return s
}

func (s *StopCodeInterpreterSessionResult) SetRequestId(v string) *StopCodeInterpreterSessionResult {
	s.RequestId = &v
	return s
}

func (s *StopCodeInterpreterSessionResult) Validate() error {
	return dara.Validate(s)
}
