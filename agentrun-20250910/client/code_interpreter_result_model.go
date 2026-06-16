// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreterResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CodeInterpreterResult
	GetCode() *string
	SetData(v *CodeInterpreter) *CodeInterpreterResult
	GetData() *CodeInterpreter
	SetRequestId(v string) *CodeInterpreterResult
	GetRequestId() *string
}

type CodeInterpreterResult struct {
	// The result code of the operation. A value of `SUCCESS` indicates success, while a failed operation returns an error type such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Details about the code interpreter.
	//
	// example:
	//
	// {}
	Data *CodeInterpreter `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request identifier, used for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CodeInterpreterResult) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreterResult) GoString() string {
	return s.String()
}

func (s *CodeInterpreterResult) GetCode() *string {
	return s.Code
}

func (s *CodeInterpreterResult) GetData() *CodeInterpreter {
	return s.Data
}

func (s *CodeInterpreterResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CodeInterpreterResult) SetCode(v string) *CodeInterpreterResult {
	s.Code = &v
	return s
}

func (s *CodeInterpreterResult) SetData(v *CodeInterpreter) *CodeInterpreterResult {
	s.Data = v
	return s
}

func (s *CodeInterpreterResult) SetRequestId(v string) *CodeInterpreterResult {
	s.RequestId = &v
	return s
}

func (s *CodeInterpreterResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
