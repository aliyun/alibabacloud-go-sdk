// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteChunkResponseBody
	GetCode() *string
	SetData(v bool) *DeleteChunkResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChunkResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteChunkResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteChunkResponseBody
	GetSuccess() *bool
}

type DeleteChunkResponseBody struct {
	// The status code.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the request is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChunkResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChunkResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChunkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChunkResponseBody) SetCode(v string) *DeleteChunkResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChunkResponseBody) SetData(v bool) *DeleteChunkResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteChunkResponseBody) SetMessage(v string) *DeleteChunkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChunkResponseBody) SetRequestId(v string) *DeleteChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChunkResponseBody) SetStatus(v string) *DeleteChunkResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteChunkResponseBody) SetSuccess(v bool) *DeleteChunkResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChunkResponseBody) Validate() error {
	return dara.Validate(s)
}
