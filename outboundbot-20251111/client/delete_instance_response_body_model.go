// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceResponseBody
	GetCode() *string
	SetData(v string) *DeleteInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInstanceResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteInstanceResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceResponseBody
	GetSuccess() *bool
}

type DeleteInstanceResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=0987654321
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetData(v string) *DeleteInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetParams(v []*string) *DeleteInstanceResponseBody {
	s.Params = v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
