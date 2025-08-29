// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckInstanceResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*CheckInstanceResourcesResponseBodyResources) *CheckInstanceResourcesResponseBody
	GetResources() []*CheckInstanceResourcesResponseBodyResources
}

type CheckInstanceResourcesResponseBody struct {
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*CheckInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceResourcesResponseBody) GetResources() []*CheckInstanceResourcesResponseBodyResources {
	return s.Resources
}

func (s *CheckInstanceResourcesResponseBody) SetRequestId(v string) *CheckInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceResourcesResponseBody) SetResources(v []*CheckInstanceResourcesResponseBodyResources) *CheckInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *CheckInstanceResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckInstanceResourcesResponseBodyResources struct {
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesResponseBodyResources) GetStatus() *string {
	return s.Status
}

func (s *CheckInstanceResourcesResponseBodyResources) GetType() *string {
	return s.Type
}

func (s *CheckInstanceResourcesResponseBodyResources) GetUri() *string {
	return s.Uri
}

func (s *CheckInstanceResourcesResponseBodyResources) SetStatus(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *CheckInstanceResourcesResponseBodyResources) SetType(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Type = &v
	return s
}

func (s *CheckInstanceResourcesResponseBodyResources) SetUri(v string) *CheckInstanceResourcesResponseBodyResources {
	s.Uri = &v
	return s
}

func (s *CheckInstanceResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
