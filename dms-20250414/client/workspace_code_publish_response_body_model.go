// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceCodePublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *WorkspaceCodePublishResponseBodyData) *WorkspaceCodePublishResponseBody
	GetData() *WorkspaceCodePublishResponseBodyData
	SetErrorCode(v string) *WorkspaceCodePublishResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *WorkspaceCodePublishResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *WorkspaceCodePublishResponseBody
	GetMessage() *string
	SetRequestId(v string) *WorkspaceCodePublishResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WorkspaceCodePublishResponseBody
	GetSuccess() *bool
}

type WorkspaceCodePublishResponseBody struct {
	// job
	Data *WorkspaceCodePublishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Failed to deploy，repo branch empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WorkspaceCodePublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceCodePublishResponseBody) GoString() string {
	return s.String()
}

func (s *WorkspaceCodePublishResponseBody) GetData() *WorkspaceCodePublishResponseBodyData {
	return s.Data
}

func (s *WorkspaceCodePublishResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WorkspaceCodePublishResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WorkspaceCodePublishResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WorkspaceCodePublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WorkspaceCodePublishResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WorkspaceCodePublishResponseBody) SetData(v *WorkspaceCodePublishResponseBodyData) *WorkspaceCodePublishResponseBody {
	s.Data = v
	return s
}

func (s *WorkspaceCodePublishResponseBody) SetErrorCode(v string) *WorkspaceCodePublishResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WorkspaceCodePublishResponseBody) SetHttpStatusCode(v int32) *WorkspaceCodePublishResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WorkspaceCodePublishResponseBody) SetMessage(v string) *WorkspaceCodePublishResponseBody {
	s.Message = &v
	return s
}

func (s *WorkspaceCodePublishResponseBody) SetRequestId(v string) *WorkspaceCodePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *WorkspaceCodePublishResponseBody) SetSuccess(v bool) *WorkspaceCodePublishResponseBody {
	s.Success = &v
	return s
}

func (s *WorkspaceCodePublishResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WorkspaceCodePublishResponseBodyData struct {
	// example:
	//
	// ws-xxxx-xxxxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s WorkspaceCodePublishResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceCodePublishResponseBodyData) GoString() string {
	return s.String()
}

func (s *WorkspaceCodePublishResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *WorkspaceCodePublishResponseBodyData) SetKey(v string) *WorkspaceCodePublishResponseBodyData {
	s.Key = &v
	return s
}

func (s *WorkspaceCodePublishResponseBodyData) Validate() error {
	return dara.Validate(s)
}
