// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMemoryResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteMemoryResponseBody
	GetRequestId() *string
}

type DeleteMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 970F08C9-EB28-5A3D-A228-D541AEC4C807
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBody) SetCode(v string) *DeleteMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
