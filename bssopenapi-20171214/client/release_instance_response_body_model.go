// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseInstanceResponseBody
	GetCode() *string
	SetData(v *ReleaseInstanceResponseBodyData) *ReleaseInstanceResponseBody
	GetData() *ReleaseInstanceResponseBodyData
	SetMessage(v string) *ReleaseInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// Invalid_Product_Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *ReleaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the execution result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A value of true indicates that the execution is complete.
	//
	// A value of false indicates that an error occurs during the execution.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseInstanceResponseBody) GetData() *ReleaseInstanceResponseBodyData {
	return s.Data
}

func (s *ReleaseInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseInstanceResponseBody) SetCode(v string) *ReleaseInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetData(v *ReleaseInstanceResponseBodyData) *ReleaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ReleaseInstanceResponseBody) SetMessage(v string) *ReleaseInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseInstanceResponseBodyData struct {
	// The site of the execution host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates whether the instance is released.
	//
	// example:
	//
	// true
	ReleaseResult *bool `json:"ReleaseResult,omitempty" xml:"ReleaseResult,omitempty"`
}

func (s ReleaseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *ReleaseInstanceResponseBodyData) GetReleaseResult() *bool {
	return s.ReleaseResult
}

func (s *ReleaseInstanceResponseBodyData) SetHostId(v string) *ReleaseInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ReleaseInstanceResponseBodyData) SetReleaseResult(v bool) *ReleaseInstanceResponseBodyData {
	s.ReleaseResult = &v
	return s
}

func (s *ReleaseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
