// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClipsBuildInResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetClipsBuildInResourceResponseBody
	GetCode() *string
	SetData(v *GetClipsBuildInResourceResponseBodyData) *GetClipsBuildInResourceResponseBody
	GetData() *GetClipsBuildInResourceResponseBodyData
	SetHttpStatusCode(v int32) *GetClipsBuildInResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetClipsBuildInResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClipsBuildInResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClipsBuildInResourceResponseBody
	GetSuccess() *bool
}

type GetClipsBuildInResourceResponseBody struct {
	// example:
	//
	// successful
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetClipsBuildInResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClipsBuildInResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClipsBuildInResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetClipsBuildInResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetClipsBuildInResourceResponseBody) GetData() *GetClipsBuildInResourceResponseBodyData {
	return s.Data
}

func (s *GetClipsBuildInResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetClipsBuildInResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClipsBuildInResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClipsBuildInResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClipsBuildInResourceResponseBody) SetCode(v string) *GetClipsBuildInResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) SetData(v *GetClipsBuildInResourceResponseBodyData) *GetClipsBuildInResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) SetHttpStatusCode(v int32) *GetClipsBuildInResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) SetMessage(v string) *GetClipsBuildInResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) SetRequestId(v string) *GetClipsBuildInResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) SetSuccess(v bool) *GetClipsBuildInResourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClipsBuildInResourceResponseBodyData struct {
	ResourceList []*string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	ResourceType *int32    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetClipsBuildInResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClipsBuildInResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClipsBuildInResourceResponseBodyData) GetResourceList() []*string {
	return s.ResourceList
}

func (s *GetClipsBuildInResourceResponseBodyData) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetClipsBuildInResourceResponseBodyData) SetResourceList(v []*string) *GetClipsBuildInResourceResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetClipsBuildInResourceResponseBodyData) SetResourceType(v int32) *GetClipsBuildInResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetClipsBuildInResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
