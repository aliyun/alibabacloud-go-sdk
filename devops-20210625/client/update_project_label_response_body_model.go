// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateProjectLabelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateProjectLabelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateProjectLabelResponseBody
	GetRequestId() *string
	SetResult(v *UpdateProjectLabelResponseBodyResult) *UpdateProjectLabelResponseBody
	GetResult() *UpdateProjectLabelResponseBodyResult
	SetSuccess(v bool) *UpdateProjectLabelResponseBody
	GetSuccess() *bool
}

type UpdateProjectLabelResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateProjectLabelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProjectLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectLabelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProjectLabelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateProjectLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectLabelResponseBody) GetResult() *UpdateProjectLabelResponseBodyResult {
	return s.Result
}

func (s *UpdateProjectLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectLabelResponseBody) SetErrorCode(v string) *UpdateProjectLabelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectLabelResponseBody) SetErrorMessage(v string) *UpdateProjectLabelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateProjectLabelResponseBody) SetRequestId(v string) *UpdateProjectLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectLabelResponseBody) SetResult(v *UpdateProjectLabelResponseBodyResult) *UpdateProjectLabelResponseBody {
	s.Result = v
	return s
}

func (s *UpdateProjectLabelResponseBody) SetSuccess(v bool) *UpdateProjectLabelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectLabelResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectLabelResponseBodyResult struct {
	// example:
	//
	// #EF433B
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 8d84d73cf315473683400760f02dbfc1
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProjectLabelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectLabelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProjectLabelResponseBodyResult) GetColor() *string {
	return s.Color
}

func (s *UpdateProjectLabelResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectLabelResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateProjectLabelResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateProjectLabelResponseBodyResult) SetColor(v string) *UpdateProjectLabelResponseBodyResult {
	s.Color = &v
	return s
}

func (s *UpdateProjectLabelResponseBodyResult) SetDescription(v string) *UpdateProjectLabelResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateProjectLabelResponseBodyResult) SetId(v string) *UpdateProjectLabelResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateProjectLabelResponseBodyResult) SetName(v string) *UpdateProjectLabelResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateProjectLabelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
