// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCmsWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutCmsWorkspaceResponseBody
	GetCode() *string
	SetData(v *PutCmsWorkspaceResponseBodyData) *PutCmsWorkspaceResponseBody
	GetData() *PutCmsWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *PutCmsWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PutCmsWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutCmsWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutCmsWorkspaceResponseBody
	GetSuccess() *bool
}

type PutCmsWorkspaceResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *PutCmsWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutCmsWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutCmsWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *PutCmsWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutCmsWorkspaceResponseBody) GetData() *PutCmsWorkspaceResponseBodyData {
	return s.Data
}

func (s *PutCmsWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PutCmsWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutCmsWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCmsWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutCmsWorkspaceResponseBody) SetCode(v string) *PutCmsWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *PutCmsWorkspaceResponseBody) SetData(v *PutCmsWorkspaceResponseBodyData) *PutCmsWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *PutCmsWorkspaceResponseBody) SetHttpStatusCode(v int32) *PutCmsWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PutCmsWorkspaceResponseBody) SetMessage(v string) *PutCmsWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *PutCmsWorkspaceResponseBody) SetRequestId(v string) *PutCmsWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCmsWorkspaceResponseBody) SetSuccess(v bool) *PutCmsWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *PutCmsWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutCmsWorkspaceResponseBodyData struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s PutCmsWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PutCmsWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutCmsWorkspaceResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCmsWorkspaceResponseBodyData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *PutCmsWorkspaceResponseBodyData) SetRequestId(v string) *PutCmsWorkspaceResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *PutCmsWorkspaceResponseBodyData) SetWorkspaceName(v string) *PutCmsWorkspaceResponseBodyData {
	s.WorkspaceName = &v
	return s
}

func (s *PutCmsWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
