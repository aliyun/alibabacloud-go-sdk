// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListPluginsResponseBodyHeaders) *ListPluginsResponseBody
	GetHeaders() *ListPluginsResponseBodyHeaders
	SetRequestId(v string) *ListPluginsResponseBody
	GetRequestId() *string
	SetResult(v []*ListPluginsResponseBodyResult) *ListPluginsResponseBody
	GetResult() []*ListPluginsResponseBodyResult
}

type ListPluginsResponseBody struct {
	// The description of the plug-in.
	Headers *ListPluginsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The return results.
	//
	// example:
	//
	// 5A5D8E74-565C-43DC-B031-29289FA9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the plug-in.
	Result []*ListPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBody) GetHeaders() *ListPluginsResponseBodyHeaders {
	return s.Headers
}

func (s *ListPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginsResponseBody) GetResult() []*ListPluginsResponseBodyResult {
	return s.Result
}

func (s *ListPluginsResponseBody) SetHeaders(v *ListPluginsResponseBodyHeaders) *ListPluginsResponseBody {
	s.Headers = v
	return s
}

func (s *ListPluginsResponseBody) SetRequestId(v string) *ListPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginsResponseBody) SetResult(v []*ListPluginsResponseBodyResult) *ListPluginsResponseBody {
	s.Result = v
	return s
}

func (s *ListPluginsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPluginsResponseBodyHeaders struct {
	// The address of the plug-in description document.
	//
	// example:
	//
	// 2
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListPluginsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListPluginsResponseBodyHeaders) SetXTotalCount(v int32) *ListPluginsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListPluginsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListPluginsResponseBodyResult struct {
	// The source type of the plug-in.
	//
	// example:
	//
	// IK analysis plug-in for Elasticsearch.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// analysis-ik
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// SYSTEM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// https://xxxx.html
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	// example:
	//
	// INSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListPluginsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListPluginsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListPluginsResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListPluginsResponseBodyResult) GetSpecificationUrl() *string {
	return s.SpecificationUrl
}

func (s *ListPluginsResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListPluginsResponseBodyResult) SetDescription(v string) *ListPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetName(v string) *ListPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetSource(v string) *ListPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetState(v string) *ListPluginsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListPluginsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
