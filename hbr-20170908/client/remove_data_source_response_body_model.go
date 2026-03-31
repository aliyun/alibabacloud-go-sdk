// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveDataSourceResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDataSourceResponseBody
	GetSuccess() *bool
}

type RemoveDataSourceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ED2BDBAC-***-***-***-495C96A63964
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDataSourceResponseBody) SetCode(v string) *RemoveDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveDataSourceResponseBody) SetMessage(v string) *RemoveDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveDataSourceResponseBody) SetRequestId(v string) *RemoveDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataSourceResponseBody) SetSuccess(v bool) *RemoveDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
