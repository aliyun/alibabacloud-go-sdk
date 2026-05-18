// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeProjectMetaResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeProjectMetaResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeProjectMetaResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeProjectMetaResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeProjectMetaResponseBody
	GetRequestId() *string
	SetResources(v *DescribeProjectMetaResponseBodyResources) *DescribeProjectMetaResponseBody
	GetResources() *DescribeProjectMetaResponseBodyResources
	SetSuccess(v bool) *DescribeProjectMetaResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeProjectMetaResponseBody
	GetTotal() *string
}

type DescribeProjectMetaResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C2061B2-3B1B-43BF-A4A4-C53426F479C0
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *DescribeProjectMetaResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeProjectMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProjectMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProjectMetaResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeProjectMetaResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeProjectMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectMetaResponseBody) GetResources() *DescribeProjectMetaResponseBodyResources {
	return s.Resources
}

func (s *DescribeProjectMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectMetaResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeProjectMetaResponseBody) SetCode(v string) *DescribeProjectMetaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetMessage(v string) *DescribeProjectMetaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageNumber(v string) *DescribeProjectMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageSize(v string) *DescribeProjectMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetRequestId(v string) *DescribeProjectMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetResources(v *DescribeProjectMetaResponseBodyResources) *DescribeProjectMetaResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetSuccess(v bool) *DescribeProjectMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetTotal(v string) *DescribeProjectMetaResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProjectMetaResponseBodyResources struct {
	Resource []*DescribeProjectMetaResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeProjectMetaResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResources) GetResource() []*DescribeProjectMetaResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeProjectMetaResponseBodyResources) SetResource(v []*DescribeProjectMetaResponseBodyResourcesResource) *DescribeProjectMetaResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeProjectMetaResponseBodyResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProjectMetaResponseBodyResourcesResource struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeProjectMetaResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetDescription() *string {
	return s.Description
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetLabels() *string {
	return s.Labels
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetDescription(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetLabels(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetNamespace(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
