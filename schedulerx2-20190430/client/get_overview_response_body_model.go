// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetOverviewResponseBody
	GetCode() *int32
	SetData(v string) *GetOverviewResponseBody
	GetData() *string
	SetMessage(v string) *GetOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOverviewResponseBody
	GetSuccess() *bool
}

type GetOverviewResponseBody struct {
	// example:
	//
	// 200
	Code *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// No access permission for the namespace [***]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOverviewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetOverviewResponseBody) GetData() *string {
	return s.Data
}

func (s *GetOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOverviewResponseBody) SetCode(v int32) *GetOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOverviewResponseBody) SetData(v string) *GetOverviewResponseBody {
	s.Data = &v
	return s
}

func (s *GetOverviewResponseBody) SetMessage(v string) *GetOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOverviewResponseBody) SetRequestId(v string) *GetOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOverviewResponseBody) SetSuccess(v bool) *GetOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}
