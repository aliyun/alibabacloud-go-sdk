// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateWorkitemV2ResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateWorkitemV2ResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateWorkitemV2ResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateWorkitemV2ResponseBody
	GetSuccess() *string
	SetWorkitemIdentifier(v string) *CreateWorkitemV2ResponseBody
	GetWorkitemIdentifier() *string
}

type CreateWorkitemV2ResponseBody struct {
	// example:
	//
	// InvalidTagGroup.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// EAE03103-5497-58D1-9169-E524DDE8604C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 11234455454355
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemV2ResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateWorkitemV2ResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateWorkitemV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkitemV2ResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateWorkitemV2ResponseBody) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemV2ResponseBody) SetErrorCode(v string) *CreateWorkitemV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemV2ResponseBody) SetErrorMessage(v string) *CreateWorkitemV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateWorkitemV2ResponseBody) SetRequestId(v string) *CreateWorkitemV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemV2ResponseBody) SetSuccess(v string) *CreateWorkitemV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemV2ResponseBody) SetWorkitemIdentifier(v string) *CreateWorkitemV2ResponseBody {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
