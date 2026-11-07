// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZnodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteZnodeResponseBodyData) *DeleteZnodeResponseBody
	GetData() *DeleteZnodeResponseBodyData
	SetErrorCode(v string) *DeleteZnodeResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteZnodeResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteZnodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteZnodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteZnodeResponseBody
	GetSuccess() *bool
}

type DeleteZnodeResponseBody struct {
	// The data overview.
	Data *DeleteZnodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// - `true`: The request succeeded.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteZnodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponseBody) GetData() *DeleteZnodeResponseBodyData {
	return s.Data
}

func (s *DeleteZnodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteZnodeResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteZnodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteZnodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteZnodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteZnodeResponseBody) SetData(v *DeleteZnodeResponseBodyData) *DeleteZnodeResponseBody {
	s.Data = v
	return s
}

func (s *DeleteZnodeResponseBody) SetErrorCode(v string) *DeleteZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetHttpCode(v string) *DeleteZnodeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetMessage(v string) *DeleteZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetRequestId(v string) *DeleteZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetSuccess(v bool) *DeleteZnodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteZnodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteZnodeResponseBodyData struct {
	// The list of node data.
	//
	// example:
	//
	// cluster
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The node list information. Valid values:
	//
	// - `true`: The information is returned successfully.
	//
	// - `false`: The information fails to be returned.
	//
	// example:
	//
	// true
	Dir *bool `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The node name.
	//
	// example:
	//
	// mse-bc1a29b0-160230875****-reg-center-0-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node path.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteZnodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteZnodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponseBodyData) GetData() *string {
	return s.Data
}

func (s *DeleteZnodeResponseBodyData) GetDir() *bool {
	return s.Dir
}

func (s *DeleteZnodeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteZnodeResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *DeleteZnodeResponseBodyData) SetData(v string) *DeleteZnodeResponseBodyData {
	s.Data = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetDir(v bool) *DeleteZnodeResponseBodyData {
	s.Dir = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetName(v string) *DeleteZnodeResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetPath(v string) *DeleteZnodeResponseBodyData {
	s.Path = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
