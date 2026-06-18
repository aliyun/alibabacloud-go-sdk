// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChunksResponseBody
	GetCode() *string
	SetData(v *ListChunksResponseBodyData) *ListChunksResponseBody
	GetData() *ListChunksResponseBodyData
	SetMessage(v string) *ListChunksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChunksResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListChunksResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListChunksResponseBody
	GetSuccess() *bool
}

type ListChunksResponseBody struct {
	// The error status code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by the operation.
	Data *ListChunksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F97A63B-xxxx-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChunksResponseBody) GetData() *ListChunksResponseBodyData {
	return s.Data
}

func (s *ListChunksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChunksResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChunksResponseBody) SetCode(v string) *ListChunksResponseBody {
	s.Code = &v
	return s
}

func (s *ListChunksResponseBody) SetData(v *ListChunksResponseBodyData) *ListChunksResponseBody {
	s.Data = v
	return s
}

func (s *ListChunksResponseBody) SetMessage(v string) *ListChunksResponseBody {
	s.Message = &v
	return s
}

func (s *ListChunksResponseBody) SetRequestId(v string) *ListChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChunksResponseBody) SetStatus(v string) *ListChunksResponseBody {
	s.Status = &v
	return s
}

func (s *ListChunksResponseBody) SetSuccess(v bool) *ListChunksResponseBody {
	s.Success = &v
	return s
}

func (s *ListChunksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChunksResponseBodyData struct {
	// The list of text chunks.
	Nodes []*ListChunksResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The total number of returned results.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChunksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyData) GetNodes() []*ListChunksResponseBodyDataNodes {
	return s.Nodes
}

func (s *ListChunksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListChunksResponseBodyData) SetNodes(v []*ListChunksResponseBodyDataNodes) *ListChunksResponseBodyData {
	s.Nodes = v
	return s
}

func (s *ListChunksResponseBodyData) SetTotal(v int64) *ListChunksResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListChunksResponseBodyData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChunksResponseBodyDataNodes struct {
	// <props="china">
	//
	// The metadata map of the text chunk.
	//
	// > The `file_path` field in the metadata map of document search knowledge bases is meaningless. Do not use it in your business code.
	//
	// > When retrieving a document search knowledge base, if a chunk contains an image, the image is returned through the `image_url` field in the metadata map, along with an expiration time.
	//
	// > When retrieving an audio/video search knowledge base, if a chunk contains audio, the audio is returned through the `audio_url` field in the metadata map, along with an expiration time.
	//
	// > When retrieving an audio/video search knowledge base, if a chunk contains video, the video is returned through the `video_url` field in the metadata map, along with an expiration time.
	//
	//
	// <props="intl">
	//
	// The metadata map of the text chunk.
	//
	// > The `file_path` field in the metadata map of document search knowledge bases is meaningless. Do not use it in your business code.
	//
	// > When retrieving a document search knowledge base, if a chunk contains an image, the image is returned through the `image_url` field in the metadata map, along with an expiration time.
	//
	// .
	//
	// example:
	//
	// {
	//
	//   "file_path": "https://bailian-***",
	//
	//   "parent": "阿里云百炼是一站式的大模型开发及应用构建平台。不论是开发者还是业务人员，都能深入参与大模型应用的设计和构建。您可以通过简单的界面操作，在 5分钟内开发出一款大模型应用，或在几小时内训练出一个专属模型，从而将更多精力专注于应用创新。",
	//
	//   "is_displayed_chunk_content": "true",
	//
	//   "image_url": [],
	//
	//   "nid": "83***",
	//
	//   "source": "0",
	//
	//   "_score": 0,
	//
	//   "title": "",
	//
	//   "doc_id": "file_24e***",
	//
	//   "content": "阿里云百炼是一站式的大模型开发及应用构建平台。不论是开发者还是业务人员，都能深入参与大模型应用的设计和构建。您可以通过简单的界面操作，在 5分钟内开发出一款大模型应用，或在几小时内训练出一个专属模型，从而将更多精力专注于应用创新。",
	//
	//   "_rc_score": 0,
	//
	//   "workspace_id": "llm-zna***",
	//
	//   "hier_title": "",
	//
	//   "doc_name": "什么是阿里云百炼",
	//
	//   "pipeline_id": "j6b***",
	//
	//   "_id": "llm-zna5***"
	//
	// }
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The similarity score of the text chunk.
	//
	// example:
	//
	// 0
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The content of the text chunk.
	//
	// example:
	//
	// 阿里云百炼是一站式的大模型开发及应用构建平台。不论是开发者还是业务人员，都能深入参与大模型应用的设计和构建。您可以通过简单的界面操作，在 5分钟内开发出一款大模型应用，或在几小时内训练出一个专属模型，从而将更多精力专注于应用创新。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListChunksResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyDataNodes) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListChunksResponseBodyDataNodes) GetScore() *float64 {
	return s.Score
}

func (s *ListChunksResponseBodyDataNodes) GetText() *string {
	return s.Text
}

func (s *ListChunksResponseBodyDataNodes) SetMetadata(v interface{}) *ListChunksResponseBodyDataNodes {
	s.Metadata = v
	return s
}

func (s *ListChunksResponseBodyDataNodes) SetScore(v float64) *ListChunksResponseBodyDataNodes {
	s.Score = &v
	return s
}

func (s *ListChunksResponseBodyDataNodes) SetText(v string) *ListChunksResponseBodyDataNodes {
	s.Text = &v
	return s
}

func (s *ListChunksResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
