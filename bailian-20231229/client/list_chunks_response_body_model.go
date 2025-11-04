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
	// The error code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
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
	// 8F97A63B-55F1-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The list of chunks.
	Nodes []*ListChunksResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The total number of chunks returned.
	//
	// example:
	//
	// 16
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
	// The metadata map of the chunk.
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The similarity score of the chunk.
	//
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The text of the chunk.
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
