// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClusterIntoServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddClusterIntoServiceMeshResponseBody
	GetCode() *string
	SetMessage(v string) *AddClusterIntoServiceMeshResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddClusterIntoServiceMeshResponseBody
	GetRequestId() *string
}

type AddClusterIntoServiceMeshResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClusterIntoServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddClusterIntoServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddClusterIntoServiceMeshResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddClusterIntoServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddClusterIntoServiceMeshResponseBody) SetCode(v string) *AddClusterIntoServiceMeshResponseBody {
	s.Code = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponseBody) SetMessage(v string) *AddClusterIntoServiceMeshResponseBody {
	s.Message = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponseBody) SetRequestId(v string) *AddClusterIntoServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
