// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterFromServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveClusterFromServiceMeshResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveClusterFromServiceMeshResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveClusterFromServiceMeshResponseBody
	GetRequestId() *string
}

type RemoveClusterFromServiceMeshResponseBody struct {
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

func (s RemoveClusterFromServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterFromServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveClusterFromServiceMeshResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveClusterFromServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetCode(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetMessage(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetRequestId(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
