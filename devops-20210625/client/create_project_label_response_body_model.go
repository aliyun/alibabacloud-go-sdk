// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateProjectLabelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateProjectLabelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateProjectLabelResponseBody
	GetRequestId() *string
	SetResult(v *CreateProjectLabelResponseBodyResult) *CreateProjectLabelResponseBody
	GetResult() *CreateProjectLabelResponseBodyResult
	SetSuccess(v bool) *CreateProjectLabelResponseBody
	GetSuccess() *bool
}

type CreateProjectLabelResponseBody struct {
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
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateProjectLabelResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProjectLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectLabelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectLabelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProjectLabelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateProjectLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectLabelResponseBody) GetResult() *CreateProjectLabelResponseBodyResult {
	return s.Result
}

func (s *CreateProjectLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProjectLabelResponseBody) SetErrorCode(v string) *CreateProjectLabelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectLabelResponseBody) SetErrorMessage(v string) *CreateProjectLabelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProjectLabelResponseBody) SetRequestId(v string) *CreateProjectLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectLabelResponseBody) SetResult(v *CreateProjectLabelResponseBodyResult) *CreateProjectLabelResponseBody {
	s.Result = v
	return s
}

func (s *CreateProjectLabelResponseBody) SetSuccess(v bool) *CreateProjectLabelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectLabelResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectLabelResponseBodyResult struct {
	// example:
	//
	// #006AD4
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// b5b5738b94954bc6aa5a293316ed1d24
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProjectLabelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectLabelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProjectLabelResponseBodyResult) GetColor() *string {
	return s.Color
}

func (s *CreateProjectLabelResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectLabelResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateProjectLabelResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateProjectLabelResponseBodyResult) SetColor(v string) *CreateProjectLabelResponseBodyResult {
	s.Color = &v
	return s
}

func (s *CreateProjectLabelResponseBodyResult) SetDescription(v string) *CreateProjectLabelResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateProjectLabelResponseBodyResult) SetId(v string) *CreateProjectLabelResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateProjectLabelResponseBodyResult) SetName(v string) *CreateProjectLabelResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateProjectLabelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
