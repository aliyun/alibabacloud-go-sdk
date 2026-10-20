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
	// The business data field returned by the operation.
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
	// The status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
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
	// The list of matched text chunks.
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
	// The metadata map of the text chunk.
	//
	// > The file_path field in the metadata map of a document search knowledge base is meaningless. Do not use it in your business code.
	//
	// > When you retrieve a document search knowledge base, if a chunk contains an image, the image is returned in the image_url field of the metadata map, along with an expiration time.
	//
	// > When you retrieve an audio and video search knowledge base, if a chunk contains audio, the audio is returned in the audio_url field of the metadata map, along with an expiration time.
	//
	// > When you retrieve an audio and video search knowledge base, if a chunk contains a video, the video is returned in the video_url field of the metadata map, along with an expiration time.
	//
	//
	// <props="intl">
	//
	// The metadata map of the text chunk.
	//
	// > The file_path field in the metadata map of a document search knowledge base is meaningless. Do not use it in your business code.
	//
	// > When you retrieve a document search knowledge base, if a chunk contains an image, the image is returned in the image_url field of the metadata map, along with an expiration time.
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
	//   "title": "Alibaba Cloud Model Studio Documentation",
	//
	//   "doc_id": "doc_***",
	//
	//   "content": "Alibaba Cloud Model Studio is a one-stop large model development platform based on Tongyi large models, industry large models, and third-party large models. It provides enterprise customers and individual developers with a complete set of model service tools and full-link application development kits, pre-installed with rich capability plugins, and offers convenient integration methods such as APIs and SDKs to efficiently build large model applications.",
	//
	//   "workspace_id": "ws_***",
	//
	//   "hier_title": "Alibaba Cloud Model Studio Documentation",
	//
	//   "doc_name": "Introduction to Alibaba Cloud Model Studio Documentation.pdpf",
	//
	//   "pipeline_id": "rhd***",
	//
	//   "_id": "ws_***"
	//
	// }
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The similarity score of the text chunk. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The content of the text chunk.
	//
	// example:
	//
	// Alibaba Cloud Model Studio is a one-stop large model development platform based on Tongyi large models, industry large models, and third-party large models. It provides enterprise customers and individual developers with a complete set of model service tools and full-link application development kits, pre-installed with rich capability plugins, and offers convenient integration methods such as APIs and SDKs to efficiently build large model applications
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
