// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDefaultCollectorConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDefaultCollectorConfigurationsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDefaultCollectorConfigurationsResponseBodyResult) *ListDefaultCollectorConfigurationsResponseBody
	GetResult() []*ListDefaultCollectorConfigurationsResponseBodyResult
}

type ListDefaultCollectorConfigurationsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8BAE3C32-8E4A-47D6-B4B0-95B5DE643BF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListDefaultCollectorConfigurationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDefaultCollectorConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDefaultCollectorConfigurationsResponseBody) GetResult() []*ListDefaultCollectorConfigurationsResponseBodyResult {
	return s.Result
}

func (s *ListDefaultCollectorConfigurationsResponseBody) SetRequestId(v string) *ListDefaultCollectorConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBody) SetResult(v []*ListDefaultCollectorConfigurationsResponseBodyResult) *ListDefaultCollectorConfigurationsResponseBody {
	s.Result = v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDefaultCollectorConfigurationsResponseBodyResult struct {
	// The content of the configuration file.
	//
	// example:
	//
	// - key: log\\n  title: Log file content\\n  description: >\\n    Contains log file lines.\\n  fields:\\n ......
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the configuration file.
	//
	// example:
	//
	// fields.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListDefaultCollectorConfigurationsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) GetFileName() *string {
	return s.FileName
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) SetContent(v string) *ListDefaultCollectorConfigurationsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) SetFileName(v string) *ListDefaultCollectorConfigurationsResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
