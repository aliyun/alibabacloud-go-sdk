// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowConstantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagConstants(v *ListTaskFlowConstantsResponseBodyDagConstants) *ListTaskFlowConstantsResponseBody
	GetDagConstants() *ListTaskFlowConstantsResponseBodyDagConstants
	SetErrorCode(v string) *ListTaskFlowConstantsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowConstantsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowConstantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowConstantsResponseBody
	GetSuccess() *bool
}

type ListTaskFlowConstantsResponseBody struct {
	// A list of constant key-value pairs for the task flow.
	DagConstants *ListTaskFlowConstantsResponseBodyDagConstants `json:"DagConstants,omitempty" xml:"DagConstants,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 028BF827-3801-5869-8548-F4A039256304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskFlowConstantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowConstantsResponseBody) GetDagConstants() *ListTaskFlowConstantsResponseBodyDagConstants {
	return s.DagConstants
}

func (s *ListTaskFlowConstantsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowConstantsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowConstantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowConstantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowConstantsResponseBody) SetDagConstants(v *ListTaskFlowConstantsResponseBodyDagConstants) *ListTaskFlowConstantsResponseBody {
	s.DagConstants = v
	return s
}

func (s *ListTaskFlowConstantsResponseBody) SetErrorCode(v string) *ListTaskFlowConstantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBody) SetErrorMessage(v string) *ListTaskFlowConstantsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBody) SetRequestId(v string) *ListTaskFlowConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBody) SetSuccess(v bool) *ListTaskFlowConstantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBody) Validate() error {
	if s.DagConstants != nil {
		if err := s.DagConstants.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowConstantsResponseBodyDagConstants struct {
	DagConstant []*ListTaskFlowConstantsResponseBodyDagConstantsDagConstant `json:"DagConstant,omitempty" xml:"DagConstant,omitempty" type:"Repeated"`
}

func (s ListTaskFlowConstantsResponseBodyDagConstants) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowConstantsResponseBodyDagConstants) GoString() string {
	return s.String()
}

func (s *ListTaskFlowConstantsResponseBodyDagConstants) GetDagConstant() []*ListTaskFlowConstantsResponseBodyDagConstantsDagConstant {
	return s.DagConstant
}

func (s *ListTaskFlowConstantsResponseBodyDagConstants) SetDagConstant(v []*ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) *ListTaskFlowConstantsResponseBodyDagConstants {
	s.DagConstant = v
	return s
}

func (s *ListTaskFlowConstantsResponseBodyDagConstants) Validate() error {
	if s.DagConstant != nil {
		for _, item := range s.DagConstant {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowConstantsResponseBodyDagConstantsDagConstant struct {
	// The constant key.
	//
	// example:
	//
	// example
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The constant value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) GoString() string {
	return s.String()
}

func (s *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) GetKey() *string {
	return s.Key
}

func (s *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) GetValue() *string {
	return s.Value
}

func (s *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) SetKey(v string) *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant {
	s.Key = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) SetValue(v string) *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant {
	s.Value = &v
	return s
}

func (s *ListTaskFlowConstantsResponseBodyDagConstantsDagConstant) Validate() error {
	return dara.Validate(s)
}
