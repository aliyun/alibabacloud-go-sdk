// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeOpenApiInfoResponseBodyData) *DescribeOpenApiInfoResponseBody
	GetData() *DescribeOpenApiInfoResponseBodyData
	SetRequestId(v string) *DescribeOpenApiInfoResponseBody
	GetRequestId() *string
}

type DescribeOpenApiInfoResponseBody struct {
	// The data returned.
	Data *DescribeOpenApiInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 358E012F-B516-599D-9ED0-A1A361CDE615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenApiInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponseBody) GetData() *DescribeOpenApiInfoResponseBodyData {
	return s.Data
}

func (s *DescribeOpenApiInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenApiInfoResponseBody) SetData(v *DescribeOpenApiInfoResponseBodyData) *DescribeOpenApiInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenApiInfoResponseBody) SetRequestId(v string) *DescribeOpenApiInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenApiInfoResponseBodyData struct {
	// The description of the API operation.
	//
	// example:
	//
	// describeEcs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input parameters of the API operation.
	//
	// example:
	//
	// {}
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The output parameters of the API operation.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The sample response parameters.
	//
	// example:
	//
	// []
	ResponseDemo *string `json:"ResponseDemo,omitempty" xml:"ResponseDemo,omitempty"`
	// The summary of the API operation.
	//
	// example:
	//
	// describeEcs
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The title of the API operation.
	//
	// example:
	//
	// describeEcs
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeOpenApiInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeOpenApiInfoResponseBodyData) GetInputParams() *string {
	return s.InputParams
}

func (s *DescribeOpenApiInfoResponseBodyData) GetOutputParams() *string {
	return s.OutputParams
}

func (s *DescribeOpenApiInfoResponseBodyData) GetResponseDemo() *string {
	return s.ResponseDemo
}

func (s *DescribeOpenApiInfoResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *DescribeOpenApiInfoResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DescribeOpenApiInfoResponseBodyData) SetDescription(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetInputParams(v string) *DescribeOpenApiInfoResponseBodyData {
	s.InputParams = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetOutputParams(v string) *DescribeOpenApiInfoResponseBodyData {
	s.OutputParams = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetResponseDemo(v string) *DescribeOpenApiInfoResponseBodyData {
	s.ResponseDemo = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetSummary(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Summary = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetTitle(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Title = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
