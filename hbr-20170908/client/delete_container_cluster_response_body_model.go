// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContainerClusterResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteContainerClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContainerClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContainerClusterResponseBody
	GetSuccess() *bool
}

type DeleteContainerClusterResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EB526A5D-1FE2-51C1-B790-1732C1DBA969
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContainerClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContainerClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContainerClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContainerClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContainerClusterResponseBody) SetCode(v string) *DeleteContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContainerClusterResponseBody) SetMessage(v string) *DeleteContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContainerClusterResponseBody) SetRequestId(v string) *DeleteContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContainerClusterResponseBody) SetSuccess(v bool) *DeleteContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContainerClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
