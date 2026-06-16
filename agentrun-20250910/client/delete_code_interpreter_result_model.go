// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeInterpreterResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCodeInterpreterResult
	GetCode() *string
	SetData(v *CodeInterpreter) *DeleteCodeInterpreterResult
	GetData() *CodeInterpreter
	SetRequestId(v string) *DeleteCodeInterpreterResult
	GetRequestId() *string
}

type DeleteCodeInterpreterResult struct {
	// Returns `SUCCESS` if the operation is successful. On failure, returns an error code, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Details of the deleted code interpreter.
	//
	// example:
	//
	// {}
	Data *CodeInterpreter `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request ID for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCodeInterpreterResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeInterpreterResult) GoString() string {
	return s.String()
}

func (s *DeleteCodeInterpreterResult) GetCode() *string {
	return s.Code
}

func (s *DeleteCodeInterpreterResult) GetData() *CodeInterpreter {
	return s.Data
}

func (s *DeleteCodeInterpreterResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCodeInterpreterResult) SetCode(v string) *DeleteCodeInterpreterResult {
	s.Code = &v
	return s
}

func (s *DeleteCodeInterpreterResult) SetData(v *CodeInterpreter) *DeleteCodeInterpreterResult {
	s.Data = v
	return s
}

func (s *DeleteCodeInterpreterResult) SetRequestId(v string) *DeleteCodeInterpreterResult {
	s.RequestId = &v
	return s
}

func (s *DeleteCodeInterpreterResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
