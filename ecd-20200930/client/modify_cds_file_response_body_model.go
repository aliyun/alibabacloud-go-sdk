// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCdsFileResponseBody
	GetCode() *string
	SetData(v string) *ModifyCdsFileResponseBody
	GetData() *string
	SetMessage(v string) *ModifyCdsFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCdsFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCdsFileResponseBody
	GetSuccess() *bool
}

type ModifyCdsFileResponseBody struct {
	// The modification result. The value success indicates that the modification is successful. If the modification failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Specifies whether data is returned.
	//
	// Valid values:
	//
	// 	- false: Data fails to be returned.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     : Data is returned.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true: The request is successful.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false: The request fails.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCdsFileResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyCdsFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdsFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCdsFileResponseBody) SetCode(v string) *ModifyCdsFileResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCdsFileResponseBody) SetData(v string) *ModifyCdsFileResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCdsFileResponseBody) SetMessage(v string) *ModifyCdsFileResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCdsFileResponseBody) SetRequestId(v string) *ModifyCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdsFileResponseBody) SetSuccess(v bool) *ModifyCdsFileResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCdsFileResponseBody) Validate() error {
	return dara.Validate(s)
}
