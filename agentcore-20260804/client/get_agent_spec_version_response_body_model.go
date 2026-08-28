// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgentSpecVersionResponseBodyData) *GetAgentSpecVersionResponseBody
	GetData() *GetAgentSpecVersionResponseBodyData
	SetRequestId(v string) *GetAgentSpecVersionResponseBody
	GetRequestId() *string
}

type GetAgentSpecVersionResponseBody struct {
	// The returned data.
	Data *GetAgentSpecVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAgentSpecVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSpecVersionResponseBody) GetData() *GetAgentSpecVersionResponseBodyData {
	return s.Data
}

func (s *GetAgentSpecVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSpecVersionResponseBody) SetData(v *GetAgentSpecVersionResponseBodyData) *GetAgentSpecVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentSpecVersionResponseBody) SetRequestId(v string) *GetAgentSpecVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSpecVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSpecVersionResponseBodyData struct {
	// The business tags.
	//
	// example:
	//
	// Sample property value
	BizTags *string `json:"bizTags,omitempty" xml:"bizTags,omitempty"`
	// The content.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource file mapping.
	Resource map[string]*DataResourceValue `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s GetAgentSpecVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentSpecVersionResponseBodyData) GetBizTags() *string {
	return s.BizTags
}

func (s *GetAgentSpecVersionResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetAgentSpecVersionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentSpecVersionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAgentSpecVersionResponseBodyData) GetResource() map[string]*DataResourceValue {
	return s.Resource
}

func (s *GetAgentSpecVersionResponseBodyData) SetBizTags(v string) *GetAgentSpecVersionResponseBodyData {
	s.BizTags = &v
	return s
}

func (s *GetAgentSpecVersionResponseBodyData) SetContent(v string) *GetAgentSpecVersionResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetAgentSpecVersionResponseBodyData) SetDescription(v string) *GetAgentSpecVersionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentSpecVersionResponseBodyData) SetName(v string) *GetAgentSpecVersionResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentSpecVersionResponseBodyData) SetResource(v map[string]*DataResourceValue) *GetAgentSpecVersionResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetAgentSpecVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
