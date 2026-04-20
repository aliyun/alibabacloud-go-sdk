// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DropUserResponseBody
	GetData() *bool
	SetErrorCode(v string) *DropUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DropUserResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DropUserResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DropUserResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DropUserResponseBody
	GetSuccess() *string
}

type DropUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EA8F0084-5831-5907-BB31-BD05D2617844
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DropUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropUserResponseBody) GoString() string {
	return s.String()
}

func (s *DropUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *DropUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DropUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DropUserResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DropUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropUserResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DropUserResponseBody) SetData(v bool) *DropUserResponseBody {
	s.Data = &v
	return s
}

func (s *DropUserResponseBody) SetErrorCode(v string) *DropUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DropUserResponseBody) SetErrorMessage(v string) *DropUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DropUserResponseBody) SetHttpStatusCode(v string) *DropUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DropUserResponseBody) SetRequestId(v string) *DropUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropUserResponseBody) SetSuccess(v string) *DropUserResponseBody {
	s.Success = &v
	return s
}

func (s *DropUserResponseBody) Validate() error {
	return dara.Validate(s)
}
