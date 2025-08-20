// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkTemplateFileIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *ListSparkTemplateFileIdsResponseBody
	GetData() []*int64
	SetRequestId(v string) *ListSparkTemplateFileIdsResponseBody
	GetRequestId() *string
}

type ListSparkTemplateFileIdsResponseBody struct {
	// The IDs of Spark template files.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkTemplateFileIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSparkTemplateFileIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *ListSparkTemplateFileIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSparkTemplateFileIdsResponseBody) SetData(v []*int64) *ListSparkTemplateFileIdsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkTemplateFileIdsResponseBody) SetRequestId(v string) *ListSparkTemplateFileIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkTemplateFileIdsResponseBody) Validate() error {
	return dara.Validate(s)
}
