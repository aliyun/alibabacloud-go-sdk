// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKibanaPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListKibanaPluginsResponseBodyHeaders) *ListKibanaPluginsResponseBody
	GetHeaders() *ListKibanaPluginsResponseBodyHeaders
	SetRequestId(v string) *ListKibanaPluginsResponseBody
	GetRequestId() *string
	SetResult(v []*ListKibanaPluginsResponseBodyResult) *ListKibanaPluginsResponseBody
	GetResult() []*ListKibanaPluginsResponseBodyResult
}

type ListKibanaPluginsResponseBody struct {
	// The request header.
	Headers *ListKibanaPluginsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 11234B4A-34CE-473B-8E61-AD95702E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the plug-ins.
	Result []*ListKibanaPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListKibanaPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBody) GetHeaders() *ListKibanaPluginsResponseBodyHeaders {
	return s.Headers
}

func (s *ListKibanaPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKibanaPluginsResponseBody) GetResult() []*ListKibanaPluginsResponseBodyResult {
	return s.Result
}

func (s *ListKibanaPluginsResponseBody) SetHeaders(v *ListKibanaPluginsResponseBodyHeaders) *ListKibanaPluginsResponseBody {
	s.Headers = v
	return s
}

func (s *ListKibanaPluginsResponseBody) SetRequestId(v string) *ListKibanaPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKibanaPluginsResponseBody) SetResult(v []*ListKibanaPluginsResponseBodyResult) *ListKibanaPluginsResponseBody {
	s.Result = v
	return s
}

func (s *ListKibanaPluginsResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKibanaPluginsResponseBodyHeaders struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListKibanaPluginsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPluginsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListKibanaPluginsResponseBodyHeaders) SetXTotalCount(v int32) *ListKibanaPluginsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListKibanaPluginsResponseBodyResult struct {
	// The description of the plug-in.
	//
	// example:
	//
	// Customize DSL statements to query data.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// bsearch_querybuilder
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source of the plug-in.
	//
	// example:
	//
	// SYSTEM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The URL of the introduction to the plug-in. The value null is supported.
	//
	// example:
	//
	// https://xxxx
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	// The installation status of the plug-in.
	//
	// example:
	//
	// INSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListKibanaPluginsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListKibanaPluginsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListKibanaPluginsResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListKibanaPluginsResponseBodyResult) GetSpecificationUrl() *string {
	return s.SpecificationUrl
}

func (s *ListKibanaPluginsResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListKibanaPluginsResponseBodyResult) SetDescription(v string) *ListKibanaPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetName(v string) *ListKibanaPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetSource(v string) *ListKibanaPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListKibanaPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetState(v string) *ListKibanaPluginsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
