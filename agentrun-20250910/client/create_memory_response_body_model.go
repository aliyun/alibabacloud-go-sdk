// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryResponseBody
	GetCode() *string
	SetData(v *CreateMemoryResponseBodyData) *CreateMemoryResponseBody
	GetData() *CreateMemoryResponseBodyData
	SetRequestId(v string) *CreateMemoryResponseBody
	GetRequestId() *string
}

type CreateMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0FB1162C-D50B-5DA7-AD04-3417ABBF133A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryResponseBody) GetData() *CreateMemoryResponseBodyData {
	return s.Data
}

func (s *CreateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryResponseBody) SetCode(v string) *CreateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryResponseBody) SetData(v *CreateMemoryResponseBodyData) *CreateMemoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryResponseBodyData struct {
	// example:
	//
	// default_workspace
	CmsWorkspaceName *string `json:"cmsWorkspaceName,omitempty" xml:"cmsWorkspaceName,omitempty"`
}

func (s CreateMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyData) GetCmsWorkspaceName() *string {
	return s.CmsWorkspaceName
}

func (s *CreateMemoryResponseBodyData) SetCmsWorkspaceName(v string) *CreateMemoryResponseBodyData {
	s.CmsWorkspaceName = &v
	return s
}

func (s *CreateMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
