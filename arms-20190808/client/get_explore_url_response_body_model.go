// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExploreUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExploreUrlResponseBody
	GetCode() *int32
	SetData(v string) *GetExploreUrlResponseBody
	GetData() *string
	SetMessage(v string) *GetExploreUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExploreUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExploreUrlResponseBody
	GetSuccess() *bool
}

type GetExploreUrlResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response parameters
	//
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The prompt information of the returned result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful:
	//
	// - true: the operation was successful
	//
	// - false: the operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetExploreUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExploreUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetExploreUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExploreUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetExploreUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExploreUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExploreUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExploreUrlResponseBody) SetCode(v int32) *GetExploreUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetExploreUrlResponseBody) SetData(v string) *GetExploreUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetExploreUrlResponseBody) SetMessage(v string) *GetExploreUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetExploreUrlResponseBody) SetRequestId(v string) *GetExploreUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExploreUrlResponseBody) SetSuccess(v bool) *GetExploreUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetExploreUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
