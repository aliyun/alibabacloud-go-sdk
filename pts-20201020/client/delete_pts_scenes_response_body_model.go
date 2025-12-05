// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePtsScenesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePtsScenesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePtsScenesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePtsScenesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePtsScenesResponseBody
	GetSuccess() *bool
}

type DeletePtsScenesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35290A5B-AB50-46BD-81E0-E316F86128C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePtsScenesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePtsScenesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePtsScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePtsScenesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePtsScenesResponseBody) SetCode(v string) *DeletePtsScenesResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetHttpStatusCode(v int32) *DeletePtsScenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetMessage(v string) *DeletePtsScenesResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetRequestId(v string) *DeletePtsScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetSuccess(v bool) *DeletePtsScenesResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePtsScenesResponseBody) Validate() error {
	return dara.Validate(s)
}
