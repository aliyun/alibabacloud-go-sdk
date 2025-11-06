// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryZnodeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryZnodeDetailResponseBodyData) *QueryZnodeDetailResponseBody
	GetData() *QueryZnodeDetailResponseBodyData
	SetErrorCode(v string) *QueryZnodeDetailResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryZnodeDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryZnodeDetailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryZnodeDetailResponseBody
	GetSuccess() *string
}

type QueryZnodeDetailResponseBody struct {
	// The details of the data.
	Data *QueryZnodeDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 58E06A0A-BD2C-47A0-99C2-B100F353****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryZnodeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryZnodeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponseBody) GetData() *QueryZnodeDetailResponseBodyData {
	return s.Data
}

func (s *QueryZnodeDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryZnodeDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryZnodeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryZnodeDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryZnodeDetailResponseBody) SetData(v *QueryZnodeDetailResponseBodyData) *QueryZnodeDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetErrorCode(v string) *QueryZnodeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetMessage(v string) *QueryZnodeDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetRequestId(v string) *QueryZnodeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetSuccess(v string) *QueryZnodeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryZnodeDetailResponseBodyData struct {
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
	// zookeeper
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// /zookeeper
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s QueryZnodeDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryZnodeDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponseBodyData) GetData() *string {
	return s.Data
}

func (s *QueryZnodeDetailResponseBodyData) GetDir() *bool {
	return s.Dir
}

func (s *QueryZnodeDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryZnodeDetailResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *QueryZnodeDetailResponseBodyData) SetData(v string) *QueryZnodeDetailResponseBodyData {
	s.Data = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetDir(v bool) *QueryZnodeDetailResponseBodyData {
	s.Dir = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetName(v string) *QueryZnodeDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetPath(v string) *QueryZnodeDetailResponseBodyData {
	s.Path = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
