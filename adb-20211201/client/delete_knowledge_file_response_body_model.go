// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteKnowledgeFileResponseBodyData) *DeleteKnowledgeFileResponseBody
	GetData() *DeleteKnowledgeFileResponseBodyData
	SetRequestId(v string) *DeleteKnowledgeFileResponseBody
	GetRequestId() *string
}

type DeleteKnowledgeFileResponseBody struct {
	// The returned data.
	Data *DeleteKnowledgeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeFileResponseBody) GetData() *DeleteKnowledgeFileResponseBodyData {
	return s.Data
}

func (s *DeleteKnowledgeFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeFileResponseBody) SetData(v *DeleteKnowledgeFileResponseBodyData) *DeleteKnowledgeFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteKnowledgeFileResponseBody) SetRequestId(v string) *DeleteKnowledgeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteKnowledgeFileResponseBodyData struct {
	// The file location. Currently, only OSS paths are supported.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKnowledgeFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeFileResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *DeleteKnowledgeFileResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteKnowledgeFileResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKnowledgeFileResponseBodyData) SetFileLocation(v string) *DeleteKnowledgeFileResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *DeleteKnowledgeFileResponseBodyData) SetMessage(v string) *DeleteKnowledgeFileResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteKnowledgeFileResponseBodyData) SetSuccess(v bool) *DeleteKnowledgeFileResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteKnowledgeFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
