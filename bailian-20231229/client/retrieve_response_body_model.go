// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetrieveResponseBody
	GetCode() *string
	SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody
	GetData() *RetrieveResponseBodyData
	SetMessage(v string) *RetrieveResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetrieveResponseBody
	GetRequestId() *string
	SetStatus(v string) *RetrieveResponseBody
	GetStatus() *string
	SetSuccess(v bool) *RetrieveResponseBody
	GetSuccess() *bool
}

type RetrieveResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by the API.
	Data *RetrieveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code of the response.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - true: The call succeeded.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetrieveResponseBody) GetData() *RetrieveResponseBodyData {
	return s.Data
}

func (s *RetrieveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetrieveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RetrieveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetrieveResponseBody) SetCode(v string) *RetrieveResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveResponseBody) SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveResponseBody) SetMessage(v string) *RetrieveResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveResponseBody) SetRequestId(v string) *RetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveResponseBody) SetStatus(v string) *RetrieveResponseBody {
	s.Status = &v
	return s
}

func (s *RetrieveResponseBody) SetSuccess(v bool) *RetrieveResponseBody {
	s.Success = &v
	return s
}

func (s *RetrieveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveResponseBodyData struct {
	// An array of retrieved text chunks.
	Nodes []*RetrieveResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RetrieveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyData) GetNodes() []*RetrieveResponseBodyDataNodes {
	return s.Nodes
}

func (s *RetrieveResponseBodyData) SetNodes(v []*RetrieveResponseBodyDataNodes) *RetrieveResponseBodyData {
	s.Nodes = v
	return s
}

func (s *RetrieveResponseBodyData) Validate() error {
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

type RetrieveResponseBodyDataNodes struct {
	// <props="china">
	//
	// A map of metadata for the text chunk.
	//
	// > For document search knowledge bases, the `file_path` field in the metadata map is not applicable and should not be used in your application code.
	//
	// > When you retrieve data from a document search knowledge base, if a text chunk contains an image, its URL is returned in the `image_url` field of the metadata map. This URL expires.
	//
	// > When you retrieve data from an audio/video search knowledge base, if a text chunk contains audio, its URL is returned in the `audio_url` field of the metadata map. This URL expires.
	//
	// > When you retrieve data from an audio/video search knowledge base, if a text chunk contains video, its URL is returned in the `video_url` field of the metadata map. This URL expires.
	//
	//
	//
	// <props="intl">
	//
	// A map of metadata for the text chunk.
	//
	// > For document search knowledge bases, the `file_path` field in the metadata map is not applicable and should not be used in your application code.
	//
	// > When you retrieve data from a document search knowledge base, if a text chunk contains an image, its URL is returned in the `image_url` field of the metadata map. This URL expires.
	//
	// example:
	//
	// {
	//
	//   "parent": "",
	//
	//   "file_path": "https://***",
	//
	//   "image_url": [
	//
	//     "http://***"
	//
	//   ],
	//
	//   "nid": "***",
	//
	//   "title": "阿里云百炼文档",
	//
	//   "doc_id": "doc_***",
	//
	//   "content": "阿里云百炼是基于通义大模型、行业大模型以及三方大模型的一站式大模型开发平台。面向企业客户和个人开发者，提供完整的模型服务工具和全链路应用开发套件，预置丰富的能力插件，提供API及SDK等便捷的集成方式，高效完成大模型应用构建",
	//
	//   "workspace_id": "ws_***",
	//
	//   "hier_title": "阿里云百炼文档",
	//
	//   "doc_name": "阿里云百炼文档介绍.pdpf",
	//
	//   "pipeline_id": "rhd***",
	//
	//   "_id": "ws_***"
	//
	// }
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The similarity score of the text chunk, ranging from 0 to 1.
	//
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The content of the text chunk.
	//
	// example:
	//
	// 阿里云百炼是基于通义大模型、行业大模型以及三方大模型的一站式大模型开发平台。面向企业客户和个人开发者，提供完整的模型服务工具和全链路应用开发套件，预置丰富的能力插件，提供API及SDK等便捷的集成方式，高效完成大模型应用构建。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RetrieveResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyDataNodes) GetMetadata() interface{} {
	return s.Metadata
}

func (s *RetrieveResponseBodyDataNodes) GetScore() *float64 {
	return s.Score
}

func (s *RetrieveResponseBodyDataNodes) GetText() *string {
	return s.Text
}

func (s *RetrieveResponseBodyDataNodes) SetMetadata(v interface{}) *RetrieveResponseBodyDataNodes {
	s.Metadata = v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetScore(v float64) *RetrieveResponseBodyDataNodes {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetText(v string) *RetrieveResponseBodyDataNodes {
	s.Text = &v
	return s
}

func (s *RetrieveResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
