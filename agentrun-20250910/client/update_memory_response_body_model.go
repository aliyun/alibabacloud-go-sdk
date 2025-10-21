// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryResponseBody
	GetCode() *string
	SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody
	GetData() *UpdateMemoryResponseBodyData
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
}

type UpdateMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C0595DB0-D1EE-55C3-8DDD-790872C7EC2F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryResponseBody) GetData() *UpdateMemoryResponseBodyData {
	return s.Data
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) SetCode(v string) *UpdateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryResponseBodyData struct {
	// example:
	//
	// default_workspace
	CmsWorkspaceName *string `json:"cmsWorkspaceName,omitempty" xml:"cmsWorkspaceName,omitempty"`
}

func (s UpdateMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyData) GetCmsWorkspaceName() *string {
	return s.CmsWorkspaceName
}

func (s *UpdateMemoryResponseBodyData) SetCmsWorkspaceName(v string) *UpdateMemoryResponseBodyData {
	s.CmsWorkspaceName = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
