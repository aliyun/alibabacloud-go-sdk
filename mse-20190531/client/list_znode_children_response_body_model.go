// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZnodeChildrenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListZnodeChildrenResponseBodyData) *ListZnodeChildrenResponseBody
	GetData() []*ListZnodeChildrenResponseBodyData
	SetErrorCode(v string) *ListZnodeChildrenResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListZnodeChildrenResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListZnodeChildrenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListZnodeChildrenResponseBody
	GetSuccess() *bool
}

type ListZnodeChildrenResponseBody struct {
	// The details of the data.
	Data []*ListZnodeChildrenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDB6CE0B-9CAF-41B5-9FEA-E08BE8E2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListZnodeChildrenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListZnodeChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponseBody) GetData() []*ListZnodeChildrenResponseBodyData {
	return s.Data
}

func (s *ListZnodeChildrenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListZnodeChildrenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListZnodeChildrenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListZnodeChildrenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListZnodeChildrenResponseBody) SetData(v []*ListZnodeChildrenResponseBodyData) *ListZnodeChildrenResponseBody {
	s.Data = v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetErrorCode(v string) *ListZnodeChildrenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetMessage(v string) *ListZnodeChildrenResponseBody {
	s.Message = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetRequestId(v string) *ListZnodeChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetSuccess(v bool) *ListZnodeChildrenResponseBody {
	s.Success = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListZnodeChildrenResponseBodyData struct {
	// The data of the node.
	//
	// example:
	//
	// cluster
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Indicates whether the node information was returned. Valid values:
	//
	// 	- `true`: The node information was returned.
	//
	// 	- `false`: The node information failed to be returned.
	//
	// example:
	//
	// true
	Dir *bool `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// mse-bc1a29b0-160230875****-reg-center-0-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// /zookeeper
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListZnodeChildrenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListZnodeChildrenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponseBodyData) GetData() *string {
	return s.Data
}

func (s *ListZnodeChildrenResponseBodyData) GetDir() *bool {
	return s.Dir
}

func (s *ListZnodeChildrenResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListZnodeChildrenResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *ListZnodeChildrenResponseBodyData) SetData(v string) *ListZnodeChildrenResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetDir(v bool) *ListZnodeChildrenResponseBodyData {
	s.Dir = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetName(v string) *ListZnodeChildrenResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetPath(v string) *ListZnodeChildrenResponseBodyData {
	s.Path = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
