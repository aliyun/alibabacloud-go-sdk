// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateZnodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateZnodeResponseBodyData) *CreateZnodeResponseBody
	GetData() *CreateZnodeResponseBodyData
	SetErrorCode(v string) *CreateZnodeResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *CreateZnodeResponseBody
	GetHttpCode() *string
	SetMessage(v string) *CreateZnodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateZnodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateZnodeResponseBody
	GetSuccess() *bool
}

type CreateZnodeResponseBody struct {
	// The details of the data.
	Data *CreateZnodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
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
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
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

func (s CreateZnodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponseBody) GetData() *CreateZnodeResponseBodyData {
	return s.Data
}

func (s *CreateZnodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateZnodeResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *CreateZnodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateZnodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateZnodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateZnodeResponseBody) SetData(v *CreateZnodeResponseBodyData) *CreateZnodeResponseBody {
	s.Data = v
	return s
}

func (s *CreateZnodeResponseBody) SetErrorCode(v string) *CreateZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateZnodeResponseBody) SetHttpCode(v string) *CreateZnodeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateZnodeResponseBody) SetMessage(v string) *CreateZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateZnodeResponseBody) SetRequestId(v string) *CreateZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateZnodeResponseBody) SetSuccess(v bool) *CreateZnodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateZnodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateZnodeResponseBodyData struct {
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
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateZnodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateZnodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponseBodyData) GetData() *string {
	return s.Data
}

func (s *CreateZnodeResponseBodyData) GetDir() *bool {
	return s.Dir
}

func (s *CreateZnodeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateZnodeResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *CreateZnodeResponseBodyData) SetData(v string) *CreateZnodeResponseBodyData {
	s.Data = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetDir(v bool) *CreateZnodeResponseBodyData {
	s.Dir = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetName(v string) *CreateZnodeResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetPath(v string) *CreateZnodeResponseBodyData {
	s.Path = &v
	return s
}

func (s *CreateZnodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
